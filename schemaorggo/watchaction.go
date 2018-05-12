package schemaorggo

import "encoding/json"

// WatchAction see : https://schema.org/WatchAction
type WatchAction struct {
	ConsumeAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	// types : Offer
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v WatchAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WatchAction"

	return json.Marshal(v)
}

func (v *WatchAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WatchAction"

	return json.Marshal(*v)
}
