<?php

class MobileApplication extends SoftwareApplication implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'MobileApplication';
	
	/**
	 * Specifies specific carrier(s) requirements for the application (e.g. an application may only work on a specific carrier network).
	 * see : https://schema.org/carrierRequirements
	 * @var string|string[]
	 */
	public var $carrier_requirements;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'MobileApplication'
		);
		
		$serialized = so_json_serialize( $this->carrier_requirements );
		if ( ! empty( $serialized ) ) {
			$out['carrierRequirements'] = $serialized;
		}
		
		return $out;
	}
}
