package schemaorg

import "encoding/json"

// BlogPosting see : https://schema.org/BlogPosting
type BlogPosting struct {

typeContext

SocialMediaPosting

// SharedContent see : https://schema.org/sharedContent
// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
SharedContent *CreativeWork `json:"sharedContent"`

}

func (v *BlogPosting) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BlogPosting"

	return json.Marshal(v)
}
