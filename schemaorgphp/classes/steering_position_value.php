<?php

// SteeringPositionValue see : https://schema.org/SteeringPositionValue
class SteeringPositionValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SteeringPositionValue';
	
	/**
	 * With properties from Enumeration see : https://schema.org/Enumeration
	 */
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from QualitativeValue see : https://schema.org/QualitativeValue
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue | \PropertyValue[]
	 */
	public var $additional_property;
	
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
	 * This ordering relation for qualitative values indicates that the subject is equal to the object.
	 * see : https://schema.org/equal
	 * @var \QualitativeValue | \QualitativeValue[]
	 */
	public var $equal;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is greater than the object.
	 * see : https://schema.org/greater
	 * @var \QualitativeValue | \QualitativeValue[]
	 */
	public var $greater;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is greater than or equal to the object.
	 * see : https://schema.org/greaterOrEqual
	 * @var \QualitativeValue | \QualitativeValue[]
	 */
	public var $greater_or_equal;
	
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
	 * This ordering relation for qualitative values indicates that the subject is lesser than the object.
	 * see : https://schema.org/lesser
	 * @var \QualitativeValue | \QualitativeValue[]
	 */
	public var $lesser;
	
	/**
	 * This ordering relation for qualitative values indicates that the subject is lesser than or equal to the object.
	 * see : https://schema.org/lesserOrEqual
	 * @var \QualitativeValue | \QualitativeValue[]
	 */
	public var $lesser_or_equal;
	
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
	 * This ordering relation for qualitative values indicates that the subject is not equal to the object.
	 * see : https://schema.org/nonEqual
	 * @var \QualitativeValue | \QualitativeValue[]
	 */
	public var $non_equal;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * Relates a term (i.e. a property, class or enumeration) to one that supersedes it.
	 * see : https://meta.schema.org/supersededBy
	 * @var \Class | \Class[] | \Enumeration | \Enumeration[] | \Property | \Property[]
	 */
	public var $superseded_by;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	/**
	 * A pointer to a secondary value that provides additional information on the original value, e.g. a reference temperature.
	 * see : https://schema.org/valueReference
	 * @var \Enumeration | \Enumeration[] | \PropertyValue | \PropertyValue[] | \QualitativeValue | \QualitativeValue[] | \QuantitativeValue | \QuantitativeValue[] | \StructuredValue | \StructuredValue[]
	 */
	public var $value_reference;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SteeringPositionValue'
		);
		
		$serialized = so_json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->equal );
		if ( ! empty( $serialized ) ) {
			$out['equal'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->greater );
		if ( ! empty( $serialized ) ) {
			$out['greater'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->greater_or_equal );
		if ( ! empty( $serialized ) ) {
			$out['greaterOrEqual'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lesser );
		if ( ! empty( $serialized ) ) {
			$out['lesser'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lesser_or_equal );
		if ( ! empty( $serialized ) ) {
			$out['lesserOrEqual'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->non_equal );
		if ( ! empty( $serialized ) ) {
			$out['nonEqual'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->superseded_by );
		if ( ! empty( $serialized ) ) {
			$out['supersededBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_reference );
		if ( ! empty( $serialized ) ) {
			$out['valueReference'] = $serialized;
		}
		
		return $out;
	}
}
