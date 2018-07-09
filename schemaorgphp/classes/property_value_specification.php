<?php

// PropertyValueSpecification see : https://schema.org/PropertyValueSpecification
class PropertyValueSpecification implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PropertyValueSpecification';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
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
	 * The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it&#39;s an ID reference to one of the current values.
	 * see : https://schema.org/defaultValue
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public var $default_value;
	
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
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * The upper value of some characteristic or property.
	 * see : https://schema.org/maxValue
	 * @var float | float[]
	 */
	public var $max_value;
	
	/**
	 * The lower value of some characteristic or property.
	 * see : https://schema.org/minValue
	 * @var float | float[]
	 */
	public var $min_value;
	
	/**
	 * Whether multiple values are allowed for the property.  Default is false.
	 * see : https://schema.org/multipleValues
	 * @var boolean | boolean[]
	 */
	public var $multiple_values;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a &quot;hidden&quot; input in an HTML form.
	 * see : https://schema.org/readonlyValue
	 * @var boolean | boolean[]
	 */
	public var $readonly_value;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	 * see : https://schema.org/stepValue
	 * @var float | float[]
	 */
	public var $step_value;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	/**
	 * Specifies the allowed range for number of characters in a literal value.
	 * see : https://schema.org/valueMaxLength
	 * @var float | float[]
	 */
	public var $value_max_length;
	
	/**
	 * Specifies the minimum allowed range for number of characters in a literal value.
	 * see : https://schema.org/valueMinLength
	 * @var float | float[]
	 */
	public var $value_min_length;
	
	/**
	 * Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML&#39;s input@name.
	 * see : https://schema.org/valueName
	 * @var string | string[]
	 */
	public var $value_name;
	
	/**
	 * Specifies a regular expression for testing literal values according to the HTML spec.
	 * see : https://schema.org/valuePattern
	 * @var string | string[]
	 */
	public var $value_pattern;
	
	/**
	 * Whether the property must be filled in to complete the action.  Default is false.
	 * see : https://schema.org/valueRequired
	 * @var boolean | boolean[]
	 */
	public var $value_required;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PropertyValueSpecification'
		);
		
		$serialized = so_json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->default_value );
		if ( ! empty( $serialized ) ) {
			$out['defaultValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->max_value );
		if ( ! empty( $serialized ) ) {
			$out['maxValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->min_value );
		if ( ! empty( $serialized ) ) {
			$out['minValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->multiple_values );
		if ( ! empty( $serialized ) ) {
			$out['multipleValues'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->readonly_value );
		if ( ! empty( $serialized ) ) {
			$out['readonlyValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->step_value );
		if ( ! empty( $serialized ) ) {
			$out['stepValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_max_length );
		if ( ! empty( $serialized ) ) {
			$out['valueMaxLength'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_min_length );
		if ( ! empty( $serialized ) ) {
			$out['valueMinLength'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_name );
		if ( ! empty( $serialized ) ) {
			$out['valueName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_pattern );
		if ( ! empty( $serialized ) ) {
			$out['valuePattern'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->value_required );
		if ( ! empty( $serialized ) ) {
			$out['valueRequired'] = $serialized;
		}
		
		return $out;
	}
}
