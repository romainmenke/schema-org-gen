package schemaorg

import "encoding/json"

// ConsumeAction see : https://schema.org/ConsumeAction
type ConsumeAction struct {

typeContext

Action

// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *ConsumeAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ConsumeAction"

	return json.Marshal(v)
}
