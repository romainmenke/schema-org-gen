package schemaorggo

import "encoding/json"

// ProductModel see : https://schema.org/ProductModel
type ProductModel struct {
	Product

	typeContext

	// IsVariantOf see : https://schema.org/isVariantOf
	// A pointer to a base product from which this product is a variant. It is safe to infer that the variant inherits all product features from the base model, unless defined locally. This is not transitive.
	IsVariantOf *ProductModel `json:"isVariantOf,omitempty"` // types : ProductModel

	// PredecessorOf see : https://schema.org/predecessorOf
	// A pointer from a previous, often discontinued variant of the product to its newer variant.
	PredecessorOf *ProductModel `json:"predecessorOf,omitempty"` // types : ProductModel

	// SuccessorOf see : https://schema.org/successorOf
	// A pointer from a newer variant of a product  to its previous, often discontinued predecessor.
	SuccessorOf *ProductModel `json:"successorOf,omitempty"` // types : ProductModel

}

func (v ProductModel) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ProductModel"

	return json.Marshal(v)
}

func (v *ProductModel) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ProductModel"

	return json.Marshal(*v)
}
