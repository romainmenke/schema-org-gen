package schemaorg

// InviteAction see : https://schema.org/InviteAction
type InviteAction struct {

CommunicateAction// Event see : /event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event string `json:"event"`

// About see : /about
// The subject matter of the content.
About string `json:"about"`

// InLanguage see : /inLanguage
// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
InLanguage interface{} `json:"inLanguage"` // types : Language Text

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action.
Recipient interface{} `json:"recipient"` // types : Audience ContactPoint Organization Person

// ActionStatus see : /actionStatus
// Indicates the current disposition of the Action.
ActionStatus string `json:"actionStatus"`

// Agent see : /agent
// The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book.
Agent interface{} `json:"agent"` // types : Organization Person

// EndTime see : /endTime
// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
// 
// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
EndTime string `json:"endTime"`

// Error see : /error
// For failed actions, more information on the cause of the failure.
Error string `json:"error"`

// Instrument see : /instrument
// The object that helped the agent perform the action. e.g. John wrote a book with a pen.
Instrument string `json:"instrument"`

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place.
Location interface{} `json:"location"` // types : Place PostalAddress Text

// Object see : /object
// The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn't). e.g. John read a book.
Object string `json:"object"`

// Participant see : /participant
// Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve.
Participant interface{} `json:"participant"` // types : Organization Person

// Result see : /result
// The result produced in the action. e.g. John wrote a book.
Result string `json:"result"`

// StartTime see : /startTime
// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
// 
// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
StartTime string `json:"startTime"`

// Target see : /target
// Indicates a target EntryPoint for an Action.
Target string `json:"target"`

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

