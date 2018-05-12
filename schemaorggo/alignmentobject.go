package schemaorggo

import "encoding/json"

// AlignmentObject see : https://schema.org/AlignmentObject
type AlignmentObject struct {
	Intangible

	typeContext

	// AlignmentType see : https://schema.org/alignmentType
	// A category of alignment between the learning resource and the framework node. Recommended values include: 'assesses', 'teaches', 'requires', 'textComplexity', 'readingLevel', 'educationalSubject', and 'educationalLevel'.
	AlignmentType string `json:"alignmentType,omitempty"` // types : Text

	// EducationalFramework see : https://schema.org/educationalFramework
	// The framework to which the resource being described is aligned.
	EducationalFramework string `json:"educationalFramework,omitempty"` // types : Text

	// TargetDescription see : https://schema.org/targetDescription
	// The description of a node in an established educational framework.
	TargetDescription string `json:"targetDescription,omitempty"` // types : Text

	// TargetName see : https://schema.org/targetName
	// The name of a node in an established educational framework.
	TargetName string `json:"targetName,omitempty"` // types : Text

	// TargetUrl see : https://schema.org/targetUrl
	// The URL of a node in an established educational framework.
	TargetUrl string `json:"targetUrl,omitempty"` // types : URL

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
