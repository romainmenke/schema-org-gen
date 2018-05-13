<?php

class BorrowAction extends TransferAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BorrowAction';
	
	/**
	 * A sub property of participant. The person that lends the object being borrowed.
	 * see : https://schema.org/lender
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $lender;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BorrowAction'
		);
		
		$serialized = so_json_serialize( $this->lender );
		if ( ! empty( $serialized ) ) {
			$out['lender'] = $serialized;
		}
		
		return $out;
	}
}
