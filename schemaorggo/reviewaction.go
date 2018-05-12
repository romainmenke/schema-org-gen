package schemaorggo

import "encoding/json"

// ReviewAction see : https://schema.org/ReviewAction
type ReviewAction struct {
	AssessAction

	typeContext

	// ResultReview see : https://schema.org/resultReview
	// A sub property of result. The review that resulted in the performing of the action.
	// types : Review
	ResultReview *Review `json:"resultReview,omitempty"`
}

func (v ReviewAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReviewAction"

	return json.Marshal(v)
}

func (v *ReviewAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ReviewAction"

	return json.Marshal(*v)
}
