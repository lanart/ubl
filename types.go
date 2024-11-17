package ubl

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

// Quantity
type Quantity struct {
	Value    float32 `xml:",chardata"`
	UnitCode string  `xml:"unitCode,attr"`
}

// InvoiceLine represents individual line items
type InvoiceLine struct {
	ID                  string   `xml:"cbc:ID"`
	InvoicedQuantity    Quantity `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount Amount   `xml:"cbc:LineExtensionAmount"`
	Item                Item     `xml:"cac:Item"`
	Price               Price    `xml:"cac:Price"`
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
