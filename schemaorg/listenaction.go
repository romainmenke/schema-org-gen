package schemaorg

import "encoding/json"

// ListenAction see : https://schema.org/ListenAction
type ListenAction struct {

typeContext

ConsumeAction

// ExpectsAcceptanceOf see : /expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *ListenAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ListenAction"

	return json.Marshal(v)
}
