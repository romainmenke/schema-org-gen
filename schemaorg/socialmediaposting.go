package schemaorg

import "encoding/json"

// SocialMediaPosting see : https://schema.org/SocialMediaPosting
type SocialMediaPosting struct {

typeContext

Article

// SharedContent see : https://schema.org/sharedContent
// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
SharedContent *CreativeWork `json:"sharedContent"`

}

func (v *SocialMediaPosting) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "SocialMediaPosting"

	return json.Marshal(v)
}
