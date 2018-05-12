package schemaorggo

import "encoding/json"

// Dataset see : https://schema.org/Dataset
type Dataset struct {
	CreativeWork

	typeContext

	// Distribution see : https://schema.org/distribution
	// A downloadable form of this dataset, at a specific location, in a specific format.
	Distribution *DataDownload `json:"distribution,omitempty"`

	// IncludedInDataCatalog see : https://schema.org/includedInDataCatalog
	// A data catalog which contains this dataset. Supersedes catalog (see: https://schema.org/catalog), includedDataCatalog (see: https://schema.org/includedDataCatalog). Inverse property: dataset (see: https://schema.org/dataset).
	IncludedInDataCatalog *DataCatalog `json:"includedInDataCatalog,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	Issn string `json:"issn,omitempty"`

	// MeasurementTechnique see : http://pending.schema.org/measurementTechnique
	// A technique or technology used in a Dataset (see: https://schema.org/Dataset) (or DataDownload (see: https://schema.org/DataDownload), DataCatalog (see: https://schema.org/DataCatalog)),
	// corresponding to the method used for measuring the corresponding variable(s) (described using variableMeasured (see: https://schema.org/variableMeasured)). This is oriented towards scientific and scholarly dataset publication but may have broader applicability; it is not intended as a full representation of measurement, but rather as a high level summary for dataset discovery.
	//
	// For example, if variableMeasured (see: https://schema.org/variableMeasured) is: molecule concentration, measurementTechnique (see: https://schema.org/measurementTechnique) could be: "mass spectrometry" or "nmr spectroscopy" or "colorimetry" or "immunofluorescence".
	//
	// If the variableMeasured (see: https://schema.org/variableMeasured) is "depression rating", the measurementTechnique (see: https://schema.org/measurementTechnique) could be "Zung Scale" or "HAM-D" or "Beck Depression Inventory".
	//
	// If there are several variableMeasured (see: https://schema.org/variableMeasured) properties recorded for some given data object, use a PropertyValue (see: https://schema.org/PropertyValue) for each variableMeasured (see: https://schema.org/variableMeasured) and attach the corresponding measurementTechnique (see: https://schema.org/measurementTechnique).
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty"` // types : Text URL

	// VariableMeasured see : http://pending.schema.org/variableMeasured
	// The variableMeasured property can indicate (repeated as necessary) the  variables that are measured in some dataset, either described as text or as pairs of identifier and description using PropertyValue.
	VariableMeasured interface{} `json:"variableMeasured,omitempty"` // types : PropertyValue Text

}

func (v Dataset) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Dataset"

	return json.Marshal(v)
}

func (v *Dataset) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Dataset"

	return json.Marshal(*v)
}
