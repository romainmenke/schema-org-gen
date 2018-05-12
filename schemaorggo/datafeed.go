package schemaorggo

import "encoding/json"

// DataFeed see : https://schema.org/DataFeed
type DataFeed struct {
	Dataset

	typeContext

	// DataFeedElement see : https://schema.org/dataFeedElement
	// An item within in a data feed. Data feeds may have many elements.
	DataFeedElement interface{} `json:"dataFeedElement,omitempty"` // types : DataFeedItem Text Thing

}

func (v DataFeed) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataFeed"

	return json.Marshal(v)
}

func (v *DataFeed) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DataFeed"

	return json.Marshal(*v)
}
