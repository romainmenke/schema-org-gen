<?php

class TipAction extends TradeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TipAction';
	
	/**
	 * A sub property of participant. The participant who is at the receiving end of the action.
	 * see : https://schema.org/recipient
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $recipient;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TipAction'
		);
		
		$serialized = so_json_serialize( $this->recipient );
		if ( ! empty( $serialized ) ) {
			$out['recipient'] = $serialized;
		}
		
		return $out;
	}
}
