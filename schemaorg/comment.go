package schemaorg

import "encoding/json"

// Comment see : https://schema.org/Comment
type Comment struct {

typeContext

CreativeWork

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

func (v *Comment) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Comment"

	return json.Marshal(v)
}
