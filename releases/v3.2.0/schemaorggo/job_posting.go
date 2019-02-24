package schemaorggo

import "encoding/json"

// JobPosting see : https://schema.org/JobPosting
type JobPosting struct {

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

	// BaseSalary see : https://schema.org/baseSalary
	// The base salary of the job or of an employee in an EmployeeRole.
	// types : Number PriceSpecification MonetaryAmount
	BaseSalary []interface{} `json:"baseSalary,omitempty"`

	// Benefits see : https://schema.org/benefits
	// Description of benefits associated with the job.
	// types : Text
	Benefits []string `json:"benefits,omitempty"`

	// DatePosted see : https://schema.org/datePosted
	// Publication date for the job posting.
	// types : Date
	DatePosted []Date `json:"datePosted,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// EducationRequirements see : https://schema.org/educationRequirements
	// Educational background needed for the position.
	// types : Text
	EducationRequirements []string `json:"educationRequirements,omitempty"`

	// EmploymentType see : https://schema.org/employmentType
	// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	// types : Text
	EmploymentType []string `json:"employmentType,omitempty"`

	// ExperienceRequirements see : https://schema.org/experienceRequirements
	// Description of skills and experience needed for the position.
	// types : Text
	ExperienceRequirements []string `json:"experienceRequirements,omitempty"`

	// HiringOrganization see : https://schema.org/hiringOrganization
	// Organization offering the job position.
	// types : Organization
	HiringOrganization []*Organization `json:"hiringOrganization,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IncentiveCompensation see : https://schema.org/incentiveCompensation
	// Description of bonus and commission compensation aspects of the job.
	// types : Text
	IncentiveCompensation []string `json:"incentiveCompensation,omitempty"`

	// Incentives see : https://schema.org/incentives
	// Description of bonus and commission compensation aspects of the job.
	// types : Text
	Incentives []string `json:"incentives,omitempty"`

	// Industry see : https://schema.org/industry
	// The industry associated with the job position.
	// types : Text
	Industry []string `json:"industry,omitempty"`

	// JobBenefits see : https://schema.org/jobBenefits
	// Description of benefits associated with the job.
	// types : Text
	JobBenefits []string `json:"jobBenefits,omitempty"`

	// JobLocation see : https://schema.org/jobLocation
	// A (typically single) geographic location associated with the job position.
	// types : Place
	JobLocation []*Place `json:"jobLocation,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// OccupationalCategory see : https://schema.org/occupationalCategory
	// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	// types : Text
	OccupationalCategory []string `json:"occupationalCategory,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Qualifications see : https://schema.org/qualifications
	// Specific qualifications required for this role.
	// types : Text
	Qualifications []string `json:"qualifications,omitempty"`

	// Responsibilities see : https://schema.org/responsibilities
	// Responsibilities associated with this role.
	// types : Text
	Responsibilities []string `json:"responsibilities,omitempty"`

	// SalaryCurrency see : https://schema.org/salaryCurrency
	// The currency (coded using [ISO 4217](http://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	// types : Text
	SalaryCurrency []string `json:"salaryCurrency,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Skills see : https://schema.org/skills
	// Skills required to fulfill this role.
	// types : Text
	Skills []string `json:"skills,omitempty"`

	// SpecialCommitments see : https://schema.org/specialCommitments
	// Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
	// types : Text
	SpecialCommitments []string `json:"specialCommitments,omitempty"`

	// Title see : https://schema.org/title
	// The title of the job.
	// types : Text
	Title []string `json:"title,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`

	// WorkHours see : https://schema.org/workHours
	// The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	// types : Text
	WorkHours []string `json:"workHours,omitempty"`
}

func (v JobPosting) MarshalJSON() ([]byte, error) {
	type Alias JobPosting

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"JobPosting\","), b[1:]...), nil
}
