package schemaorg

// PostalAddress see : https://schema.org/PostalAddress
type PostalAddress struct {

ContactPoint// AddressCountry see : /addressCountry
// The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
AddressCountry interface{} `json:"addressCountry"` // types : Country Text

// AddressLocality see : /addressLocality
// The locality. For example, Mountain View.
AddressLocality string `json:"addressLocality"`

// AddressRegion see : /addressRegion
// The region. For example, CA.
AddressRegion string `json:"addressRegion"`

// PostOfficeBoxNumber see : /postOfficeBoxNumber
// The post office box number for PO box addresses.
PostOfficeBoxNumber string `json:"postOfficeBoxNumber"`

// PostalCode see : /postalCode
// The postal code. For example, 94043.
PostalCode string `json:"postalCode"`

// StreetAddress see : /streetAddress
// The street address. For example, 1600 Amphitheatre Pkwy.
StreetAddress string `json:"streetAddress"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// AvailableLanguage see : /availableLanguage
// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
AvailableLanguage interface{} `json:"availableLanguage"` // types : Language Text

// ContactOption see : /contactOption
// An option available on this contact point (e.g. a toll-free number or support for hearing-impaired callers).
ContactOption string `json:"contactOption"`

// ContactType see : /contactType
// A person or organization can have different contact points, for different purposes. For example, a sales contact point, a PR contact point and so on. This property is used to specify the kind of contact point.
ContactType string `json:"contactType"`

// Email see : /email
// Email address.
Email string `json:"email"`

// FaxNumber see : /faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// HoursAvailable see : /hoursAvailable
// The hours during which this service or contact is available.
HoursAvailable string `json:"hoursAvailable"`

// ProductSupported see : /productSupported
// The product or service this support contact point is related to (such as product support for a particular product line). This can be a specific product or product line (e.g. "iPhone") or a general category of products or services (e.g. "smartphones").
ProductSupported interface{} `json:"productSupported"` // types : Product Text

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

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

// Address see : /address
// Physical address of the item. 
Address interface{} `json:"address"` // types : GeoCoordinates GeoShape Organization Person Place

// BillingAddress see : /billingAddress
// The billing address for the order. 
BillingAddress string `json:"billingAddress"`

// DeliveryAddress see : /deliveryAddress
// Destination address. 
DeliveryAddress string `json:"deliveryAddress"`

// GameLocation see : /gameLocation
// Real or fictional location of the game (or part of game). 
GameLocation interface{} `json:"gameLocation"` // types : Game VideoGameSeries

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place. 
Location interface{} `json:"location"` // types : Action Event Organization

// OriginAddress see : /originAddress
// Shipper's address. 
OriginAddress string `json:"originAddress"`

// ServicePostalAddress see : /servicePostalAddress
// The address for accessing the service by mail. 
ServicePostalAddress string `json:"servicePostalAddress"`

}

