package schemaorggo

import "encoding/json"

// OwnershipInfo see : https://schema.org/OwnershipInfo
type OwnershipInfo struct {
	StructuredValue

	typeContext

	// AcquiredFrom see : https://schema.org/acquiredFrom
	// The organization or person from which the product was acquired.
	// types : Organization Person
	AcquiredFrom []interface{} `json:"acquiredFrom,omitempty"`

	// OwnedFrom see : https://schema.org/ownedFrom
	// The date and time of obtaining the product.
	// types : DateTime
	OwnedFrom []DateTime `json:"ownedFrom,omitempty"`

	// OwnedThrough see : https://schema.org/ownedThrough
	// The date and time of giving up ownership on the product.
	// types : DateTime
	OwnedThrough []DateTime `json:"ownedThrough,omitempty"`

	// TypeOfGood see : https://schema.org/typeOfGood
	// The product that this structured value is referring to.
	// types : Product Service
	TypeOfGood []interface{} `json:"typeOfGood,omitempty"`
}

func (v OwnershipInfo) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AcquiredFrom
		if len(v.AcquiredFrom) == 1 {
			value = v.AcquiredFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["acquiredFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OwnedFrom
		if len(v.OwnedFrom) == 1 {
			value = v.OwnedFrom[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ownedFrom"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OwnedThrough
		if len(v.OwnedThrough) == 1 {
			value = v.OwnedThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ownedThrough"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TypeOfGood
		if len(v.TypeOfGood) == 1 {
			value = v.TypeOfGood[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["typeOfGood"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v OwnershipInfo) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "OwnershipInfo"

	return data, nil
}

func (v OwnershipInfo) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
