<?php

class JobPosting extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'JobPosting';
	
	/**
	 * The base salary of the job or of an employee in an EmployeeRole.
	 * see : https://schema.org/baseSalary
	 * @var \MonetaryAmount|\MonetaryAmount[]|float|float[]|\PriceSpecification|\PriceSpecification[]
	 */
	public var $base_salary;
	
	/**
	 * Publication date for the job posting.
	 * see : https://schema.org/datePosted
	 * @var string|string[]
	 */
	public var $date_posted;
	
	/**
	 * Educational background needed for the position or Occupation.
	 * see : https://pending.schema.org/educationRequirements
	 * @var string|string[]
	 */
	public var $education_requirements;
	
	/**
	 * Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	 * see : https://schema.org/employmentType
	 * @var string|string[]
	 */
	public var $employment_type;
	
	/**
	 * A property describing the estimated salary for a job posting based on a variety of variables including, but not limited to industry, job title, and location. The estimated salary is usually computed by outside organizations and therefore the hiring organization is not bound to this estimated salary.
	 * see : https://pending.schema.org/estimatedSalary
	 * @var \MonetaryAmount|\MonetaryAmount[]|\MonetaryAmountDistribution|\MonetaryAmountDistribution[]|float|float[]|\PriceSpecification|\PriceSpecification[]
	 */
	public var $estimated_salary;
	
	/**
	 * Description of skills and experience needed for the position or Occupation.
	 * see : https://pending.schema.org/experienceRequirements
	 * @var string|string[]
	 */
	public var $experience_requirements;
	
	/**
	 * Organization offering the job position.
	 * see : https://schema.org/hiringOrganization
	 * @var \Organization|\Organization[]
	 */
	public var $hiring_organization;
	
	/**
	 * Description of bonus and commission compensation aspects of the job. Supersedes incentives (see: https://schema.org/incentives).
	 * see : https://schema.org/incentiveCompensation
	 * @var string|string[]
	 */
	public var $incentive_compensation;
	
	/**
	 * The industry associated with the job position.
	 * see : https://schema.org/industry
	 * @var string|string[]
	 */
	public var $industry;
	
	/**
	 * Description of benefits associated with the job. Supersedes benefits (see: https://schema.org/benefits).
	 * see : https://schema.org/jobBenefits
	 * @var string|string[]
	 */
	public var $job_benefits;
	
	/**
	 * A (typically single) geographic location associated with the job position.
	 * see : https://schema.org/jobLocation
	 * @var \Place|\Place[]
	 */
	public var $job_location;
	
	/**
	 * Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	 * see : https://pending.schema.org/occupationalCategory
	 * @var string|string[]
	 */
	public var $occupational_category;
	
	/**
	 * Specific qualifications required for this role or Occupation.
	 * see : https://pending.schema.org/qualifications
	 * @var string|string[]
	 */
	public var $qualifications;
	
	/**
	 * The Occupation for the JobPosting.
	 * see : https://pending.schema.org/relevantOccupation
	 * @var \Occupation|\Occupation[]
	 */
	public var $relevant_occupation;
	
	/**
	 * Responsibilities associated with this role or Occupation.
	 * see : https://pending.schema.org/responsibilities
	 * @var string|string[]
	 */
	public var $responsibilities;
	
	/**
	 * The currency (coded using ISO 4217 (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) ) used for the main salary information in this job posting or for this employee.
	 * see : https://schema.org/salaryCurrency
	 * @var string|string[]
	 */
	public var $salary_currency;
	
	/**
	 * Skills required to fulfill this role.
	 * see : https://pending.schema.org/skills
	 * @var string|string[]
	 */
	public var $skills;
	
	/**
	 * Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
	 * see : https://schema.org/specialCommitments
	 * @var string|string[]
	 */
	public var $special_commitments;
	
	/**
	 * The title of the job.
	 * see : https://schema.org/title
	 * @var string|string[]
	 */
	public var $title;
	
	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 * @var string|string[]
	 */
	public var $valid_through;
	
	/**
	 * The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	 * see : https://schema.org/workHours
	 * @var string|string[]
	 */
	public var $work_hours;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'JobPosting'
		);
		
		$serialized = so_json_serialize( $this->base_salary );
		if ( ! empty( $serialized ) ) {
			$out['baseSalary'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_posted );
		if ( ! empty( $serialized ) ) {
			$out['datePosted'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->education_requirements );
		if ( ! empty( $serialized ) ) {
			$out['educationRequirements'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->employment_type );
		if ( ! empty( $serialized ) ) {
			$out['employmentType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->estimated_salary );
		if ( ! empty( $serialized ) ) {
			$out['estimatedSalary'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->experience_requirements );
		if ( ! empty( $serialized ) ) {
			$out['experienceRequirements'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->hiring_organization );
		if ( ! empty( $serialized ) ) {
			$out['hiringOrganization'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->incentive_compensation );
		if ( ! empty( $serialized ) ) {
			$out['incentiveCompensation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->industry );
		if ( ! empty( $serialized ) ) {
			$out['industry'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->job_benefits );
		if ( ! empty( $serialized ) ) {
			$out['jobBenefits'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->job_location );
		if ( ! empty( $serialized ) ) {
			$out['jobLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->occupational_category );
		if ( ! empty( $serialized ) ) {
			$out['occupationalCategory'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->qualifications );
		if ( ! empty( $serialized ) ) {
			$out['qualifications'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->relevant_occupation );
		if ( ! empty( $serialized ) ) {
			$out['relevantOccupation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->responsibilities );
		if ( ! empty( $serialized ) ) {
			$out['responsibilities'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->salary_currency );
		if ( ! empty( $serialized ) ) {
			$out['salaryCurrency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->skills );
		if ( ! empty( $serialized ) ) {
			$out['skills'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->special_commitments );
		if ( ! empty( $serialized ) ) {
			$out['specialCommitments'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->title );
		if ( ! empty( $serialized ) ) {
			$out['title'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->work_hours );
		if ( ! empty( $serialized ) ) {
			$out['workHours'] = $serialized;
		}
		
		return $out;
	}
}
