package schemaorg

// AlignmentObject see : https://schema.org/AlignmentObject
type AlignmentObject struct {

Intangible// AlignmentType see : /alignmentType
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

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// CoursePrerequisites see : /coursePrerequisites
// Requirements for taking the Course. May be completion of another Course (see: https://schema.org/Course) or a textual description like "permission of instructor". Requirements may be a pre-requisite competency, referenced using AlignmentObject (see: https://schema.org/AlignmentObject). 
CoursePrerequisites string `json:"coursePrerequisites"`

// EducationalAlignment see : /educationalAlignment
// An alignment to an established educational framework. 
EducationalAlignment string `json:"educationalAlignment"`

}

