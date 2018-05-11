package schemaorg

import "encoding/json"

// JobPosting see : https://schema.org/JobPosting
type JobPosting struct {

typeContext

Intangible

// BaseSalary see : https://schema.org/baseSalary
// The base salary of the job or of an employee in an EmployeeRole.
BaseSalary interface{} `json:"baseSalary"` // types : MonetaryAmount Number PriceSpecification

// DatePosted see : https://schema.org/datePosted
// Publication date for the job posting.
DatePosted interface{} `json:"datePosted"`

// EducationRequirements see : https://schema.org/educationRequirements
// Educational background needed for the position.
EducationRequirements string `json:"educationRequirements"`

// EmploymentType see : https://schema.org/employmentType
// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
EmploymentType string `json:"employmentType"`

// EstimatedSalary see : https://schema.org/estimatedSalary
// A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
EstimatedSalary interface{} `json:"estimatedSalary"` // types : MonetaryAmount Number PriceSpecification

// ExperienceRequirements see : https://schema.org/experienceRequirements
// Description of skills and experience needed for the position.
ExperienceRequirements string `json:"experienceRequirements"`

// HiringOrganization see : https://schema.org/hiringOrganization
// Organization offering the job position.
HiringOrganization *Organization `json:"hiringOrganization"`

// IncentiveCompensation see : https://schema.org/incentiveCompensation
// Description of bonus and commission compensation aspects of the job. Supersedes incentives (see: https://schema.org/incentives).
IncentiveCompensation string `json:"incentiveCompensation"`

// Industry see : https://schema.org/industry
// The industry associated with the job position.
Industry string `json:"industry"`

// JobBenefits see : https://schema.org/jobBenefits
// Description of benefits associated with the job. Supersedes benefits (see: https://schema.org/benefits).
JobBenefits string `json:"jobBenefits"`

// JobLocation see : https://schema.org/jobLocation
// A (typically single) geographic location associated with the job position.
JobLocation *Place `json:"jobLocation"`

// OccupationalCategory see : https://schema.org/occupationalCategory
// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
OccupationalCategory string `json:"occupationalCategory"`

// Qualifications see : https://schema.org/qualifications
// Specific qualifications required for this role.
Qualifications string `json:"qualifications"`

// Responsibilities see : https://schema.org/responsibilities
// Responsibilities associated with this role.
Responsibilities string `json:"responsibilities"`

// SalaryCurrency see : https://schema.org/salaryCurrency
// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
SalaryCurrency string `json:"salaryCurrency"`

// Skills see : https://schema.org/skills
// Skills required to fulfill this role.
Skills string `json:"skills"`

// SpecialCommitments see : https://schema.org/specialCommitments
// Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
SpecialCommitments string `json:"specialCommitments"`

// Title see : https://schema.org/title
// The title of the job.
Title string `json:"title"`

// ValidThrough see : https://schema.org/validThrough
// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
ValidThrough interface{} `json:"validThrough"`

// WorkHours see : https://schema.org/workHours
// The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
WorkHours string `json:"workHours"`

}

func (v *JobPosting) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "JobPosting"

	return json.Marshal(v)
}
