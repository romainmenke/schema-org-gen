package schemaorg

import "encoding/json"

// TypeAndQuantityNode see : https://schema.org/TypeAndQuantityNode
type TypeAndQuantityNode struct {

typeContext

StructuredValue

// AmountOfThisGood see : /amountOfThisGood
// The quantity of the goods included in the offer.
AmountOfThisGood float64 `json:"amountOfThisGood"`

// BusinessFunction see : /businessFunction
// The business function (e.g. sell, lease, repair, dispose) of the offer or component of a bundle (TypeAndQuantityNode). The default is http://purl.org/goodrelations/v1#Sell.
BusinessFunction *BusinessFunction `json:"businessFunction"`

// TypeOfGood see : /typeOfGood
// The product that this structured value is referring to.
TypeOfGood interface{} `json:"typeOfGood"` // types : Product Service

// UnitCode see : /unitCode
// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL. Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
UnitCode interface{} `json:"unitCode"` // types : Text URL

// UnitText see : /unitText
// A string or text indicating the unit of measurement. Useful if you cannot provide a standard unit code for
// unitCode (see: https://schema.orgunitCode).
UnitText string `json:"unitText"`

}

func (v *TypeAndQuantityNode) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TypeAndQuantityNode"

	return json.Marshal(v)
}
