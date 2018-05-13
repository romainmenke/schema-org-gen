<?php

class OpeningHoursSpecification extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OpeningHoursSpecification';
	
	/**
	 * The closing hour of the place or service on the given day(s) of the week.
	 * see : https://schema.org/closes
	 * @var string|string[]
	 */
	public var $closes;
	
	/**
	 * The day of the week for which these opening hours are valid.
	 * see : https://schema.org/dayOfWeek
	 * @var \DayOfWeek|\DayOfWeek[]
	 */
	public var $day_of_week;
	
	/**
	 * The opening hour of the place or service on the given day(s) of the week.
	 * see : https://schema.org/opens
	 * @var string|string[]
	 */
	public var $opens;
	
	/**
	 * The date when the item becomes valid.
	 * see : https://schema.org/validFrom
	 * @var string|string[]
	 */
	public var $valid_from;
	
	/**
	 * The date after when the item is not valid. For example the end of an offer, salary period, or a period of opening hours.
	 * see : https://schema.org/validThrough
	 * @var string|string[]
	 */
	public var $valid_through;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OpeningHoursSpecification'
		);
		
		$serialized = so_json_serialize( $this->closes );
		if ( ! empty( $serialized ) ) {
			$out['closes'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->day_of_week );
		if ( ! empty( $serialized ) ) {
			$out['dayOfWeek'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->opens );
		if ( ! empty( $serialized ) ) {
			$out['opens'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_from );
		if ( ! empty( $serialized ) ) {
			$out['validFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_through );
		if ( ! empty( $serialized ) ) {
			$out['validThrough'] = $serialized;
		}
		
		return $out;
	}
}
