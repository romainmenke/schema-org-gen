<?php

class ParentAudience extends PeopleAudience implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ParentAudience';
	
	/**
	 * Maximal age of the child.
	 * see : https://schema.org/childMaxAge
	 * @var float|float[]
	 */
	public var $child_max_age;
	
	/**
	 * Minimal age of the child.
	 * see : https://schema.org/childMinAge
	 * @var float|float[]
	 */
	public var $child_min_age;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ParentAudience'
		);
		
		$serialized = so_json_serialize( $this->child_max_age );
		if ( ! empty( $serialized ) ) {
			$out['childMaxAge'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->child_min_age );
		if ( ! empty( $serialized ) ) {
			$out['childMinAge'] = $serialized;
		}
		
		return $out;
	}
}
