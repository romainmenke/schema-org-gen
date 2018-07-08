<?php

class DeliveryChargeSpecification extends PriceSpecification implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DeliveryChargeSpecification';
	
	/**
	 * The delivery method(s) to which the delivery charge or payment charge specification applies.
	 * see : https://schema.org/appliesToDeliveryMethod
	 * @var \DeliveryMethod|\DeliveryMethod[]
	 */
	public var $applies_to_delivery_method;
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea|\AdministrativeArea[]|\GeoShape|\GeoShape[]|\Place|\Place[]|string|string[]
	 */
	public var $area_served;
	
	/**
	 * The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
	 * 
	 * See also ineligibleRegion (see: https://schema.org/ineligibleRegion).
	 * see : https://pending.schema.org/eligibleRegion
	 * @var \GeoShape|\GeoShape[]|\Place|\Place[]|string|string[]
	 */
	public var $eligible_region;
	
	/**
	 * The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
	 * 
	 * See also eligibleRegion (see: https://schema.org/eligibleRegion).
	 * see : https://schema.org/ineligibleRegion
	 * @var \GeoShape|\GeoShape[]|\Place|\Place[]|string|string[]
	 */
	public var $ineligible_region;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DeliveryChargeSpecification'
		);
		
		$serialized = so_json_serialize( $this->applies_to_delivery_method );
		if ( ! empty( $serialized ) ) {
			$out['appliesToDeliveryMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->eligible_region );
		if ( ! empty( $serialized ) ) {
			$out['eligibleRegion'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ineligible_region );
		if ( ! empty( $serialized ) ) {
			$out['ineligibleRegion'] = $serialized;
		}
		
		return $out;
	}
}
