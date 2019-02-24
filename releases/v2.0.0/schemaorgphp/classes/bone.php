<?php
namespace SchemaOrg;

// Bone see : https://schema.org/Bone
class Bone implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Bone';

	/**
	 * With properties from AnatomicalStructure see : https://schema.org/AnatomicalStructure
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
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * If applicable, a description of the pathophysiology associated with the anatomical system, including potential abnormal changes in the mechanical, physical, and biochemical functions of the system.
	 * see : https://schema.org/associatedPathophysiology
	 *
	 * @var string | string[]
	 */
	public $associated_pathophysiology;

	/**
	 * Location in the body of the anatomical structure.
	 * see : https://schema.org/bodyLocation
	 *
	 * @var string | string[]
	 */
	public $body_location;

	/**
	 * A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	 * see : https://schema.org/code
	 *
	 * @var \MedicalCode | \MedicalCode[]
	 */
	public $code;

	/**
	 * Other anatomical structures to which this structure is connected.
	 * see : https://schema.org/connectedTo
	 *
	 * @var \AnatomicalStructure | \AnatomicalStructure[]
	 */
	public $connected_to;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * An image containing a diagram that illustrates the structure and/or its component substructures and/or connections with other structures.
	 * see : https://schema.org/diagram
	 *
	 * @var \ImageObject | \ImageObject[]
	 */
	public $diagram;

	/**
	 * Function of the anatomical structure.
	 * see : https://schema.org/function
	 *
	 * @var string | string[]
	 */
	public $function;

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
	 * The anatomical or organ system that this structure is part of.
	 * see : https://schema.org/partOfSystem
	 *
	 * @var \AnatomicalSystem | \AnatomicalSystem[]
	 */
	public $part_of_system;

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
	 * A medical condition associated with this anatomy.
	 * see : https://schema.org/relatedCondition
	 *
	 * @var \MedicalCondition | \MedicalCondition[]
	 */
	public $related_condition;

	/**
	 * A medical therapy related to this anatomy.
	 * see : https://schema.org/relatedTherapy
	 *
	 * @var \MedicalTherapy | \MedicalTherapy[]
	 */
	public $related_therapy;

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
	 * A medical study or trial related to this entity.
	 * see : https://schema.org/study
	 *
	 * @var \MedicalStudy | \MedicalStudy[]
	 */
	public $study;

	/**
	 * Component (sub-)structure(s) that comprise this anatomical structure.
	 * see : https://schema.org/subStructure
	 *
	 * @var \AnatomicalStructure | \AnatomicalStructure[]
	 */
	public $sub_structure;

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
			'@type'    => 'Bone',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->associated_pathophysiology );
		if ( ! empty( $serialized ) ) {
			$out['associatedPathophysiology'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->body_location );
		if ( ! empty( $serialized ) ) {
			$out['bodyLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->code );
		if ( ! empty( $serialized ) ) {
			$out['code'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->connected_to );
		if ( ! empty( $serialized ) ) {
			$out['connectedTo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->diagram );
		if ( ! empty( $serialized ) ) {
			$out['diagram'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->function );
		if ( ! empty( $serialized ) ) {
			$out['function'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->part_of_system );
		if ( ! empty( $serialized ) ) {
			$out['partOfSystem'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recognizing_authority );
		if ( ! empty( $serialized ) ) {
			$out['recognizingAuthority'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->related_condition );
		if ( ! empty( $serialized ) ) {
			$out['relatedCondition'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->related_therapy );
		if ( ! empty( $serialized ) ) {
			$out['relatedTherapy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->relevant_specialty );
		if ( ! empty( $serialized ) ) {
			$out['relevantSpecialty'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->study );
		if ( ! empty( $serialized ) ) {
			$out['study'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sub_structure );
		if ( ! empty( $serialized ) ) {
			$out['subStructure'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
