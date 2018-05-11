package schemaorg

import "encoding/json"

// Answer see : https://schema.org/Answer
type Answer struct {

typeContext

Comment

// DownvoteCount see : /downvoteCount
// The number of downvotes this question, answer or comment has received from the community.
DownvoteCount int `json:"downvoteCount"`

// ParentItem see : /parentItem
// The parent of a question, answer or item in general.
ParentItem *Question `json:"parentItem"`

// UpvoteCount see : /upvoteCount
// The number of upvotes this question, answer or comment has received from the community.
UpvoteCount int `json:"upvoteCount"`

}

func (v *Answer) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Answer"

	return json.Marshal(v)
}
