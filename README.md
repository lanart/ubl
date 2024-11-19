[![Go](https://github.com/lanart/ubl/actions/workflows/go.yml/badge.svg)](https://github.com/lanart/ubl/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lanart/ubl.svg)](https://pkg.go.dev/github.com/lanart/ubl)

A Go package to create a UBL invoice for Peppol. 
This is a minimal implementation for the features we need.
Feel free to send a PR if you are missing something.


Example:

```go
inv := ubl.NewInvoice("INV-12345")

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
		TaxPercentage: 21.0,
    },
})

xmlBytes, err := inv.UblBytes()

v, err := validate.New()
defer v.Free()

err = v.ValidateBytes(xmlBytes)

os.WriteFile("invoice.xml", xmlBytes, 0644)
```
