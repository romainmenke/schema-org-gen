package schemaorggo

import "encoding/json"

// CafeOrCoffeeShop see : https://schema.org/CafeOrCoffeeShop
type CafeOrCoffeeShop struct {
	FoodEstablishment

	typeContext

	// AcceptsReservations see : https://schema.org/acceptsReservations
	// Indicates whether a FoodEstablishment accepts reservations. Values can be Boolean, an URL at which reservations can be made or (for backwards compatibility) the strings Yes or No.
	// types : Boolean Text URL
	AcceptsReservations interface{} `json:"acceptsReservations,omitempty"`

	// HasMenu see : https://schema.org/hasMenu
	// Either the actual menu as a structured representation, as text, or a URL of the menu. Supersedes menu (see: https://schema.org/menu).
	// types : Menu Text URL
	HasMenu interface{} `json:"hasMenu,omitempty"`

	// ServesCuisine see : https://schema.org/servesCuisine
	// The cuisine of the restaurant.
	// types : Text
	ServesCuisine string `json:"servesCuisine,omitempty"`

	// StarRating see : https://schema.org/starRating
	// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	// types : Rating
	StarRating *Rating `json:"starRating,omitempty"`
}

func (v CafeOrCoffeeShop) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CafeOrCoffeeShop"

	return json.Marshal(v)
}

func (v *CafeOrCoffeeShop) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CafeOrCoffeeShop"

	return json.Marshal(*v)
}
