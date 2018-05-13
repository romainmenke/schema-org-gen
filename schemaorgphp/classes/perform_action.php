<?php

class PerformAction extends PlayAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PerformAction';
	
	/**
	 * A sub property of location. The entertainment business where the action occurred.
	 * see : https://schema.org/entertainmentBusiness
	 * @var \EntertainmentBusiness|\EntertainmentBusiness[]
	 */
	public var $entertainment_business;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PerformAction'
		);
		
		$serialized = so_json_serialize( $this->entertainment_business );
		if ( ! empty( $serialized ) ) {
			$out['entertainmentBusiness'] = $serialized;
		}
		
		return $out;
	}
}
