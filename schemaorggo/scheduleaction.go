package schemaorggo

import "encoding/json"

// ScheduleAction see : https://schema.org/ScheduleAction
type ScheduleAction struct {
	PlanAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	ScheduledTime DateTime `json:"scheduledTime,omitempty"` // types : DateTime

}

func (v ScheduleAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ScheduleAction"

	return json.Marshal(v)
}

func (v *ScheduleAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ScheduleAction"

	return json.Marshal(*v)
}
