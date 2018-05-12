package schemaorg

import "encoding/json"

// MovieTheater see : https://schema.org/MovieTheater
type MovieTheater struct {

CivicStructure

typeContext

// ScreenCount see : https://schema.org/screenCount
// The number of screens in the movie theater.
ScreenCount float64 `json:"screenCount"`

}

func (v *MovieTheater) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MovieTheater"

	return json.Marshal(v)
}
