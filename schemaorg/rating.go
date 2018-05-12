package schemaorg

import "encoding/json"

// Rating see : https://schema.org/Rating
type Rating struct {

Intangible

typeContext

// Author see : https://schema.org/author
// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
Author interface{} `json:"author"` // types : Organization Person

// BestRating see : https://schema.org/bestRating
// The highest value allowed in this rating system. If bestRating is omitted, 5 is assumed.
BestRating interface{} `json:"bestRating"` // types : Number Text

// RatingValue see : https://schema.org/ratingValue
// The rating for the content.
RatingValue interface{} `json:"ratingValue"` // types : Number Text

// WorstRating see : https://schema.org/worstRating
// The lowest value allowed in this rating system. If worstRating is omitted, 1 is assumed.
WorstRating interface{} `json:"worstRating"` // types : Number Text

}

func (v *Rating) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Rating"

	return json.Marshal(v)
}
