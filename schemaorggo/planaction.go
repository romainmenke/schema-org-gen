package schemaorggo

import "encoding/json"

// PlanAction see : https://schema.org/PlanAction
type PlanAction struct {
	OrganizeAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	// types : DateTime
	ScheduledTime []DateTime `json:"scheduledTime,omitempty"`
}

func (v PlanAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.OrganizeAction.IntoMap(intop)

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

func (v PlanAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "PlanAction"

	return data, nil
}

func (v PlanAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
