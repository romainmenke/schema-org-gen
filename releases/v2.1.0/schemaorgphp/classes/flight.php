<?php
namespace SchemaOrg;

// Flight see : https://schema.org/Flight
class Flight implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Flight';

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
	 * The kind of aircraft (e.g., &quot;Boeing 747&quot;).
	 * see : https://schema.org/aircraft
	 *
	 * @var string | string[] | \Vehicle | \Vehicle[]
	 */
	public $aircraft;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * The airport where the flight terminates.
	 * see : https://schema.org/arrivalAirport
	 *
	 * @var \Airport | \Airport[]
	 */
	public $arrival_airport;

	/**
	 * Identifier of the flight&#39;s arrival gate.
	 * see : https://schema.org/arrivalGate
	 *
	 * @var string | string[]
	 */
	public $arrival_gate;

	/**
	 * Identifier of the flight&#39;s arrival terminal.
	 * see : https://schema.org/arrivalTerminal
	 *
	 * @var string | string[]
	 */
	public $arrival_terminal;

	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 *
	 * @var string | string[]
	 */
	public $arrival_time;

	/**
	 * The type of boarding policy used by the airline (e.g. zone-based or group-based).
	 * see : https://schema.org/boardingPolicy
	 *
	 * @var \BoardingPolicyType | \BoardingPolicyType[]
	 */
	public $boarding_policy;

	/**
	 * &#39;carrier&#39; is an out-dated term indicating the &#39;provider&#39; for parcel delivery and flights.
	 * see : https://schema.org/carrier
	 *
	 * @var \Organization | \Organization[]
	 */
	public $carrier;

	/**
	 * The airport where the flight originates.
	 * see : https://schema.org/departureAirport
	 *
	 * @var \Airport | \Airport[]
	 */
	public $departure_airport;

	/**
	 * Identifier of the flight&#39;s departure gate.
	 * see : https://schema.org/departureGate
	 *
	 * @var string | string[]
	 */
	public $departure_gate;

	/**
	 * Identifier of the flight&#39;s departure terminal.
	 * see : https://schema.org/departureTerminal
	 *
	 * @var string | string[]
	 */
	public $departure_terminal;

	/**
	 * The expected departure time.
	 * see : https://schema.org/departureTime
	 *
	 * @var string | string[]
	 */
	public $departure_time;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The estimated time the flight will take.
	 * see : https://schema.org/estimatedFlightDuration
	 *
	 * @var string | string[] | \Duration | \Duration[]
	 */
	public $estimated_flight_duration;

	/**
	 * The distance of the flight.
	 * see : https://schema.org/flightDistance
	 *
	 * @var string | string[] | \Distance | \Distance[]
	 */
	public $flight_distance;

	/**
	 * The unique identifier for a flight including the airline IATA code. For example, if describing United flight 110, where the IATA code for United is &#39;UA&#39;, the flightNumber is &#39;UA110&#39;.
	 * see : https://schema.org/flightNumber
	 *
	 * @var string | string[]
	 */
	public $flight_number;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * Description of the meals that will be provided or available for purchase.
	 * see : https://schema.org/mealService
	 *
	 * @var string | string[]
	 */
	public $meal_service;

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
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	 * see : https://schema.org/provider
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $provider;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.
	 * see : https://schema.org/seller
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $seller;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The time when a passenger can check into the flight online.
	 * see : https://schema.org/webCheckinTime
	 *
	 * @var string | string[]
	 */
	public $web_checkin_time;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Flight',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->aircraft );
		if ( ! empty( $serialized ) ) {
			$out['aircraft'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_airport );
		if ( ! empty( $serialized ) ) {
			$out['arrivalAirport'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_gate );
		if ( ! empty( $serialized ) ) {
			$out['arrivalGate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_terminal );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTerminal'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->boarding_policy );
		if ( ! empty( $serialized ) ) {
			$out['boardingPolicy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->carrier );
		if ( ! empty( $serialized ) ) {
			$out['carrier'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_airport );
		if ( ! empty( $serialized ) ) {
			$out['departureAirport'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_gate );
		if ( ! empty( $serialized ) ) {
			$out['departureGate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_terminal );
		if ( ! empty( $serialized ) ) {
			$out['departureTerminal'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->estimated_flight_duration );
		if ( ! empty( $serialized ) ) {
			$out['estimatedFlightDuration'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->flight_distance );
		if ( ! empty( $serialized ) ) {
			$out['flightDistance'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->flight_number );
		if ( ! empty( $serialized ) ) {
			$out['flightNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->meal_service );
		if ( ! empty( $serialized ) ) {
			$out['mealService'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->seller );
		if ( ! empty( $serialized ) ) {
			$out['seller'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->web_checkin_time );
		if ( ! empty( $serialized ) ) {
			$out['webCheckinTime'] = $serialized;
		}

		return $out;
	}
}
