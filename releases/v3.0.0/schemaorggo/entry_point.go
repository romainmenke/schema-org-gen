package schemaorggo

import "encoding/json"

// EntryPoint see : https://schema.org/EntryPoint
type EntryPoint struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// ActionApplication see : https://schema.org/actionApplication
	// An application that can complete the request.
	// types : SoftwareApplication
	ActionApplication []*SoftwareApplication `json:"actionApplication,omitempty"`

	// ActionPlatform see : https://schema.org/actionPlatform
	// The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
	// types :

	ActionPlatform []interface{} `json:"actionPlatform,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Application see : https://schema.org/application
	// An application that can complete the request.
	// types : SoftwareApplication
	Application []*SoftwareApplication `json:"application,omitempty"`

	// ContentType see : https://schema.org/contentType
	// The supported content type(s) for an EntryPoint response.
	// types : Text
	ContentType []string `json:"contentType,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// EncodingType see : https://schema.org/encodingType
	// The supported encoding type(s) for an EntryPoint request.
	// types : Text
	EncodingType []string `json:"encodingType,omitempty"`

	// HttpMethod see : https://schema.org/httpMethod
	// An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	// types : Text
	HttpMethod []string `json:"httpMethod,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// UrlTemplate see : https://schema.org/urlTemplate
	// An url template (RFC6570) that will be used to construct the target of the execution of the action.
	// types : Text
	UrlTemplate []string `json:"urlTemplate,omitempty"`
}

func (v EntryPoint) MarshalJSON() ([]byte, error) {
	type Alias EntryPoint

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"EntryPoint\","), b[1:]...), nil
}
