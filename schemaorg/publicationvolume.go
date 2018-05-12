package schemaorg

import "encoding/json"

// PublicationVolume see : https://schema.org/PublicationVolume
type PublicationVolume struct {

CreativeWork

typeContext

// PageEnd see : https://schema.org/pageEnd
// The page on which the work ends; for example "138" or "xvi".
PageEnd interface{} `json:"pageEnd"` // types : Integer Text

// PageStart see : https://schema.org/pageStart
// The page on which the work starts; for example "135" or "xiii".
PageStart interface{} `json:"pageStart"` // types : Integer Text

// Pagination see : https://schema.org/pagination
// Any description of pages that is not separated into pageStart and pageEnd; for example, "1-6, 9, 55" or "10-12, 46-49".
Pagination string `json:"pagination"`

// VolumeNumber see : https://schema.org/volumeNumber
// Identifies the volume of publication or multi-part work; for example, "iii" or "2".
VolumeNumber interface{} `json:"volumeNumber"` // types : Integer Text

}

func (v *PublicationVolume) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PublicationVolume"

	return json.Marshal(v)
}
