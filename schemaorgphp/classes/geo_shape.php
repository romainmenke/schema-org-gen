<?php

class GeoShape extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GeoShape';
	
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
	 * A box is the area enclosed by the rectangle formed by two points. The first point is the lower corner, the second point is the upper corner. A box is expressed as two points separated by a space character.
	 * see : https://schema.org/box
	 * @var string|string[]
	 */
	public var $box;
	
	/**
	 * A circle is the circular region of a specified radius centered at a specified latitude and longitude. A circle is expressed as a pair followed by a radius in meters.
	 * see : https://schema.org/circle
	 * @var string|string[]
	 */
	public var $circle;
	
	/**
	 * The elevation of a location (WGS 84 (see: https://schema.orghttps://en.wikipedia.org/wiki/World_Geodetic_System)).
	 * see : https://schema.org/elevation
	 * @var float|float[]|string|string[]
	 */
	public var $elevation;
	
	/**
	 * A line is a point-to-point path consisting of two or more points. A line is expressed as a series of two or more point objects separated by space.
	 * see : https://schema.org/line
	 * @var string|string[]
	 */
	public var $line;
	
	/**
	 * A polygon is the area enclosed by a point-to-point path for which the starting and ending points are the same. A polygon is expressed as a series of four or more space delimited points where the first and final points are identical.
	 * see : https://schema.org/polygon
	 * @var string|string[]
	 */
	public var $polygon;
	
	/**
	 * The postal code. For example, 94043.
	 * see : https://schema.org/postalCode
	 * @var string|string[]
	 */
	public var $postal_code;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GeoShape'
		);
		
		$serialized = so_json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address_country );
		if ( ! empty( $serialized ) ) {
			$out['addressCountry'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->box );
		if ( ! empty( $serialized ) ) {
			$out['box'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->circle );
		if ( ! empty( $serialized ) ) {
			$out['circle'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->elevation );
		if ( ! empty( $serialized ) ) {
			$out['elevation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->line );
		if ( ! empty( $serialized ) ) {
			$out['line'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->polygon );
		if ( ! empty( $serialized ) ) {
			$out['polygon'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->postal_code );
		if ( ! empty( $serialized ) ) {
			$out['postalCode'] = $serialized;
		}
		
		return $out;
	}
}
