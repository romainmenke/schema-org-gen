package schemaorg

// CreditCard see : https://schema.org/CreditCard
type CreditCard struct {

LoanOrCredit// MonthlyMinimumRepaymentAmount see : http://pending.schema.org/monthlyMinimumRepaymentAmount
// The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
MonthlyMinimumRepaymentAmount interface{} `json:"monthlyMinimumRepaymentAmount"` // types : MonetaryAmount Number

// Amount see : /amount
// The amount of money.
Amount interface{} `json:"amount"` // types : MonetaryAmount Number

// Currency see : /currency
// The currency in which the monetary amount is expressed (in 3-letter ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) format).
Currency string `json:"currency"`

// GracePeriod see : http://pending.schema.org/gracePeriod
// The period of time after any due date that the borrower has to fulfil its obligations before a default (failure to pay) is deemed to have occurred.
GracePeriod string `json:"gracePeriod"`

// LoanRepaymentForm see : http://pending.schema.org/loanRepaymentForm
// A form of paying back money previously borrowed from a lender. Repayment usually takes the form of periodic payments that normally include part principal plus interest in each payment.
LoanRepaymentForm string `json:"loanRepaymentForm"`

// LoanTerm see : /loanTerm
// The duration of the loan or credit agreement.
LoanTerm string `json:"loanTerm"`

// LoanType see : http://pending.schema.org/loanType
// The type of a loan or credit.
LoanType interface{} `json:"loanType"` // types : Text URL

// RecourseLoan see : http://pending.schema.org/recourseLoan
// The only way you get the money back in the event of default is the security. Recourse is where you still have the opportunity to go back to the borrower for the rest of the money.
RecourseLoan string `json:"recourseLoan"`

// RenegotiableLoan see : http://pending.schema.org/renegotiableLoan
// Whether the terms for payment of interest can be renegotiated during the life of the loan.
RenegotiableLoan string `json:"renegotiableLoan"`

// RequiredCollateral see : /requiredCollateral
// Assets required to secure loan or credit repayments. It may take form of third party pledge, goods, financial instruments (cash, securities, etc.)
RequiredCollateral interface{} `json:"requiredCollateral"` // types : Text Thing

// AnnualPercentageRate see : /annualPercentageRate
// The annual rate that is charged for borrowing (or made by investing), expressed as a single percentage number that represents the actual yearly cost of funds over the term of a loan. This includes any fees or additional costs associated with the transaction.
AnnualPercentageRate interface{} `json:"annualPercentageRate"` // types : Number QuantitativeValue

// FeesAndCommissionsSpecification see : /feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

// InterestRate see : /interestRate
// The interest rate, charged or paid, applicable to the financial product. Note: This is different from the calculated annualPercentageRate.
InterestRate interface{} `json:"interestRate"` // types : Number QuantitativeValue

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience string `json:"audience"`

// AvailableChannel see : /availableChannel
// A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
AvailableChannel string `json:"availableChannel"`

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Organization Person

// Category see : /category
// A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
Category interface{} `json:"category"` // types : PhysicalActivityCategory Text Thing

// HasOfferCatalog see : /hasOfferCatalog
// Indicates an OfferCatalog listing for this Organization, Person, or Service.
HasOfferCatalog string `json:"hasOfferCatalog"`

// HoursAvailable see : /hoursAvailable
// The hours during which this service or contact is available.
HoursAvailable string `json:"hoursAvailable"`

// IsRelatedTo see : /isRelatedTo
// A pointer to another, somehow related product (or multiple products).
IsRelatedTo interface{} `json:"isRelatedTo"` // types : Product Service

// IsSimilarTo see : /isSimilarTo
// A pointer to another, functionally similar product (or multiple products).
IsSimilarTo interface{} `json:"isSimilarTo"` // types : Product Service

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers string `json:"offers"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// ProviderMobility see : /providerMobility
// Indicates the mobility of a provided service (e.g. 'static', 'dynamic').
ProviderMobility string `json:"providerMobility"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// ServiceOutput see : /serviceOutput
// The tangible thing generated by the service, e.g. a passport, permit, etc. Supersedes produces (see: https://schema.org/produces).
ServiceOutput string `json:"serviceOutput"`

// ServiceType see : /serviceType
// The type of service being offered, e.g. veterans' benefits, emergency relief, etc.
ServiceType string `json:"serviceType"`

// TermsOfService see : http://pending.schema.org/termsOfService
// Human-readable terms of service documentation.
TermsOfService interface{} `json:"termsOfService"` // types : Text URL

// CashBack see : http://pending.schema.org/cashBack
// A cardholder benefit that pays the cardholder a small percentage of their net expenditures.
CashBack interface{} `json:"cashBack"` // types : Boolean Number

// ContactlessPayment see : http://pending.schema.org/contactlessPayment
// A secure method for consumers to purchase products or services via debit, credit or smartcards by using RFID or NFC technology.
ContactlessPayment string `json:"contactlessPayment"`

// FloorLimit see : http://pending.schema.org/floorLimit
// A floor limit is the amount of money above which credit card transactions must be authorized.
FloorLimit string `json:"floorLimit"`

// SupersededBy see : http://meta.schema.org/supersededBy
// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

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

}

