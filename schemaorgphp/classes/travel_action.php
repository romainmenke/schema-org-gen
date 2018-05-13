<?php

class TravelAction extends MoveAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TravelAction';
	
	/**
	 * The distance travelled, e.g. exercising or travelling.
	 * see : https://schema.org/distance
	 * @var \Distance|\Distance[]
	 */
	public var $distance;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TravelAction'
		);
		
		$serialized = so_json_serialize( $this->distance );
		if ( ! empty( $serialized ) ) {
			$out['distance'] = $serialized;
		}
		
		return $out;
	}
}
