<?php

namespace SchemaOrg;

// NutritionInformation see : https://schema.org/NutritionInformation
class NutritionInformation implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'NutritionInformation';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from StructuredValue see : https://schema.org/StructuredValue
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public $additional_type;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public $alternate_name;
	
	/**
	 * The number of calories.
	 * see : https://schema.org/calories
	 * @var \Energy | \Energy[]
	 */
	public $calories;
	
	/**
	 * The number of grams of carbohydrates.
	 * see : https://schema.org/carbohydrateContent
	 * @var \Mass | \Mass[]
	 */
	public $carbohydrate_content;
	
	/**
	 * The number of milligrams of cholesterol.
	 * see : https://schema.org/cholesterolContent
	 * @var \Mass | \Mass[]
	 */
	public $cholesterol_content;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public $disambiguating_description;
	
	/**
	 * The number of grams of fat.
	 * see : https://schema.org/fatContent
	 * @var \Mass | \Mass[]
	 */
	public $fat_content;
	
	/**
	 * The number of grams of fiber.
	 * see : https://schema.org/fiberContent
	 * @var \Mass | \Mass[]
	 */
	public $fiber_content;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $image;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public $name;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public $potential_action;
	
	/**
	 * The number of grams of protein.
	 * see : https://schema.org/proteinContent
	 * @var \Mass | \Mass[]
	 */
	public $protein_content;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public $same_as;
	
	/**
	 * The number of grams of saturated fat.
	 * see : https://schema.org/saturatedFatContent
	 * @var \Mass | \Mass[]
	 */
	public $saturated_fat_content;
	
	/**
	 * The serving size, in terms of the number of volume or mass.
	 * see : https://schema.org/servingSize
	 * @var string | string[]
	 */
	public $serving_size;
	
	/**
	 * The number of milligrams of sodium.
	 * see : https://schema.org/sodiumContent
	 * @var \Mass | \Mass[]
	 */
	public $sodium_content;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;
	
	/**
	 * The number of grams of sugar.
	 * see : https://schema.org/sugarContent
	 * @var \Mass | \Mass[]
	 */
	public $sugar_content;
	
	/**
	 * The number of grams of trans fat.
	 * see : https://schema.org/transFatContent
	 * @var \Mass | \Mass[]
	 */
	public $trans_fat_content;
	
	/**
	 * The number of grams of unsaturated fat.
	 * see : https://schema.org/unsaturatedFatContent
	 * @var \Mass | \Mass[]
	 */
	public $unsaturated_fat_content;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'NutritionInformation'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->calories );
		if ( ! empty( $serialized ) ) {
			$out['calories'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->carbohydrate_content );
		if ( ! empty( $serialized ) ) {
			$out['carbohydrateContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->cholesterol_content );
		if ( ! empty( $serialized ) ) {
			$out['cholesterolContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fat_content );
		if ( ! empty( $serialized ) ) {
			$out['fatContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fiber_content );
		if ( ! empty( $serialized ) ) {
			$out['fiberContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->protein_content );
		if ( ! empty( $serialized ) ) {
			$out['proteinContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->saturated_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['saturatedFatContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->serving_size );
		if ( ! empty( $serialized ) ) {
			$out['servingSize'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->sodium_content );
		if ( ! empty( $serialized ) ) {
			$out['sodiumContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->sugar_content );
		if ( ! empty( $serialized ) ) {
			$out['sugarContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->trans_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['transFatContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->unsaturated_fat_content );
		if ( ! empty( $serialized ) ) {
			$out['unsaturatedFatContent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
