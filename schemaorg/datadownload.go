package schemaorg

import "encoding/json"

// DataDownload see : https://schema.org/DataDownload
type DataDownload struct {

typeContext

MediaObject

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

}

func (v *DataDownload) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataDownload"

	return json.Marshal(v)
}
