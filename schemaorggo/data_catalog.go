package schemaorggo

import "encoding/json"

// DataCatalog see : https://schema.org/DataCatalog
type DataCatalog struct {
	CreativeWork

	typeContext

	// Dataset see : https://schema.org/dataset
	// A dataset contained in this catalog. Inverse property: includedInDataCatalog (see: https://schema.org/includedInDataCatalog).
	// types : Dataset
	Dataset []*Dataset `json:"dataset,omitempty"`

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
}

func (v DataCatalog) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Dataset
		if len(v.Dataset) == 1 {
			value = v.Dataset[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dataset"] = json.RawMessage(b)
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

	*intop = into

	return nil
}

func (v DataCatalog) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DataCatalog"

	return data, nil
}

func (v DataCatalog) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}