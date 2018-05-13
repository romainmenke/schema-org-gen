package schemaorggo

import "encoding/json"

// BorrowAction see : https://schema.org/BorrowAction
type BorrowAction struct {
	TransferAction

	typeContext

	// Lender see : https://schema.org/lender
	// A sub property of participant. The person that lends the object being borrowed.
	// types : Organization Person
	Lender []interface{} `json:"lender,omitempty"`
}

func (v BorrowAction) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.TransferAction.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Lender
		if len(v.Lender) == 1 {
			value = v.Lender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["lender"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BorrowAction) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BorrowAction"

	return data, nil
}

func (v BorrowAction) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
