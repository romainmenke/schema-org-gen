package schemaorggo

import "encoding/json"

// HowToTip see : https://schema.org/HowToTip
type HowToTip struct {
	ListItem

	typeContext

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an 'artist' in a list of 'artists')’.
	Item *Thing `json:"item,omitempty"`

	// NextItem see : https://schema.org/nextItem
	// A link to the ListItem that follows the current one.
	NextItem *ListItem `json:"nextItem,omitempty"`

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	Position interface{} `json:"position,omitempty"` // types : Integer Text

	// PreviousItem see : https://schema.org/previousItem
	// A link to the ListItem that preceeds the current one.
	PreviousItem *ListItem `json:"previousItem,omitempty"`
}

func (v HowToTip) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToTip"

	return json.Marshal(v)
}

func (v *HowToTip) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HowToTip"

	return json.Marshal(*v)
}
