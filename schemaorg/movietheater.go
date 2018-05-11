package schemaorg

import "encoding/json"

// MovieTheater see : https://schema.org/MovieTheater
type MovieTheater struct {

typeContext

CivicStructure

// ScreenCount see : /screenCount
// The number of screens in the movie theater.
ScreenCount float64 `json:"screenCount"`

}

func (v *MovieTheater) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MovieTheater"

	return json.Marshal(v)
}
