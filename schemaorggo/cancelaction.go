package schemaorggo

import "encoding/json"

// CancelAction see : https://schema.org/CancelAction
type CancelAction struct {
	PlanAction

	typeContext

	// ScheduledTime see : https://schema.org/scheduledTime
	// The time the object is scheduled to.
	// types : DateTime
	ScheduledTime DateTime `json:"scheduledTime,omitempty"`
}

func (v CancelAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CancelAction"

	return json.Marshal(v)
}

func (v *CancelAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CancelAction"

	return json.Marshal(*v)
}
