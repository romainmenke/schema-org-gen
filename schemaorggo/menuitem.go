package schemaorggo

import "encoding/json"

// MenuItem see : https://schema.org/MenuItem
type MenuItem struct {
	Intangible

	typeContext

	// MenuAddOn see : http://pending.schema.org/menuAddOn
	// Additional menu item(s) such as a side dish of salad or side order of fries that can be added to this menu item. Additionally it can be a menu section containing allowed add-on menu items for this menu item.
	// types : MenuItem MenuSection
	MenuAddOn interface{} `json:"menuAddOn,omitempty"`

	// Nutrition see : https://schema.org/nutrition
	// Nutrition information about the recipe or menu item.
	// types : NutritionInformation
	Nutrition *NutritionInformation `json:"nutrition,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers *Offer `json:"offers,omitempty"`

	// SuitableForDiet see : https://schema.org/suitableForDiet
	// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
	// types : RestrictedDiet
	SuitableForDiet *RestrictedDiet `json:"suitableForDiet,omitempty"`
}

func (v MenuItem) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MenuItem"

	return json.Marshal(v)
}

func (v *MenuItem) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MenuItem"

	return json.Marshal(*v)
}
