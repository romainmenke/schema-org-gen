package schemaorggo

import "encoding/json"

// AssignAction see : https://schema.org/AssignAction
type AssignAction struct {
	AllocateAction

	typeContext

	// Purpose see : http://health-lifesci.schema.org/purpose
	// A goal towards an action is taken. Can be concrete or abstract.
	Purpose interface{} `json:"purpose,omitempty"` // types : MedicalDevicePurpose Thing

}

func (v AssignAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AssignAction"

	return json.Marshal(v)
}

func (v *AssignAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AssignAction"

	return json.Marshal(*v)
}
