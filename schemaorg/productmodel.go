package schemaorg

import "encoding/json"

// ProductModel see : https://schema.org/ProductModel
type ProductModel struct {

typeContext

Product

// IsVariantOf see : /isVariantOf
// A pointer to a base product from which this product is a variant. It is safe to infer that the variant inherits all product features from the base model, unless defined locally. This is not transitive.
IsVariantOf *ProductModel `json:"isVariantOf"`

// PredecessorOf see : /predecessorOf
// A pointer from a previous, often discontinued variant of the product to its newer variant.
PredecessorOf *ProductModel `json:"predecessorOf"`

// SuccessorOf see : /successorOf
// A pointer from a newer variant of a product  to its previous, often discontinued predecessor.
SuccessorOf *ProductModel `json:"successorOf"`

}

func (v *ProductModel) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ProductModel"

	return json.Marshal(v)
}
