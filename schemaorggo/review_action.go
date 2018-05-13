package schemaorggo

import "encoding/json"

// ReviewAction see : https://schema.org/ReviewAction
type ReviewAction struct {
	AssessAction

	typeContext

	// ResultReview see : https://schema.org/resultReview
	// A sub property of result. The review that resulted in the performing of the action.
	// types : Review
	ResultReview []*Review `json:"resultReview,omitempty"`
}

func (v ReviewAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.AssessAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ResultReview
		if len(v.ResultReview) == 1 {
			value = v.ResultReview[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["resultReview"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ReviewAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ReviewAction"

	return data, nil
}

func (v ReviewAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
