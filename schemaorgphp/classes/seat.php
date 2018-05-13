<?php

class Seat extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Seat';
	
	/**
	 * The location of the reserved seat (e.g., 27).
	 * see : https://schema.org/seatNumber
	 * @var string|string[]
	 */
	public var $seat_number;
	
	/**
	 * The row location of the reserved seat (e.g., B).
	 * see : https://schema.org/seatRow
	 * @var string|string[]
	 */
	public var $seat_row;
	
	/**
	 * The section location of the reserved seat (e.g. Orchestra).
	 * see : https://schema.org/seatSection
	 * @var string|string[]
	 */
	public var $seat_section;
	
	/**
	 * The type/class of the seat.
	 * see : https://schema.org/seatingType
	 * @var \QualitativeValue|\QualitativeValue[]|string|string[]
	 */
	public var $seating_type;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Seat'
		);
		
		$serialized = so_json_serialize( $this->seat_number );
		if ( ! empty( $serialized ) ) {
			$out['seatNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seat_row );
		if ( ! empty( $serialized ) ) {
			$out['seatRow'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seat_section );
		if ( ! empty( $serialized ) ) {
			$out['seatSection'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seating_type );
		if ( ! empty( $serialized ) ) {
			$out['seatingType'] = $serialized;
		}
		
		return $out;
	}
}
