package schemaorggo

import "encoding/json"

// DataFeedItem see : https://schema.org/DataFeedItem
type DataFeedItem struct {
	Intangible

	typeContext

	// DateCreated see : https://schema.org/dateCreated
	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	// types : Date DateTime
	DateCreated interface{} `json:"dateCreated,omitempty"`

	// DateDeleted see : https://schema.org/dateDeleted
	// The datetime the item was removed from the DataFeed.
	// types : DateTime
	DateDeleted DateTime `json:"dateDeleted,omitempty"`

	// DateModified see : https://schema.org/dateModified
	// The date on which the CreativeWork was most recently modified or when the item&#39;s entry was modified within a DataFeed.
	// types : Date DateTime
	DateModified interface{} `json:"dateModified,omitempty"`

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	// types : Thing
	Item *Thing `json:"item,omitempty"`
}

func (v DataFeedItem) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataFeedItem"

	return json.Marshal(v)
}

func (v *DataFeedItem) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DataFeedItem"

	return json.Marshal(*v)
}
