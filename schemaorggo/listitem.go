package schemaorggo

import "encoding/json"

// ListItem see : https://schema.org/ListItem
type ListItem struct {
	Intangible

	typeContext

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	// types : Thing
	Item *Thing `json:"item,omitempty"`

	// NextItem see : https://schema.org/nextItem
	// A link to the ListItem that follows the current one.
	// types : ListItem
	NextItem *ListItem `json:"nextItem,omitempty"`

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	// types : Integer Text
	Position interface{} `json:"position,omitempty"`

	// PreviousItem see : https://schema.org/previousItem
	// A link to the ListItem that preceeds the current one.
	// types : ListItem
	PreviousItem *ListItem `json:"previousItem,omitempty"`
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
