package schemaorggo

import "encoding/json"

// AggregateOffer see : https://schema.org/AggregateOffer
type AggregateOffer struct {
	Offer

	typeContext

	// HighPrice see : https://schema.org/highPrice
	// The highest price of all offers available.
	// types : Number Text
	HighPrice interface{} `json:"highPrice,omitempty"`

	// LowPrice see : https://schema.org/lowPrice
	// The lowest price of all offers available.
	// types : Number Text
	LowPrice interface{} `json:"lowPrice,omitempty"`

	// OfferCount see : https://schema.org/offerCount
	// The number of offers for the product.
	// types : Integer
	OfferCount float64 `json:"offerCount,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers *Offer `json:"offers,omitempty"`
}

func (v AggregateOffer) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AggregateOffer"

	return json.Marshal(v)
}

func (v *AggregateOffer) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AggregateOffer"

	return json.Marshal(*v)
}
