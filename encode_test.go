package ubl_test

import (
	"testing"
	"time"

	"github.com/lanart/ubl"
	"github.com/lanart/ubl/validator"
)

func TestEncode(t *testing.T) {

	invoice := ubl.Invoice{
		Xmlns:            "urn:oasis:names:specification:ubl:schema:xsd:Invoice-2",
		Cac:              "urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2",
		Cbc:              "urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2",
		CustomizationID:  "urn:cen.eu:en16931:2017",
		ProfileID:        "urn:fdc:peppol.eu:2017:poacc:billing:01:1.0",
		ID:               "INV-12345",
		IssueDate:        time.Now().Format("2006-01-02"),
		DueDate:          time.Now().Format("2006-01-02"),
		InvoiceTypeCode:  "380",
		DocumentCurrency: "EUR",
		BuyerReference:   "INV-12345",
		SupplierParty: ubl.SupplierParty{
			Party: ubl.Party{
				PartyName: "ABC Supplies Ltd",
				PostalAddress: ubl.PostalAddress{
					StreetName: "123 Supplier Street",
					CityName:   "Supplier City",
					PostalZone: "12345",
					Country:    ubl.Country{IdentificationCode: "GB"},
				},
			},
		},
		CustomerParty: ubl.CustomerParty{
			Party: ubl.Party{
				PartyName: "XYZ Corp",
				PostalAddress: ubl.PostalAddress{
					StreetName: "789 Customer Avenue",
					CityName:   "Customer Town",
					PostalZone: "67890",
					Country:    ubl.Country{IdentificationCode: "DE"},
				},
			},
		},
		PaymentMeans: ubl.PaymentMeans{
			PaymentMeansCode: "1",
			PayeeFinancialAccount: ubl.FinancialAccount{
				ID: "BE0123456789",
				FinancialInstitutionBranch: ubl.FinancialInstitutionBranch{
					ID: "GEBABEBB",
				},
			},
		},

		TaxTotal: ubl.TaxTotal{
			TaxAmount: ubl.Amount{Value: 20.0, CurrencyID: "EUR"},
		},
		LegalMonetaryTotal: ubl.MonetaryTotal{
			LineExtensionAmount: ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
			TaxExclusiveAmount:  ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
			TaxInclusiveAmount:  ubl.Amount{Value: 120.0, CurrencyID: "EUR"},
			PayableAmount:       ubl.Amount{Value: 120.0, CurrencyID: "EUR"},
		},
		InvoiceLines: []ubl.InvoiceLine{
			{
				ID:                  "1",
				InvoicedQuantity:    10,
				LineExtensionAmount: ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
				Item:                ubl.Item{Name: "Product A", Description: "High-quality item"},
				Price:               ubl.Price{PriceAmount: ubl.Amount{Value: 10.0, CurrencyID: "EUR"}},
			},
		},
	}

	xmlBytes, err := ubl.CreateInvoice(&invoice)

	v, err := validator.New("./validator/xsd")
	if err != nil {
		t.Error(err)
	}

	defer v.Free()

	err = v.ValidateBytes(xmlBytes)
	if err != nil {
		t.Error(err)
	}

	// // Write to file or print
	// err = os.WriteFile("invoice.xml", xmlBytes, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing XML file:", err)
	// }

}
