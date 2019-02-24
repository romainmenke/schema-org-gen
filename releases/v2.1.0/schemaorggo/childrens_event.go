package schemaorggo

import "encoding/json"

// ChildrensEvent see : https://schema.org/ChildrensEvent
type ChildrensEvent struct {

	// With properties from Event see : https://schema.org/Event
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Attendee see : https://schema.org/attendee
	// A person or organization attending the event.
	// types : Organization Person
	Attendee []interface{} `json:"attendee,omitempty"`

	// Attendees see : https://schema.org/attendees
	// A person attending the event.
	// types : Organization Person
	Attendees []interface{} `json:"attendees,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DoorTime see : https://schema.org/doorTime
	// The time admission will commence.
	// types : DateTime
	DoorTime []DateTime `json:"doorTime,omitempty"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 date format&lt;/a&gt;.
	// types : Duration
	Duration []*Duration `json:"duration,omitempty"`

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 date format&lt;/a&gt;).
	// types : Date
	EndDate []Date `json:"endDate,omitempty"`

	// EventStatus see : https://schema.org/eventStatus
	// An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
	// types : EventStatusType
	EventStatus []*EventStatusType `json:"eventStatus,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the &lt;a href=&#39;http://tools.ietf.org/html/bcp47&#39;&gt;IETF BCP 47 standard&lt;/a&gt;.
	// types : Text Language
	InLanguage []interface{} `json:"inLanguage,omitempty"`

	// Location see : https://schema.org/location
	// The location of the event, organization or action.
	// types : Place PostalAddress
	Location []interface{} `json:"location,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// Organizer see : https://schema.org/organizer
	// An organizer of an Event.
	// types : Person Organization
	Organizer []interface{} `json:"organizer,omitempty"`

	// Performer see : https://schema.org/performer
	// A performer at the event&amp;#x2014;for example, a presenter, musician, musical group or actor.
	// types : Organization Person
	Performer []interface{} `json:"performer,omitempty"`

	// Performers see : https://schema.org/performers
	// The main performer or performers of the event&amp;#x2014;for example, a presenter, musician, or actor.
	// types : Organization Person
	Performers []interface{} `json:"performers,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PreviousStartDate see : https://schema.org/previousStartDate
	// Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
	// types : Date
	PreviousStartDate []Date `json:"previousStartDate,omitempty"`

	// RecordedIn see : https://schema.org/recordedIn
	// The CreativeWork that captured all or part of this Event.
	// types : CreativeWork
	RecordedIn []*CreativeWork `json:"recordedIn,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 date format&lt;/a&gt;).
	// types : Date
	StartDate []Date `json:"startDate,omitempty"`

	// SubEvent see : https://schema.org/subEvent
	// An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference.
	// types : Event
	SubEvent []*Event `json:"subEvent,omitempty"`

	// SubEvents see : https://schema.org/subEvents
	// Events that are a part of this event. For example, a conference event includes many presentations, each subEvents of the conference.
	// types : Event
	SubEvents []*Event `json:"subEvents,omitempty"`

	// SuperEvent see : https://schema.org/superEvent
	// An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent.
	// types : Event
	SuperEvent []*Event `json:"superEvent,omitempty"`

	// TypicalAgeRange see : https://schema.org/typicalAgeRange
	// The typical expected age range, e.g. &#39;7-9&#39;, &#39;11-&#39;.
	// types : Text
	TypicalAgeRange []string `json:"typicalAgeRange,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// WorkPerformed see : https://schema.org/workPerformed
	// A work performed in some event, for example a play performed in a TheaterEvent.
	// types : CreativeWork
	WorkPerformed []*CreativeWork `json:"workPerformed,omitempty"`
}

func (v ChildrensEvent) MarshalJSON() ([]byte, error) {
	type Alias ChildrensEvent

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"ChildrensEvent\","), b[1:]...), nil
}
