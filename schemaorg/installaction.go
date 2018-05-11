package schemaorg

import "encoding/json"

// InstallAction see : https://schema.org/InstallAction
type InstallAction struct {

typeContext

ConsumeAction

// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *InstallAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InstallAction"

	return json.Marshal(v)
}
