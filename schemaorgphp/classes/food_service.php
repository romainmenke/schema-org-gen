<?php
namespace SchemaOrg;

// FoodService see : https://schema.org/FoodService
class FoodService implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'FoodService';

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
	 * The geographic area where a service or offered item is provided.
	 * see : https://schema.org/areaServed
	 *
	 * @var \Place | \Place[] | \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | string | string[]
	 */
	public $area_served;

	/**
	 * An intended audience, i.e. a group for whom something was created.
	 * see : https://schema.org/audience
	 *
	 * @var \Audience | \Audience[]
	 */
	public $audience;

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
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 *
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public $brand;

	/**
	 * An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.
	 * see : https://schema.org/broker
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $broker;

	/**
	 * A category for the item. Greater signs or slashes can be used to informally indicate a category hierarchy.
	 * see : https://schema.org/category
	 *
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public $category;

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
	 * The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	 *
	 * see : https://schema.org/identifier
	 *
	 * @var string | string[] | \PropertyValue | \PropertyValue[]
	 */
	public $identifier;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * A pointer to another, somehow related product (or multiple products).
	 * see : https://schema.org/isRelatedTo
	 *
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public $is_related_to;

	/**
	 * A pointer to another, functionally similar product (or multiple products).
	 * see : https://schema.org/isSimilarTo
	 *
	 * @var \Product | \Product[] | \Service | \Service[]
	 */
	public $is_similar_to;

	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 *
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $logo;

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
	 * An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 *
	 * @var \Offer | \Offer[]
	 */
	public $offers;

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
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
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
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'FoodService',
		);

		$serialized = so_json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = so_json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}

		$serialized = so_json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = so_json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}

		$serialized = so_json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}

		$serialized = so_json_serialize( $this->available_channel );
		if ( ! empty( $serialized ) ) {
			$out['availableChannel'] = $serialized;
		}

		$serialized = so_json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}

		$serialized = so_json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}

		$serialized = so_json_serialize( $this->broker );
		if ( ! empty( $serialized ) ) {
			$out['broker'] = $serialized;
		}

		$serialized = so_json_serialize( $this->category );
		if ( ! empty( $serialized ) ) {
			$out['category'] = $serialized;
		}

		$serialized = so_json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = so_json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = so_json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}

		$serialized = so_json_serialize( $this->hours_available );
		if ( ! empty( $serialized ) ) {
			$out['hoursAvailable'] = $serialized;
		}

		$serialized = so_json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}

		$serialized = so_json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = so_json_serialize( $this->is_related_to );
		if ( ! empty( $serialized ) ) {
			$out['isRelatedTo'] = $serialized;
		}

		$serialized = so_json_serialize( $this->is_similar_to );
		if ( ! empty( $serialized ) ) {
			$out['isSimilarTo'] = $serialized;
		}

		$serialized = so_json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}

		$serialized = so_json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = so_json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = so_json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}

		$serialized = so_json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = so_json_serialize( $this->produces );
		if ( ! empty( $serialized ) ) {
			$out['produces'] = $serialized;
		}

		$serialized = so_json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = so_json_serialize( $this->provider_mobility );
		if ( ! empty( $serialized ) ) {
			$out['providerMobility'] = $serialized;
		}

		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = so_json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = so_json_serialize( $this->service_area );
		if ( ! empty( $serialized ) ) {
			$out['serviceArea'] = $serialized;
		}

		$serialized = so_json_serialize( $this->service_audience );
		if ( ! empty( $serialized ) ) {
			$out['serviceAudience'] = $serialized;
		}

		$serialized = so_json_serialize( $this->service_output );
		if ( ! empty( $serialized ) ) {
			$out['serviceOutput'] = $serialized;
		}

		$serialized = so_json_serialize( $this->service_type );
		if ( ! empty( $serialized ) ) {
			$out['serviceType'] = $serialized;
		}

		$serialized = so_json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
