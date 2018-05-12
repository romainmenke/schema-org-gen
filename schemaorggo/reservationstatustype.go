package schemaorggo

import "encoding/json"

// ReservationStatusType see : https://schema.org/ReservationStatusType
type ReservationStatusType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy,omitempty"` // types : Class Enumeration Property

}

func (v ReservationStatusType) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReservationStatusType"

	return json.Marshal(v)
}

func (v *ReservationStatusType) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ReservationStatusType"

	return json.Marshal(*v)
}
