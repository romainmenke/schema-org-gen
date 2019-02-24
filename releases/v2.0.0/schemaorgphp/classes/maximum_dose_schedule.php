<?php
namespace SchemaOrg;

// MaximumDoseSchedule see : https://schema.org/MaximumDoseSchedule
class MaximumDoseSchedule implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'MaximumDoseSchedule';

	/**
	 * With properties from DoseSchedule see : https://schema.org/DoseSchedule
	 */

	/**
	 * With properties from MedicalIntangible see : https://schema.org/MedicalIntangible
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
	 * The unit of the dose, e.g. &#39;mg&#39;.
	 * see : https://schema.org/doseUnit
	 *
	 * @var string | string[]
	 */
	public $dose_unit;

	/**
	 * The value of the dose, e.g. 500.
	 * see : https://schema.org/doseValue
	 *
	 * @var float | float[]
	 */
	public $dose_value;

	/**
	 * How often the dose is taken, e.g. &#39;daily&#39;.
	 * see : https://schema.org/frequency
	 *
	 * @var string | string[]
	 */
	public $frequency;

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
			'@type'    => 'MaximumDoseSchedule',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->code );
		if ( ! empty( $serialized ) ) {
			$out['code'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->dose_unit );
		if ( ! empty( $serialized ) ) {
			$out['doseUnit'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->dose_value );
		if ( ! empty( $serialized ) ) {
			$out['doseValue'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->frequency );
		if ( ! empty( $serialized ) ) {
			$out['frequency'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
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
