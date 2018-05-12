package schemaorggo

import "encoding/json"

// JobPosting see : https://schema.org/JobPosting
type JobPosting struct {
	Intangible

	typeContext

	// BaseSalary see : https://schema.org/baseSalary
	// The base salary of the job or of an employee in an EmployeeRole.
	// types : MonetaryAmount Number PriceSpecification
	BaseSalary interface{} `json:"baseSalary,omitempty"`

	// DatePosted see : https://schema.org/datePosted
	// Publication date for the job posting.
	// types : Date
	DatePosted Date `json:"datePosted,omitempty"`

	// EducationRequirements see : https://schema.org/educationRequirements
	// Educational background needed for the position.
	// types : Text
	EducationRequirements string `json:"educationRequirements,omitempty"`

	// EmploymentType see : https://schema.org/employmentType
	// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	// types : Text
	EmploymentType string `json:"employmentType,omitempty"`

	// EstimatedSalary see : https://schema.org/estimatedSalary
	// A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
	// types : MonetaryAmount Number PriceSpecification
	EstimatedSalary interface{} `json:"estimatedSalary,omitempty"`

	// ExperienceRequirements see : https://schema.org/experienceRequirements
	// Description of skills and experience needed for the position.
	// types : Text
	ExperienceRequirements string `json:"experienceRequirements,omitempty"`

	// HiringOrganization see : https://schema.org/hiringOrganization
	// Organization offering the job position.
	// types : Organization
	HiringOrganization *Organization `json:"hiringOrganization,omitempty"`

	// IncentiveCompensation see : https://schema.org/incentiveCompensation
	// Description of bonus and commission compensation aspects of the job. Supersedes incentives (see: https://schema.org/incentives).
	// types : Text
	IncentiveCompensation string `json:"incentiveCompensation,omitempty"`

	// Industry see : https://schema.org/industry
	// The industry associated with the job position.
	// types : Text
	Industry string `json:"industry,omitempty"`

	// JobBenefits see : https://schema.org/jobBenefits
	// Description of benefits associated with the job. Supersedes benefits (see: https://schema.org/benefits).
	// types : Text
	JobBenefits string `json:"jobBenefits,omitempty"`

	// JobLocation see : https://schema.org/jobLocation
	// A (typically single) geographic location associated with the job position.
	// types : Place
	JobLocation *Place `json:"jobLocation,omitempty"`

	// OccupationalCategory see : https://schema.org/occupationalCategory
	// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	// types : Text
	OccupationalCategory string `json:"occupationalCategory,omitempty"`

	// Qualifications see : https://schema.org/qualifications
	// Specific qualifications required for this role.
	// types : Text
	Qualifications string `json:"qualifications,omitempty"`

	// Responsibilities see : https://schema.org/responsibilities
	// Responsibilities associated with this role.
	// types : Text
	Responsibilities string `json:"responsibilities,omitempty"`

	// SalaryCurrency see : https://schema.org/salaryCurrency
	// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	// types : Text
	SalaryCurrency string `json:"salaryCurrency,omitempty"`

	// Skills see : https://schema.org/skills
	// Skills required to fulfill this role.
	// types : Text
	Skills string `json:"skills,omitempty"`

	// SpecialCommitments see : https://schema.org/specialCommitments
	// Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
	// types : Text
	SpecialCommitments string `json:"specialCommitments,omitempty"`

	// Title see : https://schema.org/title
	// The title of the job.
	// types : Text
	Title string `json:"title,omitempty"`

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough DateTime `json:"validThrough,omitempty"`

	// WorkHours see : https://schema.org/workHours
	// The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	// types : Text
	WorkHours string `json:"workHours,omitempty"`
}

func (v JobPosting) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "JobPosting"

	return json.Marshal(v)
}

func (v *JobPosting) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "JobPosting"

	return json.Marshal(*v)
}
