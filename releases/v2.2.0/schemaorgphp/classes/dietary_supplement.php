<?php
namespace SchemaOrg;

// DietarySupplement see : https://schema.org/DietarySupplement
class DietarySupplement implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'DietarySupplement';

	/**
	 * With properties from MedicalTherapy see : https://schema.org/MedicalTherapy
	 */

	/**
	 * With properties from MedicalEntity see : https://schema.org/MedicalEntity
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An active ingredient, typically chemical compounds and/or biologic substances.
	 * see : https://schema.org/activeIngredient
	 *
	 * @var string | string[]
	 */
	public $active_ingredient;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * A possible complication and/or side effect of this therapy. If it is known that an adverse outcome is serious (resulting in death, disability, or permanent damage; requiring hospitalization; or is otherwise life-threatening or requires immediate medical attention), tag it as a seriouseAdverseOutcome instead.
	 * see : https://schema.org/adverseOutcome
	 *
	 * @var \MedicalEntity | \MedicalEntity[]
	 */
	public $adverse_outcome;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * Descriptive information establishing a historical perspective on the supplement. May include the rationale for the name, the population where the supplement first came to prominence, etc.
	 * see : https://schema.org/background
	 *
	 * @var string | string[]
	 */
	public $background;

	/**
	 * A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	 * see : https://schema.org/code
	 *
	 * @var \MedicalCode | \MedicalCode[]
	 */
	public $code;

	/**
	 * A contraindication for this therapy.
	 * see : https://schema.org/contraindication
	 *
	 * @var \MedicalContraindication | \MedicalContraindication[]
	 */
	public $contraindication;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A dosage form in which this drug/supplement is available, e.g. &#39;tablet&#39;, &#39;suspension&#39;, &#39;injection&#39;.
	 * see : https://schema.org/dosageForm
	 *
	 * @var string | string[]
	 */
	public $dosage_form;

	/**
	 * A therapy that duplicates or overlaps this one.
	 * see : https://schema.org/duplicateTherapy
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $duplicate_therapy;

	/**
	 * A medical guideline related to this entity.
	 * see : https://schema.org/guideline
	 *
	 * @var \MedicalGuideline | \MedicalGuideline[]
	 */
	public $guideline;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * A factor that indicates use of this therapy for treatment and/or prevention of a condition, symptom, etc. For therapies such as drugs, indications can include both officially-approved indications as well as off-label uses. These can be distinguished by using the ApprovedIndication subtype of MedicalIndication.
	 * see : https://schema.org/indication
	 *
	 * @var \MedicalIndication | \MedicalIndication[]
	 */
	public $indication;

	/**
	 * True if this item&#39;s name is a proprietary/brand name (vs. generic name).
	 * see : https://schema.org/isProprietary
	 *
	 * @var boolean | boolean[]
	 */
	public $is_proprietary;

	/**
	 * The drug or supplement&#39;s legal status, including any controlled substance schedules that apply.
	 * see : https://schema.org/legalStatus
	 *
	 * @var \DrugLegalStatus | \DrugLegalStatus[]
	 */
	public $legal_status;

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
	 * The manufacturer of the product.
	 * see : https://schema.org/manufacturer
	 *
	 * @var \Organization | \Organization[]
	 */
	public $manufacturer;

	/**
	 * Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	 * see : https://schema.org/maximumIntake
	 *
	 * @var \MaximumDoseSchedule | \MaximumDoseSchedule[]
	 */
	public $maximum_intake;

	/**
	 * The specific biochemical interaction through which this drug or supplement produces its pharmacological effect.
	 * see : https://schema.org/mechanismOfAction
	 *
	 * @var string | string[]
	 */
	public $mechanism_of_action;

	/**
	 * The system of medicine that includes this MedicalEntity, for example &#39;evidence-based&#39;, &#39;homeopathic&#39;, &#39;chiropractic&#39;, etc.
	 * see : https://schema.org/medicineSystem
	 *
	 * @var \MedicineSystem | \MedicineSystem[]
	 */
	public $medicine_system;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * The generic name of this drug or supplement.
	 * see : https://schema.org/nonProprietaryName
	 *
	 * @var string | string[]
	 */
	public $non_proprietary_name;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	 * see : https://schema.org/recognizingAuthority
	 *
	 * @var \Organization | \Organization[]
	 */
	public $recognizing_authority;

	/**
	 * Recommended intake of this supplement for a given population as defined by a specific recommending authority.
	 * see : https://schema.org/recommendedIntake
	 *
	 * @var \RecommendedDoseSchedule | \RecommendedDoseSchedule[]
	 */
	public $recommended_intake;

	/**
	 * If applicable, a medical specialty in which this entity is relevant.
	 * see : https://schema.org/relevantSpecialty
	 *
	 * @var \MedicalSpecialty | \MedicalSpecialty[]
	 */
	public $relevant_specialty;

	/**
	 * Any potential safety concern associated with the supplement. May include interactions with other drugs and foods, pregnancy, breastfeeding, known adverse reactions, and documented efficacy of the supplement.
	 * see : https://schema.org/safetyConsideration
	 *
	 * @var string | string[]
	 */
	public $safety_consideration;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * A possible serious complication and/or serious side effect of this therapy. Serious adverse outcomes include those that are life-threatening; result in death, disability, or permanent damage; require hospitalization or prolong existing hospitalization; cause congenital anomalies or birth defects; or jeopardize the patient and may require medical or surgical intervention to prevent one of the outcomes in this definition.
	 * see : https://schema.org/seriousAdverseOutcome
	 *
	 * @var \MedicalEntity | \MedicalEntity[]
	 */
	public $serious_adverse_outcome;

	/**
	 * A medical study or trial related to this entity.
	 * see : https://schema.org/study
	 *
	 * @var \MedicalStudy | \MedicalStudy[]
	 */
	public $study;

	/**
	 * Characteristics of the population for which this is intended, or which typically uses it, e.g. &#39;adults&#39;.
	 * see : https://schema.org/targetPopulation
	 *
	 * @var string | string[]
	 */
	public $target_population;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'DietarySupplement',
		);

		$serialized = \SchemaOrg\json_serialize( $this->active_ingredient );
		if ( ! empty( $serialized ) ) {
			$out['activeIngredient'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->adverse_outcome );
		if ( ! empty( $serialized ) ) {
			$out['adverseOutcome'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->background );
		if ( ! empty( $serialized ) ) {
			$out['background'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->code );
		if ( ! empty( $serialized ) ) {
			$out['code'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contraindication );
		if ( ! empty( $serialized ) ) {
			$out['contraindication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->dosage_form );
		if ( ! empty( $serialized ) ) {
			$out['dosageForm'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->duplicate_therapy );
		if ( ! empty( $serialized ) ) {
			$out['duplicateTherapy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->guideline );
		if ( ! empty( $serialized ) ) {
			$out['guideline'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->indication );
		if ( ! empty( $serialized ) ) {
			$out['indication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_proprietary );
		if ( ! empty( $serialized ) ) {
			$out['isProprietary'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->legal_status );
		if ( ! empty( $serialized ) ) {
			$out['legalStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->manufacturer );
		if ( ! empty( $serialized ) ) {
			$out['manufacturer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->maximum_intake );
		if ( ! empty( $serialized ) ) {
			$out['maximumIntake'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->mechanism_of_action );
		if ( ! empty( $serialized ) ) {
			$out['mechanismOfAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->medicine_system );
		if ( ! empty( $serialized ) ) {
			$out['medicineSystem'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->non_proprietary_name );
		if ( ! empty( $serialized ) ) {
			$out['nonProprietaryName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recognizing_authority );
		if ( ! empty( $serialized ) ) {
			$out['recognizingAuthority'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recommended_intake );
		if ( ! empty( $serialized ) ) {
			$out['recommendedIntake'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->relevant_specialty );
		if ( ! empty( $serialized ) ) {
			$out['relevantSpecialty'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->safety_consideration );
		if ( ! empty( $serialized ) ) {
			$out['safetyConsideration'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->serious_adverse_outcome );
		if ( ! empty( $serialized ) ) {
			$out['seriousAdverseOutcome'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->study );
		if ( ! empty( $serialized ) ) {
			$out['study'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->target_population );
		if ( ! empty( $serialized ) ) {
			$out['targetPopulation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
