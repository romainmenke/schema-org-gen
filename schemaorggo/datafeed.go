package schemaorggo

import "encoding/json"

// DataFeed see : https://schema.org/DataFeed
type DataFeed struct {
	Dataset

	typeContext

	// DataFeedElement see : https://schema.org/dataFeedElement
	// An item within in a data feed. Data feeds may have many elements.
	// types : DataFeedItem Text Thing
	DataFeedElement []interface{} `json:"dataFeedElement,omitempty"`
}

func (v DataFeed) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Dataset.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.DataFeedElement
		if len(v.DataFeedElement) == 1 {
			value = v.DataFeedElement[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dataFeedElement"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v DataFeed) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DataFeed"

	return data, nil
}

func (v DataFeed) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
