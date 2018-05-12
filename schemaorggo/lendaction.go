package schemaorggo

import "encoding/json"

// LendAction see : https://schema.org/LendAction
type LendAction struct {
	TransferAction

	typeContext

	// Borrower see : https://schema.org/borrower
	// A sub property of participant. The person that borrows the object being lent.
	// types : Person
	Borrower *Person `json:"borrower,omitempty"`
}

func (v LendAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "LendAction"

	return json.Marshal(v)
}

func (v *LendAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "LendAction"

	return json.Marshal(*v)
}
