package schemaorg

import "encoding/json"

// BorrowAction see : https://schema.org/BorrowAction
type BorrowAction struct {

TransferAction

typeContext

// Lender see : https://schema.org/lender
// A sub property of participant. The person that lends the object being borrowed.
Lender interface{} `json:"lender"` // types : Organization Person

}

func (v *BorrowAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BorrowAction"

	return json.Marshal(v)
}
