package schemaorggo

import "encoding/json"

// SubscribeAction see : https://schema.org/SubscribeAction
type SubscribeAction struct {
	InteractAction

	typeContext

	// ActionStatus see : https://schema.org/actionStatus
	// Indicates the current disposition of the Action.
	// types : ActionStatusType
	ActionStatus *ActionStatusType `json:"actionStatus,omitempty"`

	// Agent see : https://schema.org/agent
	// The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book.
	// types : Organization Person
	Agent interface{} `json:"agent,omitempty"`

	// EndTime see : https://schema.org/endTime
	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	EndTime DateTime `json:"endTime,omitempty"`

	// Error see : https://schema.org/error
	// For failed actions, more information on the cause of the failure.
	// types : Thing
	Error *Thing `json:"error,omitempty"`

	// Instrument see : https://schema.org/instrument
	// The object that helped the agent perform the action. e.g. John wrote a book with a pen.
	// types : Thing
	Instrument *Thing `json:"instrument,omitempty"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
	Location interface{} `json:"location,omitempty"`

	// Object see : https://schema.org/object
	// The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn&#39;t). e.g. John read a book.
	// types : Thing
	Object *Thing `json:"object,omitempty"`

	// Participant see : https://schema.org/participant
	// Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve.
	// types : Organization Person
	Participant interface{} `json:"participant,omitempty"`

	// Result see : https://schema.org/result
	// The result produced in the action. e.g. John wrote a book.
	// types : Thing
	Result *Thing `json:"result,omitempty"`

	// StartTime see : https://schema.org/startTime
	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	StartTime DateTime `json:"startTime,omitempty"`

	// Target see : https://schema.org/target
	// Indicates a target EntryPoint for an Action.
	// types : EntryPoint
	Target *EntryPoint `json:"target,omitempty"`
}

func (v SubscribeAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SubscribeAction"

	return json.Marshal(v)
}

func (v *SubscribeAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "SubscribeAction"

	return json.Marshal(*v)
}
