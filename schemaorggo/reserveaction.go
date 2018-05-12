package schemaorggo

import "encoding/json"

// ReserveAction see : https://schema.org/ReserveAction
type ReserveAction struct {
	PlanAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	ScheduledTime DateTime `json:"scheduledTime"`
}

func (v *ReserveAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReserveAction"

	return json.Marshal(v)
}
