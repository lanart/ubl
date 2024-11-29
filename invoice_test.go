package ubl_test

import (
	"testing"

	"github.com/lanart/ubl"
	"github.com/lanart/ubl/validate"
)

func TestNewInvoice(t *testing.T) {

	inv := ubl.Invoice{
		ID:           "INV-12345",
		SupplierName: "ABC Supplies Ltd",
		SupplierVat:  "BE0123456789",
		SupplierAddress: ubl.Address{
			StreetName:  "123 Supplier Street",
			CityName:    "Supplier City",
			PostalZone:  "12345",
			CountryCode: "BE",
		},
		CustomerName: "XYZ Corp",
		CustomerVat:  "BE9876543210",
		CustomerAddress: ubl.Address{
			StreetName:  "789 Customer Avenue",
			CityName:    "Customer Town",
			PostalZone:  "67890",
			CountryCode: "BE",
		},
		Iban: "9999999999",
		Bic:  "GEBABEBB",
		Note: "You get a free sticker when you pay fast",
	}

	inv.Lines = []ubl.InvoiceLine{
		ubl.InvoiceLine{
			Quantity:      10,
			Price:         100,
			Name:          "Product A",
			Description:   "High-quality item",
			TaxPercentage: 21.0,
		},
	}

	xmlBytes, err := inv.Generate()

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
