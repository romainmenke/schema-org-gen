package schemaorggo

import "encoding/json"

// ChildrensEvent see : https://schema.org/ChildrensEvent
type ChildrensEvent struct {
	Event

	typeContext

	// About see : https://schema.org/about
	// The subject matter of the content.
	// types : Thing
	About []*Thing `json:"about,omitempty"`

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// Attendee see : https://schema.org/attendee
	// A person or organization attending the event. Supersedes attendees (see: https://schema.org/attendees).
	// types : Organization Person
	Attendee []interface{} `json:"attendee,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	// types : Audience
	Audience []*Audience `json:"audience,omitempty"`

	// Composer see : https://schema.org/composer
	// The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	// types : Organization Person
	Composer []interface{} `json:"composer,omitempty"`

	// Contributor see : https://schema.org/contributor
	// A secondary contributor to the CreativeWork or Event.
	// types : Organization Person
	Contributor []interface{} `json:"contributor,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// DoorTime see : https://schema.org/doorTime
	// The time admission will commence.
	// types : DateTime
	DoorTime []DateTime `json:"doorTime,omitempty"`

	// Duration see : https://schema.org/duration
	// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	Duration []*Duration `json:"duration,omitempty"`

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	EndDate []interface{} `json:"endDate,omitempty"`

	// EventStatus see : https://schema.org/eventStatus
	// An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
	// types : EventStatusType
	EventStatus []*EventStatusType `json:"eventStatus,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder []interface{} `json:"funder,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage []interface{} `json:"inLanguage,omitempty"`

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	// types : Boolean
	IsAccessibleForFree []bool `json:"isAccessibleForFree,omitempty"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
	Location []interface{} `json:"location,omitempty"`

	// MaximumAttendeeCapacity see : https://schema.org/maximumAttendeeCapacity
	// The total number of individuals that may attend an event or venue.
	// types : Integer
	MaximumAttendeeCapacity []float64 `json:"maximumAttendeeCapacity,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// Organizer see : https://schema.org/organizer
	// An organizer of an Event.
	// types : Organization Person
	Organizer []interface{} `json:"organizer,omitempty"`

	// Performer see : https://schema.org/performer
	// A performer at the event—for example, a presenter, musician, musical group or actor. Supersedes performers (see: https://schema.org/performers).
	// types : Organization Person
	Performer []interface{} `json:"performer,omitempty"`

	// PreviousStartDate see : https://schema.org/previousStartDate
	// Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
	// types : Date
	PreviousStartDate []Date `json:"previousStartDate,omitempty"`

	// RecordedIn see : https://schema.org/recordedIn
	// The CreativeWork that captured all or part of this Event. Inverse property: recordedAt (see: https://schema.org/recordedAt).
	// types : CreativeWork
	RecordedIn []*CreativeWork `json:"recordedIn,omitempty"`

	// RemainingAttendeeCapacity see : https://schema.org/remainingAttendeeCapacity
	// The number of attendee places for an event that remain unallocated.
	// types : Integer
	RemainingAttendeeCapacity []float64 `json:"remainingAttendeeCapacity,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor []interface{} `json:"sponsor,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate []interface{} `json:"startDate,omitempty"`

	// SubEvent see : https://schema.org/subEvent
	// An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference. Supersedes subEvents (see: https://schema.org/subEvents). Inverse property: superEvent (see: https://schema.org/superEvent).
	// types : Event
	SubEvent []*Event `json:"subEvent,omitempty"`

	// SuperEvent see : https://schema.org/superEvent
	// An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent. Inverse property: subEvent (see: https://schema.org/subEvent).
	// types : Event
	SuperEvent []*Event `json:"superEvent,omitempty"`

	// Translator see : https://schema.org/translator
	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	// types : Organization Person
	Translator []interface{} `json:"translator,omitempty"`

	// TypicalAgeRange see : https://schema.org/typicalAgeRange
	// The typical expected age range, e.g. &#39;7-9&#39;, &#39;11-&#39;.
	// types : Text
	TypicalAgeRange []string `json:"typicalAgeRange,omitempty"`

	// WorkFeatured see : https://schema.org/workFeatured
	// A work featured in some event, e.g. exhibited in an ExhibitionEvent.
	//        Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
	// types : CreativeWork
	WorkFeatured []*CreativeWork `json:"workFeatured,omitempty"`

	// WorkPerformed see : https://schema.org/workPerformed
	// A work performed in some event, for example a play performed in a TheaterEvent.
	// types : CreativeWork
	WorkPerformed []*CreativeWork `json:"workPerformed,omitempty"`
}

func (v ChildrensEvent) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Event.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.About
		if len(v.About) == 1 {
			value = v.About[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["about"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Actor
		if len(v.Actor) == 1 {
			value = v.Actor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AggregateRating
		if len(v.AggregateRating) == 1 {
			value = v.AggregateRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["aggregateRating"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Attendee
		if len(v.Attendee) == 1 {
			value = v.Attendee[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["attendee"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Audience
		if len(v.Audience) == 1 {
			value = v.Audience[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audience"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Composer
		if len(v.Composer) == 1 {
			value = v.Composer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["composer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Contributor
		if len(v.Contributor) == 1 {
			value = v.Contributor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contributor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Director
		if len(v.Director) == 1 {
			value = v.Director[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["director"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DoorTime
		if len(v.DoorTime) == 1 {
			value = v.DoorTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["doorTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Duration
		if len(v.Duration) == 1 {
			value = v.Duration[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["duration"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EndDate
		if len(v.EndDate) == 1 {
			value = v.EndDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EventStatus
		if len(v.EventStatus) == 1 {
			value = v.EventStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["eventStatus"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Funder
		if len(v.Funder) == 1 {
			value = v.Funder[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["funder"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InLanguage
		if len(v.InLanguage) == 1 {
			value = v.InLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsAccessibleForFree
		if len(v.IsAccessibleForFree) == 1 {
			value = v.IsAccessibleForFree[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isAccessibleForFree"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Location
		if len(v.Location) == 1 {
			value = v.Location[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["location"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MaximumAttendeeCapacity
		if len(v.MaximumAttendeeCapacity) == 1 {
			value = v.MaximumAttendeeCapacity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["maximumAttendeeCapacity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Offers
		if len(v.Offers) == 1 {
			value = v.Offers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Organizer
		if len(v.Organizer) == 1 {
			value = v.Organizer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["organizer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Performer
		if len(v.Performer) == 1 {
			value = v.Performer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["performer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PreviousStartDate
		if len(v.PreviousStartDate) == 1 {
			value = v.PreviousStartDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["previousStartDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RecordedIn
		if len(v.RecordedIn) == 1 {
			value = v.RecordedIn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recordedIn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RemainingAttendeeCapacity
		if len(v.RemainingAttendeeCapacity) == 1 {
			value = v.RemainingAttendeeCapacity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["remainingAttendeeCapacity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Review
		if len(v.Review) == 1 {
			value = v.Review[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["review"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Sponsor
		if len(v.Sponsor) == 1 {
			value = v.Sponsor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sponsor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StartDate
		if len(v.StartDate) == 1 {
			value = v.StartDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["startDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SubEvent
		if len(v.SubEvent) == 1 {
			value = v.SubEvent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subEvent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuperEvent
		if len(v.SuperEvent) == 1 {
			value = v.SuperEvent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["superEvent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Translator
		if len(v.Translator) == 1 {
			value = v.Translator[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["translator"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TypicalAgeRange
		if len(v.TypicalAgeRange) == 1 {
			value = v.TypicalAgeRange[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["typicalAgeRange"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkFeatured
		if len(v.WorkFeatured) == 1 {
			value = v.WorkFeatured[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workFeatured"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkPerformed
		if len(v.WorkPerformed) == 1 {
			value = v.WorkPerformed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workPerformed"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ChildrensEvent) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ChildrensEvent"

	return data, nil
}

func (v ChildrensEvent) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
