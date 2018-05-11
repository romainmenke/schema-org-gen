package schemaorg

// ServiceChannel see : https://schema.org/ServiceChannel
type ServiceChannel struct {

Intangible// AvailableLanguage see : /availableLanguage
// A language someone may use with or at the item, service or place. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also inLanguage (see: https://schema.org/inLanguage)
AvailableLanguage interface{} `json:"availableLanguage"` // types : Language Text

// ProcessingTime see : /processingTime
// Estimated processing time for the service using this channel.
ProcessingTime string `json:"processingTime"`

// ProvidesService see : /providesService
// The service provided by this channel.
ProvidesService string `json:"providesService"`

// ServiceLocation see : /serviceLocation
// The location (e.g. civic structure, local business, etc.) where a person can go to access the service.
ServiceLocation string `json:"serviceLocation"`

// ServicePhone see : /servicePhone
// The phone number to use to access the service.
ServicePhone string `json:"servicePhone"`

// ServicePostalAddress see : /servicePostalAddress
// The address for accessing the service by mail.
ServicePostalAddress string `json:"servicePostalAddress"`

// ServiceSmsNumber see : /serviceSmsNumber
// The number to access the service by text message.
ServiceSmsNumber string `json:"serviceSmsNumber"`

// ServiceUrl see : /serviceUrl
// The website to access the service.
ServiceUrl string `json:"serviceUrl"`

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

// AvailableChannel see : /availableChannel
// A means of accessing the service (e.g. a phone bank, a web site, a location, etc.). 
AvailableChannel string `json:"availableChannel"`

}

