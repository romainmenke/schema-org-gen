<?php
namespace SchemaOrg;

// PalliativeProcedure see : https://schema.org/PalliativeProcedure
class PalliativeProcedure implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'PalliativeProcedure';

	/**
	 * With properties from MedicalProcedure see : https://schema.org/MedicalProcedure
	 */

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
	 * With properties from MedicalEntity see : https://schema.org/MedicalEntity
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
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
	 * A therapy that duplicates or overlaps this one.
	 * see : https://schema.org/duplicateTherapy
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $duplicate_therapy;

	/**
	 * Typical or recommended followup care after the procedure is performed.
	 * see : https://schema.org/followup
	 *
	 * @var string | string[]
	 */
	public $followup;

	/**
	 * A medical guideline related to this entity.
	 * see : https://schema.org/guideline
	 *
	 * @var \MedicalGuideline | \MedicalGuideline[]
	 */
	public $guideline;

	/**
	 * How the procedure is performed.
	 * see : https://schema.org/howPerformed
	 *
	 * @var string | string[]
	 */
	public $how_performed;

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
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * Typical preparation that a patient must undergo before having the procedure performed.
	 * see : https://schema.org/preparation
	 *
	 * @var string | string[]
	 */
	public $preparation;

	/**
	 * The type of procedure, for example Surgical, Noninvasive, or Percutaneous.
	 * see : https://schema.org/procedureType
	 *
	 * @var \MedicalProcedureType | \MedicalProcedureType[]
	 */
	public $procedure_type;

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
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'PalliativeProcedure',
		);

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

		$serialized = \SchemaOrg\json_serialize( $this->duplicate_therapy );
		if ( ! empty( $serialized ) ) {
			$out['duplicateTherapy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->followup );
		if ( ! empty( $serialized ) ) {
			$out['followup'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->guideline );
		if ( ! empty( $serialized ) ) {
			$out['guideline'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->how_performed );
		if ( ! empty( $serialized ) ) {
			$out['howPerformed'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->indication );
		if ( ! empty( $serialized ) ) {
			$out['indication'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->preparation );
		if ( ! empty( $serialized ) ) {
			$out['preparation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->procedure_type );
		if ( ! empty( $serialized ) ) {
			$out['procedureType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recognizing_authority );
		if ( ! empty( $serialized ) ) {
			$out['recognizingAuthority'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->relevant_specialty );
		if ( ! empty( $serialized ) ) {
			$out['relevantSpecialty'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
