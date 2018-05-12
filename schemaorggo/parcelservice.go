package schemaorggo

import "encoding/json"

// ParcelService see : https://schema.org/ParcelService
type ParcelService struct {
	DeliveryMethod

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *ParcelService) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ParcelService"

	return json.Marshal(v)
}
