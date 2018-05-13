<?php

class PropertyValueSpecification extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PropertyValueSpecification';
	
	/**
	 * The default value of the input.  For properties that expect a literal, the default is a literal value, for properties that expect an object, it&#39;s an ID reference to one of the current values.
	 * see : https://schema.org/defaultValue
	 * @var string|string[]|\Thing|\Thing[]
	 */
	public var $default_value;
	
	/**
	 * The upper value of some characteristic or property.
	 * see : https://schema.org/maxValue
	 * @var float|float[]
	 */
	public var $max_value;
	
	/**
	 * The lower value of some characteristic or property.
	 * see : https://schema.org/minValue
	 * @var float|float[]
	 */
	public var $min_value;
	
	/**
	 * Whether multiple values are allowed for the property.  Default is false.
	 * see : https://schema.org/multipleValues
	 * @var boolean|boolean[]
	 */
	public var $multiple_values;
	
	/**
	 * Whether or not a property is mutable.  Default is false. Specifying this for a property that also has a value makes it act similar to a &quot;hidden&quot; input in an HTML form.
	 * see : https://schema.org/readonlyValue
	 * @var boolean|boolean[]
	 */
	public var $readonly_value;
	
	/**
	 * The stepValue attribute indicates the granularity that is expected (and required) of the value in a PropertyValueSpecification.
	 * see : https://schema.org/stepValue
	 * @var float|float[]
	 */
	public var $step_value;
	
	/**
	 * Specifies the allowed range for number of characters in a literal value.
	 * see : https://schema.org/valueMaxLength
	 * @var float|float[]
	 */
	public var $value_max_length;
	
	/**
	 * Specifies the minimum allowed range for number of characters in a literal value.
	 * see : https://schema.org/valueMinLength
	 * @var float|float[]
	 */
	public var $value_min_length;
	
	/**
	 * Indicates the name of the PropertyValueSpecification to be used in URL templates and form encoding in a manner analogous to HTML&#39;s input@name.
	 * see : https://schema.org/valueName
	 * @var string|string[]
	 */
	public var $value_name;
	
	/**
	 * Specifies a regular expression for testing literal values according to the HTML spec.
	 * see : https://schema.org/valuePattern
	 * @var string|string[]
	 */
	public var $value_pattern;
	
	/**
	 * Whether the property must be filled in to complete the action.  Default is false.
	 * see : https://schema.org/valueRequired
	 * @var boolean|boolean[]
	 */
	public var $value_required;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PropertyValueSpecification'
		);
		
		$serialized = so_json_serialize( $this->default_value );
		if ( ! empty( $serialized ) ) {
			$out['defaultValue'] = $serialized;
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
		
		$serialized = so_json_serialize( $this->readonly_value );
		if ( ! empty( $serialized ) ) {
			$out['readonlyValue'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->step_value );
		if ( ! empty( $serialized ) ) {
			$out['stepValue'] = $serialized;
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
