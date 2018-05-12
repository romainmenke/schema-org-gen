package schemaorggo

import "encoding/json"

// Menu see : https://schema.org/Menu
type Menu struct {
	CreativeWork

	typeContext

	// HasMenuItem see : https://schema.org/hasMenuItem
	// A food or drink item contained in a menu or menu section.
	// types : MenuItem
	HasMenuItem *MenuItem `json:"hasMenuItem,omitempty"`

	// HasMenuSection see : https://schema.org/hasMenuSection
	// A subgrouping of the menu (by dishes, course, serving time period, etc.).
	// types : MenuSection
	HasMenuSection *MenuSection `json:"hasMenuSection,omitempty"`
}

func (v Menu) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Menu"

	return json.Marshal(v)
}

func (v *Menu) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Menu"

	return json.Marshal(*v)
}
