package schemaorggo

import "encoding/json"

// Report see : https://schema.org/Report
type Report struct {
	Article

	typeContext

	// ReportNumber see : https://schema.org/reportNumber
	// The number or other unique designator assigned to a Report by the publishing organization.
	// types : Text
	ReportNumber []string `json:"reportNumber,omitempty"`
}

func (v Report) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Article.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ReportNumber
		if len(v.ReportNumber) == 1 {
			value = v.ReportNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reportNumber"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Report) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Report"

	return data, nil
}

func (v Report) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
