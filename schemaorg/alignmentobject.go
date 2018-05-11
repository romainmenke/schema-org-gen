package schemaorg

import "encoding/json"

// AlignmentObject see : https://schema.org/AlignmentObject
type AlignmentObject struct {

typeContext

Intangible

// AlignmentType see : /alignmentType
// A category of alignment between the learning resource and the framework node. Recommended values include: 'assesses', 'teaches', 'requires', 'textComplexity', 'readingLevel', 'educationalSubject', and 'educationalLevel'.
AlignmentType string `json:"alignmentType"`

// EducationalFramework see : /educationalFramework
// The framework to which the resource being described is aligned.
EducationalFramework string `json:"educationalFramework"`

// TargetDescription see : /targetDescription
// The description of a node in an established educational framework.
TargetDescription string `json:"targetDescription"`

// TargetName see : /targetName
// The name of a node in an established educational framework.
TargetName string `json:"targetName"`

// TargetUrl see : /targetUrl
// The URL of a node in an established educational framework.
TargetUrl string `json:"targetUrl"`

}

func (v *AlignmentObject) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AlignmentObject"

	return json.Marshal(v)
}
