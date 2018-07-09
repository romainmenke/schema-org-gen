<?php

namespace SchemaOrg;

// JobPosting see : https://schema.org/JobPosting
class JobPosting implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'JobPosting';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
	/**
	 * The base salary of the job or of an employee in an EmployeeRole.
	 * see : https://schema.org/baseSalary
	 * @var \MonetaryAmount | \MonetaryAmount[] | float | float[] | \PriceSpecification | \PriceSpecification[]
	 */
	public var $base_salary;
	
	/**
	 * Publication date for the job posting.
	 * see : https://schema.org/datePosted
	 * @var string | string[]
	 */
	public var $date_posted;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public var $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public var $disambiguating_description;
	
	/**
	 * Educational background needed for the position or Occupation.
	 * see : https://pending.schema.org/educationRequirements
	 * @var string | string[]
	 */
	public var $education_requirements;
	
	/**
	 * Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	 * see : https://schema.org/employmentType
	 * @var string | string[]
	 */
	public var $employment_type;
	
	/**
	 * A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
	 * see : https://pending.schema.org/estimatedSalary
	 * @var \MonetaryAmount | \MonetaryAmount[] | \MonetaryAmountDistribution | \MonetaryAmountDistribution[] | float | float[] | \PriceSpecification | \PriceSpecification[]
	 */
	public var $estimated_salary;
	
	/**
	 * Description of skills and experience needed for the position or Occupation.
	 * see : https://pending.schema.org/experienceRequirements
	 * @var string | string[]
	 */
	public var $experience_requirements;
	
	/**
	 * Organization offering the job position.
	 * see : https://schema.org/hiringOrganization
	 * @var \Organization | \Organization[]
	 */
	public var $hiring_organization;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public var $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $image;
	
	/**
	 * Description of bonus and commission compensation aspects of the job. Supersedes incentives (see: https://schema.org/incentives).
	 * see : https://schema.org/incentiveCompensation
	 * @var string | string[]
	 */
	public var $incentive_compensation;
	
	/**
	 * The industry associated with the job position.
	 * see : https://schema.org/industry
	 * @var string | string[]
	 */
	public var $industry;
	
	/**
	 * Description of benefits associated with the job. Supersedes benefits (see: https://schema.org/benefits).
	 * see : https://schema.org/jobBenefits
	 * @var string | string[]
	 */
	public var $job_benefits;
	
	/**
	 * A (typically single) geographic location associated with the job position.
	 * see : https://schema.org/jobLocation
	 * @var \Place | \Place[]
	 */
	public var $job_location;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	 * see : https://pending.schema.org/occupationalCategory
	 * @var string | string[]
	 */
	public var $occupational_category;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * Specific qualifications required for this role or Occupation.
	 * see : https://pending.schema.org/qualifications
	 * @var string | string[]
	 */
	public var $qualifications;
	
	/**
	 * The Occupation for the JobPosting.
	 * see : https://pending.schema.org/relevantOccupation
	 * @var \Occupation | \Occupation[]
	 */
	public var $relevant_occupation;
	
	/**
	 * Responsibilities associated with this role or Occupation.
	 * see : https://pending.schema.org/responsibilities
	 * @var string | string[]
	 */
	public var $responsibilities;
	
	/**
	 * The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	 * see : https://schema.org/salaryCurrency
	 * @var string | string[]
	 */
	public var $salary_currency;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * Skills required to fulfill this role.
	 * see : https://pending.schema.org/skills
	 * @var string | string[]
	 */
	public var $skills;
	
	/**
	 * Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
	 * see : https://schema.org/specialCommitments
	 * @var string | string[]
	 */
	public var $special_commitments;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * The title of the job.
	 * see : https://schema.org/title
	 * @var string | string[]
	 */
	public var $title;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 * @var string | string[]
	 */
	public var $valid_through;
	
	/**
	 * The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	 * see : https://schema.org/workHours
	 * @var string | string[]
	 */
	public var $work_hours;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'JobPosting'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->base_salary );
		if ( ! empty( $serialized ) ) {
			$out['baseSalary'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->date_posted );
		if ( ! empty( $serialized ) ) {
			$out['datePosted'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->education_requirements );
		if ( ! empty( $serialized ) ) {
			$out['educationRequirements'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->employment_type );
		if ( ! empty( $serialized ) ) {
			$out['employmentType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->estimated_salary );
		if ( ! empty( $serialized ) ) {
			$out['estimatedSalary'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->experience_requirements );
		if ( ! empty( $serialized ) ) {
			$out['experienceRequirements'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->hiring_organization );
		if ( ! empty( $serialized ) ) {
			$out['hiringOrganization'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->incentive_compensation );
		if ( ! empty( $serialized ) ) {
			$out['incentiveCompensation'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->industry );
		if ( ! empty( $serialized ) ) {
			$out['industry'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->job_benefits );
		if ( ! empty( $serialized ) ) {
			$out['jobBenefits'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->job_location );
		if ( ! empty( $serialized ) ) {
			$out['jobLocation'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->occupational_category );
		if ( ! empty( $serialized ) ) {
			$out['occupationalCategory'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->qualifications );
		if ( ! empty( $serialized ) ) {
			$out['qualifications'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->relevant_occupation );
		if ( ! empty( $serialized ) ) {
			$out['relevantOccupation'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->responsibilities );
		if ( ! empty( $serialized ) ) {
			$out['responsibilities'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->salary_currency );
		if ( ! empty( $serialized ) ) {
			$out['salaryCurrency'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->skills );
		if ( ! empty( $serialized ) ) {
			$out['skills'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->special_commitments );
		if ( ! empty( $serialized ) ) {
			$out['specialCommitments'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->title );
		if ( ! empty( $serialized ) ) {
			$out['title'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->work_hours );
		if ( ! empty( $serialized ) ) {
			$out['workHours'] = $serialized;
		}
		
		return $out;
	}
}
