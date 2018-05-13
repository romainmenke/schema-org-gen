package schemaorggo

import "encoding/json"

// Dataset see : https://schema.org/Dataset
type Dataset struct {
	CreativeWork

	typeContext

	// Distribution see : https://schema.org/distribution
	// A downloadable form of this dataset, at a specific location, in a specific format.
	// types : DataDownload
	Distribution []*DataDownload `json:"distribution,omitempty"`

	// IncludedInDataCatalog see : https://schema.org/includedInDataCatalog
	// A data catalog which contains this dataset. Supersedes catalog (see: https://schema.org/catalog), includedDataCatalog (see: https://schema.org/includedDataCatalog). Inverse property: dataset (see: https://schema.org/dataset).
	// types : DataCatalog
	IncludedInDataCatalog []*DataCatalog `json:"includedInDataCatalog,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn []string `json:"issn,omitempty"`

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
	MeasurementTechnique []string `json:"measurementTechnique,omitempty"`

	// VariableMeasured see : http://pending.schema.org/variableMeasured
	// The variableMeasured property can indicate (repeated as necessary) the  variables that are measured in some dataset, either described as text or as pairs of identifier and description using PropertyValue.
	// types : PropertyValue Text
	VariableMeasured []interface{} `json:"variableMeasured,omitempty"`
}

func (v Dataset) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Distribution
		if len(v.Distribution) == 1 {
			value = v.Distribution[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["distribution"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IncludedInDataCatalog
		if len(v.IncludedInDataCatalog) == 1 {
			value = v.IncludedInDataCatalog[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["includedInDataCatalog"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Issn
		if len(v.Issn) == 1 {
			value = v.Issn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["issn"] = json.RawMessage(b)
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
		var value interface{} = v.VariableMeasured
		if len(v.VariableMeasured) == 1 {
			value = v.VariableMeasured[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["variableMeasured"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Dataset) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Dataset"

	return data, nil
}

func (v Dataset) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
