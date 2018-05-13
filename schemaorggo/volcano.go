package schemaorggo

import "encoding/json"

// Volcano see : https://schema.org/Volcano
type Volcano struct {
	Landform

	typeContext

	// AdditionalProperty see : https://schema.org/additionalProperty
	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	//
	// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	// types : PropertyValue
	AdditionalProperty []*PropertyValue `json:"additionalProperty,omitempty"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress Text
	Address []interface{} `json:"address,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	// types : LocationFeatureSpecification
	AmenityFeature []*LocationFeatureSpecification `json:"amenityFeature,omitempty"`

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

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	// types : Event
	Event []*Event `json:"event,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber []string `json:"faxNumber,omitempty"`

	// Geo see : https://schema.org/geo
	// The geo coordinates of the place.
	// types : GeoCoordinates GeoShape
	Geo []interface{} `json:"geo,omitempty"`

	// GeospatiallyContains see : http://pending.schema.org/geospatiallyContains
	// Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. &quot;a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyContains []interface{} `json:"geospatiallyContains,omitempty"`

	// GeospatiallyCoveredBy see : http://pending.schema.org/geospatiallyCoveredBy
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyCoveredBy []interface{} `json:"geospatiallyCoveredBy,omitempty"`

	// GeospatiallyCovers see : http://pending.schema.org/geospatiallyCovers
	// Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. &quot;Every point of b is a point of (the interior or boundary of) a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyCovers []interface{} `json:"geospatiallyCovers,omitempty"`

	// GeospatiallyCrosses see : http://pending.schema.org/geospatiallyCrosses
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: &quot;a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyCrosses []interface{} `json:"geospatiallyCrosses,omitempty"`

	// GeospatiallyDisjoint see : http://pending.schema.org/geospatiallyDisjoint
	// Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM))
	// types : GeospatialGeometry Place
	GeospatiallyDisjoint []interface{} `json:"geospatiallyDisjoint,omitempty"`

	// GeospatiallyEquals see : http://pending.schema.org/geospatiallyEquals
	// Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM). &quot;Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other&quot; (a symmetric relationship)
	// types : GeospatialGeometry Place
	GeospatiallyEquals []interface{} `json:"geospatiallyEquals,omitempty"`

	// GeospatiallyIntersects see : http://pending.schema.org/geospatiallyIntersects
	// Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyIntersects []interface{} `json:"geospatiallyIntersects,omitempty"`

	// GeospatiallyOverlaps see : http://pending.schema.org/geospatiallyOverlaps
	// Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	// types : GeospatialGeometry Place
	GeospatiallyOverlaps []interface{} `json:"geospatiallyOverlaps,omitempty"`

	// GeospatiallyTouches see : http://pending.schema.org/geospatiallyTouches
	// Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM) )
	// types : GeospatialGeometry Place
	GeospatiallyTouches []interface{} `json:"geospatiallyTouches,omitempty"`

	// GeospatiallyWithin see : http://pending.schema.org/geospatiallyWithin
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

	// MaximumAttendeeCapacity see : https://schema.org/maximumAttendeeCapacity
	// The total number of individuals that may attend an event or venue.
	// types : Integer
	MaximumAttendeeCapacity []float64 `json:"maximumAttendeeCapacity,omitempty"`

	// OpeningHoursSpecification see : https://schema.org/openingHoursSpecification
	// The opening hours of a certain place.
	// types : OpeningHoursSpecification
	OpeningHoursSpecification []*OpeningHoursSpecification `json:"openingHoursSpecification,omitempty"`

	// Photo see : https://schema.org/photo
	// A photograph of this place. Supersedes photos (see: https://schema.org/photos).
	// types : ImageObject Photograph
	Photo []interface{} `json:"photo,omitempty"`

	// PublicAccess see : https://schema.org/publicAccess
	// A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
	// types : Boolean
	PublicAccess []bool `json:"publicAccess,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

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

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone []string `json:"telephone,omitempty"`
}

func (v Volcano) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.Landform.IntoMap(intop)

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

	*intop = into

	return nil
}

func (v Volcano) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Volcano"

	return data, nil
}

func (v Volcano) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
