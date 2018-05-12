package schemaorggo

import "encoding/json"

// PropertyValue see : https://schema.org/PropertyValue
type PropertyValue struct {
	StructuredValue

	typeContext

	// MaxValue see : https://schema.org/maxValue
	// The upper value of some characteristic or property.
	MaxValue float64 `json:"maxValue"`

	// MeasurementTechnique see : http://pending.schema.org/measurementTechnique
	// A technique or technology used in a Dataset (see: https://schema.org/Dataset) (or DataDownload (see: https://schema.org/DataDownload), DataCatalog (see: https://schema.org/DataCatalog)),
	// corresponding to the method used for measuring the corresponding variable(s) (described using variableMeasured (see: https://schema.org/variableMeasured)). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.
	//
	// For example, if variableMeasured (see: https://schema.org/variableMeasured) is: molecule concentration, measurementTechnique (see: https://schema.org/measurementTechnique) could be: "mass spectrometry" or "nmr spectroscopy" or "colorimetry" or "immunofluorescence".
	//
	// If the variableMeasured (see: https://schema.org/variableMeasured) is "depression rating", the measurementTechnique (see: https://schema.org/measurementTechnique) could be "Zung Scale" or "HAM-D" or "Beck Depression Inventory".
	//
	// If there are several variableMeasured (see: https://schema.org/variableMeasured) properties recorded for some given data object, use a PropertyValue (see: https://schema.org/PropertyValue) for each variableMeasured (see: https://schema.org/variableMeasured) and attach the corresponding measurementTechnique (see: https://schema.org/measurementTechnique).
	MeasurementTechnique interface{} `json:"measurementTechnique"` // types : Text URL

	// MinValue see : https://schema.org/minValue
	// The lower value of some characteristic or property.
	MinValue float64 `json:"minValue"`

	// PropertyID see : https://schema.org/propertyID
	// A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property. propertyID can be
	// (1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific id of the property), or (3)
	// a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
	// Standards bodies should promote a standard prefix for the identifiers of properties from their standards.
	PropertyID interface{} `json:"propertyID"` // types : Text URL

	// UnitCode see : https://schema.org/unitCode
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	UnitCode interface{} `json:"unitCode"` // types : Text URL

	// UnitText see : https://schema.org/unitText
	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	// unitCode (see: https://schema.orgunitCode).
	UnitText string `json:"unitText"`

	// Value see : https://schema.org/value
	// The value of the quantitative value or property value node.
	//
	//
	// For QuantitativeValue (see: https://schema.org/QuantitativeValue) and MonetaryAmount (see: https://schema.org/MonetaryAmount), the recommended type for values is 'Number'.
	// For PropertyValue (see: https://schema.org/PropertyValue), it can be 'Text;', 'Number', 'Boolean', or 'StructuredValue'.
	//
	//
	Value interface{} `json:"value"` // types : Boolean Number StructuredValue Text

	// ValueReference see : https://schema.org/valueReference
	// A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	ValueReference interface{} `json:"valueReference"` // types : Enumeration PropertyValue QualitativeValue QuantitativeValue StructuredValue

}

func (v *PropertyValue) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PropertyValue"

	return json.Marshal(v)
}
