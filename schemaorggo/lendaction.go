package schemaorggo

import "encoding/json"

// LendAction see : https://schema.org/LendAction
type LendAction struct {
	TransferAction

	typeContext

	// Borrower see : https://schema.org/borrower
	// A sub property of participant. The person that borrows the object being lent.
	// types : Person
	Borrower []*Person `json:"borrower,omitempty"`
}

func (v LendAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TransferAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Borrower
		if len(v.Borrower) == 1 {
			value = v.Borrower[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["borrower"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v LendAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "LendAction"

	return data, nil
}

func (v LendAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
