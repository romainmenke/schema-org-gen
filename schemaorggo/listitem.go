package schemaorggo

import "encoding/json"

// ListItem see : https://schema.org/ListItem
type ListItem struct {
	Intangible

	typeContext

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
	Item *Thing `json:"item,omitempty"` // types : Thing

	// NextItem see : https://schema.org/nextItem
	// A link to the ListItem that follows the current one.
	NextItem *ListItem `json:"nextItem,omitempty"` // types : ListItem

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	Position interface{} `json:"position,omitempty"` // types : Integer Text

	// PreviousItem see : https://schema.org/previousItem
	// A link to the ListItem that preceeds the current one.
	PreviousItem *ListItem `json:"previousItem,omitempty"` // types : ListItem

}

func (v ListItem) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ListItem"

	return json.Marshal(v)
}

func (v *ListItem) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ListItem"

	return json.Marshal(*v)
}
