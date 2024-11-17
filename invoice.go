// Package ubl contains helpers to create a UBL Invoice.
//
// This is needed for Peppol: https://docs.peppol.eu/poacc/billing/3.0/
// Specification: https://docs.peppol.eu/poacc/billing/3.0/syntax/ubl-invoice/tree/
// The result can be validated with the validate package.
package ubl

import (
	"encoding/xml"
	"math"
	"strconv"
	"time"
)

// Invoice is the root element to create a UBL invoice
type Invoice struct {
	XMLName            xml.Name      `xml:"Invoice"`
	Xmlns              string        `xml:"xmlns,attr"`
	Cac                string        `xml:"xmlns:cac,attr"`
	Cbc                string        `xml:"xmlns:cbc,attr"`
	CustomizationID    string        `xml:"cbc:CustomizationID"`
	ProfileID          string        `xml:"cbc:ProfileID"`
	ID                 string        `xml:"cbc:ID"`
	IssueDate          string        `xml:"cbc:IssueDate"`
	DueDate            string        `xml:"cbc:DueDate"`
	InvoiceTypeCode    string        `xml:"cbc:InvoiceTypeCode"`
	DocumentCurrency   string        `xml:"cbc:DocumentCurrencyCode"`
	BuyerReference     string        `xml:"cbc:BuyerReference"`
	OrderReference     string        `xml:"cac:OrderReference>cbc:ID"`
	SupplierParty      SupplierParty `xml:"cac:AccountingSupplierParty"`
	CustomerParty      CustomerParty `xml:"cac:AccountingCustomerParty"`
	PaymentMeans       PaymentMeans  `xml:"cac:PaymentMeans"`
	PaymentTerms       PaymentTerms  `xml:"cac:PaymentTerms"`
	TaxTotal           TaxTotal      `xml:"cac:TaxTotal"`
	LegalMonetaryTotal MonetaryTotal `xml:"cac:LegalMonetaryTotal"`
	InvoiceLines       []InvoiceLine `xml:"cac:InvoiceLine"`
}

// NewInvoice initializes a new Invoice struct
func NewInvoice() *Invoice {
	return &Invoice{
		Xmlns:            "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
		Cac:              "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2",
		Cbc:              "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
		CustomizationID:  "urn:cen.eu:en16931:2017",
		ProfileID:        "urn:fdc:peppol.eu:2017:poacc:billing:01:1.0",
		IssueDate:        time.Now().Format("2006-01-02"),
		DueDate:          time.Now().AddDate(0, 0, 30).Format("2006-01-02"),
		InvoiceTypeCode:  "380",
		DocumentCurrency: "EUR",
	}
}

// InitSupplier sets the supplier name and vat id
func (inv *Invoice) InitSupplier(name, companyID string) {
	inv.SupplierParty = SupplierParty{
		Party: Party{
			PartyName: name,
			PartyTaxScheme: PartyTaxScheme{
				CompanyID: companyID,
				TaxScheme: TaxScheme{
					ID: "VAT",
				},
			},
		},
	}
}

// InitCustomer sets the customer name and vat id
func (inv *Invoice) InitCustomer(name, companyID string) {

	inv.CustomerParty = CustomerParty{
		Party: Party{
			PartyName: name,
			PartyTaxScheme: PartyTaxScheme{
				CompanyID: companyID,
				TaxScheme: TaxScheme{
					ID: "VAT",
				},
			},
		},
	}
}

// InitPaymentMeans inits iban and bic
func (inv *Invoice) InitPaymentMeans(iban, bic string) {
	inv.PaymentMeans = PaymentMeans{
		PaymentMeansCode: "1",
		PayeeFinancialAccount: FinancialAccount{
			ID: iban,
			FinancialInstitutionBranch: FinancialInstitutionBranch{
				ID: bic,
			},
		},
	}
}

// InitPaymentTerms adds a note
func (inv *Invoice) InitPaymentTerms(note string) {
	inv.PaymentTerms = PaymentTerms{
		Note: note,
	}
}

// UblBytes converts the invoice to xml bytes
func (inv *Invoice) UblBytes() ([]byte, error) {
	return xml.MarshalIndent(inv, "", "  ")
}

// InvoiceLineHelper is a helper for adding lines and tax
type InvoiceLineHelper struct {
	Quantity    float64
	Price       float64
	Name        string
	Description string
}

func round(amount float64) float64 {
	return math.Round(amount*100) / 100
}

// AddLines is a helper to add InvoiceLines and Tax elements
func (inv *Invoice) AddLines(lines []InvoiceLineHelper) {
	sum := 0.0
	sumTax := 0.0
	for i, line := range lines {
		lineAmountExcl := round(line.Quantity * line.Price)
		tax := round(lineAmountExcl * 0.21)
		sum = sum + lineAmountExcl
		sumTax = sumTax + tax
		invoiceLine := InvoiceLine{
			ID:                  strconv.Itoa(i + 1),
			InvoicedQuantity:    Quantity{Value: line.Quantity, UnitCode: "ZZ"},
			LineExtensionAmount: Amount{Value: lineAmountExcl, CurrencyID: "EUR"},
			TaxTotal: TaxTotal{
				TaxAmount: Amount{Value: tax, CurrencyID: "EUR"},
			},
			Item: Item{
				Name:        line.Name,
				Description: line.Description,
				ClassifiedTaxCategory: TaxCategory{
					ID:      "S",
					Name:    "03",
					Percent: 21,
					TaxScheme: TaxScheme{
						ID: "VAT",
					},
				},
			},
			Price: Price{
				PriceAmount: Amount{
					Value:      line.Price,
					CurrencyID: "EUR",
				},
			},
		}
		inv.InvoiceLines = append(inv.InvoiceLines, invoiceLine)
	}

	inv.TaxTotal = TaxTotal{
		TaxAmount: Amount{Value: 20.0, CurrencyID: "EUR"},
	}

	total := round(sum + sumTax)

	inv.LegalMonetaryTotal = MonetaryTotal{
		LineExtensionAmount: Amount{Value: sum, CurrencyID: "EUR"},
		TaxExclusiveAmount:  Amount{Value: sum, CurrencyID: "EUR"},
		TaxInclusiveAmount:  Amount{Value: total, CurrencyID: "EUR"},
		PayableAmount:       Amount{Value: total, CurrencyID: "EUR"},
	}
}
