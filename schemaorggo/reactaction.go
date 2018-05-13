package schemaorggo

import "encoding/json"

// ReactAction see : https://schema.org/ReactAction
type ReactAction struct {
	AssessAction

	typeContext

	// ActionStatus see : https://schema.org/actionStatus
	// Indicates the current disposition of the Action.
	// types : ActionStatusType
	ActionStatus []*ActionStatusType `json:"actionStatus,omitempty"`

	// Agent see : https://schema.org/agent
	// The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book.
	// types : Organization Person
	Agent []interface{} `json:"agent,omitempty"`

	// EndTime see : https://schema.org/endTime
	// The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	EndTime []DateTime `json:"endTime,omitempty"`

	// Error see : https://schema.org/error
	// For failed actions, more information on the cause of the failure.
	// types : Thing
	Error []*Thing `json:"error,omitempty"`

	// Instrument see : https://schema.org/instrument
	// The object that helped the agent perform the action. e.g. John wrote a book with a pen.
	// types : Thing
	Instrument []*Thing `json:"instrument,omitempty"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
	Location []interface{} `json:"location,omitempty"`

	// Object see : https://schema.org/object
	// The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn&#39;t). e.g. John read a book.
	// types : Thing
	Object []*Thing `json:"object,omitempty"`

	// Participant see : https://schema.org/participant
	// Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve.
	// types : Organization Person
	Participant []interface{} `json:"participant,omitempty"`

	// Result see : https://schema.org/result
	// The result produced in the action. e.g. John wrote a book.
	// types : Thing
	Result []*Thing `json:"result,omitempty"`

	// StartTime see : https://schema.org/startTime
	// The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	//
	// Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	// types : DateTime
	StartTime []DateTime `json:"startTime,omitempty"`

	// Target see : https://schema.org/target
	// Indicates a target EntryPoint for an Action.
	// types : EntryPoint
	Target []*EntryPoint `json:"target,omitempty"`
}

func (v ReactAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.AssessAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ActionStatus
		if len(v.ActionStatus) == 1 {
			value = v.ActionStatus[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actionStatus"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Agent
		if len(v.Agent) == 1 {
			value = v.Agent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["agent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EndTime
		if len(v.EndTime) == 1 {
			value = v.EndTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["endTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Error
		if len(v.Error) == 1 {
			value = v.Error[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["error"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Instrument
		if len(v.Instrument) == 1 {
			value = v.Instrument[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["instrument"] = json.RawMessage(b)
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
		var value interface{} = v.Object
		if len(v.Object) == 1 {
			value = v.Object[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["object"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Participant
		if len(v.Participant) == 1 {
			value = v.Participant[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["participant"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Result
		if len(v.Result) == 1 {
			value = v.Result[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["result"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StartTime
		if len(v.StartTime) == 1 {
			value = v.StartTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["startTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Target
		if len(v.Target) == 1 {
			value = v.Target[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["target"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReactAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReactAction"

	return data, nil
}

func (v ReactAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
