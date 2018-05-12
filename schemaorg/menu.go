package schemaorg

import "encoding/json"

// Menu see : https://schema.org/Menu
type Menu struct {

CreativeWork

typeContext

// HasMenuItem see : https://schema.org/hasMenuItem
// A food or drink item contained in a menu or menu section.
HasMenuItem *MenuItem `json:"hasMenuItem"`

// HasMenuSection see : https://schema.org/hasMenuSection
// A subgrouping of the menu (by dishes, course, serving time period, etc.).
HasMenuSection *MenuSection `json:"hasMenuSection"`

}

func (v *Menu) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Menu"

	return json.Marshal(v)
}
