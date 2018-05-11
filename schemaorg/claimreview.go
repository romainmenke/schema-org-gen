package schemaorg

import "encoding/json"

// ClaimReview see : https://schema.org/ClaimReview
type ClaimReview struct {

typeContext

Review

// ClaimReviewed see : /claimReviewed
// A short summary of the specific claims reviewed in a ClaimReview.
ClaimReviewed string `json:"claimReviewed"`

}

func (v *ClaimReview) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ClaimReview"

	return json.Marshal(v)
}
