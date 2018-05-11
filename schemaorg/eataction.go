package schemaorg

import "encoding/json"

// EatAction see : https://schema.org/EatAction
type EatAction struct {

typeContext

ConsumeAction

// ExpectsAcceptanceOf see : /expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *EatAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EatAction"

	return json.Marshal(v)
}
