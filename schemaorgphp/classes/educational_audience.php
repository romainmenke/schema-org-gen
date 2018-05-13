<?php

class EducationalAudience extends Audience implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EducationalAudience';
	
	/**
	 * An educationalRole of an EducationalAudience.
	 * see : https://schema.org/educationalRole
	 * @var string|string[]
	 */
	public var $educational_role;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EducationalAudience'
		);
		
		$serialized = so_json_serialize( $this->educational_role );
		if ( ! empty( $serialized ) ) {
			$out['educationalRole'] = $serialized;
		}
		
		return $out;
	}
}
