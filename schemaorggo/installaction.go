package schemaorggo

import "encoding/json"

// InstallAction see : https://schema.org/InstallAction
type InstallAction struct {
	ConsumeAction

	typeContext

	// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
	// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf,omitempty"` // types : Offer

}

func (v InstallAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InstallAction"

	return json.Marshal(v)
}

func (v *InstallAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "InstallAction"

	return json.Marshal(*v)
}
