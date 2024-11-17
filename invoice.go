// Package ubl contains helpers to create a UBL Invoice.
//
// This is needed for Peppol: https://docs.peppol.eu/poacc/billing/3.0/
// Specification: https://docs.peppol.eu/poacc/billing/3.0/syntax/ubl-invoice/tree/
// The result can be validated with the validate package.
package ubl

import (
	"encoding/xml"
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

func (inv *Invoice) UblBytes() ([]byte, error) {
	return xml.MarshalIndent(inv, "", "  ")
}
