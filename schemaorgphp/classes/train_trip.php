<?php

class TrainTrip extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TrainTrip';
	
	/**
	 * The platform where the train arrives.
	 * see : https://schema.org/arrivalPlatform
	 * @var string|string[]
	 */
	public var $arrival_platform;
	
	/**
	 * The station where the train trip ends.
	 * see : https://schema.org/arrivalStation
	 * @var \TrainStation|\TrainStation[]
	 */
	public var $arrival_station;
	
	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 * @var string|string[]
	 */
	public var $arrival_time;
	
	/**
	 * The platform from which the train departs.
	 * see : https://schema.org/departurePlatform
	 * @var string|string[]
	 */
	public var $departure_platform;
	
	/**
	 * The station from which the train departs.
	 * see : https://schema.org/departureStation
	 * @var \TrainStation|\TrainStation[]
	 */
	public var $departure_station;
	
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
	
	/**
	 * The name of the train (e.g. The Orient Express).
	 * see : https://schema.org/trainName
	 * @var string|string[]
	 */
	public var $train_name;
	
	/**
	 * The unique identifier for the train.
	 * see : https://schema.org/trainNumber
	 * @var string|string[]
	 */
	public var $train_number;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TrainTrip'
		);
		
		$serialized = so_json_serialize( $this->arrival_platform );
		if ( ! empty( $serialized ) ) {
			$out['arrivalPlatform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_station );
		if ( ! empty( $serialized ) ) {
			$out['arrivalStation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_platform );
		if ( ! empty( $serialized ) ) {
			$out['departurePlatform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_station );
		if ( ! empty( $serialized ) ) {
			$out['departureStation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->train_name );
		if ( ! empty( $serialized ) ) {
			$out['trainName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->train_number );
		if ( ! empty( $serialized ) ) {
			$out['trainNumber'] = $serialized;
		}
		
		return $out;
	}
}
