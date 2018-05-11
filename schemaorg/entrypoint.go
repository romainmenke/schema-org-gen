package schemaorg

import "encoding/json"

// EntryPoint see : https://schema.org/EntryPoint
type EntryPoint struct {

typeContext

Intangible

// ActionApplication see : /actionApplication
// An application that can complete the request. Supersedes application (see: https://schema.org/application).
ActionApplication *SoftwareApplication `json:"actionApplication"`

// ActionPlatform see : /actionPlatform
// The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
ActionPlatform interface{} `json:"actionPlatform"` // types : Text URL

// ContentType see : /contentType
// The supported content type(s) for an EntryPoint response.
ContentType string `json:"contentType"`

// EncodingType see : /encodingType
// The supported encoding type(s) for an EntryPoint request.
EncodingType string `json:"encodingType"`

// HttpMethod see : /httpMethod
// An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
HttpMethod string `json:"httpMethod"`

// UrlTemplate see : /urlTemplate
// An url template (RFC6570) that will be used to construct the target of the execution of the action.
UrlTemplate string `json:"urlTemplate"`

}

func (v *EntryPoint) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "EntryPoint"

	return json.Marshal(v)
}
