package schemaorggo

import "encoding/json"

// MenuItem see : https://schema.org/MenuItem
type MenuItem struct {
	Intangible

	typeContext

	// MenuAddOn see : http://pending.schema.org/menuAddOn
	// Additional menu item(s) such as a side dish of salad or side order of fries that can be added to this menu item. Additionally it can be a menu section containing allowed add-on menu items for this menu item.
	// types : MenuItem MenuSection
	MenuAddOn []interface{} `json:"menuAddOn,omitempty"`

	// Nutrition see : https://schema.org/nutrition
	// Nutrition information about the recipe or menu item.
	// types : NutritionInformation
	Nutrition []*NutritionInformation `json:"nutrition,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// SuitableForDiet see : https://schema.org/suitableForDiet
	// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	// types : RestrictedDiet
	SuitableForDiet []*RestrictedDiet `json:"suitableForDiet,omitempty"`
}

func (v MenuItem) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.MenuAddOn
		if len(v.MenuAddOn) == 1 {
			value = v.MenuAddOn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["menuAddOn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Nutrition
		if len(v.Nutrition) == 1 {
			value = v.Nutrition[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["nutrition"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Offers
		if len(v.Offers) == 1 {
			value = v.Offers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuitableForDiet
		if len(v.SuitableForDiet) == 1 {
			value = v.SuitableForDiet[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["suitableForDiet"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MenuItem) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MenuItem"

	return data, nil
}

func (v MenuItem) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
