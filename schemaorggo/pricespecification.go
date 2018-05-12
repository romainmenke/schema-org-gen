package schemaorggo

import "encoding/json"

// PriceSpecification see : https://schema.org/PriceSpecification
type PriceSpecification struct {
	StructuredValue

	typeContext

	// EligibleQuantity see : https://schema.org/eligibleQuantity
	// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	// types : QuantitativeValue
	EligibleQuantity *QuantitativeValue `json:"eligibleQuantity,omitempty"`

	// EligibleTransactionVolume see : https://schema.org/eligibleTransactionVolume
	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	// types : PriceSpecification
	EligibleTransactionVolume *PriceSpecification `json:"eligibleTransactionVolume,omitempty"`

	// MaxPrice see : https://schema.org/maxPrice
	// The highest price if the price is a range.
	// types : Number
	MaxPrice float64 `json:"maxPrice,omitempty"`

	// MinPrice see : https://schema.org/minPrice
	// The lowest price if the price is a range.
	// types : Number
	MinPrice float64 `json:"minPrice,omitempty"`

	// Price see : https://schema.org/price
	// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
	//
	// Usage guidelines:
	//
	//
	// Use the priceCurrency (see: https://schema.org/priceCurrency) property (with ISO 4217 codes (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217#Active_codes) e.g. &quot;USD&quot;) instead of
	//   including ambiguous symbols (see: https://schema.orghttp://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign) such as &#39;$&#39; in the value.
	// Use &#39;.&#39; (Unicode &#39;FULL STOP&#39; (U+002E)) rather than &#39;,&#39; to indicate a decimal point. Avoid using these symbols as a readability separator.
	// Note that both RDFa (see: https://schema.orghttp://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute) and Microdata syntax allow the use of a &quot;content=&quot; attribute for publishing simple machine-readable values alongside more human-friendly formatting.
	// Use values from 0123456789 (Unicode &#39;DIGIT ZERO&#39; (U+0030) to &#39;DIGIT NINE&#39; (U+0039)) rather than superficially similiar Unicode symbols.
	//
	//
	// types : Number Text
	Price interface{} `json:"price,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	// types : Text
	PriceCurrency string `json:"priceCurrency,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough DateTime `json:"validThrough,omitempty"`

	// ValueAddedTaxIncluded see : https://schema.org/valueAddedTaxIncluded
	// Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
	// types : Boolean
	ValueAddedTaxIncluded bool `json:"valueAddedTaxIncluded,omitempty"`
}

func (v PriceSpecification) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PriceSpecification"

	return json.Marshal(v)
}

func (v *PriceSpecification) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PriceSpecification"

	return json.Marshal(*v)
}
