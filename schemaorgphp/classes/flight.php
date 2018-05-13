<?php

class Flight extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Flight';
	
	/**
	 * The kind of aircraft (e.g., &quot;Boeing 747&quot;).
	 * see : https://schema.org/aircraft
	 * @var string|string[]|\Vehicle|\Vehicle[]
	 */
	public var $aircraft;
	
	/**
	 * The airport where the flight terminates.
	 * see : https://schema.org/arrivalAirport
	 * @var \Airport|\Airport[]
	 */
	public var $arrival_airport;
	
	/**
	 * Identifier of the flight&#39;s arrival gate.
	 * see : https://schema.org/arrivalGate
	 * @var string|string[]
	 */
	public var $arrival_gate;
	
	/**
	 * Identifier of the flight&#39;s arrival terminal.
	 * see : https://schema.org/arrivalTerminal
	 * @var string|string[]
	 */
	public var $arrival_terminal;
	
	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 * @var string|string[]
	 */
	public var $arrival_time;
	
	/**
	 * The type of boarding policy used by the airline (e.g. zone-based or group-based).
	 * see : https://schema.org/boardingPolicy
	 * @var \BoardingPolicyType|\BoardingPolicyType[]
	 */
	public var $boarding_policy;
	
	/**
	 * The airport where the flight originates.
	 * see : https://schema.org/departureAirport
	 * @var \Airport|\Airport[]
	 */
	public var $departure_airport;
	
	/**
	 * Identifier of the flight&#39;s departure gate.
	 * see : https://schema.org/departureGate
	 * @var string|string[]
	 */
	public var $departure_gate;
	
	/**
	 * Identifier of the flight&#39;s departure terminal.
	 * see : https://schema.org/departureTerminal
	 * @var string|string[]
	 */
	public var $departure_terminal;
	
	/**
	 * The expected departure time.
	 * see : https://schema.org/departureTime
	 * @var string|string[]
	 */
	public var $departure_time;
	
	/**
	 * The estimated time the flight will take.
	 * see : https://schema.org/estimatedFlightDuration
	 * @var \Duration|\Duration[]|string|string[]
	 */
	public var $estimated_flight_duration;
	
	/**
	 * The distance of the flight.
	 * see : https://schema.org/flightDistance
	 * @var \Distance|\Distance[]|string|string[]
	 */
	public var $flight_distance;
	
	/**
	 * The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is &#39;UA&#39;, the flightNumber is &#39;UA110&#39;.
	 * see : https://schema.org/flightNumber
	 * @var string|string[]
	 */
	public var $flight_number;
	
	/**
	 * Description of the meals that will be provided or available for purchase.
	 * see : https://schema.org/mealService
	 * @var string|string[]
	 */
	public var $meal_service;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $provider;
	
	/**
	 * An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider. Supersedes merchant (see: https://schema.org/merchant), vendor (see: https://schema.org/vendor).
	 * see : https://schema.org/seller
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $seller;
	
	/**
	 * The time when a passenger can check into the flight online.
	 * see : https://schema.org/webCheckinTime
	 * @var string|string[]
	 */
	public var $web_checkin_time;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Flight'
		);
		
		$serialized = so_json_serialize( $this->aircraft );
		if ( ! empty( $serialized ) ) {
			$out['aircraft'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_airport );
		if ( ! empty( $serialized ) ) {
			$out['arrivalAirport'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_gate );
		if ( ! empty( $serialized ) ) {
			$out['arrivalGate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_terminal );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTerminal'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->boarding_policy );
		if ( ! empty( $serialized ) ) {
			$out['boardingPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_airport );
		if ( ! empty( $serialized ) ) {
			$out['departureAirport'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_gate );
		if ( ! empty( $serialized ) ) {
			$out['departureGate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_terminal );
		if ( ! empty( $serialized ) ) {
			$out['departureTerminal'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->estimated_flight_duration );
		if ( ! empty( $serialized ) ) {
			$out['estimatedFlightDuration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->flight_distance );
		if ( ! empty( $serialized ) ) {
			$out['flightDistance'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->flight_number );
		if ( ! empty( $serialized ) ) {
			$out['flightNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->meal_service );
		if ( ! empty( $serialized ) ) {
			$out['mealService'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seller );
		if ( ! empty( $serialized ) ) {
			$out['seller'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->web_checkin_time );
		if ( ! empty( $serialized ) ) {
			$out['webCheckinTime'] = $serialized;
		}
		
		return $out;
	}
}
