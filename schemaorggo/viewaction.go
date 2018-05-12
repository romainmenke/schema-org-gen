package schemaorggo

import "encoding/json"

// ViewAction see : https://schema.org/ViewAction
type ViewAction struct {
	ConsumeAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"`
}

func (v ViewAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ViewAction"

	return json.Marshal(v)
}

func (v *ViewAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ViewAction"

	return json.Marshal(*v)
}
