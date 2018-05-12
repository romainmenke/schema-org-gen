package schemaorggo

import "encoding/json"

// Answer see : https://schema.org/Answer
type Answer struct {
	Comment

	typeContext

	// DownvoteCount see : https://schema.org/downvoteCount
	// The number of downvotes this question, answer or comment has received from the community.
	DownvoteCount int `json:"downvoteCount,omitempty"`

	// ParentItem see : https://schema.org/parentItem
	// The parent of a question, answer or item in general.
	ParentItem *Question `json:"parentItem,omitempty"`

	// UpvoteCount see : https://schema.org/upvoteCount
	// The number of upvotes this question, answer or comment has received from the community.
	UpvoteCount int `json:"upvoteCount,omitempty"`
}

func (v Answer) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Answer"

	return json.Marshal(v)
}

func (v *Answer) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Answer"

	return json.Marshal(*v)
}
