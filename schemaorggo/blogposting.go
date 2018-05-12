package schemaorggo

import "encoding/json"

// BlogPosting see : https://schema.org/BlogPosting
type BlogPosting struct {
	SocialMediaPosting

	typeContext

	// SharedContent see : https://schema.org/sharedContent
	// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
	SharedContent *CreativeWork `json:"sharedContent,omitempty"`
}

func (v BlogPosting) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BlogPosting"

	return json.Marshal(v)
}

func (v *BlogPosting) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "BlogPosting"

	return json.Marshal(*v)
}
