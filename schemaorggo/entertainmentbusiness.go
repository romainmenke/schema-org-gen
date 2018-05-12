package schemaorggo

import "encoding/json"

// EntertainmentBusiness see : https://schema.org/EntertainmentBusiness
type EntertainmentBusiness struct {
	LocalBusiness

	typeContext

	// CurrenciesAccepted see : https://schema.org/currenciesAccepted
	// The currency accepted (in ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217)).
	// types : Text
	CurrenciesAccepted string `json:"currenciesAccepted,omitempty"`

	// OpeningHours see : https://schema.org/openingHours
	// The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.
	//
	//
	// Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
	// Times are specified using 24:00 time. For example, 3pm is specified as 15:00.
	// Here is an example: &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;.
	// If a business is open 7 days a week, then it can be specified as &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;.
	//
	//
	// types : Text
	OpeningHours string `json:"openingHours,omitempty"`

	// PaymentAccepted see : https://schema.org/paymentAccepted
	// Cash, credit card, etc.
	// types : Text
	PaymentAccepted string `json:"paymentAccepted,omitempty"`

	// PriceRange see : https://schema.org/priceRange
	// The price range of the business, for example $$$.
	// types : Text
	PriceRange string `json:"priceRange,omitempty"`
}

func (v EntertainmentBusiness) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EntertainmentBusiness"

	return json.Marshal(v)
}

func (v *EntertainmentBusiness) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EntertainmentBusiness"

	return json.Marshal(*v)
}
