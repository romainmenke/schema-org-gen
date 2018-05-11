package schemaorg

// ContactPoint see : https://schema.org/ContactPoint
type ContactPoint struct {

StructuredValue// AreaServed see : /areaServed
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

// BccRecipient see : /bccRecipient
// A sub property of recipient. The recipient blind copied on a message. 
BccRecipient string `json:"bccRecipient"`

// CcRecipient see : /ccRecipient
// A sub property of recipient. The recipient copied on a message. 
CcRecipient string `json:"ccRecipient"`

// ContactPoint see : /contactPoint
// A contact point for a person or organization.  Supersedes contactPoints (see: https://schema.org/contactPoints).
ContactPoint interface{} `json:"contactPoint"` // types : HealthInsurancePlan Organization Person

// Grantee see : /grantee
// The person, organization, contact point, or audience that has been granted this permission. 
Grantee string `json:"grantee"`

// HomeLocation see : /homeLocation
// A contact location for a person's residence. 
HomeLocation string `json:"homeLocation"`

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action. 
Recipient interface{} `json:"recipient"` // types : AuthorizeAction CommunicateAction DonateAction GiveAction Message PayAction ReturnAction SendAction TipAction

// ServicePhone see : /servicePhone
// The phone number to use to access the service. 
ServicePhone string `json:"servicePhone"`

// ServiceSmsNumber see : /serviceSmsNumber
// The number to access the service by text message. 
ServiceSmsNumber string `json:"serviceSmsNumber"`

// ToRecipient see : /toRecipient
// A sub property of recipient. The recipient who was directly sent the message. 
ToRecipient string `json:"toRecipient"`

// WorkLocation see : /workLocation
// A contact location for a person's place of work. 
WorkLocation string `json:"workLocation"`

}

