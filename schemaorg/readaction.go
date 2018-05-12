package schemaorg

import "encoding/json"

// ReadAction see : https://schema.org/ReadAction
type ReadAction struct {

ConsumeAction

typeContext

// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *ReadAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReadAction"

	return json.Marshal(v)
}
