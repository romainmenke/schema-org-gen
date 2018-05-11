package schemaorg

import "encoding/json"

// DataFeedItem see : https://schema.org/DataFeedItem
type DataFeedItem struct {

typeContext

Intangible

// DateCreated see : /dateCreated
// The date on which the CreativeWork was created or the item was added to a DataFeed.
DateCreated interface{} `json:"dateCreated"` // types : Date DateTime

// DateDeleted see : /dateDeleted
// The datetime the item was removed from the DataFeed.
DateDeleted interface{} `json:"dateDeleted"`

// DateModified see : /dateModified
// The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
DateModified interface{} `json:"dateModified"` // types : Date DateTime

// Item see : /item
// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
Item *Thing `json:"item"`

}

func (v *DataFeedItem) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DataFeedItem"

	return json.Marshal(v)
}
