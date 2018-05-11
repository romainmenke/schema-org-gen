package schemaorg

import "encoding/json"

// MenuSection see : https://schema.org/MenuSection
type MenuSection struct {

typeContext

CreativeWork

// HasMenuItem see : /hasMenuItem
// A food or drink item contained in a menu or menu section.
HasMenuItem *MenuItem `json:"hasMenuItem"`

// HasMenuSection see : /hasMenuSection
// A subgrouping of the menu (by dishes, course, serving time period, etc.).
HasMenuSection *MenuSection `json:"hasMenuSection"`

}

func (v *MenuSection) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MenuSection"

	return json.Marshal(v)
}
