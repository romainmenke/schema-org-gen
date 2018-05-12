package schemaorggo

import "encoding/json"

// MenuSection see : https://schema.org/MenuSection
type MenuSection struct {
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

func (v MenuSection) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MenuSection"

	return json.Marshal(v)
}

func (v *MenuSection) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MenuSection"

	return json.Marshal(*v)
}
