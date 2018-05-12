package schemaorggo

import "encoding/json"

// EatAction see : https://schema.org/EatAction
type EatAction struct {
	ConsumeAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	// types : Offer
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v EatAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EatAction"

	return json.Marshal(v)
}

func (v *EatAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EatAction"

	return json.Marshal(*v)
}
