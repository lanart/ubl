package ubl_test

import (
	"testing"

	"github.com/lanart/ubl"
	"github.com/lanart/ubl/validate"
)

func TestNewInvoice(t *testing.T) {

	inv := ubl.NewInvoice()
	inv.ID = "INV-12345"
	inv.OrderReference = "INV-12345"

	inv.InitSupplier("ABC Supplies Ltd", "BE0123456789")

	inv.SupplierParty.Party.PostalAddress = ubl.PostalAddress{
		StreetName: "123 Supplier Street",
		CityName:   "Supplier City",
		PostalZone: "12345",
		Country:    ubl.Country{IdentificationCode: "BE"},
	}

	inv.InitCustomer("XYZ Corp", "BE9876543210")

	inv.CustomerParty.Party.PostalAddress = ubl.PostalAddress{
		StreetName: "789 Customer Avenue",
		CityName:   "Customer Town",
		PostalZone: "67890",
		Country:    ubl.Country{IdentificationCode: "BE"},
	}

	inv.InitPaymentMeans("9999999999", "GEBABEBB")

	inv.InitPaymentTerms("You get a free sticker when you pay fast")

	inv.AddLines([]ubl.InvoiceLineHelper{
		ubl.InvoiceLineHelper{
			Quantity:    10,
			Price:       100,
			Name:        "Product A",
			Description: "High-quality item",
		},
	})

	xmlBytes, err := inv.UblBytes()

	v, err := validate.New()
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

func TestNewInvoiceCustom(t *testing.T) {

	inv := ubl.NewInvoice()
	inv.ID = "INV-12345"
	inv.OrderReference = "INV-12345"
	inv.SupplierParty = ubl.SupplierParty{
		Party: ubl.Party{
			PartyName: "ABC Supplies Ltd",
			PostalAddress: ubl.PostalAddress{
				StreetName: "123 Supplier Street",
				CityName:   "Supplier City",
				PostalZone: "12345",
				Country:    ubl.Country{IdentificationCode: "GB"},
			},
			PartyTaxScheme: ubl.PartyTaxScheme{
				CompanyID: "BE0123456789",
				TaxScheme: ubl.TaxScheme{
					ID: "VAT",
				},
			},
		},
	}
	inv.CustomerParty = ubl.CustomerParty{
		Party: ubl.Party{
			PartyName: "XYZ Corp",
			PostalAddress: ubl.PostalAddress{
				StreetName: "789 Customer Avenue",
				CityName:   "Customer Town",
				PostalZone: "67890",
				Country:    ubl.Country{IdentificationCode: "DE"},
			},
		},
	}
	inv.PaymentMeans = ubl.PaymentMeans{
		PaymentMeansCode: "1",
		PayeeFinancialAccount: ubl.FinancialAccount{
			ID: "9999999999",
			FinancialInstitutionBranch: ubl.FinancialInstitutionBranch{
				ID: "GEBABEBB",
			},
		},
	}

	inv.PaymentTerms = ubl.PaymentTerms{
		Note: "You get a free sticker when you pay fast",
	}

	inv.TaxTotal = ubl.TaxTotal{
		TaxAmount: ubl.Amount{Value: 20.0, CurrencyID: "EUR"},
		// TaxSubTotal: ubl.TaxSubtotal{
		// 	TaxableAmount: ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
		// 	TaxAmount:     ubl.Amount{Value: 20.0, CurrencyID: "EUR"},
		// 	TaxCategory: ubl.TaxCategory{
		// 		ID:      "S",
		// 		Name:    "03",
		// 		Percent: 21,
		// 		TaxScheme: ubl.TaxScheme{
		// 			ID: "VAT",
		// 		},
		// 	},
		// },
	}

	inv.LegalMonetaryTotal = ubl.MonetaryTotal{
		LineExtensionAmount: ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
		TaxExclusiveAmount:  ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
		TaxInclusiveAmount:  ubl.Amount{Value: 120.0, CurrencyID: "EUR"},
		PayableAmount:       ubl.Amount{Value: 120.0, CurrencyID: "EUR"},
	}
	inv.InvoiceLines = []ubl.InvoiceLine{
		{
			ID:                  "1",
			InvoicedQuantity:    ubl.Quantity{Value: 10, UnitCode: "A9"},
			LineExtensionAmount: ubl.Amount{Value: 100.0, CurrencyID: "EUR"},
			Item: ubl.Item{
				Name:        "Product A",
				Description: "High-quality item",
				ClassifiedTaxCategory: ubl.TaxCategory{
					ID:      "S",
					Name:    "03",
					Percent: 21,
					TaxScheme: ubl.TaxScheme{
						ID: "VAT",
					},
				},
			},
			Price: ubl.Price{PriceAmount: ubl.Amount{Value: 10.0, CurrencyID: "EUR"}},
		},
	}

	xmlBytes, err := inv.UblBytes()

	v, err := validate.New()
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
