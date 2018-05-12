package schemaorggo

import "encoding/json"

// AggregateRating see : https://schema.org/AggregateRating
type AggregateRating struct {
	Rating

	typeContext

	// ItemReviewed see : https://schema.org/itemReviewed
	// The item that is being reviewed/rated.
	ItemReviewed *Thing `json:"itemReviewed"`

	// RatingCount see : https://schema.org/ratingCount
	// The count of total number of ratings.
	RatingCount int `json:"ratingCount"`

	// ReviewCount see : https://schema.org/reviewCount
	// The count of total number of reviews.
	ReviewCount int `json:"reviewCount"`
}

func (v *AggregateRating) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AggregateRating"

	return json.Marshal(v)
}
