package schemaorg

import "encoding/json"

// MusicEvent see : https://schema.org/MusicEvent
type MusicEvent struct {

typeContext

Event

// About see : /about
// The subject matter of the content.
About *Thing `json:"about"`

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating *AggregateRating `json:"aggregateRating"`

// Attendee see : /attendee
// A person or organization attending the event. Supersedes attendees (see: https://schema.org/attendees).
Attendee interface{} `json:"attendee"` // types : Organization Person

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience *Audience `json:"audience"`

// Composer see : /composer
// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
Composer interface{} `json:"composer"` // types : Organization Person

// Contributor see : /contributor
// A secondary contributor to the CreativeWork or Event.
Contributor interface{} `json:"contributor"` // types : Organization Person

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// DoorTime see : /doorTime
// The time admission will commence.
DoorTime interface{} `json:"doorTime"`

// Duration see : /duration
// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
Duration *Duration `json:"duration"`

// EndDate see : /endDate
// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
EndDate interface{} `json:"endDate"` // types : Date DateTime

// EventStatus see : /eventStatus
// An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
EventStatus *EventStatusType `json:"eventStatus"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// InLanguage see : /inLanguage
// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
InLanguage interface{} `json:"inLanguage"` // types : Language Text

// IsAccessibleForFree see : /isAccessibleForFree
// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
IsAccessibleForFree bool `json:"isAccessibleForFree"`

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place.
Location interface{} `json:"location"` // types : Place PostalAddress Text

// MaximumAttendeeCapacity see : /maximumAttendeeCapacity
// The total number of individuals that may attend an event or venue.
MaximumAttendeeCapacity int `json:"maximumAttendeeCapacity"`

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers *Offer `json:"offers"`

// Organizer see : /organizer
// An organizer of an Event.
Organizer interface{} `json:"organizer"` // types : Organization Person

// Performer see : /performer
// A performer at the event—for example, a presenter, musician, musical group or actor. Supersedes performers (see: https://schema.org/performers).
Performer interface{} `json:"performer"` // types : Organization Person

// PreviousStartDate see : /previousStartDate
// Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
PreviousStartDate interface{} `json:"previousStartDate"`

// RecordedIn see : /recordedIn
// The CreativeWork that captured all or part of this Event. Inverse property: recordedAt (see: https://schema.org/recordedAt).
RecordedIn *CreativeWork `json:"recordedIn"`

// RemainingAttendeeCapacity see : /remainingAttendeeCapacity
// The number of attendee places for an event that remain unallocated.
RemainingAttendeeCapacity int `json:"remainingAttendeeCapacity"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review *Review `json:"review"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// StartDate see : /startDate
// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
StartDate interface{} `json:"startDate"` // types : Date DateTime

// SubEvent see : /subEvent
// An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference. Supersedes subEvents (see: https://schema.org/subEvents). Inverse property: superEvent (see: https://schema.org/superEvent).
SubEvent *Event `json:"subEvent"`

// SuperEvent see : /superEvent
// An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent. Inverse property: subEvent (see: https://schema.org/subEvent).
SuperEvent *Event `json:"superEvent"`

// Translator see : /translator
// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
Translator interface{} `json:"translator"` // types : Organization Person

// TypicalAgeRange see : /typicalAgeRange
// The typical expected age range, e.g. '7-9', '11-'.
TypicalAgeRange string `json:"typicalAgeRange"`

// WorkFeatured see : /workFeatured
// A work featured in some event, e.g. exhibited in an ExhibitionEvent.
//        Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
WorkFeatured *CreativeWork `json:"workFeatured"`

// WorkPerformed see : /workPerformed
// A work performed in some event, for example a play performed in a TheaterEvent.
WorkPerformed *CreativeWork `json:"workPerformed"`

}

func (v *MusicEvent) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MusicEvent"

	return json.Marshal(v)
}
