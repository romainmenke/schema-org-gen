package schemaorggo

import "encoding/json"

// DataFeedItem see : https://schema.org/DataFeedItem
type DataFeedItem struct {
	Intangible

	typeContext

	// DateCreated see : https://schema.org/dateCreated
	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	// types : Date DateTime
	DateCreated []interface{} `json:"dateCreated,omitempty"`

	// DateDeleted see : https://schema.org/dateDeleted
	// The datetime the item was removed from the DataFeed.
	// types : DateTime
	DateDeleted []DateTime `json:"dateDeleted,omitempty"`

	// DateModified see : https://schema.org/dateModified
	// The date on which the CreativeWork was most recently modified or when the item&#39;s entry was modified within a DataFeed.
	// types : Date DateTime
	DateModified []interface{} `json:"dateModified,omitempty"`

	// Item see : https://schema.org/item
	// An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	// types : Thing
	Item []*Thing `json:"item,omitempty"`
}

func (v DataFeedItem) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.DateCreated
		if len(v.DateCreated) == 1 {
			value = v.DateCreated[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateCreated"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateDeleted
		if len(v.DateDeleted) == 1 {
			value = v.DateDeleted[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateDeleted"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateModified
		if len(v.DateModified) == 1 {
			value = v.DateModified[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateModified"] = json.RawMessage(b)
		}
	}

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

	*intop = into

	return nil
}

func (v DataFeedItem) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "DataFeedItem"

	return data, nil
}

func (v DataFeedItem) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
