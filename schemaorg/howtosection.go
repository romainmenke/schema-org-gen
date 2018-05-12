package schemaorg

import "encoding/json"

// HowToSection see : https://schema.org/HowToSection
type HowToSection struct {

ItemList

typeContext

// Steps see : https://schema.org/steps
// The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
Steps interface{} `json:"steps"` // types : CreativeWork ItemList Text

}

func (v *HowToSection) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToSection"

	return json.Marshal(v)
}
