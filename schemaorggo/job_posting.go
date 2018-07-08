package schemaorggo

import "encoding/json"

// JobPosting see : https://schema.org/JobPosting
type JobPosting struct {
	Intangible

	typeContext

	// BaseSalary see : https://schema.org/baseSalary
	// The base salary of the job or of an employee in an EmployeeRole.
	// types : MonetaryAmount Number PriceSpecification
	BaseSalary []interface{} `json:"baseSalary,omitempty"`

	// DatePosted see : https://schema.org/datePosted
	// Publication date for the job posting.
	// types : Date
	DatePosted []Date `json:"datePosted,omitempty"`

	// EducationRequirements see : https://pending.schema.org/educationRequirements
	// Educational background needed for the position or Occupation.
	// types : Text
	EducationRequirements []string `json:"educationRequirements,omitempty"`

	// EmploymentType see : https://schema.org/employmentType
	// Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	// types : Text
	EmploymentType []string `json:"employmentType,omitempty"`

	// EstimatedSalary see : https://pending.schema.org/estimatedSalary
	// A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
	// types : MonetaryAmount MonetaryAmountDistribution Number PriceSpecification
	EstimatedSalary []interface{} `json:"estimatedSalary,omitempty"`

	// ExperienceRequirements see : https://pending.schema.org/experienceRequirements
	// Description of skills and experience needed for the position or Occupation.
	// types : Text
	ExperienceRequirements []string `json:"experienceRequirements,omitempty"`

	// HiringOrganization see : https://schema.org/hiringOrganization
	// Organization offering the job position.
	// types : Organization
	HiringOrganization []*Organization `json:"hiringOrganization,omitempty"`

	// IncentiveCompensation see : https://schema.org/incentiveCompensation
	// Description of bonus and commission compensation aspects of the job. Supersedes incentives (see: https://schema.org/incentives).
	// types : Text
	IncentiveCompensation []string `json:"incentiveCompensation,omitempty"`

	// Industry see : https://schema.org/industry
	// The industry associated with the job position.
	// types : Text
	Industry []string `json:"industry,omitempty"`

	// JobBenefits see : https://schema.org/jobBenefits
	// Description of benefits associated with the job. Supersedes benefits (see: https://schema.org/benefits).
	// types : Text
	JobBenefits []string `json:"jobBenefits,omitempty"`

	// JobLocation see : https://schema.org/jobLocation
	// A (typically single) geographic location associated with the job position.
	// types : Place
	JobLocation []*Place `json:"jobLocation,omitempty"`

	// OccupationalCategory see : https://pending.schema.org/occupationalCategory
	// Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	// types : Text
	OccupationalCategory []string `json:"occupationalCategory,omitempty"`

	// Qualifications see : https://pending.schema.org/qualifications
	// Specific qualifications required for this role or Occupation.
	// types : Text
	Qualifications []string `json:"qualifications,omitempty"`

	// RelevantOccupation see : https://pending.schema.org/relevantOccupation
	// The Occupation for the JobPosting.
	// types : Occupation
	RelevantOccupation []interface{} `json:"relevantOccupation,omitempty"`

	// Responsibilities see : https://pending.schema.org/responsibilities
	// Responsibilities associated with this role or Occupation.
	// types : Text
	Responsibilities []string `json:"responsibilities,omitempty"`

	// SalaryCurrency see : https://schema.org/salaryCurrency
	// The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	// types : Text
	SalaryCurrency []string `json:"salaryCurrency,omitempty"`

	// Skills see : https://pending.schema.org/skills
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

	// ValidThrough see : https://schema.org/validThrough
	// The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	// types : DateTime
	ValidThrough []DateTime `json:"validThrough,omitempty"`

	// WorkHours see : https://schema.org/workHours
	// The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	// types : Text
	WorkHours []string `json:"workHours,omitempty"`
}

func (v JobPosting) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Intangible.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.BaseSalary
		if len(v.BaseSalary) == 1 {
			value = v.BaseSalary[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["baseSalary"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DatePosted
		if len(v.DatePosted) == 1 {
			value = v.DatePosted[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["datePosted"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EducationRequirements
		if len(v.EducationRequirements) == 1 {
			value = v.EducationRequirements[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["educationRequirements"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EmploymentType
		if len(v.EmploymentType) == 1 {
			value = v.EmploymentType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["employmentType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EstimatedSalary
		if len(v.EstimatedSalary) == 1 {
			value = v.EstimatedSalary[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["estimatedSalary"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExperienceRequirements
		if len(v.ExperienceRequirements) == 1 {
			value = v.ExperienceRequirements[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["experienceRequirements"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HiringOrganization
		if len(v.HiringOrganization) == 1 {
			value = v.HiringOrganization[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hiringOrganization"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IncentiveCompensation
		if len(v.IncentiveCompensation) == 1 {
			value = v.IncentiveCompensation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["incentiveCompensation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Industry
		if len(v.Industry) == 1 {
			value = v.Industry[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["industry"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.JobBenefits
		if len(v.JobBenefits) == 1 {
			value = v.JobBenefits[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["jobBenefits"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.JobLocation
		if len(v.JobLocation) == 1 {
			value = v.JobLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["jobLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OccupationalCategory
		if len(v.OccupationalCategory) == 1 {
			value = v.OccupationalCategory[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["occupationalCategory"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Qualifications
		if len(v.Qualifications) == 1 {
			value = v.Qualifications[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["qualifications"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.RelevantOccupation
		if len(v.RelevantOccupation) == 1 {
			value = v.RelevantOccupation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["relevantOccupation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Responsibilities
		if len(v.Responsibilities) == 1 {
			value = v.Responsibilities[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["responsibilities"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SalaryCurrency
		if len(v.SalaryCurrency) == 1 {
			value = v.SalaryCurrency[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["salaryCurrency"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Skills
		if len(v.Skills) == 1 {
			value = v.Skills[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["skills"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SpecialCommitments
		if len(v.SpecialCommitments) == 1 {
			value = v.SpecialCommitments[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["specialCommitments"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Title
		if len(v.Title) == 1 {
			value = v.Title[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["title"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ValidThrough
		if len(v.ValidThrough) == 1 {
			value = v.ValidThrough[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["validThrough"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkHours
		if len(v.WorkHours) == 1 {
			value = v.WorkHours[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workHours"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v JobPosting) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "JobPosting"

	return data, nil
}

func (v JobPosting) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
