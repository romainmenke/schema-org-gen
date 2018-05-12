package schemaorggo

import "encoding/json"

// AlignmentObject see : https://schema.org/AlignmentObject
type AlignmentObject struct {
	Intangible

	typeContext

	// AlignmentType see : https://schema.org/alignmentType
	// A category of alignment between the learning resource and the framework node. Recommended values include: &#39;assesses&#39;, &#39;teaches&#39;, &#39;requires&#39;, &#39;textComplexity&#39;, &#39;readingLevel&#39;, &#39;educationalSubject&#39;, and &#39;educationalLevel&#39;.
	// types : Text
	AlignmentType string `json:"alignmentType,omitempty"`

	// EducationalFramework see : https://schema.org/educationalFramework
	// The framework to which the resource being described is aligned.
	// types : Text
	EducationalFramework string `json:"educationalFramework,omitempty"`

	// TargetDescription see : https://schema.org/targetDescription
	// The description of a node in an established educational framework.
	// types : Text
	TargetDescription string `json:"targetDescription,omitempty"`

	// TargetName see : https://schema.org/targetName
	// The name of a node in an established educational framework.
	// types : Text
	TargetName string `json:"targetName,omitempty"`

	// TargetUrl see : https://schema.org/targetUrl
	// The URL of a node in an established educational framework.
	// types : URL
	TargetUrl string `json:"targetUrl,omitempty"`
}

func (v AlignmentObject) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AlignmentObject"

	return json.Marshal(v)
}

func (v *AlignmentObject) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AlignmentObject"

	return json.Marshal(*v)
}
