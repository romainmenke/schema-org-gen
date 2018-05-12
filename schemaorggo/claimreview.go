package schemaorggo

import "encoding/json"

// ClaimReview see : https://schema.org/ClaimReview
type ClaimReview struct {
	Review

	typeContext

	// ClaimReviewed see : https://schema.org/claimReviewed
	// A short summary of the specific claims reviewed in a ClaimReview.
	ClaimReviewed string `json:"claimReviewed,omitempty"` // types : Text

}

func (v ClaimReview) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ClaimReview"

	return json.Marshal(v)
}

func (v *ClaimReview) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ClaimReview"

	return json.Marshal(*v)
}
