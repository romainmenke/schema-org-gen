package schemaorggo

import "encoding/json"

// DataFeedItem see : https://schema.org/DataFeedItem
type DataFeedItem struct {
	Intangible

	typeContext

	// DateCreated see : https://schema.org/dateCreated
	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	DateCreated interface{} `json:"dateCreated,omitempty"` // types : Date DateTime

	// DateDeleted see : https://schema.org/dateDeleted
	// The datetime the item was removed from the DataFeed.
	DateDeleted DateTime `json:"dateDeleted,omitempty"` // types : DateTime

	// DateModified see : https://schema.org/dateModified
	// The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
	DateModified interface{} `json:"dateModified,omitempty"` // types : Date DateTime

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
	Item *Thing `json:"item,omitempty"` // types : Thing

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
