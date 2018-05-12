package schemaorggo

import "encoding/json"

// CookAction see : https://schema.org/CookAction
type CookAction struct {
	CreateAction

	typeContext

	// FoodEstablishment see : https://schema.org/foodEstablishment
	// A sub property of location. The specific food establishment where the action occurred.
	// types : FoodEstablishment Place
	FoodEstablishment interface{} `json:"foodEstablishment,omitempty"`

	// FoodEvent see : https://schema.org/foodEvent
	// A sub property of location. The specific food event where the action occurred.
	// types : FoodEvent
	FoodEvent *FoodEvent `json:"foodEvent,omitempty"`

	// Recipe see : https://schema.org/recipe
	// A sub property of instrument. The recipe/instructions used to perform the action.
	// types : Recipe
	Recipe *Recipe `json:"recipe,omitempty"`
}

func (v CookAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CookAction"

	return json.Marshal(v)
}

func (v *CookAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CookAction"

	return json.Marshal(*v)
}
