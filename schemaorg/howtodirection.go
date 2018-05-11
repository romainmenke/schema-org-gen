package schemaorg

import "encoding/json"

// HowToDirection see : https://schema.org/HowToDirection
type HowToDirection struct {

typeContext

ListItem

// AfterMedia see : /afterMedia
// A media object representing the circumstances after performing this direction.
AfterMedia *MediaObject `json:"afterMedia"`

// BeforeMedia see : /beforeMedia
// A media object representing the circumstances before performing this direction.
BeforeMedia *MediaObject `json:"beforeMedia"`

// DuringMedia see : /duringMedia
// A media object representing the circumstances while performing this direction.
DuringMedia *MediaObject `json:"duringMedia"`

// PerformTime see : /performTime
// The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
PerformTime *Duration `json:"performTime"`

// PrepTime see : /prepTime
// The length of time it takes to prepare the items to be used in instructions or a direction, in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
PrepTime *Duration `json:"prepTime"`

// Supply see : /supply
// A sub-property of instrument. A supply consumed when performing instructions or a direction.
Supply interface{} `json:"supply"` // types : HowToSupply Text

// Tool see : /tool
// A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
Tool interface{} `json:"tool"` // types : HowToTool Text

// TotalTime see : /totalTime
// The total time required to perform instructions or a direction (including time to prepare the supplies), in ISO 8601 duration format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
TotalTime *Duration `json:"totalTime"`

}

func (v *HowToDirection) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "HowToDirection"

	return json.Marshal(v)
}
