package schemaorggo

import "encoding/json"

// Review see : https://schema.org/Review
type Review struct {
	CreativeWork

	typeContext

	// ItemReviewed see : https://schema.org/itemReviewed
	// The item that is being reviewed/rated.
	ItemReviewed *Thing `json:"itemReviewed,omitempty"`

	// ReviewBody see : https://schema.org/reviewBody
	// The actual body of the review.
	ReviewBody string `json:"reviewBody,omitempty"`

	// ReviewRating see : https://schema.org/reviewRating
	// The rating given in this review. Note that reviews can themselves be rated. The reviewRating applies to rating given by the review. The aggregateRating (see: https://schema.org/aggregateRating) property applies to the review itself, as a creative work.
	ReviewRating *Rating `json:"reviewRating,omitempty"`
}

func (v Review) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Review"

	return json.Marshal(v)
}

func (v *Review) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Review"

	return json.Marshal(*v)
}
