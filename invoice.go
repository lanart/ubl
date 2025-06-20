// Package ubl is a basic implementation to create a UBL Invoice.
//
// This is needed for Peppol: https://docs.peppol.eu/poacc/billing/3.0/
// Specification: https://docs.peppol.eu/poacc/billing/3.0/syntax/ubl-invoice/tree/
// The result can be validated with the validate package.
package ubl

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Invoice struct {
	xml                *xmlInvoice
	ID                 string
	SupplierName       string
	SupplierVat        string
	SupplierAddress    Address
	CustomerName       string
	CustomerVat        string
	CustomerAddress    Address
	Iban               string
	Bic                string
	Note               string
	Lines              []InvoiceLine
	PdfInvoiceFilename string
}

type InvoiceLine struct {
	Quantity      float64
	Price         float64
	TaxPercentage float64
	Name          string
	Description   string
}

type Address struct {
	StreetName  string
	CityName    string
	PostalZone  string
	CountryCode string
}

func (inv *Invoice) Generate() ([]byte, error) {

	inv.xml = &xmlInvoice{
		Xmlns:            "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
		Cac:              "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2",
		Cbc:              "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
		CustomizationID:  "urn:cen.eu:en16931:2017#conformant#urn:UBL.BE:1.0.0.20180214",
		ProfileID:        "urn:fdc:peppol.eu:2017:poacc:billing:01:1.0",
		IssueDate:        time.Now().Format("2006-01-02"),
		DueDate:          time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
		InvoiceTypeCode:  "380",
		DocumentCurrency: "EUR",
		ID:               inv.ID,
		OrderReference:   inv.ID,
	}

	inv.xml.SupplierParty = xmlSupplierParty{
		Party: xmlParty{
			EndpointID: xmlEndpointID{

				Value:    inv.SupplierVat,
				SchemeID: "9925",
			},
			PartyName:        inv.SupplierName,
			RegistrationName: inv.SupplierName,
			PartyTaxScheme: xmlPartyTaxScheme{
				CompanyID: inv.SupplierVat,
				TaxScheme: xmlTaxScheme{
					ID: "VAT",
				},
			},
		},
	}

	inv.xml.SupplierParty.Party.PostalAddress = xmlPostalAddress{
		StreetName: inv.SupplierAddress.StreetName,
		CityName:   inv.SupplierAddress.CityName,
		PostalZone: inv.SupplierAddress.PostalZone,
		Country:    xmlCountry{IdentificationCode: inv.SupplierAddress.CountryCode},
	}

	inv.xml.CustomerParty.Party.PostalAddress = xmlPostalAddress{
		StreetName: inv.CustomerAddress.StreetName,
		CityName:   inv.CustomerAddress.CityName,
		PostalZone: inv.CustomerAddress.PostalZone,
		Country:    xmlCountry{IdentificationCode: inv.CustomerAddress.CountryCode},
	}

	inv.xml.CustomerParty = xmlCustomerParty{
		Party: xmlParty{
			EndpointID: xmlEndpointID{

				Value:    inv.CustomerVat,
				SchemeID: "9925",
			},
			PartyName:        inv.CustomerName,
			RegistrationName: inv.CustomerName,
			PartyTaxScheme: xmlPartyTaxScheme{
				CompanyID: inv.CustomerVat,
				TaxScheme: xmlTaxScheme{
					ID: "VAT",
				},
			},
			PostalAddress: xmlPostalAddress{
				Country: xmlCountry{
					IdentificationCode: "BE",
				},
			},
		},
	}

	inv.xml.PaymentMeans = xmlPaymentMeans{
		PaymentMeansCode: "1",
		PayeeFinancialAccount: xmlFinancialAccount{
			ID: inv.Iban,
			FinancialInstitutionBranch: xmlFinancialInstitutionBranch{
				ID: inv.Bic,
			},
		},
	}

	if inv.Note != "" {
		inv.xml.PaymentTerms = xmlPaymentTerms{
			Note: inv.Note,
		}
	}

	inv.addLines()

	if inv.PdfInvoiceFilename != "" {
		err := inv.addAttachment(inv.PdfInvoiceFilename, "Invoice")
		if err != nil {
			return nil, fmt.Errorf("add attachment: %w", err)
		}
	}

	return xml.MarshalIndent(inv.xml, "", "  ")
}

func (inv *Invoice) addAttachment(filename, description string) error {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	mime := http.DetectContentType(bytes)

	encoded := base64.StdEncoding.EncodeToString(bytes)

	inv.xml.AdditionalDocumentReference = []xmlDocumentReference{
		{
			ID:                  "UBL.BE",
			DocumentDescription: "CommercialInvoice",
		},
		{
			ID:                  inv.ID,
			DocumentDescription: description,
			Attachment: []xmlAttachment{
				{xmlEmbeddedDocumentBinaryObject{
					Value:    encoded,
					MimeCode: mime,
					Filename: filename,
				}},
			},
		},
	}

	return nil

}

func round(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func (inv *Invoice) addLines() {
	sum := 0.0
	sumTax := 0.0
	taxPercentage := 21.0
	for i, line := range inv.Lines {
		taxPercentage = line.TaxPercentage
		lineAmountExcl := round(line.Quantity * line.Price)
		tax := round(lineAmountExcl * 0.21)
		sum = sum + lineAmountExcl
		sumTax = sumTax + tax
		invoiceLine := xmlInvoiceLine{
			ID:                  strconv.Itoa(i + 1),
			InvoicedQuantity:    xmlQuantity{Value: line.Quantity, UnitCode: "ZZ"},
			LineExtensionAmount: xmlAmount{Value: lineAmountExcl, CurrencyID: "EUR"},
			TaxTotal: xmlTaxTotal{
				TaxAmount: xmlAmount{Value: tax, CurrencyID: "EUR"},
			},
			Item: xmlItem{
				Name:        line.Name,
				Description: line.Description,
				ClassifiedTaxCategory: xmlTaxCategory{
					ID:      "S",
					Name:    "03",
					Percent: line.TaxPercentage,
					TaxScheme: xmlTaxScheme{
						ID: "VAT",
					},
				},
			},
			Price: xmlPrice{
				PriceAmount: xmlAmount{
					Value:      line.Price,
					CurrencyID: "EUR",
				},
			},
		}
		inv.xml.InvoiceLines = append(inv.xml.InvoiceLines, invoiceLine)
	}

	inv.xml.TaxTotal = xmlTaxTotal{
		TaxAmount: xmlAmount{Value: sumTax, CurrencyID: "EUR"},
		TaxSubtotal: []xmlTaxSubtotal{
			{
				TaxableAmount: xmlAmount{
					Value:      sum,
					CurrencyID: "EUR",
				},
				TaxAmount: xmlAmount{Value: sumTax, CurrencyID: "EUR"},
				TaxCategory: xmlTaxCategory{
					ID:      "S",
					Name:    "03",
					Percent: taxPercentage,
					TaxScheme: xmlTaxScheme{
						ID: "VAT",
					},
				},
			},
		},
	}

	total := round(sum + sumTax)

	inv.xml.LegalMonetaryTotal = xmlMonetaryTotal{
		LineExtensionAmount: xmlAmount{Value: sum, CurrencyID: "EUR"},
		TaxExclusiveAmount:  xmlAmount{Value: sum, CurrencyID: "EUR"},
		TaxInclusiveAmount:  xmlAmount{Value: total, CurrencyID: "EUR"},
		PayableAmount:       xmlAmount{Value: total, CurrencyID: "EUR"},
	}
}
