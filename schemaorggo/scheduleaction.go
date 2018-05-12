package schemaorggo

import "encoding/json"

// ScheduleAction see : https://schema.org/ScheduleAction
type ScheduleAction struct {
	PlanAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	ScheduledTime DateTime `json:"scheduledTime"`
}

func (v *ScheduleAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ScheduleAction"

	return json.Marshal(v)
}
