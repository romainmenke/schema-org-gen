package schemaorggo

import "encoding/json"

// PlanAction see : https://schema.org/PlanAction
type PlanAction struct {
	OrganizeAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	ScheduledTime DateTime `json:"scheduledTime"`
}

func (v *PlanAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PlanAction"

	return json.Marshal(v)
}
