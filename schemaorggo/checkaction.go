package schemaorggo

import "encoding/json"

// CheckAction see : https://schema.org/CheckAction
type CheckAction struct {
	FindAction

	typeContext

	// ActionStatus see : https://schema.org/actionStatus
	// Indicates the current disposition of the Action.
	ActionStatus *ActionStatusType `json:"actionStatus"`

	// Agent see : https://schema.org/agent
	// The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book.
	Agent interface{} `json:"agent"` // types : Organization Person

	// EndTime see : https://schema.org/endTime
	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	EndTime interface{} `json:"endTime"`

	// Error see : https://schema.org/error
	// For failed actions, more information on the cause of the failure.
	Error *Thing `json:"error"`

	// Instrument see : https://schema.org/instrument
	// The object that helped the agent perform the action. e.g. John wrote a book with a pen.
	Instrument *Thing `json:"instrument"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	Location interface{} `json:"location"` // types : Place PostalAddress Text

	// Object see : https://schema.org/object
	// The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn't). e.g. John read a book.
	Object *Thing `json:"object"`

	// Participant see : https://schema.org/participant
	// Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve.
	Participant interface{} `json:"participant"` // types : Organization Person

	// Result see : https://schema.org/result
	// The result produced in the action. e.g. John wrote a book.
	Result *Thing `json:"result"`

	// StartTime see : https://schema.org/startTime
	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	StartTime interface{} `json:"startTime"`

	// Target see : https://schema.org/target
	// Indicates a target EntryPoint for an Action.
	Target *EntryPoint `json:"target"`
}

func (v *CheckAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CheckAction"

	return json.Marshal(v)
}
