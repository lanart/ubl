package ubl

import "encoding/xml"

type xmlInvoice struct {
	XMLName                     xml.Name               `xml:"Invoice"`
	Xmlns                       string                 `xml:"xmlns,attr"`
	Cac                         string                 `xml:"xmlns:cac,attr"`
	Cbc                         string                 `xml:"xmlns:cbc,attr"`
	CustomizationID             string                 `xml:"cbc:CustomizationID"`
	ProfileID                   string                 `xml:"cbc:ProfileID"`
	ID                          string                 `xml:"cbc:ID"`
	IssueDate                   string                 `xml:"cbc:IssueDate"`
	DueDate                     string                 `xml:"cbc:DueDate"`
	InvoiceTypeCode             string                 `xml:"cbc:InvoiceTypeCode"`
	DocumentCurrency            string                 `xml:"cbc:DocumentCurrencyCode"`
	BuyerReference              string                 `xml:"cbc:BuyerReference,omitempty"`
	OrderReference              string                 `xml:"cac:OrderReference>cbc:ID"`
	AdditionalDocumentReference []xmlDocumentReference `xml:"cac:AdditionalDocumentReference"`
	SupplierParty               xmlSupplierParty       `xml:"cac:AccountingSupplierParty"`
	CustomerParty               xmlCustomerParty       `xml:"cac:AccountingCustomerParty"`
	PaymentMeans                xmlPaymentMeans        `xml:"cac:PaymentMeans"`
	PaymentTerms                xmlPaymentTerms        `xml:"cac:PaymentTerms"`
	TaxTotal                    xmlTaxTotal            `xml:"cac:TaxTotal"`
	LegalMonetaryTotal          xmlMonetaryTotal       `xml:"cac:LegalMonetaryTotal"`
	InvoiceLines                []xmlInvoiceLine       `xml:"cac:InvoiceLine"`
}

type xmlDocumentReference struct {
	ID                  string          `xml:"cbc:ID"`
	DocumentDescription string          `xml:"cbc:DocumentDescription"`
	Attachment          []xmlAttachment `xml:"cac:Attachment"`
}

type xmlAttachment struct {
	EmbeddedDocumentBinaryObject xmlEmbeddedDocumentBinaryObject `xml:"cbc:EmbeddedDocumentBinaryObject"`
}

type xmlEmbeddedDocumentBinaryObject struct {
	Value    string `xml:",chardata"`
	MimeCode string `xml:"mimeCode,attr"`
	Filename string `xml:"filename,attr"`
}

type xmlSupplierParty struct {
	Party xmlParty `xml:"cac:Party"`
}

type xmlCustomerParty struct {
	Party xmlParty `xml:"cac:Party"`
}

type xmlEndpointID struct {
	Value    string `xml:",chardata"`
	SchemeID string `xml:"schemeID,attr"`
}

type xmlParty struct {
	EndpointID       xmlEndpointID     `xml:"cbc:EndpointID"`
	PartyName        string            `xml:"cac:PartyName>cbc:Name"`
	PostalAddress    xmlPostalAddress  `xml:"cac:PostalAddress"`
	PartyTaxScheme   xmlPartyTaxScheme `xml:"cac:PartyTaxScheme"`
	RegistrationName string            `xml:"cac:PartyLegalEntity>cbc:RegistrationName"`
}

type xmlPostalAddress struct {
	StreetName string     `xml:"cbc:StreetName,omitempty"`
	CityName   string     `xml:"cbc:CityName,omitempty"`
	PostalZone string     `xml:"cbc:PostalZone,omitempty"`
	Country    xmlCountry `xml:"cac:Country,omitempty"`
}

type xmlPartyTaxScheme struct {
	CompanyID string       `xml:"cbc:CompanyID"`
	TaxScheme xmlTaxScheme `xml:"cac:TaxScheme"`
}

type xmlCountry struct {
	IdentificationCode string `xml:"cbc:IdentificationCode,omitempty"`
}

type xmlPaymentMeans struct {
	PaymentMeansCode      string              `xml:"cbc:PaymentMeansCode"`
	PayeeFinancialAccount xmlFinancialAccount `xml:"cac:PayeeFinancialAccount"`
}

type xmlFinancialAccount struct {
	ID                         string                        `xml:"cbc:ID"`
	FinancialInstitutionBranch xmlFinancialInstitutionBranch `xml:"cac:FinancialInstitutionBranch"`
}

type xmlFinancialInstitutionBranch struct {
	ID string `xml:"cbc:ID"`
}

type xmlPaymentTerms struct {
	Note string `xml:"cbc:Note"`
}

type xmlTaxTotal struct {
	TaxAmount   xmlAmount        `xml:"cbc:TaxAmount"`
	TaxSubtotal []xmlTaxSubtotal `xml:"cac:TaxSubtotal"`
}

type xmlTaxSubtotal struct {
	TaxableAmount xmlAmount      `xml:"cbc:TaxableAmount"`
	TaxAmount     xmlAmount      `xml:"cbc:TaxAmount"`
	TaxCategory   xmlTaxCategory `xml:"cac:TaxCategory"`
}

type xmlMonetaryTotal struct {
	LineExtensionAmount xmlAmount `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount  xmlAmount `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount  xmlAmount `xml:"cbc:TaxInclusiveAmount"`
	PayableAmount       xmlAmount `xml:"cbc:PayableAmount"`
}

type xmlAmount struct {
	Value      float64 `xml:",chardata"`
	CurrencyID string  `xml:"currencyID,attr"`
}

// Possible values for the unitcode:
// https://docs.peppol.eu/poacc/billing/3.0/codelist/UNECERec20/
type xmlQuantity struct {
	Value    float64 `xml:",chardata"`
	UnitCode string  `xml:"unitCode,attr"`
}

type xmlInvoiceLine struct {
	ID                  string      `xml:"cbc:ID"`
	InvoicedQuantity    xmlQuantity `xml:"cbc:InvoicedQuantity"`
	LineExtensionAmount xmlAmount   `xml:"cbc:LineExtensionAmount"`
	TaxTotal            xmlTaxTotal `xml:"cac:TaxTotal"`
	Item                xmlItem     `xml:"cac:Item"`
	Price               xmlPrice    `xml:"cac:Price"`
}

type xmlItem struct {
	Description           string         `xml:"cbc:Description"`
	Name                  string         `xml:"cbc:Name"`
	ClassifiedTaxCategory xmlTaxCategory `xml:"cac:ClassifiedTaxCategory"`
}

type xmlTaxCategory struct {
	ID        string       `xml:"cbc:ID"`
	Name      string       `xml:"cbc:Name"`
	Percent   float64      `xml:"cbc:Percent"`
	TaxScheme xmlTaxScheme `xml:"cac:TaxScheme"`
}

type xmlTaxScheme struct {
	ID string `xml:"cbc:ID"`
}

type xmlPrice struct {
	PriceAmount xmlAmount `xml:"cbc:PriceAmount"`
}
