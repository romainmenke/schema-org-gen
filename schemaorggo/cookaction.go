package schemaorggo

import "encoding/json"

// CookAction see : https://schema.org/CookAction
type CookAction struct {
	CreateAction

	typeContext

	// FoodEstablishment see : https://schema.org/foodEstablishment
	// A sub property of location. The specific food establishment where the action occurred.
	// types : FoodEstablishment Place
	FoodEstablishment []interface{} `json:"foodEstablishment,omitempty"`

	// FoodEvent see : https://schema.org/foodEvent
	// A sub property of location. The specific food event where the action occurred.
	// types : FoodEvent
	FoodEvent []*FoodEvent `json:"foodEvent,omitempty"`

	// Recipe see : https://schema.org/recipe
	// A sub property of instrument. The recipe/instructions used to perform the action.
	// types : Recipe
	Recipe []*Recipe `json:"recipe,omitempty"`
}

func (v CookAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreateAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.FoodEstablishment
		if len(v.FoodEstablishment) == 1 {
			value = v.FoodEstablishment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["foodEstablishment"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FoodEvent
		if len(v.FoodEvent) == 1 {
			value = v.FoodEvent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["foodEvent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Recipe
		if len(v.Recipe) == 1 {
			value = v.Recipe[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recipe"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v CookAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "CookAction"

	return data, nil
}

func (v CookAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
