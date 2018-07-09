package schemaorggo

import "encoding/json"

// StadiumOrArena see : https://schema.org/StadiumOrArena
type StadiumOrArena struct {
	typeContext

	// With properties from CivicStructure see : https://schema.org/CivicStructure
	//

	// With properties from LocalBusiness see : https://schema.org/LocalBusiness
	//

	// With properties from Organization see : https://schema.org/Organization
	//

	// With properties from Place see : https://schema.org/Place
	//

	// With properties from SportsActivityLocation see : https://schema.org/SportsActivityLocation
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// ActionableFeedbackPolicy see : https://pending.schema.org/actionableFeedbackPolicy
	// For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
	// types : CreativeWork URL
	ActionableFeedbackPolicy []interface{} `json:"actionableFeedbackPolicy,omitempty"`

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

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	// types : Person
	Alumni []*Person `json:"alumni,omitempty"`

	// AmenityFeature see : https://schema.org/amenityFeature
	// An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	// types : LocationFeatureSpecification
	AmenityFeature []*LocationFeatureSpecification `json:"amenityFeature,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	// types : AdministrativeArea GeoShape Place Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award []string `json:"award,omitempty"`

	// BranchCode see : https://schema.org/branchCode
	// A short textual code (also called &quot;store code&quot;) that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
	//
	// For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code &quot;3047&quot; is a branchCode for a particular branch.
	// types : Text
	BranchCode []string `json:"branchCode,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand []interface{} `json:"brand,omitempty"`

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	// types : ContactPoint
	ContactPoint []*ContactPoint `json:"contactPoint,omitempty"`

	// ContainedInPlace see : https://schema.org/containedInPlace
	// The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
	// types : Place
	ContainedInPlace []*Place `json:"containedInPlace,omitempty"`

	// ContainsPlace see : https://schema.org/containsPlace
	// The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
	// types : Place
	ContainsPlace []*Place `json:"containsPlace,omitempty"`

	// CorrectionsPolicy see : https://pending.schema.org/correctionsPolicy
	// For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
	// types : CreativeWork URL
	CorrectionsPolicy []interface{} `json:"correctionsPolicy,omitempty"`

	// CurrenciesAccepted see : https://schema.org/currenciesAccepted
	// The currency accepted.
	//
	// Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	// types : Text
	CurrenciesAccepted []string `json:"currenciesAccepted,omitempty"`

	// Department see : https://schema.org/department
	// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	// types : Organization
	Department []*Organization `json:"department,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// DissolutionDate see : https://schema.org/dissolutionDate
	// The date that this organization was dissolved.
	// types : Date
	DissolutionDate []Date `json:"dissolutionDate,omitempty"`

	// DiversityPolicy see : https://pending.schema.org/diversityPolicy
	// Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
	// types : CreativeWork URL
	DiversityPolicy []interface{} `json:"diversityPolicy,omitempty"`

	// DiversityStaffingReport see : https://pending.schema.org/diversityStaffingReport
	// For an Organization (see: https://schema.org/Organization) (often but not necessarily a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a report on staffing diversity issues. In a news context this might be for example ASNE or RTDNA (US) reports, or self-reported.
	// types : Article URL
	DiversityStaffingReport []interface{} `json:"diversityStaffingReport,omitempty"`

	// Duns see : https://schema.org/duns
	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	// types : Text
	Duns []string `json:"duns,omitempty"`

	// Email see : https://schema.org/email
	// Email address.
	// types : Text
	Email []string `json:"email,omitempty"`

	// Employee see : https://schema.org/employee
	// Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
	// types : Person
	Employee []*Person `json:"employee,omitempty"`

	// EthicsPolicy see : https://pending.schema.org/ethicsPolicy
	// Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
	// types : CreativeWork URL
	EthicsPolicy []interface{} `json:"ethicsPolicy,omitempty"`

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	// types : Event
	Event []*Event `json:"event,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber []string `json:"faxNumber,omitempty"`

	// Founder see : https://schema.org/founder
	// A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
	// types : Person
	Founder []*Person `json:"founder,omitempty"`

	// FoundingDate see : https://schema.org/foundingDate
	// The date that this organization was founded.
	// types : Date
	FoundingDate []Date `json:"foundingDate,omitempty"`

	// FoundingLocation see : https://schema.org/foundingLocation
	// The place where the Organization was founded.
	// types : Place
	FoundingLocation []*Place `json:"foundingLocation,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder []interface{} `json:"funder,omitempty"`

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

	// HasOfferCatalog see : https://schema.org/hasOfferCatalog
	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	// types : OfferCatalog
	HasOfferCatalog []*OfferCatalog `json:"hasOfferCatalog,omitempty"`

	// HasPOS see : https://schema.org/hasPOS
	// Points-of-Sales operated by the organization or person.
	// types : Place
	HasPOS []*Place `json:"hasPOS,omitempty"`

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

	// KnowsAbout see : https://pending.schema.org/knowsAbout
	// Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a topic that is known about - suggesting possible expertise but not implying it. We do not distinguish skill levels here, or yet relate this to educational content, events, objectives or JobPosting (see: https://schema.org/JobPosting) descriptions.
	// types : Text Thing URL
	KnowsAbout []interface{} `json:"knowsAbout,omitempty"`

	// KnowsLanguage see : https://pending.schema.org/knowsLanguage
	// Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a known language. We do not distinguish skill levels or reading/writing/speaking/signing here. Use language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	// types : Language Text
	KnowsLanguage []interface{} `json:"knowsLanguage,omitempty"`

	// LegalName see : https://schema.org/legalName
	// The official name of the organization, e.g. the registered company name.
	// types : Text
	LegalName []string `json:"legalName,omitempty"`

	// LeiCode see : https://schema.org/leiCode
	// An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
	// types : Text
	LeiCode []string `json:"leiCode,omitempty"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
	Location []interface{} `json:"location,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo []interface{} `json:"logo,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	// types : Offer
	MakesOffer []*Offer `json:"makesOffer,omitempty"`

	// MaximumAttendeeCapacity see : https://schema.org/maximumAttendeeCapacity
	// The total number of individuals that may attend an event or venue.
	// types : Integer
	MaximumAttendeeCapacity []float64 `json:"maximumAttendeeCapacity,omitempty"`

	// Member see : https://schema.org/member
	// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	// types : Organization Person
	Member []interface{} `json:"member,omitempty"`

	// MemberOf see : https://schema.org/memberOf
	// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	// types : Organization ProgramMembership
	MemberOf []interface{} `json:"memberOf,omitempty"`

	// Naics see : https://schema.org/naics
	// The North American Industry Classification System (NAICS) code for a particular organization or business person.
	// types : Text
	Naics []string `json:"naics,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NumberOfEmployees see : https://schema.org/numberOfEmployees
	// The number of employees in an organization e.g. business.
	// types : QuantitativeValue
	NumberOfEmployees []*QuantitativeValue `json:"numberOfEmployees,omitempty"`

	// OpeningHours see : https://schema.org/openingHours
	// The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.
	//
	//
	// Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
	// Times are specified using 24:00 time. For example, 3pm is specified as 15:00.
	// Here is an example: &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;.
	// If a business is open 7 days a week, then it can be specified as &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;.
	//
	//
	// types : Text
	OpeningHours []string `json:"openingHours,omitempty"`

	// OpeningHoursSpecification see : https://schema.org/openingHoursSpecification
	// The opening hours of a certain place.
	// types : OpeningHoursSpecification
	OpeningHoursSpecification []*OpeningHoursSpecification `json:"openingHoursSpecification,omitempty"`

	// OwnershipFundingInfo see : https://pending.schema.org/ownershipFundingInfo
	// For an Organization (see: https://schema.org/Organization) (often but not necessarily a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a description of organizational ownership structure; funding and grants. In a news/media setting, this is with particular reference to editorial independence.   Note that the funder (see: https://schema.org/funder) is also available and can be used to make basic funder information machine-readable.
	// types : AboutPage CreativeWork Text URL
	OwnershipFundingInfo []interface{} `json:"ownershipFundingInfo,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns []interface{} `json:"owns,omitempty"`

	// ParentOrganization see : https://schema.org/parentOrganization
	// The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
	// types : Organization
	ParentOrganization []*Organization `json:"parentOrganization,omitempty"`

	// PaymentAccepted see : https://schema.org/paymentAccepted
	// Cash, Credit Card, Cryptocurrency, Local Exchange Tradings System, etc.
	// types : Text
	PaymentAccepted []string `json:"paymentAccepted,omitempty"`

	// Photo see : https://schema.org/photo
	// A photograph of this place. Supersedes photos (see: https://schema.org/photos).
	// types : ImageObject Photograph
	Photo []interface{} `json:"photo,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PriceRange see : https://schema.org/priceRange
	// The price range of the business, for example $$$.
	// types : Text
	PriceRange []string `json:"priceRange,omitempty"`

	// PublicAccess see : https://schema.org/publicAccess
	// A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
	// types : Boolean
	PublicAccess []bool `json:"publicAccess,omitempty"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	// types : CreativeWork URL
	PublishingPrinciples []interface{} `json:"publishingPrinciples,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	// types : Demand
	Seeks []*Demand `json:"seeks,omitempty"`

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

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor []interface{} `json:"sponsor,omitempty"`

	// SubOrganization see : https://schema.org/subOrganization
	// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific &#39;department&#39; property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
	// types : Organization
	SubOrganization []*Organization `json:"subOrganization,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// TaxID see : https://schema.org/taxID
	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	// types : Text
	TaxID []string `json:"taxID,omitempty"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone []string `json:"telephone,omitempty"`

	// UnnamedSourcesPolicy see : https://pending.schema.org/unnamedSourcesPolicy
	// For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
	// types : CreativeWork URL
	UnnamedSourcesPolicy []interface{} `json:"unnamedSourcesPolicy,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// VatID see : https://schema.org/vatID
	// The Value-added Tax ID of the organization or person.
	// types : Text
	VatID []string `json:"vatID,omitempty"`
}

func (v StadiumOrArena) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.ActionableFeedbackPolicy
		if len(v.ActionableFeedbackPolicy) == 1 {
			value = v.ActionableFeedbackPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actionableFeedbackPolicy"] = json.RawMessage(b)
		}
	}

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
		var value interface{} = v.Alumni
		if len(v.Alumni) == 1 {
			value = v.Alumni[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alumni"] = json.RawMessage(b)
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
		var value interface{} = v.AreaServed
		if len(v.AreaServed) == 1 {
			value = v.AreaServed[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["areaServed"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Award
		if len(v.Award) == 1 {
			value = v.Award[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["award"] = json.RawMessage(b)
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
		var value interface{} = v.Brand
		if len(v.Brand) == 1 {
			value = v.Brand[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["brand"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContactPoint
		if len(v.ContactPoint) == 1 {
			value = v.ContactPoint[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contactPoint"] = json.RawMessage(b)
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
		var value interface{} = v.CorrectionsPolicy
		if len(v.CorrectionsPolicy) == 1 {
			value = v.CorrectionsPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["correctionsPolicy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CurrenciesAccepted
		if len(v.CurrenciesAccepted) == 1 {
			value = v.CurrenciesAccepted[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["currenciesAccepted"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Department
		if len(v.Department) == 1 {
			value = v.Department[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["department"] = json.RawMessage(b)
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
		var value interface{} = v.DissolutionDate
		if len(v.DissolutionDate) == 1 {
			value = v.DissolutionDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dissolutionDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DiversityPolicy
		if len(v.DiversityPolicy) == 1 {
			value = v.DiversityPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["diversityPolicy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DiversityStaffingReport
		if len(v.DiversityStaffingReport) == 1 {
			value = v.DiversityStaffingReport[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["diversityStaffingReport"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Duns
		if len(v.Duns) == 1 {
			value = v.Duns[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["duns"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Email
		if len(v.Email) == 1 {
			value = v.Email[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["email"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Employee
		if len(v.Employee) == 1 {
			value = v.Employee[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["employee"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EthicsPolicy
		if len(v.EthicsPolicy) == 1 {
			value = v.EthicsPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ethicsPolicy"] = json.RawMessage(b)
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
		var value interface{} = v.Founder
		if len(v.Founder) == 1 {
			value = v.Founder[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["founder"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FoundingDate
		if len(v.FoundingDate) == 1 {
			value = v.FoundingDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["foundingDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FoundingLocation
		if len(v.FoundingLocation) == 1 {
			value = v.FoundingLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["foundingLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Funder
		if len(v.Funder) == 1 {
			value = v.Funder[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["funder"] = json.RawMessage(b)
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
		var value interface{} = v.HasOfferCatalog
		if len(v.HasOfferCatalog) == 1 {
			value = v.HasOfferCatalog[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasOfferCatalog"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasPOS
		if len(v.HasPOS) == 1 {
			value = v.HasPOS[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasPOS"] = json.RawMessage(b)
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
		var value interface{} = v.KnowsAbout
		if len(v.KnowsAbout) == 1 {
			value = v.KnowsAbout[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["knowsAbout"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.KnowsLanguage
		if len(v.KnowsLanguage) == 1 {
			value = v.KnowsLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["knowsLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LegalName
		if len(v.LegalName) == 1 {
			value = v.LegalName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["legalName"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LeiCode
		if len(v.LeiCode) == 1 {
			value = v.LeiCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["leiCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Location
		if len(v.Location) == 1 {
			value = v.Location[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["location"] = json.RawMessage(b)
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
		var value interface{} = v.MakesOffer
		if len(v.MakesOffer) == 1 {
			value = v.MakesOffer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["makesOffer"] = json.RawMessage(b)
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
		var value interface{} = v.Member
		if len(v.Member) == 1 {
			value = v.Member[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["member"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MemberOf
		if len(v.MemberOf) == 1 {
			value = v.MemberOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["memberOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Naics
		if len(v.Naics) == 1 {
			value = v.Naics[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["naics"] = json.RawMessage(b)
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
		var value interface{} = v.NumberOfEmployees
		if len(v.NumberOfEmployees) == 1 {
			value = v.NumberOfEmployees[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfEmployees"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.OpeningHours
		if len(v.OpeningHours) == 1 {
			value = v.OpeningHours[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["openingHours"] = json.RawMessage(b)
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
		var value interface{} = v.OwnershipFundingInfo
		if len(v.OwnershipFundingInfo) == 1 {
			value = v.OwnershipFundingInfo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["ownershipFundingInfo"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Owns
		if len(v.Owns) == 1 {
			value = v.Owns[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["owns"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ParentOrganization
		if len(v.ParentOrganization) == 1 {
			value = v.ParentOrganization[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["parentOrganization"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PaymentAccepted
		if len(v.PaymentAccepted) == 1 {
			value = v.PaymentAccepted[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["paymentAccepted"] = json.RawMessage(b)
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
		var value interface{} = v.PriceRange
		if len(v.PriceRange) == 1 {
			value = v.PriceRange[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["priceRange"] = json.RawMessage(b)
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
		var value interface{} = v.PublishingPrinciples
		if len(v.PublishingPrinciples) == 1 {
			value = v.PublishingPrinciples[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publishingPrinciples"] = json.RawMessage(b)
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
		var value interface{} = v.Seeks
		if len(v.Seeks) == 1 {
			value = v.Seeks[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["seeks"] = json.RawMessage(b)
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
		var value interface{} = v.Sponsor
		if len(v.Sponsor) == 1 {
			value = v.Sponsor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sponsor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SubOrganization
		if len(v.SubOrganization) == 1 {
			value = v.SubOrganization[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["subOrganization"] = json.RawMessage(b)
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
		var value interface{} = v.TaxID
		if len(v.TaxID) == 1 {
			value = v.TaxID[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["taxID"] = json.RawMessage(b)
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
		var value interface{} = v.UnnamedSourcesPolicy
		if len(v.UnnamedSourcesPolicy) == 1 {
			value = v.UnnamedSourcesPolicy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["unnamedSourcesPolicy"] = json.RawMessage(b)
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

	{
		var value interface{} = v.VatID
		if len(v.VatID) == 1 {
			value = v.VatID[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["vatID"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v StadiumOrArena) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "StadiumOrArena"

	return data, nil
}

func (v StadiumOrArena) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
