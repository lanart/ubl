package ubl

type DocumentReference struct {
	ID                  string       `xml:"cbc:ID"`
	DocumentDescription string       `xml:"cbc:DocumentDescription"`
	Attachment          []Attachment `xml:"cac:Attachment"`
}

type Attachment struct {
	EmbeddedDocumentBinaryObject EmbeddedDocumentBinaryObject `xml:"cbc:EmbeddedDocumentBinaryObject"`
}

type EmbeddedDocumentBinaryObject struct {
	Value    string `xml:",chardata"`
	MimeCode string `xml:"mimeCode,attr"`
	Filename string `xml:"filename,attr"`
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
	PartyName      string         `xml:"cac:PartyName>cbc:Name"`
	PostalAddress  PostalAddress  `xml:"cac:PostalAddress"`
	PartyTaxScheme PartyTaxScheme `xml:"cac:PartyTaxScheme"`
}

// PostalAddress represents address details
type PostalAddress struct {
	StreetName string  `xml:"cbc:StreetName"`
	CityName   string  `xml:"cbc:CityName"`
	PostalZone string  `xml:"cbc:PostalZone"`
	Country    Country `xml:"cac:Country"`
}

// PartyTaxScheme represents the company tax ID
type PartyTaxScheme struct {
	CompanyID string    `xml:"cbc:CompanyID"`
	TaxScheme TaxScheme `xml:"cac:TaxScheme"`
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
	TaxAmount   Amount        `xml:"cbc:TaxAmount"`
	TaxSubtotal []TaxSubtotal `xml:"cac:TaxSubtotal"`
}

// TaxSubtotal represents more detail about the tax
type TaxSubtotal struct {
	TaxableAmount Amount      `xml:"cbc:TaxableAmount"`
	TaxAmount     Amount      `xml:"cbc:TaxAmount"`
	TaxCategory   TaxCategory `xml:"cac:TaxCategory"`
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
	Value      float64 `xml:",chardata"`
	CurrencyID string  `xml:"currencyID,attr"`
}

// Quantity with the unitcode.
// Possible values for the unitcode:
// https://docs.peppol.eu/poacc/billing/3.0/codelist/UNECERec20/
type Quantity struct {
	Value    float64 `xml:",chardata"`
	UnitCode string  `xml:"unitCode,attr"`
}

// InvoiceLine represents individual line items
type InvoiceLine struct {
	ID                  string   `xml:"cbc:ID"`
	InvoicedQuantity    Quantity `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount Amount   `xml:"cbc:LineExtensionAmount"`
	TaxTotal            TaxTotal `xml:"cac:TaxTotal"`
	Item                Item     `xml:"cac:Item"`
	Price               Price    `xml:"cac:Price"`
}

// Item represents an item being invoiced
type Item struct {
	Description           string      `xml:"cbc:Description"`
	Name                  string      `xml:"cbc:Name"`
	ClassifiedTaxCategory TaxCategory `xml:"cac:ClassifiedTaxCategory"`
}

// TaxCategory for tax information
type TaxCategory struct {
	ID        string    `xml:"cbc:ID"`
	Name      string    `xml:"cbc:Name"`
	Percent   float64   `xml:"cbc:Percent"`
	TaxScheme TaxScheme `xml:"cac:TaxScheme"`
}

// TaxScheme = VAT
type TaxScheme struct {
	ID string `xml:"cbc:ID"`
}

// Price represents pricing details for an item
type Price struct {
	PriceAmount Amount `xml:"cbc:PriceAmount"`
}
