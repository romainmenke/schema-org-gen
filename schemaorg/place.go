package schemaorg

// Place see : https://schema.org/Place
type Place struct {

Thing// AdditionalProperty see : /additionalProperty
// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
// 
// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
AdditionalProperty string `json:"additionalProperty"`

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// AmenityFeature see : /amenityFeature
// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
AmenityFeature string `json:"amenityFeature"`

// BranchCode see : /branchCode
// A short textual code (also called "store code") that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
// 
// For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code "3047" is a branchCode for a particular branch.
BranchCode string `json:"branchCode"`

// ContainedInPlace see : /containedInPlace
// The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
ContainedInPlace string `json:"containedInPlace"`

// ContainsPlace see : /containsPlace
// The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
ContainsPlace string `json:"containsPlace"`

// Event see : /event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event string `json:"event"`

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
IsAccessibleForFree string `json:"isAccessibleForFree"`

// IsicV4 see : /isicV4
// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
IsicV4 string `json:"isicV4"`

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// MaximumAttendeeCapacity see : /maximumAttendeeCapacity
// The total number of individuals that may attend an event or venue.
MaximumAttendeeCapacity string `json:"maximumAttendeeCapacity"`

// OpeningHoursSpecification see : /openingHoursSpecification
// The opening hours of a certain place.
OpeningHoursSpecification string `json:"openingHoursSpecification"`

// Photo see : /photo
// A photograph of this place. Supersedes photos (see: https://schema.org/photos).
Photo interface{} `json:"photo"` // types : ImageObject Photograph

// PublicAccess see : /publicAccess
// A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
PublicAccess string `json:"publicAccess"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// SmokingAllowed see : /smokingAllowed
// Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
SmokingAllowed string `json:"smokingAllowed"`

// SpecialOpeningHoursSpecification see : /specialOpeningHoursSpecification
// The special opening hours of a certain place.
// 
// Use this to explicitly override general opening hours brought in scope by openingHoursSpecification (see: https://schema.org/openingHoursSpecification) or openingHours (see: https://schema.org/openingHours).
SpecialOpeningHoursSpecification string `json:"specialOpeningHoursSpecification"`

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

// Identifier see : /identifier
// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
Identifier interface{} `json:"identifier"` // types : PropertyValue Text URL

// Image see : /image
// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
Image interface{} `json:"image"` // types : ImageObject URL

// MainEntityOfPage see : /mainEntityOfPage
// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
MainEntityOfPage interface{} `json:"mainEntityOfPage"` // types : CreativeWork URL

// Name see : /name
// The name of the item.
Name string `json:"name"`

// PotentialAction see : /potentialAction
// Indicates a potential Action, which describes an idealized action in which this thing would play an 'object' role.
PotentialAction string `json:"potentialAction"`

// SameAs see : /sameAs
// URL of a reference Web page that unambiguously indicates the item's identity. E.g. the URL of the item's Wikipedia page, Wikidata entry, or official website.
SameAs string `json:"sameAs"`

// Url see : /url
// URL of the item.
Url string `json:"url"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided.  Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : ContactPoint DeliveryChargeSpecification Demand Offer Organization Service

// AvailableAtOrFrom see : /availableAtOrFrom
// The place(s) from which the offer can be obtained (e.g. store locations). 
AvailableAtOrFrom interface{} `json:"availableAtOrFrom"` // types : Demand Offer

// BirthPlace see : /birthPlace
// The place where the person was born. 
BirthPlace string `json:"birthPlace"`

// ContainedInPlace see : /containedInPlace
// The basic containment relation between a place and one that contains it.  Supersedes containedIn (see: https://schema.org/containedIn). inverse property: containsPlace (see: https://schema.org/containsPlace).
ContainedInPlace string `json:"containedInPlace"`

// ContainsPlace see : /containsPlace
// The basic containment relation between a place and another that it contains.  inverse property: containedInPlace (see: https://schema.org/containedInPlace).
ContainsPlace string `json:"containsPlace"`

// ContentLocation see : /contentLocation
// The location depicted or described in the content. For example, the location in a photograph or painting. 
ContentLocation string `json:"contentLocation"`

// DeathPlace see : /deathPlace
// The place where the person died. 
DeathPlace string `json:"deathPlace"`

// DropoffLocation see : /dropoffLocation
// Where a rental car can be dropped off. 
DropoffLocation string `json:"dropoffLocation"`

// EligibleRegion see : /eligibleRegion
// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is valid.
// 
// See also ineligibleRegion (see: https://schema.org/ineligibleRegion). 
EligibleRegion interface{} `json:"eligibleRegion"` // types : DeliveryChargeSpecification Demand Offer

// ExerciseCourse see : /exerciseCourse
// A sub property of location. The course where this action was taken.  Supersedes course (see: https://schema.org/course).
ExerciseCourse string `json:"exerciseCourse"`

// FoodEstablishment see : /foodEstablishment
// A sub property of location. The specific food establishment where the action occurred. 
FoodEstablishment string `json:"foodEstablishment"`

// FoundingLocation see : /foundingLocation
// The place where the Organization was founded. 
FoundingLocation string `json:"foundingLocation"`

// FromLocation see : /fromLocation
// A sub property of location. The original location of the object or the agent before the action. 
FromLocation interface{} `json:"fromLocation"` // types : ExerciseAction MoveAction TransferAction

// GameLocation see : /gameLocation
// Real or fictional location of the game (or part of game). 
GameLocation interface{} `json:"gameLocation"` // types : Game VideoGameSeries

// HasPOS see : /hasPOS
// Points-of-Sales operated by the organization or person. 
HasPOS interface{} `json:"hasPOS"` // types : Organization Person

// HomeLocation see : /homeLocation
// A contact location for a person's residence. 
HomeLocation string `json:"homeLocation"`

// IneligibleRegion see : /ineligibleRegion
// The ISO 3166-1 (ISO 3166-1 alpha-2) or ISO 3166-2 code, the place, or the GeoShape for the geo-political region(s) for which the offer or delivery charge specification is not valid, e.g. a region where the transaction is not allowed.
// 
// See also eligibleRegion (see: https://schema.org/eligibleRegion). 
IneligibleRegion interface{} `json:"ineligibleRegion"` // types : DeliveryChargeSpecification Demand Offer

// JobLocation see : /jobLocation
// A (typically single) geographic location associated with the job position. 
JobLocation string `json:"jobLocation"`

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place. 
Location interface{} `json:"location"` // types : Action Event Organization

// LocationCreated see : /locationCreated
// The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork. 
LocationCreated string `json:"locationCreated"`

// PickupLocation see : /pickupLocation
// Where a taxi will pick up a passenger or a rental car can be picked up. 
PickupLocation interface{} `json:"pickupLocation"` // types : RentalCarReservation TaxiReservation

// RegionsAllowed see : /regionsAllowed
// The regions where the media is allowed. If not specified, then it's assumed to be allowed everywhere. Specify the countries in ISO 3166 format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_3166). 
RegionsAllowed string `json:"regionsAllowed"`

// ServiceLocation see : /serviceLocation
// The location (e.g. civic structure, local business, etc.) where a person can go to access the service. 
ServiceLocation string `json:"serviceLocation"`

// SpatialCoverage see : /spatialCoverage
// The spatialCoverage of a CreativeWork indicates the place(s) which are the focus of the content. It is a subproperty of
//       contentLocation intended primarily for more technical and detailed materials. For example with a Dataset, it indicates
//       areas that the dataset describes: a dataset of New York weather would have spatialCoverage which was the place: the state of New York.  Supersedes spatial (see: https://schema.org/spatial).
SpatialCoverage string `json:"spatialCoverage"`

// ToLocation see : /toLocation
// A sub property of location. The final location of the object or the agent after the action. 
ToLocation interface{} `json:"toLocation"` // types : ExerciseAction InsertAction MoveAction TransferAction

// WorkLocation see : /workLocation
// A contact location for a person's place of work. 
WorkLocation string `json:"workLocation"`

}

