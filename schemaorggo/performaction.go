package schemaorggo

import "encoding/json"

// PerformAction see : https://schema.org/PerformAction
type PerformAction struct {
	PlayAction

	typeContext

	// EntertainmentBusiness see : https://schema.org/entertainmentBusiness
	// A sub property of location. The entertainment business where the action occurred.
	EntertainmentBusiness *EntertainmentBusiness `json:"entertainmentBusiness,omitempty"` // types : EntertainmentBusiness

}

func (v PerformAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PerformAction"

	return json.Marshal(v)
}

func (v *PerformAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PerformAction"

	return json.Marshal(*v)
}
