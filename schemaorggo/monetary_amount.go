package schemaorggo

import "encoding/json"

// MonetaryAmount see : https://schema.org/MonetaryAmount
type MonetaryAmount struct {
	typeContext

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from StructuredValue see : https://schema.org/StructuredValue
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

	// Currency see : https://pending.schema.org/currency
	// The currency in which the monetary amount is expressed.
	//
	// Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	// types : Text
	Currency []string `json:"currency,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	// types : PropertyValue Text URL
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	// types : ImageObject URL
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MaxValue see : https://schema.org/maxValue
	// The upper value of some characteristic or property.
	// types : Number
	MaxValue []float64 `json:"maxValue,omitempty"`

	// MinValue see : https://schema.org/minValue
	// The lower value of some characteristic or property.
	// types : Number
	MinValue []float64 `json:"minValue,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// ValidFrom see : https://schema.org/validFrom
	// The date when the item becomes valid.
	// types : DateTime
	ValidFrom []DateTime `json:"validFrom,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`

	// Value see : https://schema.org/value
	// The value of the quantitative value or property value node.
	//
	//
	// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is &#39;Number&#39;.
	// For PropertyValue (see: https://schema.org/PropertyValue), it can be &#39;Text;&#39;, &#39;Number&#39;, &#39;Boolean&#39;, or &#39;StructuredValue&#39;.
	//
	//
	// types : Boolean Number StructuredValue Text
	Value []interface{} `json:"value,omitempty"`
}

func (v MonetaryAmount) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.AdditionalType
		if len(v.AdditionalType) == 1 {
			value = v.AdditionalType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AlternateName
		if len(v.AlternateName) == 1 {
			value = v.AlternateName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alternateName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Currency
		if len(v.Currency) == 1 {
			value = v.Currency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["currency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Description
		if len(v.Description) == 1 {
			value = v.Description[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["description"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DisambiguatingDescription
		if len(v.DisambiguatingDescription) == 1 {
			value = v.DisambiguatingDescription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["disambiguatingDescription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Identifier
		if len(v.Identifier) == 1 {
			value = v.Identifier[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["identifier"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Image
		if len(v.Image) == 1 {
			value = v.Image[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["image"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MainEntityOfPage
		if len(v.MainEntityOfPage) == 1 {
			value = v.MainEntityOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mainEntityOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MaxValue
		if len(v.MaxValue) == 1 {
			value = v.MaxValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["maxValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MinValue
		if len(v.MinValue) == 1 {
			value = v.MinValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["minValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Name
		if len(v.Name) == 1 {
			value = v.Name[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["name"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PotentialAction
		if len(v.PotentialAction) == 1 {
			value = v.PotentialAction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["potentialAction"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SameAs
		if len(v.SameAs) == 1 {
			value = v.SameAs[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sameAs"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SubjectOf
		if len(v.SubjectOf) == 1 {
			value = v.SubjectOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subjectOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Url
		if len(v.Url) == 1 {
			value = v.Url[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["url"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidFrom
		if len(v.ValidFrom) == 1 {
			value = v.ValidFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidThrough
		if len(v.ValidThrough) == 1 {
			value = v.ValidThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validThrough"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Value
		if len(v.Value) == 1 {
			value = v.Value[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["value"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v MonetaryAmount) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "MonetaryAmount"

	return data, nil
}

func (v MonetaryAmount) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
