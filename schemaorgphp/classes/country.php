<?php

namespace SchemaOrg;

// Country see : https://schema.org/Country
class Country implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Country';
	
	/**
	 * With properties from AdministrativeArea see : https://schema.org/AdministrativeArea
	 */
	
	/**
	 * With properties from Place see : https://schema.org/Place
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue | \PropertyValue[]
	 */
	public var $additional_property;
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 * @var \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public var $address;
	
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
	 * An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	 * see : https://schema.org/amenityFeature
	 * @var \LocationFeatureSpecification | \LocationFeatureSpecification[]
	 */
	public var $amenity_feature;
	
	/**
	 * A short textual code (also called &quot;store code&quot;) that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
	 * 
	 * For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code &quot;3047&quot; is a branchCode for a particular branch.
	 * see : https://schema.org/branchCode
	 * @var string | string[]
	 */
	public var $branch_code;
	
	/**
	 * The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
	 * see : https://schema.org/containedInPlace
	 * @var \Place | \Place[]
	 */
	public var $contained_in_place;
	
	/**
	 * The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
	 * see : https://schema.org/containsPlace
	 * @var \Place | \Place[]
	 */
	public var $contains_place;
	
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
	 * Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	 * see : https://schema.org/event
	 * @var \Event | \Event[]
	 */
	public var $event;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string | string[]
	 */
	public var $fax_number;
	
	/**
	 * The geo coordinates of the place.
	 * see : https://schema.org/geo
	 * @var \GeoCoordinates | \GeoCoordinates[] | \GeoShape | \GeoShape[]
	 */
	public var $geo;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. &quot;a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyContains
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_contains;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCoveredBy
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_covered_by;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. &quot;Every point of b is a point of (the interior or boundary of) a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCovers
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_covers;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: &quot;a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCrosses
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_crosses;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM))
	 * see : https://pending.schema.org/geospatiallyDisjoint
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_disjoint;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM). &quot;Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other&quot; (a symmetric relationship)
	 * see : https://pending.schema.org/geospatiallyEquals
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_equals;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyIntersects
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_intersects;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyOverlaps
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_overlaps;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM) )
	 * see : https://pending.schema.org/geospatiallyTouches
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_touches;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to one that contains it, i.e. it is inside (i.e. within) its interior. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyWithin
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public var $geospatially_within;
	
	/**
	 * The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 * @var string | string[]
	 */
	public var $global_location_number;
	
	/**
	 * A URL to a map of the place. Supersedes map (see: https://schema.org/map), maps (see: https://schema.org/maps).
	 * see : https://schema.org/hasMap
	 * @var \Map | \Map[] | string | string[]
	 */
	public var $has_map;
	
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
	 * A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	 * see : https://schema.org/isAccessibleForFree
	 * @var boolean | boolean[]
	 */
	public var $is_accessible_for_free;
	
	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 * @var string | string[]
	 */
	public var $isic_v_4;
	
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
	 * The total number of individuals that may attend an event or venue.
	 * see : https://schema.org/maximumAttendeeCapacity
	 * @var integer | integer[]
	 */
	public var $maximum_attendee_capacity;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * The opening hours of a certain place.
	 * see : https://schema.org/openingHoursSpecification
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public var $opening_hours_specification;
	
	/**
	 * A photograph of this place. Supersedes photos (see: https://schema.org/photos).
	 * see : https://schema.org/photo
	 * @var \ImageObject | \ImageObject[] | \Photograph | \Photograph[]
	 */
	public var $photo;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
	 * see : https://schema.org/publicAccess
	 * @var boolean | boolean[]
	 */
	public var $public_access;
	
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
	 * Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
	 * see : https://schema.org/smokingAllowed
	 * @var boolean | boolean[]
	 */
	public var $smoking_allowed;
	
	/**
	 * The special opening hours of a certain place.
	 * 
	 * Use this to explicitly override general opening hours brought in scope by openingHoursSpecification (see: https://schema.org/openingHoursSpecification) or openingHours (see: https://schema.org/openingHours).
	 * see : https://schema.org/specialOpeningHoursSpecification
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public var $special_opening_hours_specification;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string | string[]
	 */
	public var $telephone;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Country'
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
		
		$serialized = \SchemaOrg\json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geo );
		if ( ! empty( $serialized ) ) {
			$out['geo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_contains );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyContains'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_covered_by );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyCoveredBy'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_covers );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyCovers'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_crosses );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyCrosses'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_disjoint );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyDisjoint'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_equals );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyEquals'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_intersects );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyIntersects'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_overlaps );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyOverlaps'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_touches );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyTouches'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geospatially_within );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyWithin'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->maximum_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['maximumAttendeeCapacity'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['openingHoursSpecification'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->photo );
		if ( ! empty( $serialized ) ) {
			$out['photo'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
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
