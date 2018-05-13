package schemaorggo

import "encoding/json"

// AggregateOffer see : https://schema.org/AggregateOffer
type AggregateOffer struct {
	Offer

	typeContext

	// HighPrice see : https://schema.org/highPrice
	// The highest price of all offers available.
	// types : Number Text
	HighPrice []interface{} `json:"highPrice,omitempty"`

	// LowPrice see : https://schema.org/lowPrice
	// The lowest price of all offers available.
	// types : Number Text
	LowPrice []interface{} `json:"lowPrice,omitempty"`

	// OfferCount see : https://schema.org/offerCount
	// The number of offers for the product.
	// types : Integer
	OfferCount []float64 `json:"offerCount,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`
}

func (v AggregateOffer) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Offer.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.HighPrice
		if len(v.HighPrice) == 1 {
			value = v.HighPrice[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["highPrice"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LowPrice
		if len(v.LowPrice) == 1 {
			value = v.LowPrice[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lowPrice"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OfferCount
		if len(v.OfferCount) == 1 {
			value = v.OfferCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offerCount"] = json.RawMessage(b)
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

	*intop = into

	return nil
}

func (v AggregateOffer) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AggregateOffer"

	return data, nil
}

func (v AggregateOffer) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
