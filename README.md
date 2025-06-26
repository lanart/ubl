[![Go](https://github.com/lanart/ubl/actions/workflows/go.yml/badge.svg)](https://github.com/lanart/ubl/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lanart/ubl.svg)](https://pkg.go.dev/github.com/lanart/ubl)

A Go package to create a UBL (UBL.BE) invoice for Peppol. 
This is a minimal implementation for the features we need.
Feel free to send a PR if you are missing something.

You can validate with schematron like this:
```
docker run --rm -e UBL_BE=true -e PLAIN_TEXT=true -v ./invoice.xml:/app/invoice.xml:ro ghcr.io/roel4d/peppol_schematron:latest
```

Example:

```go
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
defer v.Free()

err = v.ValidateBytes(xmlBytes)

os.WriteFile("invoice.xml", xmlBytes, 0644)
```

