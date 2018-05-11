package schemaorg

import "encoding/json"

// CookAction see : https://schema.org/CookAction
type CookAction struct {

typeContext

CreateAction

// FoodEstablishment see : /foodEstablishment
// A sub property of location. The specific food establishment where the action occurred.
FoodEstablishment interface{} `json:"foodEstablishment"` // types : FoodEstablishment Place

// FoodEvent see : /foodEvent
// A sub property of location. The specific food event where the action occurred.
FoodEvent *FoodEvent `json:"foodEvent"`

// Recipe see : /recipe
// A sub property of instrument. The recipe/instructions used to perform the action.
Recipe *Recipe `json:"recipe"`

}

func (v *CookAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CookAction"

	return json.Marshal(v)
}
