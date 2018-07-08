<?php

class Ticket extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Ticket';
	
	/**
	 * The date the ticket was issued.
	 * see : https://schema.org/dateIssued
	 * @var string|string[]
	 */
	public var $date_issued;
	
	/**
	 * The organization issuing the ticket or permit.
	 * see : https://schema.org/issuedBy
	 * @var \Organization|\Organization[]
	 */
	public var $issued_by;
	
	/**
	 * The currency of the price, or a price component when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	 * 
	 * Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	 * see : https://schema.org/priceCurrency
	 * @var string|string[]
	 */
	public var $price_currency;
	
	/**
	 * The unique identifier for the ticket.
	 * see : https://schema.org/ticketNumber
	 * @var string|string[]
	 */
	public var $ticket_number;
	
	/**
	 * Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
	 * see : https://schema.org/ticketToken
	 * @var string|string[]|string|string[]
	 */
	public var $ticket_token;
	
	/**
	 * The seat associated with the ticket.
	 * see : https://schema.org/ticketedSeat
	 * @var \Seat|\Seat[]
	 */
	public var $ticketed_seat;
	
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
			'@type' => 'Ticket'
		);
		
		$serialized = so_json_serialize( $this->date_issued );
		if ( ! empty( $serialized ) ) {
			$out['dateIssued'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->issued_by );
		if ( ! empty( $serialized ) ) {
			$out['issuedBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->price_currency );
		if ( ! empty( $serialized ) ) {
			$out['priceCurrency'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ticket_number );
		if ( ! empty( $serialized ) ) {
			$out['ticketNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ticket_token );
		if ( ! empty( $serialized ) ) {
			$out['ticketToken'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ticketed_seat );
		if ( ! empty( $serialized ) ) {
			$out['ticketedSeat'] = $serialized;
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
