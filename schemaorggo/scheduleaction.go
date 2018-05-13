package schemaorggo

import "encoding/json"

// ScheduleAction see : https://schema.org/ScheduleAction
type ScheduleAction struct {
	PlanAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	// types : DateTime
	ScheduledTime []DateTime `json:"scheduledTime,omitempty"`
}

func (v ScheduleAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PlanAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ScheduledTime
		if len(v.ScheduledTime) == 1 {
			value = v.ScheduledTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["scheduledTime"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ScheduleAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ScheduleAction"

	return data, nil
}

func (v ScheduleAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
