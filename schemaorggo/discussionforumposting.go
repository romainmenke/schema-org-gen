package schemaorggo

import "encoding/json"

// DiscussionForumPosting see : https://schema.org/DiscussionForumPosting
type DiscussionForumPosting struct {
	SocialMediaPosting

	typeContext

	// SharedContent see : https://schema.org/sharedContent
	// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
	SharedContent *CreativeWork `json:"sharedContent,omitempty"` // types : CreativeWork

}

func (v DiscussionForumPosting) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DiscussionForumPosting"

	return json.Marshal(v)
}

func (v *DiscussionForumPosting) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DiscussionForumPosting"

	return json.Marshal(*v)
}
