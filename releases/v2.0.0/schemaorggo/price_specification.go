package schemaorggo

import "encoding/json"

// PriceSpecification see : https://schema.org/PriceSpecification
type PriceSpecification struct {

	// With properties from StructuredValue see : https://schema.org/StructuredValue
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// EligibleQuantity see : https://schema.org/eligibleQuantity
	// The interval and unit of measurement of ordering quantities for which the offer or price specification is valid. This allows e.g. specifying that a certain freight charge is valid only for a certain quantity.
	// types : QuantitativeValue
	EligibleQuantity []*QuantitativeValue `json:"eligibleQuantity,omitempty"`

	// EligibleTransactionVolume see : https://schema.org/eligibleTransactionVolume
	// The transaction volume, in a monetary unit, for which the offer or price specification is valid, e.g. for indicating a minimal purchasing volume, to express free shipping above a certain order volume, or to limit the acceptance of credit cards to purchases to a certain minimal amount.
	// types : PriceSpecification
	EligibleTransactionVolume []*PriceSpecification `json:"eligibleTransactionVolume,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MaxPrice see : https://schema.org/maxPrice
	// The highest price if the price is a range.
	// types : Number
	MaxPrice []float64 `json:"maxPrice,omitempty"`

	// MinPrice see : https://schema.org/minPrice
	// The lowest price if the price is a range.
	// types : Number
	MinPrice []float64 `json:"minPrice,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Price see : https://schema.org/price
	// The offer price of a product, or of a price component when attached to PriceSpecification and its subtypes.
	// &lt;br /&gt;
	// &lt;br /&gt;
	//       Usage guidelines:
	// &lt;br /&gt;
	// &lt;ul&gt;
	// &lt;li&gt;Use the &lt;a href=&quot;/priceCurrency&quot;&gt;priceCurrency&lt;/a&gt; property (with &lt;a href=&quot;http://en.wikipedia.org/wiki/ISO_4217#Active_codes&quot;&gt;ISO 4217 codes&lt;/a&gt; e.g. &quot;USD&quot;) instead of
	//       including &lt;a href=&quot;http://en.wikipedia.org/wiki/Dollar_sign#Currencies_that_use_the_dollar_or_peso_sign&quot;&gt;ambiguous symbols&lt;/a&gt; such as &#39;$&#39; in the value.
	// &lt;/li&gt;
	// &lt;li&gt;
	//       Use &#39;.&#39; (Unicode &#39;FULL STOP&#39; (U+002E)) rather than &#39;,&#39; to indicate a decimal point. Avoid using these symbols as a readability separator.
	// &lt;/li&gt;
	// &lt;li&gt;
	//       Note that both &lt;a href=&quot;http://www.w3.org/TR/xhtml-rdfa-primer/#using-the-content-attribute&quot;&gt;RDFa&lt;/a&gt; and Microdata syntax allow the use of a &quot;content=&quot; attribute for publishing simple machine-readable values
	//       alongside more human-friendly formatting.
	// &lt;/li&gt;
	// &lt;li&gt;
	//       Use values from 0123456789 (Unicode &#39;DIGIT ZERO&#39; (U+0030) to &#39;DIGIT NINE&#39; (U+0039)) rather than superficially similiar Unicode symbols.
	// &lt;/li&gt;
	// &lt;/ul&gt;
	//
	// types : Number Text
	Price []interface{} `json:"price,omitempty"`

	// PriceCurrency see : https://schema.org/priceCurrency
	// The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification and its subtypes.
	// types : Text
	PriceCurrency []string `json:"priceCurrency,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom []DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The end of the validity of offer, price specification, or opening hours data.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`

	// ValueAddedTaxIncluded see : https://schema.org/valueAddedTaxIncluded
	// Specifies whether the applicable value-added tax (VAT) is included in the price specification or not.
	// types : Boolean
	ValueAddedTaxIncluded []bool `json:"valueAddedTaxIncluded,omitempty"`
}

func (v PriceSpecification) MarshalJSON() ([]byte, error) {
	type Alias PriceSpecification

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"PriceSpecification\","), b[1:]...), nil
}
