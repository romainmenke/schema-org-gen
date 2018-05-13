package schemaorggo

import "encoding/json"

// ItemList see : https://schema.org/ItemList
type ItemList struct {
	Intangible

	typeContext

	// ItemListElement see : https://schema.org/itemListElement
	// For itemListElement values, you can use simple strings (e.g. &quot;Peter&quot;, &quot;Paul&quot;, &quot;Mary&quot;), existing entities, or use ListItem.
	//
	// Text values are best if the elements in the list are plain strings. Existing entities are best for a simple, unordered list of existing things in your data. ListItem is used with ordered lists when you want to provide additional context about the element in that list or when the same item might be in different places in different lists.
	//
	// Note: The order of elements in your mark-up is not sufficient for indicating the order or elements.  Use ListItem with a &#39;position&#39; property in such cases.
	// types : ListItem Text Thing
	ItemListElement []interface{} `json:"itemListElement,omitempty"`

	// ItemListOrder see : https://schema.org/itemListOrder
	// Type of ordering (e.g. Ascending, Descending, Unordered).
	// types : ItemListOrderType Text
	ItemListOrder []interface{} `json:"itemListOrder,omitempty"`

	// NumberOfItems see : https://schema.org/numberOfItems
	// The number of items in an ItemList. Note that some descriptions might not fully describe all items in a list (e.g., multi-page pagination); in such cases, the numberOfItems would be for the entire list.
	// types : Integer
	NumberOfItems []float64 `json:"numberOfItems,omitempty"`
}

func (v ItemList) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ItemListElement
		if len(v.ItemListElement) == 1 {
			value = v.ItemListElement[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itemListElement"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ItemListOrder
		if len(v.ItemListOrder) == 1 {
			value = v.ItemListOrder[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itemListOrder"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfItems
		if len(v.NumberOfItems) == 1 {
			value = v.NumberOfItems[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfItems"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ItemList) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ItemList"

	return data, nil
}

func (v ItemList) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
