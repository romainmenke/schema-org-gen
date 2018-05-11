package schemaorg

import "encoding/json"

// Volcano see : https://schema.org/Volcano
type Volcano struct {

typeContext

Landform

// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty *PropertyValue `json:"additionalProperty"`

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating *AggregateRating `json:"aggregateRating"`

// AmenityFeature see : /amenityFeature
// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
AmenityFeature *LocationFeatureSpecification `json:"amenityFeature"`

// BranchCode see : /branchCode
// A short textual code (also called "store code") that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
// 
// For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code "3047" is a branchCode for a particular branch.
BranchCode string `json:"branchCode"`

// ContainedInPlace see : /containedInPlace
// The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
ContainedInPlace *Place `json:"containedInPlace"`

// ContainsPlace see : /containsPlace
// The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
ContainsPlace *Place `json:"containsPlace"`

// Event see : /event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event *Event `json:"event"`

// FaxNumber see : /faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// Geo see : /geo
// The geo coordinates of the place.
Geo interface{} `json:"geo"` // types : GeoCoordinates GeoShape

// GeospatiallyContains see : http://pending.schema.org/geospatiallyContains
// Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. "a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a". As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyContains interface{} `json:"geospatiallyContains"` // types : GeospatialGeometry Place

// GeospatiallyCoveredBy see : http://pending.schema.org/geospatiallyCoveredBy
// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyCoveredBy interface{} `json:"geospatiallyCoveredBy"` // types : GeospatialGeometry Place

// GeospatiallyCovers see : http://pending.schema.org/geospatiallyCovers
// Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. "Every point of b is a point of (the interior or boundary of) a". As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyCovers interface{} `json:"geospatiallyCovers"` // types : GeospatialGeometry Place

// GeospatiallyCrosses see : http://pending.schema.org/geospatiallyCrosses
// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: "a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them". As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyCrosses interface{} `json:"geospatiallyCrosses"` // types : GeospatialGeometry Place

// GeospatiallyDisjoint see : http://pending.schema.org/geospatiallyDisjoint
// Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries." (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM))
GeospatiallyDisjoint interface{} `json:"geospatiallyDisjoint"` // types : GeospatialGeometry Place

// GeospatiallyEquals see : http://pending.schema.org/geospatiallyEquals
// Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM). "Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other" (a symmetric relationship)
GeospatiallyEquals interface{} `json:"geospatiallyEquals"` // types : GeospatialGeometry Place

// GeospatiallyIntersects see : http://pending.schema.org/geospatiallyIntersects
// Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyIntersects interface{} `json:"geospatiallyIntersects"` // types : GeospatialGeometry Place

// GeospatiallyOverlaps see : http://pending.schema.org/geospatiallyOverlaps
// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyOverlaps interface{} `json:"geospatiallyOverlaps"` // types : GeospatialGeometry Place

// GeospatiallyTouches see : http://pending.schema.org/geospatiallyTouches
// Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points." (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM) )
GeospatiallyTouches interface{} `json:"geospatiallyTouches"` // types : GeospatialGeometry Place

// GeospatiallyWithin see : http://pending.schema.org/geospatiallyWithin
// Represents a relationship between two geometries (or the places they represent), relating a geometry to one that contains it, i.e. it is inside (i.e. within) its interior. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
GeospatiallyWithin interface{} `json:"geospatiallyWithin"` // types : GeospatialGeometry Place

// GlobalLocationNumber see : /globalLocationNumber
// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
GlobalLocationNumber string `json:"globalLocationNumber"`

// HasMap see : /hasMap
// A URL to a map of the place. Supersedes map (see: https://schema.org/map), maps (see: https://schema.org/maps).
HasMap interface{} `json:"hasMap"` // types : Map URL

// IsAccessibleForFree see : /isAccessibleForFree
// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
IsAccessibleForFree bool `json:"isAccessibleForFree"`

// IsicV4 see : /isicV4
// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
IsicV4 string `json:"isicV4"`

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// MaximumAttendeeCapacity see : /maximumAttendeeCapacity
// The total number of individuals that may attend an event or venue.
MaximumAttendeeCapacity int `json:"maximumAttendeeCapacity"`

// OpeningHoursSpecification see : /openingHoursSpecification
// The opening hours of a certain place.
OpeningHoursSpecification *OpeningHoursSpecification `json:"openingHoursSpecification"`

// Photo see : /photo
// A photograph of this place. Supersedes photos (see: https://schema.org/photos).
Photo interface{} `json:"photo"` // types : ImageObject Photograph

// PublicAccess see : /publicAccess
// A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
PublicAccess bool `json:"publicAccess"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review *Review `json:"review"`

// SmokingAllowed see : /smokingAllowed
// Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
SmokingAllowed bool `json:"smokingAllowed"`

// SpecialOpeningHoursSpecification see : /specialOpeningHoursSpecification
// The special opening hours of a certain place.
// 
// Use this to explicitly override general opening hours brought in scope by openingHoursSpecification (see: https://schema.org/openingHoursSpecification) or openingHours (see: https://schema.org/openingHours).
SpecialOpeningHoursSpecification *OpeningHoursSpecification `json:"specialOpeningHoursSpecification"`

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

}

func (v *Volcano) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Volcano"

	return json.Marshal(v)
}
