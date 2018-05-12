package schemaorggo

import "encoding/json"

// HowTo see : https://schema.org/HowTo
type HowTo struct {
	CreativeWork

	typeContext

	// EstimatedCost see : https://schema.org/estimatedCost
	// The estimated cost of the supply or supplies consumed when performing instructions.
	// types : MonetaryAmount Text
	EstimatedCost interface{} `json:"estimatedCost,omitempty"`

	// PerformTime see : https://schema.org/performTime
	// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PerformTime *Duration `json:"performTime,omitempty"`

	// PrepTime see : https://schema.org/prepTime
	// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PrepTime *Duration `json:"prepTime,omitempty"`

	// Steps see : https://schema.org/steps
	// The steps in the form of a single item (text, document, video, etc.) or an ordered list with HowToStep and/or HowToSection items.
	// types : CreativeWork ItemList Text
	Steps interface{} `json:"steps,omitempty"`

	// Supply see : https://schema.org/supply
	// A sub-property of instrument. A supply consumed when performing instructions or a direction.
	// types : HowToSupply Text
	Supply interface{} `json:"supply,omitempty"`

	// Tool see : https://schema.org/tool
	// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
	// types : HowToTool Text
	Tool interface{} `json:"tool,omitempty"`

	// TotalTime see : https://schema.org/totalTime
	// The total time required to perform instructions or a direction (including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	TotalTime *Duration `json:"totalTime,omitempty"`

	// Yield see : https://schema.org/yield
	// The quantity that results by performing instructions. For example, a paper airplane, 10 personalized candles.
	// types : QuantitativeValue Text
	Yield interface{} `json:"yield,omitempty"`
}

func (v HowTo) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowTo"

	return json.Marshal(v)
}

func (v *HowTo) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HowTo"

	return json.Marshal(*v)
}
