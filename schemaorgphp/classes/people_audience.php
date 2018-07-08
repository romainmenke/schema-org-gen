<?php

class PeopleAudience extends Audience implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PeopleAudience';
	
	/**
	 * Specifying the health condition(s) of a patient, medical study, or other target audience.
	 * see : https://health-lifesci.schema.org/healthCondition
	 * @var \MedicalCondition|\MedicalCondition[]
	 */
	public var $health_condition;
	
	/**
	 * Audiences defined by a person&#39;s gender.
	 * see : https://schema.org/requiredGender
	 * @var string|string[]
	 */
	public var $required_gender;
	
	/**
	 * Audiences defined by a person&#39;s maximum age.
	 * see : https://schema.org/requiredMaxAge
	 * @var integer|integer[]
	 */
	public var $required_max_age;
	
	/**
	 * Audiences defined by a person&#39;s minimum age.
	 * see : https://schema.org/requiredMinAge
	 * @var integer|integer[]
	 */
	public var $required_min_age;
	
	/**
	 * The gender of the person or audience.
	 * see : https://schema.org/suggestedGender
	 * @var string|string[]
	 */
	public var $suggested_gender;
	
	/**
	 * Maximal age recommended for viewing content.
	 * see : https://schema.org/suggestedMaxAge
	 * @var float|float[]
	 */
	public var $suggested_max_age;
	
	/**
	 * Minimal age recommended for viewing content.
	 * see : https://schema.org/suggestedMinAge
	 * @var float|float[]
	 */
	public var $suggested_min_age;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PeopleAudience'
		);
		
		$serialized = so_json_serialize( $this->health_condition );
		if ( ! empty( $serialized ) ) {
			$out['healthCondition'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->required_gender );
		if ( ! empty( $serialized ) ) {
			$out['requiredGender'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->required_max_age );
		if ( ! empty( $serialized ) ) {
			$out['requiredMaxAge'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->required_min_age );
		if ( ! empty( $serialized ) ) {
			$out['requiredMinAge'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->suggested_gender );
		if ( ! empty( $serialized ) ) {
			$out['suggestedGender'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->suggested_max_age );
		if ( ! empty( $serialized ) ) {
			$out['suggestedMaxAge'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->suggested_min_age );
		if ( ! empty( $serialized ) ) {
			$out['suggestedMinAge'] = $serialized;
		}
		
		return $out;
	}
}
