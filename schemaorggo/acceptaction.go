package schemaorggo

import "encoding/json"

// AcceptAction see : https://schema.org/AcceptAction
type AcceptAction struct {
	AllocateAction

	typeContext

	// Purpose see : http://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	Purpose interface{} `json:"purpose,omitempty"` // types : MedicalDevicePurpose Thing

}

func (v AcceptAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AcceptAction"

	return json.Marshal(v)
}

func (v *AcceptAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AcceptAction"

	return json.Marshal(*v)
}
