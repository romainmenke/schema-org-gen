package schemaorggo

import "encoding/json"

// MovieTheater see : https://schema.org/MovieTheater
type MovieTheater struct {
	CivicStructure

	typeContext

	// ScreenCount see : https://schema.org/screenCount
	// The number of screens in the movie theater.
	ScreenCount float64 `json:"screenCount,omitempty"` // types : Number

}

func (v MovieTheater) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MovieTheater"

	return json.Marshal(v)
}

func (v *MovieTheater) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MovieTheater"

	return json.Marshal(*v)
}
