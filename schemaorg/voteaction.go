package schemaorg

import "encoding/json"

// VoteAction see : https://schema.org/VoteAction
type VoteAction struct {

typeContext

ChooseAction

// Candidate see : https://schema.org/candidate
// A sub property of object. The candidate subject of this action.
Candidate *Person `json:"candidate"`

}

func (v *VoteAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VoteAction"

	return json.Marshal(v)
}
