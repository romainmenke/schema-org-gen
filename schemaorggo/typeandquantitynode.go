package schemaorggo

import "encoding/json"

// TypeAndQuantityNode see : https://schema.org/TypeAndQuantityNode
type TypeAndQuantityNode struct {
	StructuredValue

	typeContext

	// AmountOfThisGood see : https://schema.org/amountOfThisGood
	// The quantity of the goods included in the offer.
	// types : Number
	AmountOfThisGood []float64 `json:"amountOfThisGood,omitempty"`

	// BusinessFunction see : https://schema.org/businessFunction
	// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	// types : BusinessFunction
	BusinessFunction []*BusinessFunction `json:"businessFunction,omitempty"`

	// TypeOfGood see : https://schema.org/typeOfGood
	// The product that this structured value is referring to.
	// types : Product Service
	TypeOfGood []interface{} `json:"typeOfGood,omitempty"`

	// UnitCode see : https://schema.org/unitCode
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	// types : Text URL
	UnitCode []string `json:"unitCode,omitempty"`

	// UnitText see : https://schema.org/unitText
	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	// unitCode (see: https://schema.orgunitCode).
	// types : Text
	UnitText []string `json:"unitText,omitempty"`
}

func (v TypeAndQuantityNode) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.StructuredValue.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.AmountOfThisGood
		if len(v.AmountOfThisGood) == 1 {
			value = v.AmountOfThisGood[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["amountOfThisGood"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BusinessFunction
		if len(v.BusinessFunction) == 1 {
			value = v.BusinessFunction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["businessFunction"] = json.RawMessage(b)
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

	{
		var value interface{} = v.UnitCode
		if len(v.UnitCode) == 1 {
			value = v.UnitCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unitCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UnitText
		if len(v.UnitText) == 1 {
			value = v.UnitText[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unitText"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v TypeAndQuantityNode) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "TypeAndQuantityNode"

	return data, nil
}

func (v TypeAndQuantityNode) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
