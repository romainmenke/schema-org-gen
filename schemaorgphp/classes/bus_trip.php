<?php

class BusTrip extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BusTrip';
	
	/**
	 * The stop or station from which the bus arrives.
	 * see : https://schema.org/arrivalBusStop
	 * @var \BusStation|\BusStation[]|\BusStop|\BusStop[]
	 */
	public var $arrival_bus_stop;
	
	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 * @var string|string[]
	 */
	public var $arrival_time;
	
	/**
	 * The name of the bus (e.g. Bolt Express).
	 * see : https://schema.org/busName
	 * @var string|string[]
	 */
	public var $bus_name;
	
	/**
	 * The unique identifier for the bus.
	 * see : https://schema.org/busNumber
	 * @var string|string[]
	 */
	public var $bus_number;
	
	/**
	 * The stop or station from which the bus departs.
	 * see : https://schema.org/departureBusStop
	 * @var \BusStation|\BusStation[]|\BusStop|\BusStop[]
	 */
	public var $departure_bus_stop;
	
	/**
	 * The expected departure time.
	 * see : https://schema.org/departureTime
	 * @var string|string[]
	 */
	public var $departure_time;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $provider;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BusTrip'
		);
		
		$serialized = so_json_serialize( $this->arrival_bus_stop );
		if ( ! empty( $serialized ) ) {
			$out['arrivalBusStop'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->bus_name );
		if ( ! empty( $serialized ) ) {
			$out['busName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->bus_number );
		if ( ! empty( $serialized ) ) {
			$out['busNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_bus_stop );
		if ( ! empty( $serialized ) ) {
			$out['departureBusStop'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		return $out;
	}
}
