<?php

class School extends EducationalOrganization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'School';
	
	/**
	 * Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	 * see : https://schema.org/alumni
	 * @var \Person|\Person[]
	 */
	public var $alumni;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'School'
		);
		
		$serialized = so_json_serialize( $this->alumni );
		if ( ! empty( $serialized ) ) {
			$out['alumni'] = $serialized;
		}
		
		return $out;
	}
}
