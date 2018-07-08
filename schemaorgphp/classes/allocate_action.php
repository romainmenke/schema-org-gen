<?php

class AllocateAction extends OrganizeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AllocateAction';
	
	/**
	 * A goal towards an action is taken. Can be concrete or abstract.
	 * see : https://health-lifesci.schema.org/purpose
	 * @var \MedicalDevicePurpose|\MedicalDevicePurpose[]|\Thing|\Thing[]
	 */
	public var $purpose;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AllocateAction'
		);
		
		$serialized = so_json_serialize( $this->purpose );
		if ( ! empty( $serialized ) ) {
			$out['purpose'] = $serialized;
		}
		
		return $out;
	}
}
