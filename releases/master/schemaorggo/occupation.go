package schemaorggo

import "encoding/json"

// Occupation see : https://schema.org/Occupation
type Occupation struct {

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// EducationRequirements see : https://schema.org/educationRequirements
	// Educational background needed for the position or Occupation.
	// types : Text
	EducationRequirements []string `json:"educationRequirements,omitempty"`

	// EstimatedSalary see : https://schema.org/estimatedSalary
	// An estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. Estimated salaries  are often computed by outside organizations rather than the hiring organization, who may not have committed to the estimated value.
	// types : MonetaryAmountDistribution
	EstimatedSalary []*MonetaryAmountDistribution `json:"estimatedSalary,omitempty"`

	// ExperienceRequirements see : https://schema.org/experienceRequirements
	// Description of skills and experience needed for the position or Occupation.
	// types : Text
	ExperienceRequirements []string `json:"experienceRequirements,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// OccupationLocation see : https://schema.org/occupationLocation
	//  The region/country for which this occupational description is appropriate. Note that educational requirements and qualifications can vary between jurisdictions.
	// types : AdministrativeArea
	OccupationLocation []*AdministrativeArea `json:"occupationLocation,omitempty"`

	// OccupationalCategory see : https://schema.org/occupationalCategory
	// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	// types : Text
	OccupationalCategory []string `json:"occupationalCategory,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Qualifications see : https://schema.org/qualifications
	// Specific qualifications required for this role or Occupation.
	// types : Text
	Qualifications []string `json:"qualifications,omitempty"`

	// Responsibilities see : https://schema.org/responsibilities
	// Responsibilities associated with this role or Occupation.
	// types : Text
	Responsibilities []string `json:"responsibilities,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Skills see : https://schema.org/skills
	// Skills required to fulfill this role or in this Occupation.
	// types : Text
	Skills []string `json:"skills,omitempty"`

	// SubjectOf see : https://schema.org/subjectOf
	// A CreativeWork or Event about this Thing..
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v Occupation) MarshalJSON() ([]byte, error) {
	type Alias Occupation

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Occupation\","), b[1:]...), nil
}