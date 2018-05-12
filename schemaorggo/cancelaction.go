package schemaorggo

import "encoding/json"

// CancelAction see : https://schema.org/CancelAction
type CancelAction struct {
	PlanAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	ScheduledTime DateTime `json:"scheduledTime"`
}

func (v *CancelAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CancelAction"

	return json.Marshal(v)
}
