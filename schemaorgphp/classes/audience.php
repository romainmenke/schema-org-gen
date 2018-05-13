<?php

class Audience extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Audience';
	
	/**
	 * The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	 * see : https://schema.org/audienceType
	 * @var string|string[]
	 */
	public var $audience_type;
	
	/**
	 * The geographic area associated with the audience.
	 * see : https://schema.org/geographicArea
	 * @var \AdministrativeArea|\AdministrativeArea[]
	 */
	public var $geographic_area;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Audience'
		);
		
		$serialized = so_json_serialize( $this->audience_type );
		if ( ! empty( $serialized ) ) {
			$out['audienceType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geographic_area );
		if ( ! empty( $serialized ) ) {
			$out['geographicArea'] = $serialized;
		}
		
		return $out;
	}
}
