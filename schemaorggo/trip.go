package schemaorggo

import "encoding/json"

// Trip see : https://schema.org/Trip
type Trip struct {
	Intangible

	typeContext

	// ArrivalTime see : https://schema.org/arrivalTime
	// The expected arrival time.
	// types : DateTime
	ArrivalTime []DateTime `json:"arrivalTime,omitempty"`

	// DepartureTime see : https://schema.org/departureTime
	// The expected departure time.
	// types : DateTime
	DepartureTime []DateTime `json:"departureTime,omitempty"`

	// HasPart see : https://schema.org/hasPart
	// Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense). Inverse property: isPartOf (see: https://schema.org/isPartOf).
	// types : CreativeWork Trip
	HasPart []interface{} `json:"hasPart,omitempty"`

	// IsPartOf see : https://schema.org/isPartOf
	// Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of. Inverse property: hasPart (see: https://schema.org/hasPart).
	// types : CreativeWork Trip
	IsPartOf []interface{} `json:"isPartOf,omitempty"`

	// Itinerary see : https://pending.schema.org/itinerary
	// Destination(s) ( Place (see: https://schema.org/Place) ) that make up a trip. For a trip where destination order is important use ItemList (see: https://schema.org/ItemList) to specify that order (see examples).
	// types : ItemList Place
	Itinerary []interface{} `json:"itinerary,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`
}

func (v Trip) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ArrivalTime
		if len(v.ArrivalTime) == 1 {
			value = v.ArrivalTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["arrivalTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DepartureTime
		if len(v.DepartureTime) == 1 {
			value = v.DepartureTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["departureTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasPart
		if len(v.HasPart) == 1 {
			value = v.HasPart[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasPart"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsPartOf
		if len(v.IsPartOf) == 1 {
			value = v.IsPartOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isPartOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Itinerary
		if len(v.Itinerary) == 1 {
			value = v.Itinerary[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itinerary"] = json.RawMessage(b)
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
		var value interface{} = v.Provider
		if len(v.Provider) == 1 {
			value = v.Provider[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["provider"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Trip) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Trip"

	return data, nil
}

func (v Trip) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
