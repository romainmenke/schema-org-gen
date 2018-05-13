<?php

class EndorseAction extends ReactAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EndorseAction';
	
	/**
	 * A sub property of participant. The person/organization being supported.
	 * see : https://schema.org/endorsee
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $endorsee;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EndorseAction'
		);
		
		$serialized = so_json_serialize( $this->endorsee );
		if ( ! empty( $serialized ) ) {
			$out['endorsee'] = $serialized;
		}
		
		return $out;
	}
}
