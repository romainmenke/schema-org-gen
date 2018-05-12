package schemaorggo

import "encoding/json"

// DataCatalog see : https://schema.org/DataCatalog
type DataCatalog struct {
	CreativeWork

	typeContext

	// Dataset see : https://schema.org/dataset
	// A dataset contained in this catalog. Inverse property: includedInDataCatalog (see: https://schema.org/includedInDataCatalog).
	Dataset *Dataset `json:"dataset,omitempty"` // types : Dataset

	// MeasurementTechnique see : http://pending.schema.org/measurementTechnique
	// A technique or technology used in a Dataset (see: https://schema.org/Dataset) (or DataDownload (see: https://schema.org/DataDownload), DataCatalog (see: https://schema.org/DataCatalog)),
	// corresponding to the method used for measuring the corresponding variable(s) (described using variableMeasured (see: https://schema.org/variableMeasured)). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.
	//
	// For example, if variableMeasured (see: https://schema.org/variableMeasured) is: molecule concentration, measurementTechnique (see: https://schema.org/measurementTechnique) could be: "mass spectrometry" or "nmr spectroscopy" or "colorimetry" or "immunofluorescence".
	//
	// If the variableMeasured (see: https://schema.org/variableMeasured) is "depression rating", the measurementTechnique (see: https://schema.org/measurementTechnique) could be "Zung Scale" or "HAM-D" or "Beck Depression Inventory".
	//
	// If there are several variableMeasured (see: https://schema.org/variableMeasured) properties recorded for some given data object, use a PropertyValue (see: https://schema.org/PropertyValue) for each variableMeasured (see: https://schema.org/variableMeasured) and attach the corresponding measurementTechnique (see: https://schema.org/measurementTechnique).
	MeasurementTechnique string `json:"measurementTechnique,omitempty"` // types : Text URL

}

func (v DataCatalog) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataCatalog"

	return json.Marshal(v)
}

func (v *DataCatalog) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DataCatalog"

	return json.Marshal(*v)
}
