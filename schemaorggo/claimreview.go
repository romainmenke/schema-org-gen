package schemaorggo

import "encoding/json"

// ClaimReview see : https://schema.org/ClaimReview
type ClaimReview struct {
	Review

	typeContext

	// ClaimReviewed see : https://schema.org/claimReviewed
	// A short summary of the specific claims reviewed in a ClaimReview.
	// types : Text
	ClaimReviewed []string `json:"claimReviewed,omitempty"`
}

func (v ClaimReview) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Review.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ClaimReviewed
		if len(v.ClaimReviewed) == 1 {
			value = v.ClaimReviewed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["claimReviewed"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ClaimReview) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ClaimReview"

	return data, nil
}

func (v ClaimReview) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
