<?php

namespace SchemaOrg;

// TaxiReservation see : https://schema.org/TaxiReservation
class TaxiReservation implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TaxiReservation';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Reservation see : https://schema.org/Reservation
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public $additional_type;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public $alternate_name;
	
	/**
	 * The date and time the reservation was booked.
	 * see : https://schema.org/bookingTime
	 * @var string | string[]
	 */
	public $booking_time;
	
	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	 * see : https://schema.org/broker
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $broker;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public $disambiguating_description;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $image;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;
	
	/**
	 * The date and time the reservation was modified.
	 * see : https://schema.org/modifiedTime
	 * @var string | string[]
	 */
	public $modified_time;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public $name;
	
	/**
	 * Number of people the reservation should accommodate.
	 * see : https://schema.org/partySize
	 * @var integer | integer[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public $party_size;
	
	/**
	 * Where a taxi will pick up a passenger or a rental car can be picked up.
	 * see : https://schema.org/pickupLocation
	 * @var \Place | \Place[]
	 */
	public $pickup_location;
	
	/**
	 * When a taxi will pickup a passenger or a rental car can be picked up.
	 * see : https://schema.org/pickupTime
	 * @var string | string[]
	 */
	public $pickup_time;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public $potential_action;
	
	/**
	 * The currency of the price, or a price component when attached to PriceSpecification (see: https://schema.org/PriceSpecification) and its subtypes.
	 * 
	 * Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	 * see : https://schema.org/priceCurrency
	 * @var string | string[]
	 */
	public $price_currency;
	
	/**
	 * Any membership in a frequent flyer, hotel loyalty program, etc. being applied to the reservation.
	 * see : https://schema.org/programMembershipUsed
	 * @var \ProgramMembership | \ProgramMembership[]
	 */
	public $program_membership_used;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $provider;
	
	/**
	 * The thing -- flight, event, restaurant,etc. being reserved.
	 * see : https://schema.org/reservationFor
	 * @var \Thing | \Thing[]
	 */
	public $reservation_for;
	
	/**
	 * A unique identifier for the reservation.
	 * see : https://schema.org/reservationId
	 * @var string | string[]
	 */
	public $reservation_id;
	
	/**
	 * The current status of the reservation.
	 * see : https://schema.org/reservationStatus
	 * @var \ReservationStatusType | \ReservationStatusType[]
	 */
	public $reservation_status;
	
	/**
	 * A ticket associated with the reservation.
	 * see : https://schema.org/reservedTicket
	 * @var \Ticket | \Ticket[]
	 */
	public $reserved_ticket;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public $same_as;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;
	
	/**
	 * The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	 * see : https://schema.org/totalPrice
	 * @var float | float[] | \PriceSpecification | \PriceSpecification[] | string | string[]
	 */
	public $total_price;
	
	/**
	 * The person or organization the reservation or ticket is for.
	 * see : https://schema.org/underName
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $under_name;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TaxiReservation'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->booking_time );
		if ( ! empty( $serialized ) ) {
			$out['bookingTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->modified_time );
		if ( ! empty( $serialized ) ) {
			$out['modifiedTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->party_size );
		if ( ! empty( $serialized ) ) {
			$out['partySize'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->pickup_location );
		if ( ! empty( $serialized ) ) {
			$out['pickupLocation'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->pickup_time );
		if ( ! empty( $serialized ) ) {
			$out['pickupTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->price_currency );
		if ( ! empty( $serialized ) ) {
			$out['priceCurrency'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->program_membership_used );
		if ( ! empty( $serialized ) ) {
			$out['programMembershipUsed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->reservation_for );
		if ( ! empty( $serialized ) ) {
			$out['reservationFor'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->reservation_id );
		if ( ! empty( $serialized ) ) {
			$out['reservationId'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->reservation_status );
		if ( ! empty( $serialized ) ) {
			$out['reservationStatus'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->reserved_ticket );
		if ( ! empty( $serialized ) ) {
			$out['reservedTicket'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->total_price );
		if ( ! empty( $serialized ) ) {
			$out['totalPrice'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->under_name );
		if ( ! empty( $serialized ) ) {
			$out['underName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
