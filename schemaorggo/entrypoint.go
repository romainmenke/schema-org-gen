package schemaorggo

import "encoding/json"

// EntryPoint see : https://schema.org/EntryPoint
type EntryPoint struct {
	Intangible

	typeContext

	// ActionApplication see : https://schema.org/actionApplication
	// An application that can complete the request. Supersedes application (see: https://schema.org/application).
	// types : SoftwareApplication
	ActionApplication *SoftwareApplication `json:"actionApplication,omitempty"`

	// ActionPlatform see : https://schema.org/actionPlatform
	// The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
	// types : Text URL
	ActionPlatform string `json:"actionPlatform,omitempty"`

	// ContentType see : https://schema.org/contentType
	// The supported content type(s) for an EntryPoint response.
	// types : Text
	ContentType string `json:"contentType,omitempty"`

	// EncodingType see : https://schema.org/encodingType
	// The supported encoding type(s) for an EntryPoint request.
	// types : Text
	EncodingType string `json:"encodingType,omitempty"`

	// HttpMethod see : https://schema.org/httpMethod
	// An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	// types : Text
	HttpMethod string `json:"httpMethod,omitempty"`

	// UrlTemplate see : https://schema.org/urlTemplate
	// An url template (RFC6570) that will be used to construct the target of the execution of the action.
	// types : Text
	UrlTemplate string `json:"urlTemplate,omitempty"`
}

func (v EntryPoint) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EntryPoint"

	return json.Marshal(v)
}

func (v *EntryPoint) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "EntryPoint"

	return json.Marshal(*v)
}
