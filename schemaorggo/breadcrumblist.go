package schemaorggo

import "encoding/json"

// BreadcrumbList see : https://schema.org/BreadcrumbList
type BreadcrumbList struct {
	ItemList

	typeContext

	// ItemListElement see : https://schema.org/itemListElement
	// For itemListElement values, you can use simple strings (e.g. &quot;Peter&quot;, &quot;Paul&quot;, &quot;Mary&quot;), existing entities, or use ListItem.
	//
	// Text values are best if the elements in the list are plain strings. Existing entities are best for a simple, unordered list of existing things in your data. ListItem is used with ordered lists when you want to provide additional context about the element in that list or when the same item might be in different places in different lists.
	//
	// Note: The order of elements in your mark-up is not sufficient for indicating the order or elements.  Use ListItem with a &#39;position&#39; property in such cases.
	// types : ListItem Text Thing
	ItemListElement interface{} `json:"itemListElement,omitempty"`

	// ItemListOrder see : https://schema.org/itemListOrder
	// Type of ordering (e.g. Ascending, Descending, Unordered).
	// types : ItemListOrderType Text
	ItemListOrder interface{} `json:"itemListOrder,omitempty"`

	// NumberOfItems see : https://schema.org/numberOfItems
	// The number of items in an ItemList. Note that some descriptions might not fully describe all items in a list (e.g., multi-page pagination); in such cases, the numberOfItems would be for the entire list.
	// types : Integer
	NumberOfItems float64 `json:"numberOfItems,omitempty"`
}

func (v BreadcrumbList) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BreadcrumbList"

	return json.Marshal(v)
}

func (v *BreadcrumbList) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BreadcrumbList"

	return json.Marshal(*v)
}
