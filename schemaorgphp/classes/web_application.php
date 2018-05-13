<?php

class WebApplication extends SoftwareApplication implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'WebApplication';
	
	/**
	 * Specifies browser requirements in human-readable text. For example, &#39;requires HTML5 support&#39;.
	 * see : https://schema.org/browserRequirements
	 * @var string|string[]
	 */
	public var $browser_requirements;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'WebApplication'
		);
		
		$serialized = so_json_serialize( $this->browser_requirements );
		if ( ! empty( $serialized ) ) {
			$out['browserRequirements'] = $serialized;
		}
		
		return $out;
	}
}
