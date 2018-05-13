<?php

class GeoCoordinates extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GeoCoordinates';
	
	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 * @var \PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $address;
	
	/**
	 * The country. For example, USA. You can also provide the two-letter ISO 3166-1 alpha-2 country code (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166-1).
	 * see : https://schema.org/addressCountry
	 * @var \Country|\Country[]|string|string[]
	 */
	public var $address_country;
	
	/**
	 * The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	 * see : https://schema.org/elevation
	 * @var float|float[]|string|string[]
	 */
	public var $elevation;
	
	/**
	 * The latitude of a location. For example 37.42242 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	 * see : https://schema.org/latitude
	 * @var float|float[]|string|string[]
	 */
	public var $latitude;
	
	/**
	 * The longitude of a location. For example -122.08585 (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	 * see : https://schema.org/longitude
	 * @var float|float[]|string|string[]
	 */
	public var $longitude;
	
	/**
	 * The postal code. For example, 94043.
	 * see : https://schema.org/postalCode
	 * @var string|string[]
	 */
	public var $postal_code;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GeoCoordinates'
		);
		
		$serialized = so_json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address_country );
		if ( ! empty( $serialized ) ) {
			$out['addressCountry'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->elevation );
		if ( ! empty( $serialized ) ) {
			$out['elevation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->latitude );
		if ( ! empty( $serialized ) ) {
			$out['latitude'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->longitude );
		if ( ! empty( $serialized ) ) {
			$out['longitude'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->postal_code );
		if ( ! empty( $serialized ) ) {
			$out['postalCode'] = $serialized;
		}
		
		return $out;
	}
}
