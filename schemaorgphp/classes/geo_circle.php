<?php

class GeoCircle extends GeoShape implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GeoCircle';
	
	/**
	 * Indicates the GeoCoordinates at the centre of a GeoShape e.g. GeoCircle.
	 * see : https://schema.org/geoMidpoint
	 * @var \GeoCoordinates|\GeoCoordinates[]
	 */
	public var $geo_midpoint;
	
	/**
	 * Indicates the approximate radius of a GeoCircle (metres unless indicated otherwise via Distance notation).
	 * see : https://schema.org/geoRadius
	 * @var \Distance|\Distance[]|float|float[]|string|string[]
	 */
	public var $geo_radius;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GeoCircle'
		);
		
		$serialized = so_json_serialize( $this->geo_midpoint );
		if ( ! empty( $serialized ) ) {
			$out['geoMidpoint'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geo_radius );
		if ( ! empty( $serialized ) ) {
			$out['geoRadius'] = $serialized;
		}
		
		return $out;
	}
}
