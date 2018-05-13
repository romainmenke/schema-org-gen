<?php

class AcceptAction extends AllocateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AcceptAction';
	
	/**
	 * A goal towards an action is taken. Can be concrete or abstract.
	 * see : http://health-lifesci.schema.org/purpose
	 * @var \MedicalDevicePurpose|\MedicalDevicePurpose[]|\Thing|\Thing[]
	 */
	public var $purpose;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AcceptAction'
		);
		
		$serialized = so_json_serialize( $this->purpose );
		if ( ! empty( $serialized ) ) {
			$out['purpose'] = $serialized;
		}
		
		return $out;
	}
}
