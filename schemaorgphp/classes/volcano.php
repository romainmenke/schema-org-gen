<?php

class Volcano extends Landform implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Volcano';
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue|\PropertyValue[]
	 */
	public var $additional_property;
	
	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 * @var \PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $address;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating|\AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	 * see : https://schema.org/amenityFeature
	 * @var \LocationFeatureSpecification|\LocationFeatureSpecification[]
	 */
	public var $amenity_feature;
	
	/**
	 * A short textual code (also called &quot;store code&quot;) that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
	 * 
	 * For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code &quot;3047&quot; is a branchCode for a particular branch.
	 * see : https://schema.org/branchCode
	 * @var string|string[]
	 */
	public var $branch_code;
	
	/**
	 * The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
	 * see : https://schema.org/containedInPlace
	 * @var \Place|\Place[]
	 */
	public var $contained_in_place;
	
	/**
	 * The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
	 * see : https://schema.org/containsPlace
	 * @var \Place|\Place[]
	 */
	public var $contains_place;
	
	/**
	 * Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	 * see : https://schema.org/event
	 * @var \Event|\Event[]
	 */
	public var $event;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string|string[]
	 */
	public var $fax_number;
	
	/**
	 * The geo coordinates of the place.
	 * see : https://schema.org/geo
	 * @var \GeoCoordinates|\GeoCoordinates[]|\GeoShape|\GeoShape[]
	 */
	public var $geo;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. &quot;a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyContains
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_contains;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCoveredBy
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_covered_by;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. &quot;Every point of b is a point of (the interior or boundary of) a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCovers
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_covers;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: &quot;a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCrosses
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_crosses;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM))
	 * see : https://pending.schema.org/geospatiallyDisjoint
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_disjoint;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM). &quot;Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other&quot; (a symmetric relationship)
	 * see : https://pending.schema.org/geospatiallyEquals
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_equals;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyIntersects
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_intersects;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyOverlaps
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_overlaps;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM) )
	 * see : https://pending.schema.org/geospatiallyTouches
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_touches;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to one that contains it, i.e. it is inside (i.e. within) its interior. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyWithin
	 * @var \GeospatialGeometry|\GeospatialGeometry[]|\Place|\Place[]
	 */
	public var $geospatially_within;
	
	/**
	 * The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 * @var string|string[]
	 */
	public var $global_location_number;
	
	/**
	 * A URL to a map of the place. Supersedes map (see: https://schema.org/map), maps (see: https://schema.org/maps).
	 * see : https://schema.org/hasMap
	 * @var \Map|\Map[]|string|string[]
	 */
	public var $has_map;
	
	/**
	 * A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	 * see : https://schema.org/isAccessibleForFree
	 * @var boolean|boolean[]
	 */
	public var $is_accessible_for_free;
	
	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 * @var string|string[]
	 */
	public var $isic_v_4;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject|\ImageObject[]|string|string[]
	 */
	public var $logo;
	
	/**
	 * The total number of individuals that may attend an event or venue.
	 * see : https://schema.org/maximumAttendeeCapacity
	 * @var integer|integer[]
	 */
	public var $maximum_attendee_capacity;
	
	/**
	 * The opening hours of a certain place.
	 * see : https://schema.org/openingHoursSpecification
	 * @var \OpeningHoursSpecification|\OpeningHoursSpecification[]
	 */
	public var $opening_hours_specification;
	
	/**
	 * A photograph of this place. Supersedes photos (see: https://schema.org/photos).
	 * see : https://schema.org/photo
	 * @var \ImageObject|\ImageObject[]|\Photograph|\Photograph[]
	 */
	public var $photo;
	
	/**
	 * A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
	 * see : https://schema.org/publicAccess
	 * @var boolean|boolean[]
	 */
	public var $public_access;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review|\Review[]
	 */
	public var $review;
	
	/**
	 * Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
	 * see : https://schema.org/smokingAllowed
	 * @var boolean|boolean[]
	 */
	public var $smoking_allowed;
	
	/**
	 * The special opening hours of a certain place.
	 * 
	 * Use this to explicitly override general opening hours brought in scope by openingHoursSpecification (see: https://schema.org/openingHoursSpecification) or openingHours (see: https://schema.org/openingHours).
	 * see : https://schema.org/specialOpeningHoursSpecification
	 * @var \OpeningHoursSpecification|\OpeningHoursSpecification[]
	 */
	public var $special_opening_hours_specification;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string|string[]
	 */
	public var $telephone;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Volcano'
		);
		
		$serialized = so_json_serialize( $this->additional_property );
		if ( ! empty( $serialized ) ) {
			$out['additionalProperty'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->amenity_feature );
		if ( ! empty( $serialized ) ) {
			$out['amenityFeature'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->branch_code );
		if ( ! empty( $serialized ) ) {
			$out['branchCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contained_in_place );
		if ( ! empty( $serialized ) ) {
			$out['containedInPlace'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contains_place );
		if ( ! empty( $serialized ) ) {
			$out['containsPlace'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->event );
		if ( ! empty( $serialized ) ) {
			$out['event'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geo );
		if ( ! empty( $serialized ) ) {
			$out['geo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_contains );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyContains'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_covered_by );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyCoveredBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_covers );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyCovers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_crosses );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyCrosses'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_disjoint );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyDisjoint'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_equals );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyEquals'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_intersects );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyIntersects'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_overlaps );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyOverlaps'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_touches );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyTouches'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->geospatially_within );
		if ( ! empty( $serialized ) ) {
			$out['geospatiallyWithin'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->global_location_number );
		if ( ! empty( $serialized ) ) {
			$out['globalLocationNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_map );
		if ( ! empty( $serialized ) ) {
			$out['hasMap'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_accessible_for_free );
		if ( ! empty( $serialized ) ) {
			$out['isAccessibleForFree'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->isic_v_4 );
		if ( ! empty( $serialized ) ) {
			$out['isicV4'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->maximum_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['maximumAttendeeCapacity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['openingHoursSpecification'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->photo );
		if ( ! empty( $serialized ) ) {
			$out['photo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->public_access );
		if ( ! empty( $serialized ) ) {
			$out['publicAccess'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->smoking_allowed );
		if ( ! empty( $serialized ) ) {
			$out['smokingAllowed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->special_opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['specialOpeningHoursSpecification'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}
		
		return $out;
	}
}
