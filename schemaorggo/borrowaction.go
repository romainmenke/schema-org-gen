package schemaorggo

import "encoding/json"

// BorrowAction see : https://schema.org/BorrowAction
type BorrowAction struct {
	TransferAction

	typeContext

	// Lender see : https://schema.org/lender
	// A sub property of participant. The person that lends the object being borrowed.
	// types : Organization Person
	Lender interface{} `json:"lender,omitempty"`
}

func (v BorrowAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BorrowAction"

	return json.Marshal(v)
}

func (v *BorrowAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BorrowAction"

	return json.Marshal(*v)
}
