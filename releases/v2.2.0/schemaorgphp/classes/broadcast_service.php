<?php
namespace SchemaOrg;

// BroadcastService see : https://schema.org/BroadcastService
class BroadcastService implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'BroadcastService';

	/**
	 * With properties from Service see : https://schema.org/Service
	 */

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
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 *
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public $aggregate_rating;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * The area within which users can expect to reach the broadcast service.
	 * see : https://schema.org/area
	 *
	 * @var \Place | \Place[]
	 */
	public $area;

	/**
	 * The geographic area where a service or offered item is provided.
	 * see : https://schema.org/areaServed
	 *
	 * @var \Place | \Place[] | \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | string | string[]
	 */
	public $area_served;

	/**
	 * A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
	 * see : https://schema.org/availableChannel
	 *
	 * @var \ServiceChannel | \ServiceChannel[]
	 */
	public $available_channel;

	/**
	 * An award won by or for this item.
	 * see : https://schema.org/award
	 *
	 * @var string | string[]
	 */
	public $award;

	/**
	 * The media network(s) whose content is broadcast on this station.
	 * see : https://schema.org/broadcastAffiliateOf
	 *
	 * @var \Organization | \Organization[]
	 */
	public $broadcast_affiliate_of;

	/**
	 * The name displayed in the channel guide. For many US affiliates, it is the network name.
	 * see : https://schema.org/broadcastDisplayName
	 *
	 * @var string | string[]
	 */
	public $broadcast_display_name;

	/**
	 * The organization owning or operating the broadcast service.
	 * see : https://schema.org/broadcaster
	 *
	 * @var \Organization | \Organization[]
	 */
	public $broadcaster;

	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://schema.org/category
	 *
	 * @var \PhysicalActivityCategory | \PhysicalActivityCategory[] | string | string[] | \Thing | \Thing[]
	 */
	public $category;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 *
	 * @var \OfferCatalog | \OfferCatalog[]
	 */
	public $has_offer_catalog;

	/**
	 * The hours during which this service or contact is available.
	 * see : https://schema.org/hoursAvailable
	 *
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public $hours_available;

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
	 * An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 *
	 * @var \Offer | \Offer[]
	 */
	public $offers;

	/**
	 * A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	 * see : https://schema.org/parentService
	 *
	 * @var \BroadcastService | \BroadcastService[]
	 */
	public $parent_service;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * The tangible thing generated by the service, e.g. a passport, permit, etc.
	 * see : https://schema.org/produces
	 *
	 * @var \Thing | \Thing[]
	 */
	public $produces;

	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	 * see : https://schema.org/provider
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $provider;

	/**
	 * Indicates the mobility of a provided service (e.g. &#39;static&#39;, &#39;dynamic&#39;).
	 * see : https://schema.org/providerMobility
	 *
	 * @var string | string[]
	 */
	public $provider_mobility;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The geographic area where the service is provided.
	 * see : https://schema.org/serviceArea
	 *
	 * @var \Place | \Place[] | \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[]
	 */
	public $service_area;

	/**
	 * The audience eligible for this service.
	 * see : https://schema.org/serviceAudience
	 *
	 * @var \Audience | \Audience[]
	 */
	public $service_audience;

	/**
	 * The tangible thing generated by the service, e.g. a passport, permit, etc.
	 * see : https://schema.org/serviceOutput
	 *
	 * @var \Thing | \Thing[]
	 */
	public $service_output;

	/**
	 * The type of service being offered, e.g. veterans&#39; benefits, emergency relief, etc.
	 * see : https://schema.org/serviceType
	 *
	 * @var string | string[]
	 */
	public $service_type;

	/**
	 * The timezone in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 format&lt;/a&gt; for which the service bases its broadcasts.
	 * see : https://schema.org/broadcastTimezone
	 *
	 * @var string | string[]
	 */
	public $timezone;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	 * see : https://schema.org/videoFormat
	 *
	 * @var string | string[]
	 */
	public $video_format;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'BroadcastService',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->area );
		if ( ! empty( $serialized ) ) {
			$out['area'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->available_channel );
		if ( ! empty( $serialized ) ) {
			$out['availableChannel'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->broadcast_affiliate_of );
		if ( ! empty( $serialized ) ) {
			$out['broadcastAffiliateOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->broadcast_display_name );
		if ( ! empty( $serialized ) ) {
			$out['broadcastDisplayName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->broadcaster );
		if ( ! empty( $serialized ) ) {
			$out['broadcaster'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->parent_service );
		if ( ! empty( $serialized ) ) {
			$out['parentService'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->produces );
		if ( ! empty( $serialized ) ) {
			$out['produces'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->provider_mobility );
		if ( ! empty( $serialized ) ) {
			$out['providerMobility'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->service_area );
		if ( ! empty( $serialized ) ) {
			$out['serviceArea'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->service_audience );
		if ( ! empty( $serialized ) ) {
			$out['serviceAudience'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->service_output );
		if ( ! empty( $serialized ) ) {
			$out['serviceOutput'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->service_type );
		if ( ! empty( $serialized ) ) {
			$out['serviceType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->timezone );
		if ( ! empty( $serialized ) ) {
			$out['timezone'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->video_format );
		if ( ! empty( $serialized ) ) {
			$out['videoFormat'] = $serialized;
		}

		return $out;
	}
}
