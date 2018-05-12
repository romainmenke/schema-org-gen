package schemaorggo

import "encoding/json"

// DataDownload see : https://schema.org/DataDownload
type DataDownload struct {
	MediaObject

	typeContext

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
}

func (v DataDownload) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataDownload"

	return json.Marshal(v)
}

func (v *DataDownload) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DataDownload"

	return json.Marshal(*v)
}
