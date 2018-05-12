package schemaorggo

import "encoding/json"

// ChildrensEvent see : https://schema.org/ChildrensEvent
type ChildrensEvent struct {
	Event

	typeContext

	// About see : https://schema.org/about
	// The subject matter of the content.
	// types : Thing
	About *Thing `json:"about,omitempty"`

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty"`

	// Attendee see : https://schema.org/attendee
	// A person or organization attending the event. Supersedes attendees (see: https://schema.org/attendees).
	// types : Organization Person
	Attendee interface{} `json:"attendee,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience *Audience `json:"audience,omitempty"`

	// Composer see : https://schema.org/composer
	// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	// types : Organization Person
	Composer interface{} `json:"composer,omitempty"`

	// Contributor see : https://schema.org/contributor
	// A secondary contributor to the CreativeWork or Event.
	// types : Organization Person
	Contributor interface{} `json:"contributor,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// DoorTime see : https://schema.org/doorTime
	// The time admission will commence.
	// types : DateTime
	DoorTime DateTime `json:"doorTime,omitempty"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	Duration *Duration `json:"duration,omitempty"`

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate interface{} `json:"endDate,omitempty"`

	// EventStatus see : https://schema.org/eventStatus
	// An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
	// types : EventStatusType
	EventStatus *EventStatusType `json:"eventStatus,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder interface{} `json:"funder,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage interface{} `json:"inLanguage,omitempty"`

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	// types : Boolean
	IsAccessibleForFree bool `json:"isAccessibleForFree,omitempty"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
	Location interface{} `json:"location,omitempty"`

	// MaximumAttendeeCapacity see : https://schema.org/maximumAttendeeCapacity
	// The total number of individuals that may attend an event or venue.
	// types : Integer
	MaximumAttendeeCapacity float64 `json:"maximumAttendeeCapacity,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers *Offer `json:"offers,omitempty"`

	// Organizer see : https://schema.org/organizer
	// An organizer of an Event.
	// types : Organization Person
	Organizer interface{} `json:"organizer,omitempty"`

	// Performer see : https://schema.org/performer
	// A performer at the event—for example, a presenter, musician, musical group or actor. Supersedes performers (see: https://schema.org/performers).
	// types : Organization Person
	Performer interface{} `json:"performer,omitempty"`

	// PreviousStartDate see : https://schema.org/previousStartDate
	// Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
	// types : Date
	PreviousStartDate Date `json:"previousStartDate,omitempty"`

	// RecordedIn see : https://schema.org/recordedIn
	// The CreativeWork that captured all or part of this Event. Inverse property: recordedAt (see: https://schema.org/recordedAt).
	// types : CreativeWork
	RecordedIn *CreativeWork `json:"recordedIn,omitempty"`

	// RemainingAttendeeCapacity see : https://schema.org/remainingAttendeeCapacity
	// The number of attendee places for an event that remain unallocated.
	// types : Integer
	RemainingAttendeeCapacity float64 `json:"remainingAttendeeCapacity,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review *Review `json:"review,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor interface{} `json:"sponsor,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate interface{} `json:"startDate,omitempty"`

	// SubEvent see : https://schema.org/subEvent
	// An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference. Supersedes subEvents (see: https://schema.org/subEvents). Inverse property: superEvent (see: https://schema.org/superEvent).
	// types : Event
	SubEvent *Event `json:"subEvent,omitempty"`

	// SuperEvent see : https://schema.org/superEvent
	// An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent. Inverse property: subEvent (see: https://schema.org/subEvent).
	// types : Event
	SuperEvent *Event `json:"superEvent,omitempty"`

	// Translator see : https://schema.org/translator
	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	// types : Organization Person
	Translator interface{} `json:"translator,omitempty"`

	// TypicalAgeRange see : https://schema.org/typicalAgeRange
	// The typical expected age range, e.g. &#39;7-9&#39;, &#39;11-&#39;.
	// types : Text
	TypicalAgeRange string `json:"typicalAgeRange,omitempty"`

	// WorkFeatured see : https://schema.org/workFeatured
	// A work featured in some event, e.g. exhibited in an ExhibitionEvent.
	//        Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
	// types : CreativeWork
	WorkFeatured *CreativeWork `json:"workFeatured,omitempty"`

	// WorkPerformed see : https://schema.org/workPerformed
	// A work performed in some event, for example a play performed in a TheaterEvent.
	// types : CreativeWork
	WorkPerformed *CreativeWork `json:"workPerformed,omitempty"`
}

func (v ChildrensEvent) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ChildrensEvent"

	return json.Marshal(v)
}

func (v *ChildrensEvent) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ChildrensEvent"

	return json.Marshal(*v)
}
