package schemaorg

// BookStore see : https://schema.org/BookStore
type BookStore struct {

Store// CurrenciesAccepted see : /currenciesAccepted
// The currency accepted (in ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217)).
CurrenciesAccepted string `json:"currenciesAccepted"`

// OpeningHours see : /openingHours
// The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas ',' separating each day. Day or time ranges are specified using a hyphen '-'.
// 
// 
// Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
// Times are specified using 24:00 time. For example, 3pm is specified as 15:00. 
// Here is an example: <time itemprop="openingHours" datetime="Tu,Th 16:00-20:00">Tuesdays and Thursdays 4-8pm</time>.
// If a business is open 7 days a week, then it can be specified as <time itemprop="openingHours" datetime="Mo-Su">Monday through Sunday, all day</time>.
// 
// 
OpeningHours string `json:"openingHours"`

// PaymentAccepted see : /paymentAccepted
// Cash, credit card, etc.
PaymentAccepted string `json:"paymentAccepted"`

// PriceRange see : /priceRange
// The price range of the business, for example $$$.
PriceRange string `json:"priceRange"`

// ActionableFeedbackPolicy see : http://pending.schema.org/actionableFeedbackPolicy
// For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
ActionableFeedbackPolicy interface{} `json:"actionableFeedbackPolicy"` // types : CreativeWork URL

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// Alumni see : /alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni string `json:"alumni"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// ContactPoint see : /contactPoint
// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
ContactPoint string `json:"contactPoint"`

// CorrectionsPolicy see : http://pending.schema.org/correctionsPolicy
// For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
CorrectionsPolicy interface{} `json:"correctionsPolicy"` // types : CreativeWork URL

// Department see : /department
// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
Department string `json:"department"`

// DissolutionDate see : /dissolutionDate
// The date that this organization was dissolved.
DissolutionDate string `json:"dissolutionDate"`

// DiversityPolicy see : http://pending.schema.org/diversityPolicy
// Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
DiversityPolicy interface{} `json:"diversityPolicy"` // types : CreativeWork URL

// Duns see : /duns
// The Dun & Bradstreet DUNS number for identifying an organization or business person.
Duns string `json:"duns"`

// Email see : /email
// Email address.
Email string `json:"email"`

// Employee see : /employee
// Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
Employee string `json:"employee"`

// EthicsPolicy see : http://pending.schema.org/ethicsPolicy
// Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
EthicsPolicy interface{} `json:"ethicsPolicy"` // types : CreativeWork URL

// Event see : /event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event string `json:"event"`

// FaxNumber see : /faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// Founder see : /founder
// A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
Founder string `json:"founder"`

// FoundingDate see : /foundingDate
// The date that this organization was founded.
FoundingDate string `json:"foundingDate"`

// FoundingLocation see : /foundingLocation
// The place where the Organization was founded.
FoundingLocation string `json:"foundingLocation"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// GlobalLocationNumber see : /globalLocationNumber
// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
GlobalLocationNumber string `json:"globalLocationNumber"`

// HasOfferCatalog see : /hasOfferCatalog
// Indicates an OfferCatalog listing for this Organization, Person, or Service.
HasOfferCatalog string `json:"hasOfferCatalog"`

// HasPOS see : /hasPOS
// Points-of-Sales operated by the organization or person.
HasPOS string `json:"hasPOS"`

// IsicV4 see : /isicV4
// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
IsicV4 string `json:"isicV4"`

// LegalName see : /legalName
// The official name of the organization, e.g. the registered company name.
LegalName string `json:"legalName"`

// LeiCode see : /leiCode
// An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
LeiCode string `json:"leiCode"`

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place.
Location interface{} `json:"location"` // types : Place PostalAddress Text

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// MakesOffer see : /makesOffer
// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
MakesOffer string `json:"makesOffer"`

// Member see : /member
// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
Member interface{} `json:"member"` // types : Organization Person

// MemberOf see : /memberOf
// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
MemberOf interface{} `json:"memberOf"` // types : Organization ProgramMembership

// Naics see : /naics
// The North American Industry Classification System (NAICS) code for a particular organization or business person.
Naics string `json:"naics"`

// NumberOfEmployees see : /numberOfEmployees
// The number of employees in an organization e.g. business.
NumberOfEmployees string `json:"numberOfEmployees"`

// Owns see : /owns
// Products owned by the organization or person.
Owns interface{} `json:"owns"` // types : OwnershipInfo Product

// ParentOrganization see : /parentOrganization
// The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
ParentOrganization string `json:"parentOrganization"`

// PublishingPrinciples see : /publishingPrinciples
// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
// 
// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// Seeks see : /seeks
// A pointer to products or services sought by the organization or person (demand).
Seeks string `json:"seeks"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// SubOrganization see : /subOrganization
// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific 'department' property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
SubOrganization string `json:"subOrganization"`

// TaxID see : /taxID
// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
TaxID string `json:"taxID"`

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

// UnnamedSourcesPolicy see : http://pending.schema.org/unnamedSourcesPolicy
// For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
UnnamedSourcesPolicy interface{} `json:"unnamedSourcesPolicy"` // types : CreativeWork URL

// VatID see : /vatID
// The Value-added Tax ID of the organization or person.
VatID string `json:"vatID"`

// AdditionalProperty see : /additionalProperty
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

}

