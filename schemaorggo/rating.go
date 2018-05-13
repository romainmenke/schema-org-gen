package schemaorggo

import "encoding/json"

// Rating see : https://schema.org/Rating
type Rating struct {
	Intangible

	typeContext

	// Author see : https://schema.org/author
	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	// types : Organization Person
	Author []interface{} `json:"author,omitempty"`

	// BestRating see : https://schema.org/bestRating
	// The highest value allowed in this rating system. If bestRating is omitted, 5 is assumed.
	// types : Number Text
	BestRating []interface{} `json:"bestRating,omitempty"`

	// RatingValue see : https://schema.org/ratingValue
	// The rating for the content.
	// types : Number Text
	RatingValue []interface{} `json:"ratingValue,omitempty"`

	// WorstRating see : https://schema.org/worstRating
	// The lowest value allowed in this rating system. If worstRating is omitted, 1 is assumed.
	// types : Number Text
	WorstRating []interface{} `json:"worstRating,omitempty"`
}

func (v Rating) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.Author
		if len(v.Author) == 1 {
			value = v.Author[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["author"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BestRating
		if len(v.BestRating) == 1 {
			value = v.BestRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bestRating"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RatingValue
		if len(v.RatingValue) == 1 {
			value = v.RatingValue[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ratingValue"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorstRating
		if len(v.WorstRating) == 1 {
			value = v.WorstRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["worstRating"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Rating) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Rating"

	return data, nil
}

func (v Rating) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
