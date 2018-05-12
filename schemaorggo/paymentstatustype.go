package schemaorggo

import "encoding/json"

// PaymentStatusType see : https://schema.org/PaymentStatusType
type PaymentStatusType struct {
	Enumeration

	typeContext

	// SupersededBy see : http://meta.schema.org/supersededBy
	// Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	SupersededBy interface{} `json:"supersededBy"` // types : Class Enumeration Property

}

func (v *PaymentStatusType) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PaymentStatusType"

	return json.Marshal(v)
}
