package schemaorg

import "encoding/json"

// AggregateRating see : https://schema.org/AggregateRating
type AggregateRating struct {

typeContext

Rating

// ItemReviewed see : /itemReviewed
// The item that is being reviewed/rated.
ItemReviewed *Thing `json:"itemReviewed"`

// RatingCount see : /ratingCount
// The count of total number of ratings.
RatingCount int `json:"ratingCount"`

// ReviewCount see : /reviewCount
// The count of total number of reviews.
ReviewCount int `json:"reviewCount"`

}

func (v *AggregateRating) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AggregateRating"

	return json.Marshal(v)
}
