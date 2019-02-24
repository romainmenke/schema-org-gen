<?php
namespace SchemaOrg;

// DefenceEstablishment see : https://schema.org/DefenceEstablishment
class DefenceEstablishment implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'DefenceEstablishment';

	/**
	 * With properties from GovernmentBuilding see : https://schema.org/GovernmentBuilding
	 */

	/**
	 * With properties from CivicStructure see : https://schema.org/CivicStructure
	 */

	/**
	 * With properties from Place see : https://schema.org/Place
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Place see : https://schema.org/Place
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.\n\nNote: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 *
	 * see : https://schema.org/additionalProperty
	 *
	 * @var \PropertyValue | \PropertyValue[]
	 */
	public $additional_property;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 *
	 * @var \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $address;

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
	 * An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	 * see : https://schema.org/amenityFeature
	 *
	 * @var \LocationFeatureSpecification | \LocationFeatureSpecification[]
	 */
	public $amenity_feature;

	/**
	 * A short textual code (also called &quot;store code&quot;) that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.\n\nFor example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code &quot;3047&quot; is a branchCode for a particular branch.
	 *
	 * see : https://schema.org/branchCode
	 *
	 * @var string | string[]
	 */
	public $branch_code;

	/**
	 * The basic containment relation between a place and one that contains it.
	 * see : https://schema.org/containedIn
	 *
	 * @var \Place | \Place[]
	 */
	public $contained_in;

	/**
	 * The basic containment relation between a place and one that contains it.
	 * see : https://schema.org/containedInPlace
	 *
	 * @var \Place | \Place[]
	 */
	public $contained_in_place;

	/**
	 * The basic containment relation between a place and another that it contains.
	 * see : https://schema.org/containsPlace
	 *
	 * @var \Place | \Place[]
	 */
	public $contains_place;

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
	 * Upcoming or past event associated with this place, organization, or action.
	 * see : https://schema.org/event
	 *
	 * @var \Event | \Event[]
	 */
	public $event;

	/**
	 * Upcoming or past events associated with this place or organization.
	 * see : https://schema.org/events
	 *
	 * @var \Event | \Event[]
	 */
	public $events;

	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 *
	 * @var string | string[]
	 */
	public $fax_number;

	/**
	 * The geo coordinates of the place.
	 * see : https://schema.org/geo
	 *
	 * @var \GeoCoordinates | \GeoCoordinates[] | \GeoShape | \GeoShape[]
	 */
	public $geo;

	/**
	 * The [Global Location Number](http://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 *
	 * @var string | string[]
	 */
	public $global_location_number;

	/**
	 * A URL to a map of the place.
	 * see : https://schema.org/hasMap
	 *
	 * @var string | string[] | \Map | \Map[]
	 */
	public $has_map;

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
	 * A flag to signal that the item, event, or place is accessible for free.
	 * see : https://schema.org/isAccessibleForFree
	 *
	 * @var boolean | boolean[]
	 */
	public $is_accessible_for_free;

	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 *
	 * @var string | string[]
	 */
	public $isic_v_4;

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
	 * A URL to a map of the place.
	 * see : https://schema.org/map
	 *
	 * @var string | string[]
	 */
	public $map;

	/**
	 * A URL to a map of the place.
	 * see : https://schema.org/maps
	 *
	 * @var string | string[]
	 */
	public $maps;

	/**
	 * The total number of individuals that may attend an event or venue.
	 * see : https://schema.org/maximumAttendeeCapacity
	 *
	 * @var integer | integer[]
	 */
	public $maximum_attendee_capacity;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.\n\n* Days are specified using the following two-letter combinations: ```Mo```, ```Tu```, ```We```, ```Th```, ```Fr```, ```Sa```, ```Su```.\n* Times are specified using 24:00 time. For example, 3pm is specified as ```15:00```. \n* Here is an example: &lt;code&gt;&amp;lt;time itemprop=&quot;openingHours&quot; datetime=&amp;quot;Tu,Th 16:00-20:00&amp;quot;&amp;gt;Tuesdays and Thursdays 4-8pm&amp;lt;/time&amp;gt;&lt;/code&gt;.\n* If a business is open 7 days a week, then it can be specified as &lt;code&gt;&amp;lt;time itemprop=&amp;quot;openingHours&amp;quot; datetime=&amp;quot;Mo-Su&amp;quot;&amp;gt;Monday through Sunday, all day&amp;lt;/time&amp;gt;&lt;/code&gt;.
	 * see : https://schema.org/openingHours
	 *
	 * @var string | string[]
	 */
	public $opening_hours;

	/**
	 * The opening hours of a certain place.
	 * see : https://schema.org/openingHoursSpecification
	 *
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public $opening_hours_specification;

	/**
	 * A photograph of this place.
	 * see : https://schema.org/photo
	 *
	 * @var \ImageObject | \ImageObject[] | \Photograph | \Photograph[]
	 */
	public $photo;

	/**
	 * Photographs of this place.
	 * see : https://schema.org/photos
	 *
	 * @var \ImageObject | \ImageObject[] | \Photograph | \Photograph[]
	 */
	public $photos;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * A flag to signal that the [[Place]] is open to public visitors.  If this property is omitted there is no assumed default boolean value
	 * see : https://schema.org/publicAccess
	 *
	 * @var boolean | boolean[]
	 */
	public $public_access;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * Review of the item.
	 * see : https://schema.org/reviews
	 *
	 * @var \Review | \Review[]
	 */
	public $reviews;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
	 * see : https://schema.org/smokingAllowed
	 *
	 * @var boolean | boolean[]
	 */
	public $smoking_allowed;

	/**
	 * The special opening hours of a certain place.\n\nUse this to explicitly override general opening hours brought in scope by [[openingHoursSpecification]] or [[openingHours]].
	 *
	 * see : https://schema.org/specialOpeningHoursSpecification
	 *
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public $special_opening_hours_specification;

	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 *
	 * @var string | string[]
	 */
	public $telephone;

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
			'@type'    => 'DefenceEstablishment',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->amenity_feature );
		if ( ! empty( $serialized ) ) {
			$out['amenityFeature'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->branch_code );
		if ( ! empty( $serialized ) ) {
			$out['branchCode'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contained_in );
		if ( ! empty( $serialized ) ) {
			$out['containedIn'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contained_in_place );
		if ( ! empty( $serialized ) ) {
			$out['containedInPlace'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contains_place );
		if ( ! empty( $serialized ) ) {
			$out['containsPlace'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->event );
		if ( ! empty( $serialized ) ) {
			$out['event'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->events );
		if ( ! empty( $serialized ) ) {
			$out['events'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->geo );
		if ( ! empty( $serialized ) ) {
			$out['geo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->global_location_number );
		if ( ! empty( $serialized ) ) {
			$out['globalLocationNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->has_map );
		if ( ! empty( $serialized ) ) {
			$out['hasMap'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_accessible_for_free );
		if ( ! empty( $serialized ) ) {
			$out['isAccessibleForFree'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->isic_v_4 );
		if ( ! empty( $serialized ) ) {
			$out['isicV4'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->map );
		if ( ! empty( $serialized ) ) {
			$out['map'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->maps );
		if ( ! empty( $serialized ) ) {
			$out['maps'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->maximum_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['maximumAttendeeCapacity'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->opening_hours );
		if ( ! empty( $serialized ) ) {
			$out['openingHours'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['openingHoursSpecification'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->photo );
		if ( ! empty( $serialized ) ) {
			$out['photo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->photos );
		if ( ! empty( $serialized ) ) {
			$out['photos'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->public_access );
		if ( ! empty( $serialized ) ) {
			$out['publicAccess'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->reviews );
		if ( ! empty( $serialized ) ) {
			$out['reviews'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->smoking_allowed );
		if ( ! empty( $serialized ) ) {
			$out['smokingAllowed'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->special_opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['specialOpeningHoursSpecification'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
