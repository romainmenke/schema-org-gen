package schemaorg

import "encoding/json"

// WearAction see : https://schema.org/WearAction
type WearAction struct {

typeContext

UseAction

// ExpectsAcceptanceOf see : /expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *WearAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WearAction"

	return json.Marshal(v)
}
