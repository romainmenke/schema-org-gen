package schemaorggo

import "encoding/json"

// IceCreamShop see : https://schema.org/IceCreamShop
type IceCreamShop struct {
	FoodEstablishment

	typeContext

	// AcceptsReservations see : https://schema.org/acceptsReservations
	// Indicates whether a FoodEstablishment accepts reservations. Values can be Boolean, an URL at which reservations can be made or (for backwards compatibility) the strings Yes or No.
	// types : Boolean Text URL
	AcceptsReservations []interface{} `json:"acceptsReservations,omitempty"`

	// HasMenu see : https://schema.org/hasMenu
	// Either the actual menu as a structured representation, as text, or a URL of the menu. Supersedes menu (see: https://schema.org/menu).
	// types : Menu Text URL
	HasMenu []interface{} `json:"hasMenu,omitempty"`

	// ServesCuisine see : https://schema.org/servesCuisine
	// The cuisine of the restaurant.
	// types : Text
	ServesCuisine []string `json:"servesCuisine,omitempty"`

	// StarRating see : https://schema.org/starRating
	// An official rating for a lodging business or food establishment, e.g. from national associations or standards bodies. Use the author property to indicate the rating organization, e.g. as an Organization with name such as (e.g. HOTREC, DEHOGA, WHR, or Hotelstars).
	// types : Rating
	StarRating []*Rating `json:"starRating,omitempty"`
}

func (v IceCreamShop) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.FoodEstablishment.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AcceptsReservations
		if len(v.AcceptsReservations) == 1 {
			value = v.AcceptsReservations[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["acceptsReservations"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasMenu
		if len(v.HasMenu) == 1 {
			value = v.HasMenu[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasMenu"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ServesCuisine
		if len(v.ServesCuisine) == 1 {
			value = v.ServesCuisine[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["servesCuisine"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.StarRating
		if len(v.StarRating) == 1 {
			value = v.StarRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["starRating"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v IceCreamShop) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "IceCreamShop"

	return data, nil
}

func (v IceCreamShop) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
