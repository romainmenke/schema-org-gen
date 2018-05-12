package schemaorg

import "encoding/json"

// DiscussionForumPosting see : https://schema.org/DiscussionForumPosting
type DiscussionForumPosting struct {

SocialMediaPosting

typeContext

// SharedContent see : https://schema.org/sharedContent
// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
SharedContent *CreativeWork `json:"sharedContent"`

}

func (v *DiscussionForumPosting) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DiscussionForumPosting"

	return json.Marshal(v)
}
