<?php

class Reservation extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Reservation';
	
	/**
	 * The date and time the reservation was booked.
	 * see : https://schema.org/bookingTime
	 * @var string|string[]
	 */
	public var $booking_time;
	
	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	 * see : https://schema.org/broker
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $broker;
	
	/**
	 * The date and time the reservation was modified.
	 * see : https://schema.org/modifiedTime
	 * @var string|string[]
	 */
	public var $modified_time;
	
	/**
	 * The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	 * see : https://schema.org/priceCurrency
	 * @var string|string[]
	 */
	public var $price_currency;
	
	/**
	 * Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
	 * see : https://schema.org/programMembershipUsed
	 * @var \ProgramMembership|\ProgramMembership[]
	 */
	public var $program_membership_used;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $provider;
	
	/**
	 * The thing -- flight, event, restaurant,etc. being reserved.
	 * see : https://schema.org/reservationFor
	 * @var \Thing|\Thing[]
	 */
	public var $reservation_for;
	
	/**
	 * A unique identifier for the reservation.
	 * see : https://schema.org/reservationId
	 * @var string|string[]
	 */
	public var $reservation_id;
	
	/**
	 * The current status of the reservation.
	 * see : https://schema.org/reservationStatus
	 * @var \ReservationStatusType|\ReservationStatusType[]
	 */
	public var $reservation_status;
	
	/**
	 * A ticket associated with the reservation.
	 * see : https://schema.org/reservedTicket
	 * @var \Ticket|\Ticket[]
	 */
	public var $reserved_ticket;
	
	/**
	 * The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	 * see : https://schema.org/totalPrice
	 * @var float|float[]|\PriceSpecification|\PriceSpecification[]|string|string[]
	 */
	public var $total_price;
	
	/**
	 * The person or organization the reservation or ticket is for.
	 * see : https://schema.org/underName
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $under_name;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Reservation'
		);
		
		$serialized = so_json_serialize( $this->booking_time );
		if ( ! empty( $serialized ) ) {
			$out['bookingTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->modified_time );
		if ( ! empty( $serialized ) ) {
			$out['modifiedTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price_currency );
		if ( ! empty( $serialized ) ) {
			$out['priceCurrency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->program_membership_used );
		if ( ! empty( $serialized ) ) {
			$out['programMembershipUsed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->reservation_for );
		if ( ! empty( $serialized ) ) {
			$out['reservationFor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->reservation_id );
		if ( ! empty( $serialized ) ) {
			$out['reservationId'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->reservation_status );
		if ( ! empty( $serialized ) ) {
			$out['reservationStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->reserved_ticket );
		if ( ! empty( $serialized ) ) {
			$out['reservedTicket'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->total_price );
		if ( ! empty( $serialized ) ) {
			$out['totalPrice'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->under_name );
		if ( ! empty( $serialized ) ) {
			$out['underName'] = $serialized;
		}
		
		return $out;
	}
}
