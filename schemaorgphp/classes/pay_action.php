<?php

class PayAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PayAction';
	
	/**
	 * A goal towards an action is taken. Can be concrete or abstract.
	 * see : http://health-lifesci.schema.org/purpose
	 * @var \MedicalDevicePurpose|\MedicalDevicePurpose[]|\Thing|\Thing[]
	 */
	public var $purpose;
	
	/**
	 * A sub property of participant. The participant who is at the receiving end of the action.
	 * see : https://schema.org/recipient
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $recipient;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PayAction'
		);
		
		$serialized = so_json_serialize( $this->purpose );
		if ( ! empty( $serialized ) ) {
			$out['purpose'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipient );
		if ( ! empty( $serialized ) ) {
			$out['recipient'] = $serialized;
		}
		
		return $out;
	}
}
