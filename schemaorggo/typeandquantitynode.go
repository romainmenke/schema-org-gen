package schemaorggo

import "encoding/json"

// TypeAndQuantityNode see : https://schema.org/TypeAndQuantityNode
type TypeAndQuantityNode struct {
	StructuredValue

	typeContext

	// AmountOfThisGood see : https://schema.org/amountOfThisGood
	// The quantity of the goods included in the offer.
	// types : Number
	AmountOfThisGood float64 `json:"amountOfThisGood,omitempty"`

	// BusinessFunction see : https://schema.org/businessFunction
	// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
	// types : BusinessFunction
	BusinessFunction *BusinessFunction `json:"businessFunction,omitempty"`

	// TypeOfGood see : https://schema.org/typeOfGood
	// The product that this structured value is referring to.
	// types : Product Service
	TypeOfGood interface{} `json:"typeOfGood,omitempty"`

	// UnitCode see : https://schema.org/unitCode
	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	// types : Text URL
	UnitCode string `json:"unitCode,omitempty"`

	// UnitText see : https://schema.org/unitText
	// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
	// unitCode (see: https://schema.orgunitCode).
	// types : Text
	UnitText string `json:"unitText,omitempty"`
}

func (v TypeAndQuantityNode) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TypeAndQuantityNode"

	return json.Marshal(v)
}

func (v *TypeAndQuantityNode) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "TypeAndQuantityNode"

	return json.Marshal(*v)
}
