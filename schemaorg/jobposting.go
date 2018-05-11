package schemaorg

// JobPosting see : https://schema.org/JobPosting
type JobPosting struct {

Intangible// BaseSalary see : /baseSalary
// The base salary of the job or of an employee in an EmployeeRole.
BaseSalary interface{} `json:"baseSalary"` // types : MonetaryAmount Number PriceSpecification

// DatePosted see : /datePosted
// Publication date for the job posting.
DatePosted string `json:"datePosted"`

// EducationRequirements see : /educationRequirements
// Educational background needed for the position.
EducationRequirements string `json:"educationRequirements"`

// EmploymentType see : /employmentType
// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
EmploymentType string `json:"employmentType"`

// EstimatedSalary see : /estimatedSalary
// A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
EstimatedSalary interface{} `json:"estimatedSalary"` // types : MonetaryAmount Number PriceSpecification

// ExperienceRequirements see : /experienceRequirements
// Description of skills and experience needed for the position.
ExperienceRequirements string `json:"experienceRequirements"`

// HiringOrganization see : /hiringOrganization
// Organization offering the job position.
HiringOrganization string `json:"hiringOrganization"`

// IncentiveCompensation see : /incentiveCompensation
// Description of bonus and commission compensation aspects of the job. Supersedes incentives (see: https://schema.org/incentives).
IncentiveCompensation string `json:"incentiveCompensation"`

// Industry see : /industry
// The industry associated with the job position.
Industry string `json:"industry"`

// JobBenefits see : /jobBenefits
// Description of benefits associated with the job. Supersedes benefits (see: https://schema.org/benefits).
JobBenefits string `json:"jobBenefits"`

// JobLocation see : /jobLocation
// A (typically single) geographic location associated with the job position.
JobLocation string `json:"jobLocation"`

// OccupationalCategory see : /occupationalCategory
// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
OccupationalCategory string `json:"occupationalCategory"`

// Qualifications see : /qualifications
// Specific qualifications required for this role.
Qualifications string `json:"qualifications"`

// Responsibilities see : /responsibilities
// Responsibilities associated with this role.
Responsibilities string `json:"responsibilities"`

// SalaryCurrency see : /salaryCurrency
// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
SalaryCurrency string `json:"salaryCurrency"`

// Skills see : /skills
// Skills required to fulfill this role.
Skills string `json:"skills"`

// SpecialCommitments see : /specialCommitments
// Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
SpecialCommitments string `json:"specialCommitments"`

// Title see : /title
// The title of the job.
Title string `json:"title"`

// ValidThrough see : /validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough string `json:"validThrough"`

// WorkHours see : /workHours
// The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
WorkHours string `json:"workHours"`

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

}

