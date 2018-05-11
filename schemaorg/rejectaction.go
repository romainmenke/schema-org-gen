package schemaorg

import "encoding/json"

// RejectAction see : https://schema.org/RejectAction
type RejectAction struct {

typeContext

AllocateAction

// Purpose see : https://schema.orghttp://health-lifesci.schema.org/purpose
// A goal towards an action is taken. Can be concrete or abstract.
Purpose interface{} `json:"purpose"` // types : MedicalDevicePurpose Thing

}

func (v *RejectAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "RejectAction"

	return json.Marshal(v)
}
