package schemaorg

import "encoding/json"

// MenuItem see : https://schema.org/MenuItem
type MenuItem struct {

typeContext

Intangible

// MenuAddOn see : http://pending.schema.org/menuAddOn
// Additional menu item(s) such as a side dish of salad or side order of fries that can be added to this menu item. Additionally it can be a menu section containing allowed add-on menu items for this menu item.
MenuAddOn interface{} `json:"menuAddOn"` // types : MenuItem MenuSection

// Nutrition see : /nutrition
// Nutrition information about the recipe or menu item.
Nutrition *NutritionInformation `json:"nutrition"`

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers *Offer `json:"offers"`

// SuitableForDiet see : /suitableForDiet
// Indicates a dietary restriction or guideline for which this recipe or menu item is suitable, e.g. diabetic, halal etc.
SuitableForDiet *RestrictedDiet `json:"suitableForDiet"`

}

func (v *MenuItem) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MenuItem"

	return json.Marshal(v)
}
