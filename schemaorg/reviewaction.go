package schemaorg

import "encoding/json"

// ReviewAction see : https://schema.org/ReviewAction
type ReviewAction struct {

typeContext

AssessAction

// ResultReview see : /resultReview
// A sub property of result. The review that resulted in the performing of the action.
ResultReview *Review `json:"resultReview"`

}

func (v *ReviewAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ReviewAction"

	return json.Marshal(v)
}
