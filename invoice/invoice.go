package invoice

import (
	"encoding/xml"
	"time"

	"github.com/lanart/ubl"
)

type Invoice struct {
	XMLName            xml.Name          `xml:"Invoice"`
	Xmlns              string            `xml:"xmlns,attr"`
	Cac                string            `xml:"xmlns:cac,attr"`
	Cbc                string            `xml:"xmlns:cbc,attr"`
	CustomizationID    string            `xml:"cbc:CustomizationID"`
	ProfileID          string            `xml:"cbc:ProfileID"`
	ID                 string            `xml:"cbc:ID"`
	IssueDate          string            `xml:"cbc:IssueDate"`
	DueDate            string            `xml:"cbc:DueDate"`
	InvoiceTypeCode    string            `xml:"cbc:InvoiceTypeCode"`
	DocumentCurrency   string            `xml:"cbc:DocumentCurrencyCode"`
	BuyerReference     string            `xml:"cbc:BuyerReference"`
	SupplierParty      ubl.SupplierParty `xml:"cac:AccountingSupplierParty"`
	CustomerParty      ubl.CustomerParty `xml:"cac:AccountingCustomerParty"`
	PaymentMeans       ubl.PaymentMeans  `xml:"cac:PaymentMeans"`
	PaymentTerms       ubl.PaymentTerms  `xml:"cac:PaymentTerms"`
	TaxTotal           ubl.TaxTotal      `xml:"cac:TaxTotal"`
	LegalMonetaryTotal ubl.MonetaryTotal `xml:"cac:LegalMonetaryTotal"`
	InvoiceLines       []ubl.InvoiceLine `xml:"cac:InvoiceLine"`
}

func New() *Invoice {
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
