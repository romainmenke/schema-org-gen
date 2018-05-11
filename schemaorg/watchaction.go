package schemaorg

import "encoding/json"

// WatchAction see : https://schema.org/WatchAction
type WatchAction struct {

typeContext

ConsumeAction

// ExpectsAcceptanceOf see : /expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *WatchAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WatchAction"

	return json.Marshal(v)
}
