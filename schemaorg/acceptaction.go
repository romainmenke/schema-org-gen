package schemaorg

import "encoding/json"

// AcceptAction see : https://schema.org/AcceptAction
type AcceptAction struct {

typeContext

AllocateAction

// Purpose see : https://schema.orghttp://health-lifesci.schema.org/purpose
// A goal towards an action is taken. Can be concrete or abstract.
Purpose interface{} `json:"purpose"` // types : MedicalDevicePurpose Thing

}

func (v *AcceptAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AcceptAction"

	return json.Marshal(v)
}
