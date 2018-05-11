package schemaorg

import "encoding/json"

// DrinkAction see : https://schema.org/DrinkAction
type DrinkAction struct {

typeContext

ConsumeAction

// ExpectsAcceptanceOf see : https://schema.org/expectsAcceptanceOf
// An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
ExpectsAcceptanceOf *Offer `json:"expectsAcceptanceOf"`

}

func (v *DrinkAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DrinkAction"

	return json.Marshal(v)
}
