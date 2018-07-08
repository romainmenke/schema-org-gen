<?php

class Trip extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Trip';
	
	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 * @var string|string[]
	 */
	public var $arrival_time;
	
	/**
	 * The expected departure time.
	 * see : https://schema.org/departureTime
	 * @var string|string[]
	 */
	public var $departure_time;
	
	/**
	 * Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense). Inverse property: isPartOf (see: https://schema.org/isPartOf).
	 * see : https://schema.org/hasPart
	 * @var \CreativeWork|\CreativeWork[]|\Trip|\Trip[]
	 */
	public var $has_part;
	
	/**
	 * Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of. Inverse property: hasPart (see: https://schema.org/hasPart).
	 * see : https://schema.org/isPartOf
	 * @var \CreativeWork|\CreativeWork[]|\Trip|\Trip[]
	 */
	public var $is_part_of;
	
	/**
	 * Destination(s) ( Place (see: https://schema.org/Place) ) that make up a trip. For a trip where destination order is important use ItemList (see: https://schema.org/ItemList) to specify that order (see examples).
	 * see : https://pending.schema.org/itinerary
	 * @var \ItemList|\ItemList[]|\Place|\Place[]
	 */
	public var $itinerary;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer|\Offer[]
	 */
	public var $offers;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $provider;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Trip'
		);
		
		$serialized = so_json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_part );
		if ( ! empty( $serialized ) ) {
			$out['hasPart'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_part_of );
		if ( ! empty( $serialized ) ) {
			$out['isPartOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->itinerary );
		if ( ! empty( $serialized ) ) {
			$out['itinerary'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		return $out;
	}
}
