package schemaorggo

import "encoding/json"

// MenuSection see : https://schema.org/MenuSection
type MenuSection struct {
	CreativeWork

	typeContext

	// HasMenuItem see : https://schema.org/hasMenuItem
	// A food or drink item contained in a menu or menu section.
	HasMenuItem *MenuItem `json:"hasMenuItem"`

	// HasMenuSection see : https://schema.org/hasMenuSection
	// A subgrouping of the menu (by dishes, course, serving time period, etc.).
	HasMenuSection *MenuSection `json:"hasMenuSection"`
}

func (v *MenuSection) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MenuSection"

	return json.Marshal(v)
}
