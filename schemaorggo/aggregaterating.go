package schemaorggo

import "encoding/json"

// AggregateRating see : https://schema.org/AggregateRating
type AggregateRating struct {
	Rating

	typeContext

	// ItemReviewed see : https://schema.org/itemReviewed
	// The item that is being reviewed/rated.
	// types : Thing
	ItemReviewed []*Thing `json:"itemReviewed,omitempty"`

	// RatingCount see : https://schema.org/ratingCount
	// The count of total number of ratings.
	// types : Integer
	RatingCount []float64 `json:"ratingCount,omitempty"`

	// ReviewCount see : https://schema.org/reviewCount
	// The count of total number of reviews.
	// types : Integer
	ReviewCount []float64 `json:"reviewCount,omitempty"`
}

func (v AggregateRating) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Rating.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.ItemReviewed
		if len(v.ItemReviewed) == 1 {
			value = v.ItemReviewed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["itemReviewed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RatingCount
		if len(v.RatingCount) == 1 {
			value = v.RatingCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ratingCount"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReviewCount
		if len(v.ReviewCount) == 1 {
			value = v.ReviewCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reviewCount"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v AggregateRating) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "AggregateRating"

	return data, nil
}

func (v AggregateRating) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
