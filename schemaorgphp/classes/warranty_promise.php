<?php

class WarrantyPromise extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'WarrantyPromise';
	
	/**
	 * The duration of the warranty promise. Common unitCode values are ANN for year, MON for months, or DAY for days.
	 * see : https://schema.org/durationOfWarranty
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $duration_of_warranty;
	
	/**
	 * The scope of the warranty promise.
	 * see : https://schema.org/warrantyScope
	 * @var \WarrantyScope|\WarrantyScope[]
	 */
	public var $warranty_scope;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'WarrantyPromise'
		);
		
		$serialized = so_json_serialize( $this->duration_of_warranty );
		if ( ! empty( $serialized ) ) {
			$out['durationOfWarranty'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->warranty_scope );
		if ( ! empty( $serialized ) ) {
			$out['warrantyScope'] = $serialized;
		}
		
		return $out;
	}
}
