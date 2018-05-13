package schemaorggo

import "encoding/json"

// HowToTip see : https://schema.org/HowToTip
type HowToTip struct {
	ListItem

	typeContext

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	// types : Thing
	Item []*Thing `json:"item,omitempty"`

	// NextItem see : https://schema.org/nextItem
	// A link to the ListItem that follows the current one.
	// types : ListItem
	NextItem []*ListItem `json:"nextItem,omitempty"`

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	// types : Integer Text
	Position []interface{} `json:"position,omitempty"`

	// PreviousItem see : https://schema.org/previousItem
	// A link to the ListItem that preceeds the current one.
	// types : ListItem
	PreviousItem []*ListItem `json:"previousItem,omitempty"`
}

func (v HowToTip) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.ListItem.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Item
		if len(v.Item) == 1 {
			value = v.Item[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["item"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NextItem
		if len(v.NextItem) == 1 {
			value = v.NextItem[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["nextItem"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Position
		if len(v.Position) == 1 {
			value = v.Position[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["position"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PreviousItem
		if len(v.PreviousItem) == 1 {
			value = v.PreviousItem[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["previousItem"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v HowToTip) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "HowToTip"

	return data, nil
}

func (v HowToTip) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
