package schemaorg

import "encoding/json"

// DataFeed see : https://schema.org/DataFeed
type DataFeed struct {

Dataset

typeContext

// DataFeedElement see : https://schema.org/dataFeedElement
// An item within in a data feed. Data feeds may have many elements.
DataFeedElement interface{} `json:"dataFeedElement"` // types : DataFeedItem Text Thing

}

func (v *DataFeed) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataFeed"

	return json.Marshal(v)
}
