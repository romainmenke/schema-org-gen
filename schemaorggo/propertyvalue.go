package schemaorggo

import "encoding/json"

// PropertyValue see : https://schema.org/PropertyValue
type PropertyValue struct {
	StructuredValue

	typeContext

	// MaxValue see : https://schema.org/maxValue
	// The upper value of some characteristic or property.
	// types : Number
	MaxValue float64 `json:"maxValue,omitempty"`

	// MeasurementTechnique see : http://pending.schema.org/measurementTechnique
	// A technique or technology used in a Dataset (see: https://schema.org/Dataset) (or DataDownload (see: https://schema.org/DataDownload), DataCatalog (see: https://schema.org/DataCatalog)),
	// corresponding to the method used for measuring the corresponding variable(s) (described using variableMeasured (see: https://schema.org/variableMeasured)). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.
	//
	// For example, if variableMeasured (see: https://schema.org/variableMeasured) is: molecule concentration, measurementTechnique (see: https://schema.org/measurementTechnique) could be: &quot;mass spectrometry&quot; or &quot;nmr spectroscopy&quot; or &quot;colorimetry&quot; or &quot;immunofluorescence&quot;.
	//
	// If the variableMeasured (see: https://schema.org/variableMeasured) is &quot;depression rating&quot;, the measurementTechnique (see: https://schema.org/measurementTechnique) could be &quot;Zung Scale&quot; or &quot;HAM-D&quot; or &quot;Beck Depression Inventory&quot;.
	//
	// If there are several variableMeasured (see: https://schema.org/variableMeasured) properties recorded for some given data object, use a PropertyValue (see: https://schema.org/PropertyValue) for each variableMeasured (see: https://schema.org/variableMeasured) and attach the corresponding measurementTechnique (see: https://schema.org/measurementTechnique).
	// types : Text URL
	MeasurementTechnique string `json:"measurementTechnique,omitempty"`

	// MinValue see : https://schema.org/minValue
	// The lower value of some characteristic or property.
	// types : Number
	MinValue float64 `json:"minValue,omitempty"`

	// PropertyID see : https://schema.org/propertyID
	// A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be
	// (1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific id of the property), or (3)
	// a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
	// Standards bodies should promote a standard prefix for the identifiers of properties from their standards.
	// types : Text URL
	PropertyID string `json:"propertyID,omitempty"`

	// UnitCode see : https://schema.org/unitCode
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	// types : Text URL
	UnitCode string `json:"unitCode,omitempty"`

	// UnitText see : https://schema.org/unitText
	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	// unitCode (see: https://schema.orgunitCode).
	// types : Text
	UnitText string `json:"unitText,omitempty"`

	// Value see : https://schema.org/value
	// The value of the quantitative value or property value node.
	//
	//
	// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is &#39;Number&#39;.
	// For PropertyValue (see: https://schema.org/PropertyValue), it can be &#39;Text;&#39;, &#39;Number&#39;, &#39;Boolean&#39;, or &#39;StructuredValue&#39;.
	//
	//
	// types : Boolean Number StructuredValue Text
	Value interface{} `json:"value,omitempty"`

	// ValueReference see : https://schema.org/valueReference
	// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	// types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue
	ValueReference interface{} `json:"valueReference,omitempty"`
}

func (v PropertyValue) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PropertyValue"

	return json.Marshal(v)
}

func (v *PropertyValue) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PropertyValue"

	return json.Marshal(*v)
}
