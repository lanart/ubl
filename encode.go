package ubl

import (
	"encoding/xml"
)

// Invoice represents the root element of the UBL Invoice
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
	SupplierParty      SupplierParty `xml:"cac:AccountingSupplierParty"`
	CustomerParty      CustomerParty `xml:"cac:AccountingCustomerParty"`
	PaymentMeans       PaymentMeans  `xml:"cac:PaymentMeans"`
	PaymentTerms       PaymentTerms  `xml:"cac:PaymentTerms"`
	TaxTotal           TaxTotal      `xml:"cac:TaxTotal"`
	LegalMonetaryTotal MonetaryTotal `xml:"cac:LegalMonetaryTotal"`
	InvoiceLines       []InvoiceLine `xml:"cac:InvoiceLine"`
}

// SupplierParty represents the supplier's details
type SupplierParty struct {
	Party Party `xml:"cac:Party"`
}

// CustomerParty represents the customer's details
type CustomerParty struct {
	Party Party `xml:"cac:Party"`
}

// Party represents a general party structure
type Party struct {
	PartyName     string        `xml:"cac:PartyName>cbc:Name"`
	PostalAddress PostalAddress `xml:"cac:PostalAddress"`
}

// PostalAddress represents address details
type PostalAddress struct {
	StreetName string  `xml:"cbc:StreetName"`
	CityName   string  `xml:"cbc:CityName"`
	PostalZone string  `xml:"cbc:PostalZone"`
	Country    Country `xml:"cac:Country"`
}

// Country represents country details
type Country struct {
	IdentificationCode string `xml:"cbc:IdentificationCode"`
}

// PaymentMeans represents the payment method details for the invoice
type PaymentMeans struct {
	PaymentMeansCode      string           `xml:"cbc:PaymentMeansCode"`
	PayeeFinancialAccount FinancialAccount `xml:"cac:PayeeFinancialAccount"`
}

// FinancialAccount represents the bank account details (ID = IBAN)
type FinancialAccount struct {
	ID                         string                     `xml:"cbc:ID"`
	FinancialInstitutionBranch FinancialInstitutionBranch `xml:"cac:FinancialInstitutionBranch"`
}

// FinancialInstitutionBranch (ID = BIC)
type FinancialInstitutionBranch struct {
	ID string `xml:"cbc:ID"`
}

// PaymentTerms represents the payment terms for the invoice
type PaymentTerms struct {
	Note string `xml:"cbc:Note"`
}

// TaxTotal represents the tax total for the invoice
type TaxTotal struct {
	TaxAmount Amount `xml:"cbc:TaxAmount"`
}

// MonetaryTotal represents the total monetary amount
type MonetaryTotal struct {
	LineExtensionAmount Amount `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount  Amount `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount  Amount `xml:"cbc:TaxInclusiveAmount"`
	PayableAmount       Amount `xml:"cbc:PayableAmount"`
}

// Amount represents a monetary amount with a currency attribute
type Amount struct {
	Value      float32 `xml:",chardata"`
	CurrencyID string  `xml:"currencyID,attr"`
}

// InvoiceLine represents individual line items
type InvoiceLine struct {
	ID                  string  `xml:"cbc:ID"`
	InvoicedQuantity    float32 `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount Amount  `xml:"cbc:LineExtensionAmount"`
	Item                Item    `xml:"cac:Item"`
	Price               Price   `xml:"cac:Price"`
}

// Item represents an item being invoiced
type Item struct {
	Description string `xml:"cbc:Description"`
	Name        string `xml:"cbc:Name"`
}

// Price represents pricing details for an item
type Price struct {
	PriceAmount Amount `xml:"cbc:PriceAmount"`
}

func CreateInvoice(invoice *Invoice) ([]byte, error) {
	return xml.MarshalIndent(invoice, "", "  ")
}
