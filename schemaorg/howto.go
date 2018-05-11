package schemaorg

import "encoding/json"

// HowTo see : https://schema.org/HowTo
type HowTo struct {

typeContext

CreativeWork

// EstimatedCost see : /estimatedCost
// The estimated cost of the supply or supplies consumed when performing instructions.
EstimatedCost interface{} `json:"estimatedCost"` // types : MonetaryAmount Text

// PerformTime see : /performTime
// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
PerformTime *Duration `json:"performTime"`

// PrepTime see : /prepTime
// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
PrepTime *Duration `json:"prepTime"`

// Steps see : /steps
// The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
Steps interface{} `json:"steps"` // types : CreativeWork ItemList Text

// Supply see : /supply
// A sub-property of instrument. A supply consumed when performing instructions or a direction.
Supply interface{} `json:"supply"` // types : HowToSupply Text

// Tool see : /tool
// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
Tool interface{} `json:"tool"` // types : HowToTool Text

// TotalTime see : /totalTime
// The total time required to perform instructions or a direction (including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
TotalTime *Duration `json:"totalTime"`

// Yield see : /yield
// The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles.
Yield interface{} `json:"yield"` // types : QuantitativeValue Text

}

func (v *HowTo) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowTo"

	return json.Marshal(v)
}
