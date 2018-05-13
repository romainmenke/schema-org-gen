<?php

class IndividualProduct extends Product implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'IndividualProduct';
	
	/**
	 * The serial number or any alphanumeric identifier of a particular product. When attached to an offer, it is a shortcut for the serial number of the product included in the offer.
	 * see : https://schema.org/serialNumber
	 * @var string|string[]
	 */
	public var $serial_number;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'IndividualProduct'
		);
		
		$serialized = so_json_serialize( $this->serial_number );
		if ( ! empty( $serialized ) ) {
			$out['serialNumber'] = $serialized;
		}
		
		return $out;
	}
}
