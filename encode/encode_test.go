package encode

import (
	"testing"
	"time"

	"github.com/lanart/ubl/validate"
)

func TestEncode(t *testing.T) {

	invoice := UBLInvoice{
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
		SupplierParty: SupplierParty{
			Party: Party{
				PartyName: "ABC Supplies Ltd",
				PostalAddress: PostalAddress{
					StreetName: "123 Supplier Street",
					CityName:   "Supplier City",
					PostalZone: "12345",
					Country:    Country{IdentificationCode: "GB"},
				},
			},
		},
		CustomerParty: CustomerParty{
			Party: Party{
				PartyName: "XYZ Corp",
				PostalAddress: PostalAddress{
					StreetName: "789 Customer Avenue",
					CityName:   "Customer Town",
					PostalZone: "67890",
					Country:    Country{IdentificationCode: "DE"},
				},
			},
		},
		PaymentMeans: PaymentMeans{
			PaymentMeansCode: "1",
			PayeeFinancialAccount: FinancialAccount{
				ID: "BE0123456789",
				FinancialInstitutionBranch: FinancialInstitutionBranch{
					ID: "GEBABEBB",
				},
			},
		},

		TaxTotal: TaxTotal{
			TaxAmount: Amount{Value: 20.0, CurrencyID: "EUR"},
		},
		LegalMonetaryTotal: MonetaryTotal{
			LineExtensionAmount: Amount{Value: 100.0, CurrencyID: "EUR"},
			TaxExclusiveAmount:  Amount{Value: 100.0, CurrencyID: "EUR"},
			TaxInclusiveAmount:  Amount{Value: 120.0, CurrencyID: "EUR"},
			PayableAmount:       Amount{Value: 120.0, CurrencyID: "EUR"},
		},
		InvoiceLines: []InvoiceLine{
			{
				ID:                  "1",
				InvoicedQuantity:    10,
				LineExtensionAmount: Amount{Value: 100.0, CurrencyID: "EUR"},
				Item:                Item{Name: "Product A", Description: "High-quality item"},
				Price:               Price{PriceAmount: Amount{Value: 10.0, CurrencyID: "EUR"}},
			},
		},
	}

	xmlBytes, err := CreateInvoice(&invoice)

	err = validate.Init("../validate/xsd")
	if err != nil {
		t.Error(err)
	}

	defer validate.Free()

	err = validate.Bytes(xmlBytes)
	if err != nil {
		t.Error(err)
	}

	// // Write to file or print
	// err = os.WriteFile("invoice.xml", xmlBytes, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing XML file:", err)
	// }

}
