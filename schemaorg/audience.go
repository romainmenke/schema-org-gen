package schemaorg

// Audience see : https://schema.org/Audience
type Audience struct {

Intangible// AudienceType see : /audienceType
// The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
AudienceType string `json:"audienceType"`

// GeographicArea see : /geographicArea
// The geographic area associated with the audience.
GeographicArea string `json:"geographicArea"`

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

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created.  Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience interface{} `json:"audience"` // types : CreativeWork Event LodgingBusiness PlayAction Product Service

// Grantee see : /grantee
// The person, organization, contact point, or audience that has been granted this permission. 
Grantee string `json:"grantee"`

// PermitAudience see : /permitAudience
// The target audience for this permit. 
PermitAudience string `json:"permitAudience"`

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action. 
Recipient interface{} `json:"recipient"` // types : AuthorizeAction CommunicateAction DonateAction GiveAction Message PayAction ReturnAction SendAction TipAction

// Sender see : /sender
// A sub property of participant. The participant who is at the sending end of the action. 
Sender interface{} `json:"sender"` // types : Message ReceiveAction

// ToRecipient see : /toRecipient
// A sub property of recipient. The recipient who was directly sent the message. 
ToRecipient string `json:"toRecipient"`

// TouristType see : /touristType
// Attraction suitable for type(s) of tourist. eg. Children, visitors from a particular country, etc. 
TouristType string `json:"touristType"`

}

