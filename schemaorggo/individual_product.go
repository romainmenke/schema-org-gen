package schemaorggo

import "encoding/json"

// IndividualProduct see : https://schema.org/IndividualProduct
type IndividualProduct struct {
	Product

	typeContext

	// SerialNumber see : https://schema.org/serialNumber
	// The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	// types : Text
	SerialNumber []string `json:"serialNumber,omitempty"`
}

func (v IndividualProduct) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Product.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.SerialNumber
		if len(v.SerialNumber) == 1 {
			value = v.SerialNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["serialNumber"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v IndividualProduct) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "IndividualProduct"

	return data, nil
}

func (v IndividualProduct) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
