package schemaorggo

import "encoding/json"

// AnimalShelter see : https://schema.org/AnimalShelter
type AnimalShelter struct {
	LocalBusiness

	typeContext

	// CurrenciesAccepted see : https://schema.org/currenciesAccepted
	// The currency accepted (in ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217)).
	CurrenciesAccepted string `json:"currenciesAccepted"`

	// OpeningHours see : https://schema.org/openingHours
	// The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas ',' separating each day. Day or time ranges are specified using a hyphen '-'.
	//
	//
	// Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
	// Times are specified using 24:00 time. For example, 3pm is specified as 15:00.
	// Here is an example: <time itemprop="openingHours" datetime="Tu,Th 16:00-20:00">Tuesdays and Thursdays 4-8pm</time>.
	// If a business is open 7 days a week, then it can be specified as <time itemprop="openingHours" datetime="Mo-Su">Monday through Sunday, all day</time>.
	//
	//
	OpeningHours string `json:"openingHours"`

	// PaymentAccepted see : https://schema.org/paymentAccepted
	// Cash, credit card, etc.
	PaymentAccepted string `json:"paymentAccepted"`

	// PriceRange see : https://schema.org/priceRange
	// The price range of the business, for example $$$.
	PriceRange string `json:"priceRange"`
}

func (v *AnimalShelter) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AnimalShelter"

	return json.Marshal(v)
}
