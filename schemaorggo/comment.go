package schemaorggo

import "encoding/json"

// Comment see : https://schema.org/Comment
type Comment struct {
	CreativeWork

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

func (v Comment) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Comment"

	return json.Marshal(v)
}

func (v *Comment) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Comment"

	return json.Marshal(*v)
}
