package schemaorggo

import "encoding/json"

// EntryPoint see : https://schema.org/EntryPoint
type EntryPoint struct {
	Intangible

	typeContext

	// ActionApplication see : https://schema.org/actionApplication
	// An application that can complete the request. Supersedes application (see: https://schema.org/application).
	// types : SoftwareApplication
	ActionApplication []*SoftwareApplication `json:"actionApplication,omitempty"`

	// ActionPlatform see : https://schema.org/actionPlatform
	// The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
	// types : Text URL
	ActionPlatform []string `json:"actionPlatform,omitempty"`

	// ContentType see : https://schema.org/contentType
	// The supported content type(s) for an EntryPoint response.
	// types : Text
	ContentType []string `json:"contentType,omitempty"`

	// EncodingType see : https://schema.org/encodingType
	// The supported encoding type(s) for an EntryPoint request.
	// types : Text
	EncodingType []string `json:"encodingType,omitempty"`

	// HttpMethod see : https://schema.org/httpMethod
	// An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	// types : Text
	HttpMethod []string `json:"httpMethod,omitempty"`

	// UrlTemplate see : https://schema.org/urlTemplate
	// An url template (RFC6570) that will be used to construct the target of the execution of the action.
	// types : Text
	UrlTemplate []string `json:"urlTemplate,omitempty"`
}

func (v EntryPoint) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.ActionApplication
		if len(v.ActionApplication) == 1 {
			value = v.ActionApplication[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actionApplication"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ActionPlatform
		if len(v.ActionPlatform) == 1 {
			value = v.ActionPlatform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actionPlatform"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContentType
		if len(v.ContentType) == 1 {
			value = v.ContentType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contentType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EncodingType
		if len(v.EncodingType) == 1 {
			value = v.EncodingType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["encodingType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HttpMethod
		if len(v.HttpMethod) == 1 {
			value = v.HttpMethod[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["httpMethod"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.UrlTemplate
		if len(v.UrlTemplate) == 1 {
			value = v.UrlTemplate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["urlTemplate"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v EntryPoint) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "EntryPoint"

	return data, nil
}

func (v EntryPoint) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
