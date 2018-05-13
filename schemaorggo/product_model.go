package schemaorggo

import "encoding/json"

// ProductModel see : https://schema.org/ProductModel
type ProductModel struct {
	Product

	typeContext

	// IsVariantOf see : https://schema.org/isVariantOf
	// A pointer to a base product from which this product is a variant. It is safe to infer that the variant inherits all product features from the base model, unless defined locally. This is not transitive.
	// types : ProductModel
	IsVariantOf []*ProductModel `json:"isVariantOf,omitempty"`

	// PredecessorOf see : https://schema.org/predecessorOf
	// A pointer from a previous, often discontinued variant of the product to its newer variant.
	// types : ProductModel
	PredecessorOf []*ProductModel `json:"predecessorOf,omitempty"`

	// SuccessorOf see : https://schema.org/successorOf
	// A pointer from a newer variant of a product  to its previous, often discontinued predecessor.
	// types : ProductModel
	SuccessorOf []*ProductModel `json:"successorOf,omitempty"`
}

func (v ProductModel) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Product.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.IsVariantOf
		if len(v.IsVariantOf) == 1 {
			value = v.IsVariantOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isVariantOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PredecessorOf
		if len(v.PredecessorOf) == 1 {
			value = v.PredecessorOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["predecessorOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SuccessorOf
		if len(v.SuccessorOf) == 1 {
			value = v.SuccessorOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["successorOf"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ProductModel) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ProductModel"

	return data, nil
}

func (v ProductModel) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
