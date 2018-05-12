package schemaorg

import "encoding/json"

// AggregateOffer see : https://schema.org/AggregateOffer
type AggregateOffer struct {

Offer

typeContext

// HighPrice see : https://schema.org/highPrice
// The highest price of all offers available.
HighPrice interface{} `json:"highPrice"` // types : Number Text

// LowPrice see : https://schema.org/lowPrice
// The lowest price of all offers available.
LowPrice interface{} `json:"lowPrice"` // types : Number Text

// OfferCount see : https://schema.org/offerCount
// The number of offers for the product.
OfferCount int `json:"offerCount"`

// Offers see : https://schema.org/offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers *Offer `json:"offers"`

}

func (v *AggregateOffer) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AggregateOffer"

	return json.Marshal(v)
}
