<?php

namespace SchemaOrg;

// BroadcastService see : https://schema.org/BroadcastService
class BroadcastService implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'BroadcastService';
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Service see : https://schema.org/Service
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | \Place | \Place[] | string | string[]
	 */
	public var $area_served;
	
	/**
	 * An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	 * see : https://schema.org/audience
	 * @var \Audience | \Audience[]
	 */
	public var $audience;
	
	/**
	 * A means of accessing the service (e.g. a phone bank, a web site, a location, etc.).
	 * see : https://schema.org/availableChannel
	 * @var \ServiceChannel | \ServiceChannel[]
	 */
	public var $available_channel;
	
	/**
	 * An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	 * see : https://schema.org/award
	 * @var string | string[]
	 */
	public var $award;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public var $brand;
	
	/**
	 * The media network(s) whose content is broadcast on this station.
	 * see : https://schema.org/broadcastAffiliateOf
	 * @var \Organization | \Organization[]
	 */
	public var $broadcast_affiliate_of;
	
	/**
	 * The name displayed in the channel guide. For many US affiliates, it is the network name.
	 * see : https://schema.org/broadcastDisplayName
	 * @var string | string[]
	 */
	public var $broadcast_display_name;
	
	/**
	 * The frequency used for over-the-air broadcasts. Numeric values or simple ranges e.g. 87-99. In addition a shortcut idiom is supported for frequences of AM and FM radio channels, e.g. &quot;87 FM&quot;.
	 * see : https://pending.schema.org/broadcastFrequency
	 * @var \BroadcastFrequencySpecification | \BroadcastFrequencySpecification[] | string | string[]
	 */
	public var $broadcast_frequency;
	
	/**
	 * The timezone in ISO 8601 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601) for which the service bases its broadcasts
	 * see : https://schema.org/broadcastTimezone
	 * @var string | string[]
	 */
	public var $broadcast_timezone;
	
	/**
	 * The organization owning or operating the broadcast service.
	 * see : https://schema.org/broadcaster
	 * @var \Organization | \Organization[]
	 */
	public var $broadcaster;
	
	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred. Supersedes bookingAgent (see: https://schema.org/bookingAgent).
	 * see : https://schema.org/broker
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $broker;
	
	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://pending.schema.org/category
	 * @var \PhysicalActivityCategory | \PhysicalActivityCategory[] | string | string[] | \Thing | \Thing[]
	 */
	public var $category;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public var $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public var $disambiguating_description;
	
	/**
	 * A broadcast channel of a broadcast service. Inverse property: providesBroadcastService (see: https://schema.org/providesBroadcastService).
	 * see : https://pending.schema.org/hasBroadcastChannel
	 * @var \BroadcastChannel | \BroadcastChannel[]
	 */
	public var $has_broadcast_channel;
	
	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 * @var \OfferCatalog | \OfferCatalog[]
	 */
	public var $has_offer_catalog;
	
	/**
	 * The hours during which this service or contact is available.
	 * see : https://schema.org/hoursAvailable
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public var $hours_available;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public var $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $image;
	
	/**
	 * A pointer to another, somehow related product (or multiple products).
	 * see : https://schema.org/isRelatedTo
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public var $is_related_to;
	
	/**
	 * A pointer to another, functionally similar product (or multiple products).
	 * see : https://schema.org/isSimilarTo
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public var $is_similar_to;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $logo;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer | \Offer[]
	 */
	public var $offers;
	
	/**
	 * A broadcast service to which the broadcast service may belong to such as regional variations of a national channel.
	 * see : https://schema.org/parentService
	 * @var \BroadcastService | \BroadcastService[]
	 */
	public var $parent_service;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	 * see : https://schema.org/provider
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $provider;
	
	/**
	 * Indicates the mobility of a provided service (e.g. &#39;static&#39;, &#39;dynamic&#39;).
	 * see : https://schema.org/providerMobility
	 * @var string | string[]
	 */
	public var $provider_mobility;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review | \Review[]
	 */
	public var $review;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The tangible thing generated by the service, e.g. a passport, permit, etc. Supersedes produces (see: https://schema.org/produces).
	 * see : https://schema.org/serviceOutput
	 * @var \Thing | \Thing[]
	 */
	public var $service_output;
	
	/**
	 * The type of service being offered, e.g. veterans&#39; benefits, emergency relief, etc.
	 * see : https://schema.org/serviceType
	 * @var string | string[]
	 */
	public var $service_type;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * Human-readable terms of service documentation.
	 * see : https://pending.schema.org/termsOfService
	 * @var string | string[]
	 */
	public var $terms_of_service;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	/**
	 * The type of screening or video broadcast used (e.g. IMAX, 3D, SD, HD, etc.).
	 * see : https://schema.org/videoFormat
	 * @var string | string[]
	 */
	public var $video_format;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'BroadcastService'
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
		
		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->available_channel );
		if ( ! empty( $serialized ) ) {
			$out['availableChannel'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broadcast_affiliate_of );
		if ( ! empty( $serialized ) ) {
			$out['broadcastAffiliateOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broadcast_display_name );
		if ( ! empty( $serialized ) ) {
			$out['broadcastDisplayName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broadcast_frequency );
		if ( ! empty( $serialized ) ) {
			$out['broadcastFrequency'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broadcast_timezone );
		if ( ! empty( $serialized ) ) {
			$out['broadcastTimezone'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broadcaster );
		if ( ! empty( $serialized ) ) {
			$out['broadcaster'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->has_broadcast_channel );
		if ( ! empty( $serialized ) ) {
			$out['hasBroadcastChannel'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->is_related_to );
		if ( ! empty( $serialized ) ) {
			$out['isRelatedTo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->is_similar_to );
		if ( ! empty( $serialized ) ) {
			$out['isSimilarTo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->service_output );
		if ( ! empty( $serialized ) ) {
			$out['serviceOutput'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->service_type );
		if ( ! empty( $serialized ) ) {
			$out['serviceType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->terms_of_service );
		if ( ! empty( $serialized ) ) {
			$out['termsOfService'] = $serialized;
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
