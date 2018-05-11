package schemaorg

import "encoding/json"

// DownloadAction see : https://schema.org/DownloadAction
type DownloadAction struct {

typeContext

TransferAction

// FromLocation see : https://schema.org/fromLocation
// A sub property of location. The original location of the object or the agent before the action.
FromLocation *Place `json:"fromLocation"`

// ToLocation see : https://schema.org/toLocation
// A sub property of location. The final location of the object or the agent after the action.
ToLocation *Place `json:"toLocation"`

}

func (v *DownloadAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DownloadAction"

	return json.Marshal(v)
}
