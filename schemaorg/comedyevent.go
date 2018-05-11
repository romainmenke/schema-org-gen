package schemaorg

// ComedyEvent see : https://schema.org/ComedyEvent
type ComedyEvent struct {

Event// About see : /about
// The subject matter of the content.
About string `json:"about"`

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor string `json:"actor"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// Attendee see : /attendee
// A person or organization attending the event. Supersedes attendees (see: https://schema.org/attendees).
Attendee interface{} `json:"attendee"` // types : Organization Person

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience string `json:"audience"`

// Composer see : /composer
// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
Composer interface{} `json:"composer"` // types : Organization Person

// Contributor see : /contributor
// A secondary contributor to the CreativeWork or Event.
Contributor interface{} `json:"contributor"` // types : Organization Person

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director string `json:"director"`

// DoorTime see : /doorTime
// The time admission will commence.
DoorTime string `json:"doorTime"`

// Duration see : /duration
// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
Duration string `json:"duration"`

// EndDate see : /endDate
// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
EndDate interface{} `json:"endDate"` // types : Date DateTime

// EventStatus see : /eventStatus
// An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
EventStatus string `json:"eventStatus"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// InLanguage see : /inLanguage
// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
InLanguage interface{} `json:"inLanguage"` // types : Language Text

// IsAccessibleForFree see : /isAccessibleForFree
// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
IsAccessibleForFree string `json:"isAccessibleForFree"`

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place.
Location interface{} `json:"location"` // types : Place PostalAddress Text

// MaximumAttendeeCapacity see : /maximumAttendeeCapacity
// The total number of individuals that may attend an event or venue.
MaximumAttendeeCapacity string `json:"maximumAttendeeCapacity"`

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers string `json:"offers"`

// Organizer see : /organizer
// An organizer of an Event.
Organizer interface{} `json:"organizer"` // types : Organization Person

// Performer see : /performer
// A performer at the event—for example, a presenter, musician, musical group or actor. Supersedes performers (see: https://schema.org/performers).
Performer interface{} `json:"performer"` // types : Organization Person

// PreviousStartDate see : /previousStartDate
// Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
PreviousStartDate string `json:"previousStartDate"`

// RecordedIn see : /recordedIn
// The CreativeWork that captured all or part of this Event. Inverse property: recordedAt (see: https://schema.org/recordedAt).
RecordedIn string `json:"recordedIn"`

// RemainingAttendeeCapacity see : /remainingAttendeeCapacity
// The number of attendee places for an event that remain unallocated.
RemainingAttendeeCapacity string `json:"remainingAttendeeCapacity"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// StartDate see : /startDate
// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
StartDate interface{} `json:"startDate"` // types : Date DateTime

// SubEvent see : /subEvent
// An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference. Supersedes subEvents (see: https://schema.org/subEvents). Inverse property: superEvent (see: https://schema.org/superEvent).
SubEvent string `json:"subEvent"`

// SuperEvent see : /superEvent
// An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent. Inverse property: subEvent (see: https://schema.org/subEvent).
SuperEvent string `json:"superEvent"`

// Translator see : /translator
// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
Translator interface{} `json:"translator"` // types : Organization Person

// TypicalAgeRange see : /typicalAgeRange
// The typical expected age range, e.g. '7-9', '11-'.
TypicalAgeRange string `json:"typicalAgeRange"`

// WorkFeatured see : /workFeatured
// A work featured in some event, e.g. exhibited in an ExhibitionEvent.
//        Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
WorkFeatured string `json:"workFeatured"`

// WorkPerformed see : /workPerformed
// A work performed in some event, for example a play performed in a TheaterEvent.
WorkPerformed string `json:"workPerformed"`

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

