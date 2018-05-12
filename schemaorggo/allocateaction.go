package schemaorggo

import "encoding/json"

// AllocateAction see : https://schema.org/AllocateAction
type AllocateAction struct {
	OrganizeAction

	typeContext

	// Purpose see : https://schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	Purpose interface{} `json:"purpose,omitempty"` // types : MedicalDevicePurpose Thing

}

func (v AllocateAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AllocateAction"

	return json.Marshal(v)
}

func (v *AllocateAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AllocateAction"

	return json.Marshal(*v)
}
