<?php

class LendAction extends TransferAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'LendAction';
	
	/**
	 * A sub property of participant. The person that borrows the object being lent.
	 * see : https://schema.org/borrower
	 * @var \Person|\Person[]
	 */
	public var $borrower;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'LendAction'
		);
		
		$serialized = so_json_serialize( $this->borrower );
		if ( ! empty( $serialized ) ) {
			$out['borrower'] = $serialized;
		}
		
		return $out;
	}
}
