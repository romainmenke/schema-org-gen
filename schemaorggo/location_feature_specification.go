package schemaorggo

import "encoding/json"

// LocationFeatureSpecification see : https://schema.org/LocationFeatureSpecification
type LocationFeatureSpecification struct {
	typeContext

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from PropertyValue see : https://schema.org/PropertyValue
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

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// HoursAvailable see : https://schema.org/hoursAvailable
	// The hours during which this service or contact is available.
	// types : OpeningHoursSpecification
	HoursAvailable []*OpeningHoursSpecification `json:"hoursAvailable,omitempty"`

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

	// MeasurementTechnique see : https://pending.schema.org/measurementTechnique
	// A technique or technology used in a Dataset (see: https://schema.org/Dataset) (or DataDownload (see: https://schema.org/DataDownload), DataCatalog (see: https://schema.org/DataCatalog)),
	// corresponding to the method used for measuring the corresponding variable(s) (described using variableMeasured (see: https://schema.org/variableMeasured)). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.
	//
	// For example, if variableMeasured (see: https://schema.org/variableMeasured) is: molecule concentration, measurementTechnique (see: https://schema.org/measurementTechnique) could be: &quot;mass spectrometry&quot; or &quot;nmr spectroscopy&quot; or &quot;colorimetry&quot; or &quot;immunofluorescence&quot;.
	//
	// If the variableMeasured (see: https://schema.org/variableMeasured) is &quot;depression rating&quot;, the measurementTechnique (see: https://schema.org/measurementTechnique) could be &quot;Zung Scale&quot; or &quot;HAM-D&quot; or &quot;Beck Depression Inventory&quot;.
	//
	// If there are several variableMeasured (see: https://schema.org/variableMeasured) properties recorded for some given data object, use a PropertyValue (see: https://schema.org/PropertyValue) for each variableMeasured (see: https://schema.org/variableMeasured) and attach the corresponding measurementTechnique (see: https://schema.org/measurementTechnique).
	// types : Text URL
	MeasurementTechnique []string `json:"measurementTechnique,omitempty"`

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

	// PropertyID see : https://schema.org/propertyID
	// A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be
	// (1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific id of the property), or (3)
	// a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
	// Standards bodies should promote a standard prefix for the identifiers of properties from their standards.
	// types : Text URL
	PropertyID []string `json:"propertyID,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// UnitCode see : https://schema.org/unitCode
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	// types : Text URL
	UnitCode []string `json:"unitCode,omitempty"`

	// UnitText see : https://schema.org/unitText
	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	// unitCode (see: https://schema.orgunitCode).
	// types : Text
	UnitText []string `json:"unitText,omitempty"`

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

	// ValueReference see : https://schema.org/valueReference
	// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	// types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue
	ValueReference []interface{} `json:"valueReference,omitempty"`
}

func (v LocationFeatureSpecification) intoMap(intop *map[string]interface{}) error {
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
		var value interface{} = v.HoursAvailable
		if len(v.HoursAvailable) == 1 {
			value = v.HoursAvailable[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hoursAvailable"] = json.RawMessage(b)
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
		var value interface{} = v.MeasurementTechnique
		if len(v.MeasurementTechnique) == 1 {
			value = v.MeasurementTechnique[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["measurementTechnique"] = json.RawMessage(b)
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
		var value interface{} = v.PropertyID
		if len(v.PropertyID) == 1 {
			value = v.PropertyID[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["propertyID"] = json.RawMessage(b)
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
		var value interface{} = v.UnitCode
		if len(v.UnitCode) == 1 {
			value = v.UnitCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unitCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UnitText
		if len(v.UnitText) == 1 {
			value = v.UnitText[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unitText"] = json.RawMessage(b)
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

	{
		var value interface{} = v.ValueReference
		if len(v.ValueReference) == 1 {
			value = v.ValueReference[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["valueReference"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v LocationFeatureSpecification) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "LocationFeatureSpecification"

	return data, nil
}

func (v LocationFeatureSpecification) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
