package schemaorggo

import "encoding/json"

// DrinkAction see : https://schema.org/DrinkAction
type DrinkAction struct {
	ConsumeAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	// types : Offer
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v DrinkAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DrinkAction"

	return json.Marshal(v)
}

func (v *DrinkAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DrinkAction"

	return json.Marshal(*v)
}
