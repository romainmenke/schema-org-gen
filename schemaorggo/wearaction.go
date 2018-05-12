package schemaorggo

import "encoding/json"

// WearAction see : https://schema.org/WearAction
type WearAction struct {
	UseAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"` // types : Offer

}

func (v WearAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WearAction"

	return json.Marshal(v)
}

func (v *WearAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "WearAction"

	return json.Marshal(*v)
}
