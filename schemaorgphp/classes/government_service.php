<?php

class GovernmentService extends Service implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GovernmentService';
	
	/**
	 * The operating organization, if different from the provider.  This enables the representation of services that are provided by an organization, but operated by another organization like a subcontractor.
	 * see : https://schema.org/serviceOperator
	 * @var \Organization|\Organization[]
	 */
	public var $service_operator;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GovernmentService'
		);
		
		$serialized = so_json_serialize( $this->service_operator );
		if ( ! empty( $serialized ) ) {
			$out['serviceOperator'] = $serialized;
		}
		
		return $out;
	}
}
