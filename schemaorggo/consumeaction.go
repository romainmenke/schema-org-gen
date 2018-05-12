package schemaorggo

import "encoding/json"

// ConsumeAction see : https://schema.org/ConsumeAction
type ConsumeAction struct {
	Action

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	// types : Offer
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v ConsumeAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ConsumeAction"

	return json.Marshal(v)
}

func (v *ConsumeAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ConsumeAction"

	return json.Marshal(*v)
}
