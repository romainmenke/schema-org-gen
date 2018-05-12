package schemaorggo

import "encoding/json"

// RejectAction see : https://schema.org/RejectAction
type RejectAction struct {
	AllocateAction

	typeContext

	// Purpose see : http://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	// types : MedicalDevicePurpose Thing
	Purpose interface{} `json:"purpose,omitempty"`
}

func (v RejectAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RejectAction"

	return json.Marshal(v)
}

func (v *RejectAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "RejectAction"

	return json.Marshal(*v)
}
