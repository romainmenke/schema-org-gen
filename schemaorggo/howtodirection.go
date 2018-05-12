package schemaorggo

import "encoding/json"

// HowToDirection see : https://schema.org/HowToDirection
type HowToDirection struct {
	ListItem

	typeContext

	// AfterMedia see : https://schema.org/afterMedia
	// A media object representing the circumstances after performing this direction.
	// types : MediaObject
	AfterMedia *MediaObject `json:"afterMedia,omitempty"`

	// BeforeMedia see : https://schema.org/beforeMedia
	// A media object representing the circumstances before performing this direction.
	// types : MediaObject
	BeforeMedia *MediaObject `json:"beforeMedia,omitempty"`

	// DuringMedia see : https://schema.org/duringMedia
	// A media object representing the circumstances while performing this direction.
	// types : MediaObject
	DuringMedia *MediaObject `json:"duringMedia,omitempty"`

	// PerformTime see : https://schema.org/performTime
	// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PerformTime *Duration `json:"performTime,omitempty"`

	// PrepTime see : https://schema.org/prepTime
	// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	// types : Duration
	PrepTime *Duration `json:"prepTime,omitempty"`

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
}

func (v HowToDirection) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToDirection"

	return json.Marshal(v)
}

func (v *HowToDirection) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "HowToDirection"

	return json.Marshal(*v)
}
