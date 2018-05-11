package schemaorg

import "encoding/json"

// QuoteAction see : https://schema.org/QuoteAction
type QuoteAction struct {

typeContext

TradeAction

// Price see : https://schema.org/price
// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
// 
// Usage guidelines:
// 
// 
// Use the priceCurrency (see: https://schema.org/priceCurrency) property (with ISO 4217 codes (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217#Active_codes) e.g. "USD") instead of
//   including ambiguous symbols (see: https://schema.orghttp://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as '$' in the value.
// Use '.' (Unicode 'FULL STOP' (U+002E)) rather than ',' to indicate a decimal point. Avoid using these symbols as a readability separator.
// Note that both RDFa (see: https://schema.orghttp://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a "content=" attribute for publishing simple machine-readable values alongside more human-friendly formatting.
// Use values from 0123456789 (Unicode 'DIGIT ZERO' (U+0030) to 'DIGIT NINE' (U+0039)) rather than superficially similiar Unicode symbols.
// 
// 
Price interface{} `json:"price"` // types : Number Text

// PriceSpecification see : https://schema.org/priceSpecification
// One or more detailed price specifications, indicating the unit price and delivery or payment charges.
PriceSpecification *PriceSpecification `json:"priceSpecification"`

}

func (v *QuoteAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "QuoteAction"

	return json.Marshal(v)
}
