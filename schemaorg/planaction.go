package schemaorg

import "encoding/json"

// PlanAction see : https://schema.org/PlanAction
type PlanAction struct {

typeContext

OrganizeAction

// ScheduledTime see : https://schema.org/scheduledTime
// The time the object is scheduled to.
ScheduledTime interface{} `json:"scheduledTime"`

}

func (v *PlanAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PlanAction"

	return json.Marshal(v)
}
