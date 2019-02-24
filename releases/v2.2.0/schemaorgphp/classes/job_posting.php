<?php
namespace SchemaOrg;

// JobPosting see : https://schema.org/JobPosting
class JobPosting implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'JobPosting';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * The base salary of the job or of an employee in an EmployeeRole.
	 * see : https://schema.org/baseSalary
	 *
	 * @var float | float[] | \PriceSpecification | \PriceSpecification[]
	 */
	public $base_salary;

	/**
	 * Description of benefits associated with the job.
	 * see : https://schema.org/benefits
	 *
	 * @var string | string[]
	 */
	public $benefits;

	/**
	 * Publication date for the job posting.
	 * see : https://schema.org/datePosted
	 *
	 * @var string | string[]
	 */
	public $date_posted;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * Educational background needed for the position.
	 * see : https://schema.org/educationRequirements
	 *
	 * @var string | string[]
	 */
	public $education_requirements;

	/**
	 * Type of employment (e.g. full-time, part-time, contract, temporary, seasonal, internship).
	 * see : https://schema.org/employmentType
	 *
	 * @var string | string[]
	 */
	public $employment_type;

	/**
	 * Description of skills and experience needed for the position.
	 * see : https://schema.org/experienceRequirements
	 *
	 * @var string | string[]
	 */
	public $experience_requirements;

	/**
	 * Organization offering the job position.
	 * see : https://schema.org/hiringOrganization
	 *
	 * @var \Organization | \Organization[]
	 */
	public $hiring_organization;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Description of bonus and commission compensation aspects of the job.
	 * see : https://schema.org/incentiveCompensation
	 *
	 * @var string | string[]
	 */
	public $incentive_compensation;

	/**
	 * Description of bonus and commission compensation aspects of the job.
	 * see : https://schema.org/incentives
	 *
	 * @var string | string[]
	 */
	public $incentives;

	/**
	 * The industry associated with the job position.
	 * see : https://schema.org/industry
	 *
	 * @var string | string[]
	 */
	public $industry;

	/**
	 * Description of benefits associated with the job.
	 * see : https://schema.org/jobBenefits
	 *
	 * @var string | string[]
	 */
	public $job_benefits;

	/**
	 * A (typically single) geographic location associated with the job position.
	 * see : https://schema.org/jobLocation
	 *
	 * @var \Place | \Place[]
	 */
	public $job_location;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * Category or categories describing the job. Use BLS O*NET-SOC taxonomy: http://www.onetcenter.org/taxonomy.html. Ideally includes textual label and formal code, with the property repeated for each applicable value.
	 * see : https://schema.org/occupationalCategory
	 *
	 * @var string | string[]
	 */
	public $occupational_category;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * Specific qualifications required for this role.
	 * see : https://schema.org/qualifications
	 *
	 * @var string | string[]
	 */
	public $qualifications;

	/**
	 * Responsibilities associated with this role.
	 * see : https://schema.org/responsibilities
	 *
	 * @var string | string[]
	 */
	public $responsibilities;

	/**
	 * The currency (coded using ISO 4217, http://en.wikipedia.org/wiki/ISO_4217 ) used for the main salary information in this job posting or for this employee.
	 * see : https://schema.org/salaryCurrency
	 *
	 * @var string | string[]
	 */
	public $salary_currency;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * Skills required to fulfill this role.
	 * see : https://schema.org/skills
	 *
	 * @var string | string[]
	 */
	public $skills;

	/**
	 * Any special commitments associated with this job posting. Valid entries include VeteranCommit, MilitarySpouseCommit, etc.
	 * see : https://schema.org/specialCommitments
	 *
	 * @var string | string[]
	 */
	public $special_commitments;

	/**
	 * The title of the job.
	 * see : https://schema.org/title
	 *
	 * @var string | string[]
	 */
	public $title;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The typical working hours for this job (e.g. 1st shift, night shift, 8am-5pm).
	 * see : https://schema.org/workHours
	 *
	 * @var string | string[]
	 */
	public $work_hours;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'JobPosting',
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

		$serialized = \SchemaOrg\json_serialize( $this->benefits );
		if ( ! empty( $serialized ) ) {
			$out['benefits'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->date_posted );
		if ( ! empty( $serialized ) ) {
			$out['datePosted'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->education_requirements );
		if ( ! empty( $serialized ) ) {
			$out['educationRequirements'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->employment_type );
		if ( ! empty( $serialized ) ) {
			$out['employmentType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->experience_requirements );
		if ( ! empty( $serialized ) ) {
			$out['experienceRequirements'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->hiring_organization );
		if ( ! empty( $serialized ) ) {
			$out['hiringOrganization'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->incentive_compensation );
		if ( ! empty( $serialized ) ) {
			$out['incentiveCompensation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->incentives );
		if ( ! empty( $serialized ) ) {
			$out['incentives'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->title );
		if ( ! empty( $serialized ) ) {
			$out['title'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->work_hours );
		if ( ! empty( $serialized ) ) {
			$out['workHours'] = $serialized;
		}

		return $out;
	}
}
