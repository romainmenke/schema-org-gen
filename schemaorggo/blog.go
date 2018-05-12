package schemaorggo

import "encoding/json"

// Blog see : https://schema.org/Blog
type Blog struct {
	CreativeWork

	typeContext

	// BlogPost see : https://schema.org/blogPost
	// A posting that is part of this blog. Supersedes blogPosts (see: https://schema.org/blogPosts).
	// types : BlogPosting
	BlogPost *BlogPosting `json:"blogPost,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn string `json:"issn,omitempty"`
}

func (v Blog) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Blog"

	return json.Marshal(v)
}

func (v *Blog) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Blog"

	return json.Marshal(*v)
}
