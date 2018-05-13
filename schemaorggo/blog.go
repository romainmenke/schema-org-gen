package schemaorggo

import "encoding/json"

// Blog see : https://schema.org/Blog
type Blog struct {
	CreativeWork

	typeContext

	// BlogPost see : https://schema.org/blogPost
	// A posting that is part of this blog. Supersedes blogPosts (see: https://schema.org/blogPosts).
	// types : BlogPosting
	BlogPost []*BlogPosting `json:"blogPost,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn []string `json:"issn,omitempty"`
}

func (v Blog) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BlogPost
		if len(v.BlogPost) == 1 {
			value = v.BlogPost[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["blogPost"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Issn
		if len(v.Issn) == 1 {
			value = v.Issn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["issn"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Blog) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Blog"

	return data, nil
}

func (v Blog) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
