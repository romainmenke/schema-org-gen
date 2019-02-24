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
	// types : Number PriceSpecification
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
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

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

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
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
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
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
	// The currency (coded using ISO 4217, http://en.wikipedia.org/wiki/ISO_4217 ) used for the main salary information in this job posting or for this employee.
	// types : Text
	SalaryCurrency []string `json:"salaryCurrency,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
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
