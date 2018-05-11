package schemaorg

import "encoding/json"

// FastFoodRestaurant see : https://schema.org/FastFoodRestaurant
type FastFoodRestaurant struct {

typeContext

FoodEstablishment

// AcceptsReservations see : https://schema.org/acceptsReservations
// Indicates whether a FoodEstablishment accepts reservations. Values can be Boolean, an URL at which reservations can be made or (for backwards compatibility) the strings Yes or No.
AcceptsReservations interface{} `json:"acceptsReservations"` // types : Boolean Text URL

// HasMenu see : https://schema.org/hasMenu
// Either the actual menu as a structured representation, as text, or a URL of the menu. Supersedes menu (see: https://schema.org/menu).
HasMenu interface{} `json:"hasMenu"` // types : Menu Text URL

// ServesCuisine see : https://schema.org/servesCuisine
// The cuisine of the restaurant.
ServesCuisine string `json:"servesCuisine"`

// StarRating see : https://schema.org/starRating
// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
StarRating *Rating `json:"starRating"`

}

func (v *FastFoodRestaurant) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FastFoodRestaurant"

	return json.Marshal(v)
}
