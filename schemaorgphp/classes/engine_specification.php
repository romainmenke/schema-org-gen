<?php

namespace SchemaOrg;

// EngineSpecification see : https://schema.org/EngineSpecification
class EngineSpecification implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EngineSpecification';
	
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
	 * The volume swept by all of the pistons inside the cylinders of an internal combustion engine in a single movement. 
	 * 
	 * Typical unit code(s): CMQ for cubic centimeter, LTR for liters, INQ for cubic inches
	 * * Note 1: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	 * * Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * see : https://auto.schema.org/engineDisplacement
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $engine_displacement;
	
	/**
	 * The power of the vehicle&#39;s engine.
	 *     Typical unit code(s): KWT for kilowatt, BHP for brake horsepower, N12 for metric horsepower (PS, with 1 PS = 735,49875 W)
	 * 
	 * 
	 * Note 1: There are many different ways of measuring an engine&#39;s power. For an overview, see  http://en.wikipedia.org/wiki/Horsepower#Engine (see: https://schema.orghttp://en.wikipedia.org/wiki/Horsepower#Engine_power_test_codes)powertest_codes.
	 * Note 2: You can link to information about how the given value has been determined using the valueReference (see: https://schema.org/valueReference) property.
	 * Note 3: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/enginePower
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $engine_power;
	
	/**
	 * The type of engine or engines powering the vehicle.
	 * see : https://auto.schema.org/engineType
	 * @var \QualitativeValue | \QualitativeValue[] | string | string[]
	 */
	public $engine_type;
	
	/**
	 * The type of fuel suitable for the engine or engines of the vehicle. If the vehicle has only one engine, this property can be attached directly to the vehicle.
	 * see : https://schema.org/fuelType
	 * @var \QualitativeValue | \QualitativeValue[] | string | string[]
	 */
	public $fuel_type;
	
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
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public $same_as;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;
	
	/**
	 * The torque (turning force) of the vehicle&#39;s engine.
	 * 
	 * Typical unit code(s): NU for newton metre (N m), F17 for pound-force per foot, or F48 for pound-force per inch
	 * 
	 * 
	 * Note 1: You can link to information about how the given value has been determined (e.g. reference RPM) using the valueReference (see: https://schema.org/valueReference) property.
	 * Note 2: You can use minValue (see: https://schema.org/minValue) and maxValue (see: https://schema.org/maxValue) to indicate ranges.
	 * 
	 * 
	 * see : https://auto.schema.org/torque
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $torque;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EngineSpecification'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->engine_displacement );
		if ( ! empty( $serialized ) ) {
			$out['engineDisplacement'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->engine_power );
		if ( ! empty( $serialized ) ) {
			$out['enginePower'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->engine_type );
		if ( ! empty( $serialized ) ) {
			$out['engineType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fuel_type );
		if ( ! empty( $serialized ) ) {
			$out['fuelType'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->torque );
		if ( ! empty( $serialized ) ) {
			$out['torque'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
