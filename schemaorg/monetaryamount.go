package schemaorg

// MonetaryAmount see : https://schema.org/MonetaryAmount
type MonetaryAmount struct {

StructuredValue// Currency see : /currency
// The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
Currency string `json:"currency"`

// MaxValue see : /maxValue
// The upper value of some characteristic or property.
MaxValue string `json:"maxValue"`

// MinValue see : /minValue
// The lower value of some characteristic or property.
MinValue string `json:"minValue"`

// ValidFrom see : /validFrom
// The date when the item becomes valid.
ValidFrom string `json:"validFrom"`

// ValidThrough see : /validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough string `json:"validThrough"`

// Value see : /value
// The value of the quantitative value or property value node.
// 
// 
// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is 'Number'.
// For PropertyValue (see: https://schema.org/PropertyValue), it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.
// 
// 
Value interface{} `json:"value"` // types : Boolean Number StructuredValue Text

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// Amount see : /amount
// The amount of money. 
Amount interface{} `json:"amount"` // types : DatedMoneySpecification InvestmentOrDeposit LoanOrCredit MoneyTransfer

// BaseSalary see : /baseSalary
// The base salary of the job or of an employee in an EmployeeRole. 
BaseSalary interface{} `json:"baseSalary"` // types : EmployeeRole JobPosting

// EstimatedCost see : /estimatedCost
// The estimated cost of the supply or supplies consumed when performing instructions. 
EstimatedCost interface{} `json:"estimatedCost"` // types : HowTo HowToSupply

// MinimumPaymentDue see : /minimumPaymentDue
// The minimum payment required at this time. 
MinimumPaymentDue string `json:"minimumPaymentDue"`

// NetWorth see : /netWorth
// The total financial value of the person as calculated by subtracting assets from liabilities. 
NetWorth string `json:"netWorth"`

// TotalPaymentDue see : /totalPaymentDue
// The total amount due. 
TotalPaymentDue string `json:"totalPaymentDue"`

}

