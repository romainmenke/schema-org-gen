package schemaorggo

import "encoding/json"

// Suite see : https://schema.org/Suite
type Suite struct {
	typeContext

	// With properties from Accommodation see : https://schema.org/Accommodation
	//

	// With properties from Place see : https://schema.org/Place
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalProperty see : https://schema.org/additionalProperty
	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	//
	// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	// types : PropertyValue
	AdditionalProperty []*PropertyValue `json:"additionalProperty,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress Text
	Address []interface{} `json:"address,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	// types : LocationFeatureSpecification
	AmenityFeature []*LocationFeatureSpecification `json:"amenityFeature,omitempty"`

	// Bed see : https://schema.org/bed
	// The type of bed or beds included in the accommodation. For the single case of just one bed of a certain type, you use bed directly with a text.
	//       If you want to indicate the quantity of a certain kind of bed, use an instance of BedDetails. For more detailed information, use the amenityFeature property.
	// types : BedDetails BedType Text
	Bed []interface{} `json:"bed,omitempty"`

	// BranchCode see : https://schema.org/branchCode
	// A short textual code (also called &quot;store code&quot;) that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
	//
	// For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code &quot;3047&quot; is a branchCode for a particular branch.
	// types : Text
	BranchCode []string `json:"branchCode,omitempty"`

	// ContainedInPlace see : https://schema.org/containedInPlace
	// The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
	// types : Place
	ContainedInPlace []*Place `json:"containedInPlace,omitempty"`

	// ContainsPlace see : https://schema.org/containsPlace
	// The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
	// types : Place
	ContainsPlace []*Place `json:"containsPlace,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	// types : Event
	Event []*Event `json:"event,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber []string `json:"faxNumber,omitempty"`

	// FloorSize see : https://schema.org/floorSize
	// The size of the accommodation, e.g. in square meter or squarefoot.
	// Typical unit code(s): MTK for square meter, FTK for square foot, or YDK for square yard
	// types : QuantitativeValue
	FloorSize []*QuantitativeValue `json:"floorSize,omitempty"`

	// Geo see : https://schema.org/geo
	// The geo coordinates of the place.
	// types : GeoCoordinates GeoShape
	Geo []interface{} `json:"geo,omitempty"`

	// GeospatiallyContains see : https://pending.schema.org/geospatiallyContains
	// Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. &quot;a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyContains []interface{} `json:"geospatiallyContains,omitempty"`

	// GeospatiallyCoveredBy see : https://pending.schema.org/geospatiallyCoveredBy
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyCoveredBy []interface{} `json:"geospatiallyCoveredBy,omitempty"`

	// GeospatiallyCovers see : https://pending.schema.org/geospatiallyCovers
	// Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. &quot;Every point of b is a point of (the interior or boundary of) a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyCovers []interface{} `json:"geospatiallyCovers,omitempty"`

	// GeospatiallyCrosses see : https://pending.schema.org/geospatiallyCrosses
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: &quot;a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyCrosses []interface{} `json:"geospatiallyCrosses,omitempty"`

	// GeospatiallyDisjoint see : https://pending.schema.org/geospatiallyDisjoint
	// Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM))
	// types : GeospatialGeometry Place
	GeospatiallyDisjoint []interface{} `json:"geospatiallyDisjoint,omitempty"`

	// GeospatiallyEquals see : https://pending.schema.org/geospatiallyEquals
	// Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM). &quot;Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other&quot; (a symmetric relationship)
	// types : GeospatialGeometry Place
	GeospatiallyEquals []interface{} `json:"geospatiallyEquals,omitempty"`

	// GeospatiallyIntersects see : https://pending.schema.org/geospatiallyIntersects
	// Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyIntersects []interface{} `json:"geospatiallyIntersects,omitempty"`

	// GeospatiallyOverlaps see : https://pending.schema.org/geospatiallyOverlaps
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyOverlaps []interface{} `json:"geospatiallyOverlaps,omitempty"`

	// GeospatiallyTouches see : https://pending.schema.org/geospatiallyTouches
	// Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM) )
	// types : GeospatialGeometry Place
	GeospatiallyTouches []interface{} `json:"geospatiallyTouches,omitempty"`

	// GeospatiallyWithin see : https://pending.schema.org/geospatiallyWithin
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to one that contains it, i.e. it is inside (i.e. within) its interior. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyWithin []interface{} `json:"geospatiallyWithin,omitempty"`

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	// types : Text
	GlobalLocationNumber []string `json:"globalLocationNumber,omitempty"`

	// HasMap see : https://schema.org/hasMap
	// A URL to a map of the place. Supersedes map (see: https://schema.org/map), maps (see: https://schema.org/maps).
	// types : Map URL
	HasMap []interface{} `json:"hasMap,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	// types : PropertyValue Text URL
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	// types : ImageObject URL
	Image []interface{} `json:"image,omitempty"`

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	// types : Boolean
	IsAccessibleForFree []bool `json:"isAccessibleForFree,omitempty"`

	// IsicV4 see : https://schema.org/isicV4
	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	// types : Text
	IsicV4 []string `json:"isicV4,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo []interface{} `json:"logo,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MaximumAttendeeCapacity see : https://schema.org/maximumAttendeeCapacity
	// The total number of individuals that may attend an event or venue.
	// types : Integer
	MaximumAttendeeCapacity []float64 `json:"maximumAttendeeCapacity,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NumberOfRooms see : https://schema.org/numberOfRooms
	// The number of rooms (excluding bathrooms and closets) of the accommodation or lodging business.
	// Typical unit code(s): ROM for room or C62 for no unit. The type of room can be put in the unitText property of the QuantitativeValue.
	// types : Number QuantitativeValue
	NumberOfRooms []interface{} `json:"numberOfRooms,omitempty"`

	// Occupancy see : https://schema.org/occupancy
	// The allowed total occupancy for the accommodation in persons (including infants etc). For individual accommodations, this is not necessarily the legal maximum but defines the permitted usage as per the contractual agreement (e.g. a double room used by a single person).
	// Typical unit code(s): C62 for person
	// types : QuantitativeValue
	Occupancy []*QuantitativeValue `json:"occupancy,omitempty"`

	// OpeningHoursSpecification see : https://schema.org/openingHoursSpecification
	// The opening hours of a certain place.
	// types : OpeningHoursSpecification
	OpeningHoursSpecification []*OpeningHoursSpecification `json:"openingHoursSpecification,omitempty"`

	// PermittedUsage see : https://schema.org/permittedUsage
	// Indications regarding the permitted usage of the accommodation.
	// types : Text
	PermittedUsage []string `json:"permittedUsage,omitempty"`

	// PetsAllowed see : https://schema.org/petsAllowed
	// Indicates whether pets are allowed to enter the accommodation or lodging business. More detailed information can be put in a text value.
	// types : Boolean Text
	PetsAllowed []interface{} `json:"petsAllowed,omitempty"`

	// Photo see : https://schema.org/photo
	// A photograph of this place. Supersedes photos (see: https://schema.org/photos).
	// types : ImageObject Photograph
	Photo []interface{} `json:"photo,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PublicAccess see : https://schema.org/publicAccess
	// A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
	// types : Boolean
	PublicAccess []bool `json:"publicAccess,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SmokingAllowed see : https://schema.org/smokingAllowed
	// Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
	// types : Boolean
	SmokingAllowed []bool `json:"smokingAllowed,omitempty"`

	// SpecialOpeningHoursSpecification see : https://schema.org/specialOpeningHoursSpecification
	// The special opening hours of a certain place.
	//
	// Use this to explicitly override general opening hours brought in scope by openingHoursSpecification (see: https://schema.org/openingHoursSpecification) or openingHours (see: https://schema.org/openingHours).
	// types : OpeningHoursSpecification
	SpecialOpeningHoursSpecification []*OpeningHoursSpecification `json:"specialOpeningHoursSpecification,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone []string `json:"telephone,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`
}

func (v Suite) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.AdditionalProperty
		if len(v.AdditionalProperty) == 1 {
			value = v.AdditionalProperty[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalProperty"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AdditionalType
		if len(v.AdditionalType) == 1 {
			value = v.AdditionalType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Address
		if len(v.Address) == 1 {
			value = v.Address[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["address"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AggregateRating
		if len(v.AggregateRating) == 1 {
			value = v.AggregateRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["aggregateRating"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AlternateName
		if len(v.AlternateName) == 1 {
			value = v.AlternateName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alternateName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AmenityFeature
		if len(v.AmenityFeature) == 1 {
			value = v.AmenityFeature[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["amenityFeature"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Bed
		if len(v.Bed) == 1 {
			value = v.Bed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["bed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BranchCode
		if len(v.BranchCode) == 1 {
			value = v.BranchCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["branchCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContainedInPlace
		if len(v.ContainedInPlace) == 1 {
			value = v.ContainedInPlace[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["containedInPlace"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContainsPlace
		if len(v.ContainsPlace) == 1 {
			value = v.ContainsPlace[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["containsPlace"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Description
		if len(v.Description) == 1 {
			value = v.Description[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["description"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DisambiguatingDescription
		if len(v.DisambiguatingDescription) == 1 {
			value = v.DisambiguatingDescription[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["disambiguatingDescription"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Event
		if len(v.Event) == 1 {
			value = v.Event[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["event"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FaxNumber
		if len(v.FaxNumber) == 1 {
			value = v.FaxNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["faxNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FloorSize
		if len(v.FloorSize) == 1 {
			value = v.FloorSize[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["floorSize"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Geo
		if len(v.Geo) == 1 {
			value = v.Geo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geo"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyContains
		if len(v.GeospatiallyContains) == 1 {
			value = v.GeospatiallyContains[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyContains"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyCoveredBy
		if len(v.GeospatiallyCoveredBy) == 1 {
			value = v.GeospatiallyCoveredBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyCoveredBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyCovers
		if len(v.GeospatiallyCovers) == 1 {
			value = v.GeospatiallyCovers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyCovers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyCrosses
		if len(v.GeospatiallyCrosses) == 1 {
			value = v.GeospatiallyCrosses[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyCrosses"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyDisjoint
		if len(v.GeospatiallyDisjoint) == 1 {
			value = v.GeospatiallyDisjoint[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyDisjoint"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyEquals
		if len(v.GeospatiallyEquals) == 1 {
			value = v.GeospatiallyEquals[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyEquals"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyIntersects
		if len(v.GeospatiallyIntersects) == 1 {
			value = v.GeospatiallyIntersects[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyIntersects"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyOverlaps
		if len(v.GeospatiallyOverlaps) == 1 {
			value = v.GeospatiallyOverlaps[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyOverlaps"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyTouches
		if len(v.GeospatiallyTouches) == 1 {
			value = v.GeospatiallyTouches[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyTouches"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GeospatiallyWithin
		if len(v.GeospatiallyWithin) == 1 {
			value = v.GeospatiallyWithin[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["geospatiallyWithin"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GlobalLocationNumber
		if len(v.GlobalLocationNumber) == 1 {
			value = v.GlobalLocationNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["globalLocationNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasMap
		if len(v.HasMap) == 1 {
			value = v.HasMap[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasMap"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Identifier
		if len(v.Identifier) == 1 {
			value = v.Identifier[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["identifier"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Image
		if len(v.Image) == 1 {
			value = v.Image[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["image"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsAccessibleForFree
		if len(v.IsAccessibleForFree) == 1 {
			value = v.IsAccessibleForFree[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isAccessibleForFree"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsicV4
		if len(v.IsicV4) == 1 {
			value = v.IsicV4[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isicV4"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Logo
		if len(v.Logo) == 1 {
			value = v.Logo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["logo"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MainEntityOfPage
		if len(v.MainEntityOfPage) == 1 {
			value = v.MainEntityOfPage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mainEntityOfPage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MaximumAttendeeCapacity
		if len(v.MaximumAttendeeCapacity) == 1 {
			value = v.MaximumAttendeeCapacity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["maximumAttendeeCapacity"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Name
		if len(v.Name) == 1 {
			value = v.Name[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["name"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfRooms
		if len(v.NumberOfRooms) == 1 {
			value = v.NumberOfRooms[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfRooms"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Occupancy
		if len(v.Occupancy) == 1 {
			value = v.Occupancy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["occupancy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OpeningHoursSpecification
		if len(v.OpeningHoursSpecification) == 1 {
			value = v.OpeningHoursSpecification[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["openingHoursSpecification"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PermittedUsage
		if len(v.PermittedUsage) == 1 {
			value = v.PermittedUsage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["permittedUsage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PetsAllowed
		if len(v.PetsAllowed) == 1 {
			value = v.PetsAllowed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["petsAllowed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Photo
		if len(v.Photo) == 1 {
			value = v.Photo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["photo"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PotentialAction
		if len(v.PotentialAction) == 1 {
			value = v.PotentialAction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["potentialAction"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PublicAccess
		if len(v.PublicAccess) == 1 {
			value = v.PublicAccess[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publicAccess"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Review
		if len(v.Review) == 1 {
			value = v.Review[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["review"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SameAs
		if len(v.SameAs) == 1 {
			value = v.SameAs[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sameAs"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SmokingAllowed
		if len(v.SmokingAllowed) == 1 {
			value = v.SmokingAllowed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["smokingAllowed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SpecialOpeningHoursSpecification
		if len(v.SpecialOpeningHoursSpecification) == 1 {
			value = v.SpecialOpeningHoursSpecification[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["specialOpeningHoursSpecification"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SubjectOf
		if len(v.SubjectOf) == 1 {
			value = v.SubjectOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subjectOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Telephone
		if len(v.Telephone) == 1 {
			value = v.Telephone[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["telephone"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Url
		if len(v.Url) == 1 {
			value = v.Url[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["url"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Suite) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Suite"

	return data, nil
}

func (v Suite) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
