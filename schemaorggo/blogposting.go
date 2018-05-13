package schemaorggo

import "encoding/json"

// BlogPosting see : https://schema.org/BlogPosting
type BlogPosting struct {
	SocialMediaPosting

	typeContext

	// SharedContent see : https://schema.org/sharedContent
	// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
	// types : CreativeWork
	SharedContent []*CreativeWork `json:"sharedContent,omitempty"`
}

func (v BlogPosting) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.SocialMediaPosting.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.SharedContent
		if len(v.SharedContent) == 1 {
			value = v.SharedContent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sharedContent"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v BlogPosting) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "BlogPosting"

	return data, nil
}

func (v BlogPosting) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
