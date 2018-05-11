package schemaorg

import "encoding/json"

// UseAction see : https://schema.org/UseAction
type UseAction struct {

typeContext

ConsumeAction

// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *UseAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "UseAction"

	return json.Marshal(v)
}
