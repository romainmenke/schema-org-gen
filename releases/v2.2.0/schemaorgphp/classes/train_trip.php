<?php
namespace SchemaOrg;

// TrainTrip see : https://schema.org/TrainTrip
class TrainTrip implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'TrainTrip';

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
	 * The platform where the train arrives.
	 * see : https://schema.org/arrivalPlatform
	 *
	 * @var string | string[]
	 */
	public $arrival_platform;

	/**
	 * The station where the train trip ends.
	 * see : https://schema.org/arrivalStation
	 *
	 * @var \TrainStation | \TrainStation[]
	 */
	public $arrival_station;

	/**
	 * The expected arrival time.
	 * see : https://schema.org/arrivalTime
	 *
	 * @var string | string[]
	 */
	public $arrival_time;

	/**
	 * The platform from which the train departs.
	 * see : https://schema.org/departurePlatform
	 *
	 * @var string | string[]
	 */
	public $departure_platform;

	/**
	 * The station from which the train departs.
	 * see : https://schema.org/departureStation
	 *
	 * @var \TrainStation | \TrainStation[]
	 */
	public $departure_station;

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
	 * The name of the train (e.g. The Orient Express).
	 * see : https://schema.org/trainName
	 *
	 * @var string | string[]
	 */
	public $train_name;

	/**
	 * The unique identifier for the train.
	 * see : https://schema.org/trainNumber
	 *
	 * @var string | string[]
	 */
	public $train_number;

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
			'@type'    => 'TrainTrip',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_platform );
		if ( ! empty( $serialized ) ) {
			$out['arrivalPlatform'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_station );
		if ( ! empty( $serialized ) ) {
			$out['arrivalStation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->arrival_time );
		if ( ! empty( $serialized ) ) {
			$out['arrivalTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_platform );
		if ( ! empty( $serialized ) ) {
			$out['departurePlatform'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_station );
		if ( ! empty( $serialized ) ) {
			$out['departureStation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->departure_time );
		if ( ! empty( $serialized ) ) {
			$out['departureTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->train_name );
		if ( ! empty( $serialized ) ) {
			$out['trainName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->train_number );
		if ( ! empty( $serialized ) ) {
			$out['trainNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
