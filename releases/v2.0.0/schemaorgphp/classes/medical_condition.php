<?php
namespace SchemaOrg;

// MedicalCondition see : https://schema.org/MedicalCondition
class MedicalCondition implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'MedicalCondition';

	/**
	 * With properties from MedicalEntity see : https://schema.org/MedicalEntity
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
	 * The anatomy of the underlying organ system or structures associated with this entity.
	 * see : https://schema.org/associatedAnatomy
	 *
	 * @var \AnatomicalStructure | \AnatomicalStructure[] | \AnatomicalSystem | \AnatomicalSystem[] | \SuperficialAnatomy | \SuperficialAnatomy[]
	 */
	public $associated_anatomy;

	/**
	 * An underlying cause. More specifically, one of the causative agent(s) that are most directly responsible for the pathophysiologic process that eventually results in the occurrence.
	 * see : https://schema.org/cause
	 *
	 * @var \MedicalCause | \MedicalCause[]
	 */
	public $cause;

	/**
	 * A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	 * see : https://schema.org/code
	 *
	 * @var \MedicalCode | \MedicalCode[]
	 */
	public $code;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * One of a set of differential diagnoses for the condition. Specifically, a closely-related or competing diagnosis typically considered later in the cognitive process whereby this medical condition is distinguished from others most likely responsible for a similar collection of signs and symptoms to reach the most parsimonious diagnosis or diagnoses in a patient.
	 * see : https://schema.org/differentialDiagnosis
	 *
	 * @var \DDxElement | \DDxElement[]
	 */
	public $differential_diagnosis;

	/**
	 * The characteristics of associated patients, such as age, gender, race etc.
	 * see : https://schema.org/epidemiology
	 *
	 * @var string | string[]
	 */
	public $epidemiology;

	/**
	 * The likely outcome in either the short term or long term of the medical condition.
	 * see : https://schema.org/expectedPrognosis
	 *
	 * @var string | string[]
	 */
	public $expected_prognosis;

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
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	 *       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	 *       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	 *       between the page and the primary entity.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       Related properties include sameAs, about, and url.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	 *       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	 *       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	 *       serves more to clarify which of several entities is the main one for that page.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	 *       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	 *       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	 *       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	 *       describes some other entity. For example, one web page may display a news article about a particular person.
	 *       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	 *       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

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
	 * The expected progression of the condition if it is not treated and allowed to progress naturally.
	 * see : https://schema.org/naturalProgression
	 *
	 * @var string | string[]
	 */
	public $natural_progression;

	/**
	 * Changes in the normal mechanical, physical, and biochemical functions that are associated with this activity or condition.
	 * see : https://schema.org/pathophysiology
	 *
	 * @var string | string[]
	 */
	public $pathophysiology;

	/**
	 * A possible unexpected and unfavorable evolution of a medical condition. Complications may include worsening of the signs or symptoms of the disease, extension of the condition to other organ systems, etc.
	 * see : https://schema.org/possibleComplication
	 *
	 * @var string | string[]
	 */
	public $possible_complication;

	/**
	 * A possible treatment to address this condition, sign or symptom.
	 * see : https://schema.org/possibleTreatment
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $possible_treatment;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * A preventative therapy used to prevent an initial occurrence of the medical condition, such as vaccination.
	 * see : https://schema.org/primaryPrevention
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $primary_prevention;

	/**
	 * If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	 * see : https://schema.org/recognizingAuthority
	 *
	 * @var \Organization | \Organization[]
	 */
	public $recognizing_authority;

	/**
	 * If applicable, a medical specialty in which this entity is relevant.
	 * see : https://schema.org/relevantSpecialty
	 *
	 * @var \MedicalSpecialty | \MedicalSpecialty[]
	 */
	public $relevant_specialty;

	/**
	 * A modifiable or non-modifiable factor that increases the risk of a patient contracting this condition, e.g. age,  coexisting condition.
	 * see : https://schema.org/riskFactor
	 *
	 * @var \MedicalRiskFactor | \MedicalRiskFactor[]
	 */
	public $risk_factor;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * A preventative therapy used to prevent reoccurrence of the medical condition after an initial episode of the condition.
	 * see : https://schema.org/secondaryPrevention
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $secondary_prevention;

	/**
	 * A sign or symptom of this condition. Signs are objective or physically observable manifestations of the medical condition while symptoms are the subjective experience of the medical condition.
	 * see : https://schema.org/signOrSymptom
	 *
	 * @var \MedicalSignOrSymptom | \MedicalSignOrSymptom[]
	 */
	public $sign_or_symptom;

	/**
	 * The stage of the condition, if applicable.
	 * see : https://schema.org/stage
	 *
	 * @var \MedicalConditionStage | \MedicalConditionStage[]
	 */
	public $stage;

	/**
	 * A medical study or trial related to this entity.
	 * see : https://schema.org/study
	 *
	 * @var \MedicalStudy | \MedicalStudy[]
	 */
	public $study;

	/**
	 * A more specific type of the condition, where applicable, for example &#39;Type 1 Diabetes&#39;, &#39;Type 2 Diabetes&#39;, or &#39;Gestational Diabetes&#39; for Diabetes.
	 * see : https://schema.org/subtype
	 *
	 * @var string | string[]
	 */
	public $subtype;

	/**
	 * A medical test typically performed given this condition.
	 * see : https://schema.org/typicalTest
	 *
	 * @var \MedicalTest | \MedicalTest[]
	 */
	public $typical_test;

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
			'@type'    => 'MedicalCondition',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->associated_anatomy );
		if ( ! empty( $serialized ) ) {
			$out['associatedAnatomy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->cause );
		if ( ! empty( $serialized ) ) {
			$out['cause'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->code );
		if ( ! empty( $serialized ) ) {
			$out['code'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->differential_diagnosis );
		if ( ! empty( $serialized ) ) {
			$out['differentialDiagnosis'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->epidemiology );
		if ( ! empty( $serialized ) ) {
			$out['epidemiology'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->expected_prognosis );
		if ( ! empty( $serialized ) ) {
			$out['expectedPrognosis'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->guideline );
		if ( ! empty( $serialized ) ) {
			$out['guideline'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->medicine_system );
		if ( ! empty( $serialized ) ) {
			$out['medicineSystem'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->natural_progression );
		if ( ! empty( $serialized ) ) {
			$out['naturalProgression'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->pathophysiology );
		if ( ! empty( $serialized ) ) {
			$out['pathophysiology'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->possible_complication );
		if ( ! empty( $serialized ) ) {
			$out['possibleComplication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->possible_treatment );
		if ( ! empty( $serialized ) ) {
			$out['possibleTreatment'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->primary_prevention );
		if ( ! empty( $serialized ) ) {
			$out['primaryPrevention'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recognizing_authority );
		if ( ! empty( $serialized ) ) {
			$out['recognizingAuthority'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->relevant_specialty );
		if ( ! empty( $serialized ) ) {
			$out['relevantSpecialty'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->risk_factor );
		if ( ! empty( $serialized ) ) {
			$out['riskFactor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->secondary_prevention );
		if ( ! empty( $serialized ) ) {
			$out['secondaryPrevention'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sign_or_symptom );
		if ( ! empty( $serialized ) ) {
			$out['signOrSymptom'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->stage );
		if ( ! empty( $serialized ) ) {
			$out['stage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->study );
		if ( ! empty( $serialized ) ) {
			$out['study'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->subtype );
		if ( ! empty( $serialized ) ) {
			$out['subtype'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->typical_test );
		if ( ! empty( $serialized ) ) {
			$out['typicalTest'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
