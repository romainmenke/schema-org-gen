package schemaorg

import "encoding/json"

// LendAction see : https://schema.org/LendAction
type LendAction struct {

typeContext

TransferAction

// Borrower see : https://schema.org/borrower
// A sub property of participant. The person that borrows the object being lent.
Borrower *Person `json:"borrower"`

}

func (v *LendAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LendAction"

	return json.Marshal(v)
}
