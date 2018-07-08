package schemaorggo

import "encoding/json"

// Review see : https://schema.org/Review
type Review struct {
	CreativeWork

	typeContext

	// ItemReviewed see : https://schema.org/itemReviewed
	// The item that is being reviewed/rated.
	// types : Thing
	ItemReviewed []*Thing `json:"itemReviewed,omitempty"`

	// ReviewAspect see : https://pending.schema.org/reviewAspect
	// This Review or Rating is relevant to this part or facet of the itemReviewed.
	// types : Text
	ReviewAspect []string `json:"reviewAspect,omitempty"`

	// ReviewBody see : https://schema.org/reviewBody
	// The actual body of the review.
	// types : Text
	ReviewBody []string `json:"reviewBody,omitempty"`

	// ReviewRating see : https://schema.org/reviewRating
	// The rating given in this review. Note that reviews can themselves be rated. The reviewRating applies to rating given by the review. The aggregateRating (see: https://schema.org/aggregateRating) property applies to the review itself, as a creative work.
	// types : Rating
	ReviewRating []*Rating `json:"reviewRating,omitempty"`
}

func (v Review) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

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
		var value interface{} = v.ReviewAspect
		if len(v.ReviewAspect) == 1 {
			value = v.ReviewAspect[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reviewAspect"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReviewBody
		if len(v.ReviewBody) == 1 {
			value = v.ReviewBody[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reviewBody"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReviewRating
		if len(v.ReviewRating) == 1 {
			value = v.ReviewRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["reviewRating"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Review) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Review"

	return data, nil
}

func (v Review) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
