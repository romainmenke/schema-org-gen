<?php
namespace SchemaOrg;

// Ticket see : https://schema.org/Ticket
class Ticket implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Ticket';

	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * The date the ticket was issued.
	 * see : https://schema.org/dateIssued
	 *
	 * @var string | string[]
	 */
	public $date_issued;

	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 *
	 * @var string | string[]
	 */
	public $disambiguating_description;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * The organization issuing the ticket or permit.
	 * see : https://schema.org/issuedBy
	 *
	 * @var \Organization | \Organization[]
	 */
	public $issued_by;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * The currency (in 3-letter ISO 4217 format) of the price or a price component, when attached to [[PriceSpecification]] and its subtypes.
	 * see : https://schema.org/priceCurrency
	 *
	 * @var string | string[]
	 */
	public $price_currency;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The unique identifier for the ticket.
	 * see : https://schema.org/ticketNumber
	 *
	 * @var string | string[]
	 */
	public $ticket_number;

	/**
	 * Reference to an asset (e.g., Barcode, QR code image or PDF) usable for entrance.
	 * see : https://schema.org/ticketToken
	 *
	 * @var string | string[]
	 */
	public $ticket_token;

	/**
	 * The seat associated with the ticket.
	 * see : https://schema.org/ticketedSeat
	 *
	 * @var \Seat | \Seat[]
	 */
	public $ticketed_seat;

	/**
	 * The total price for the reservation or ticket, including applicable taxes, shipping, etc.
	 * see : https://schema.org/totalPrice
	 *
	 * @var float | float[] | string | string[] | \PriceSpecification | \PriceSpecification[]
	 */
	public $total_price;

	/**
	 * The person or organization the reservation or ticket is for.
	 * see : https://schema.org/underName
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $under_name;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Ticket',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->date_issued );
		if ( ! empty( $serialized ) ) {
			$out['dateIssued'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->issued_by );
		if ( ! empty( $serialized ) ) {
			$out['issuedBy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->price_currency );
		if ( ! empty( $serialized ) ) {
			$out['priceCurrency'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->ticket_number );
		if ( ! empty( $serialized ) ) {
			$out['ticketNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->ticket_token );
		if ( ! empty( $serialized ) ) {
			$out['ticketToken'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->ticketed_seat );
		if ( ! empty( $serialized ) ) {
			$out['ticketedSeat'] = $serialized;
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
