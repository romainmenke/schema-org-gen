<?php

namespace SchemaOrg;

// LocalBusiness see : https://schema.org/LocalBusiness
class LocalBusiness implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'LocalBusiness';
	
	/**
	 * With properties from Organization see : https://schema.org/Organization
	 */
	
	/**
	 * With properties from Place see : https://schema.org/Place
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
	 * see : https://pending.schema.org/actionableFeedbackPolicy
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $actionable_feedback_policy;
	
	/**
	 * A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org.
	 * 
	 * Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	 * see : https://schema.org/additionalProperty
	 * @var \PropertyValue | \PropertyValue[]
	 */
	public $additional_property;
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public $additional_type;
	
	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 * @var \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $address;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public $aggregate_rating;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public $alternate_name;
	
	/**
	 * Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	 * see : https://schema.org/alumni
	 * @var \Person | \Person[]
	 */
	public $alumni;
	
	/**
	 * An amenity feature (e.g. a characteristic or service) of the Accommodation. This generic property does not make a statement about whether the feature is included in an offer for the main accommodation or available at extra costs.
	 * see : https://schema.org/amenityFeature
	 * @var \LocationFeatureSpecification | \LocationFeatureSpecification[]
	 */
	public $amenity_feature;
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | \Place | \Place[] | string | string[]
	 */
	public $area_served;
	
	/**
	 * An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	 * see : https://schema.org/award
	 * @var string | string[]
	 */
	public $award;
	
	/**
	 * A short textual code (also called &quot;store code&quot;) that uniquely identifies a place of business. The code is typically assigned by the parentOrganization and used in structured URLs.
	 * 
	 * For example, in the URL http://www.starbucks.co.uk/store-locator/etc/detail/3047 the code &quot;3047&quot; is a branchCode for a particular branch.
	 * see : https://schema.org/branchCode
	 * @var string | string[]
	 */
	public $branch_code;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public $brand;
	
	/**
	 * A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	 * see : https://schema.org/contactPoint
	 * @var \ContactPoint | \ContactPoint[]
	 */
	public $contact_point;
	
	/**
	 * The basic containment relation between a place and one that contains it. Supersedes containedIn (see: https://schema.org/containedIn). Inverse property: containsPlace (see: https://schema.org/containsPlace).
	 * see : https://schema.org/containedInPlace
	 * @var \Place | \Place[]
	 */
	public $contained_in_place;
	
	/**
	 * The basic containment relation between a place and another that it contains. Inverse property: containedInPlace (see: https://schema.org/containedInPlace).
	 * see : https://schema.org/containsPlace
	 * @var \Place | \Place[]
	 */
	public $contains_place;
	
	/**
	 * For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
	 * see : https://pending.schema.org/correctionsPolicy
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $corrections_policy;
	
	/**
	 * The currency accepted.
	 * 
	 * Use standard formats: ISO 4217 currency format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_4217) e.g. &quot;USD&quot;; Ticker symbol (see: https://schema.orghttps://en.wikipedia.org/wiki/List_of_cryptocurrencies) for cryptocurrencies e.g. &quot;BTC&quot;; well known names for Local Exchange Tradings Systems (see: https://schema.orghttps://en.wikipedia.org/wiki/Local_exchange_trading_system) (LETS) and other currency types e.g. &quot;Ithaca HOUR&quot;.
	 * see : https://schema.org/currenciesAccepted
	 * @var string | string[]
	 */
	public $currencies_accepted;
	
	/**
	 * A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	 * see : https://schema.org/department
	 * @var \Organization | \Organization[]
	 */
	public $department;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public $disambiguating_description;
	
	/**
	 * The date that this organization was dissolved.
	 * see : https://schema.org/dissolutionDate
	 * @var string | string[]
	 */
	public $dissolution_date;
	
	/**
	 * Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
	 * see : https://pending.schema.org/diversityPolicy
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $diversity_policy;
	
	/**
	 * For an Organization (see: https://schema.org/Organization) (often but not necessarily a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a report on staffing diversity issues. In a news context this might be for example ASNE or RTDNA (US) reports, or self-reported.
	 * see : https://pending.schema.org/diversityStaffingReport
	 * @var \Article | \Article[] | string | string[]
	 */
	public $diversity_staffing_report;
	
	/**
	 * The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	 * see : https://schema.org/duns
	 * @var string | string[]
	 */
	public $duns;
	
	/**
	 * Email address.
	 * see : https://schema.org/email
	 * @var string | string[]
	 */
	public $email;
	
	/**
	 * Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
	 * see : https://schema.org/employee
	 * @var \Person | \Person[]
	 */
	public $employee;
	
	/**
	 * Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
	 * see : https://pending.schema.org/ethicsPolicy
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $ethics_policy;
	
	/**
	 * Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	 * see : https://schema.org/event
	 * @var \Event | \Event[]
	 */
	public $event;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string | string[]
	 */
	public $fax_number;
	
	/**
	 * A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
	 * see : https://schema.org/founder
	 * @var \Person | \Person[]
	 */
	public $founder;
	
	/**
	 * The date that this organization was founded.
	 * see : https://schema.org/foundingDate
	 * @var string | string[]
	 */
	public $founding_date;
	
	/**
	 * The place where the Organization was founded.
	 * see : https://schema.org/foundingLocation
	 * @var \Place | \Place[]
	 */
	public $founding_location;
	
	/**
	 * A person or organization that supports (sponsors) something through some kind of financial contribution.
	 * see : https://schema.org/funder
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $funder;
	
	/**
	 * The geo coordinates of the place.
	 * see : https://schema.org/geo
	 * @var \GeoCoordinates | \GeoCoordinates[] | \GeoShape | \GeoShape[]
	 */
	public $geo;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a containing geometry to a contained geometry. &quot;a contains b iff no points of b lie in the exterior of a, and at least one point of the interior of b lies in the interior of a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyContains
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_contains;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that covers it. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCoveredBy
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_covered_by;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a covering geometry to a covered geometry. &quot;Every point of b is a point of (the interior or boundary of) a&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCovers
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_covers;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that crosses it: &quot;a crosses b: they have some but not all interior points in common, and the dimension of the intersection is less than that of at least one of them&quot;. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyCrosses
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_crosses;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) are topologically disjoint: they have no point in common. They form a set of disconnected geometries.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM))
	 * see : https://pending.schema.org/geospatiallyDisjoint
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_disjoint;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) are topologically equal, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM). &quot;Two geometries are topologically equal if their interiors intersect and no part of the interior or boundary of one geometry intersects the exterior of the other&quot; (a symmetric relationship)
	 * see : https://pending.schema.org/geospatiallyEquals
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_equals;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) have at least one point in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyIntersects
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_intersects;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to another that geospatially overlaps it, i.e. they have some but not all points in common. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyOverlaps
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_overlaps;
	
	/**
	 * Represents spatial relations in which two geometries (or the places they represent) touch: they have at least one boundary point in common, but no interior points.&quot; (a symmetric relationship, as defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM) )
	 * see : https://pending.schema.org/geospatiallyTouches
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_touches;
	
	/**
	 * Represents a relationship between two geometries (or the places they represent), relating a geometry to one that contains it, i.e. it is inside (i.e. within) its interior. As defined in DE-9IM (see: https://schema.orghttps://en.wikipedia.org/wiki/DE-9IM).
	 * see : https://pending.schema.org/geospatiallyWithin
	 * @var \GeospatialGeometry | \GeospatialGeometry[] | \Place | \Place[]
	 */
	public $geospatially_within;
	
	/**
	 * The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 * @var string | string[]
	 */
	public $global_location_number;
	
	/**
	 * A URL to a map of the place. Supersedes map (see: https://schema.org/map), maps (see: https://schema.org/maps).
	 * see : https://schema.org/hasMap
	 * @var \Map | \Map[] | string | string[]
	 */
	public $has_map;
	
	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 * @var \OfferCatalog | \OfferCatalog[]
	 */
	public $has_offer_catalog;
	
	/**
	 * Points-of-Sales operated by the organization or person.
	 * see : https://schema.org/hasPOS
	 * @var \Place | \Place[]
	 */
	public $has_pos;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $image;
	
	/**
	 * A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	 * see : https://schema.org/isAccessibleForFree
	 * @var boolean | boolean[]
	 */
	public $is_accessible_for_free;
	
	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 * @var string | string[]
	 */
	public $isic_v_4;
	
	/**
	 * Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a topic that is known about - suggesting possible expertise but not implying it. We do not distinguish skill levels here, or yet relate this to educational content, events, objectives or JobPosting (see: https://schema.org/JobPosting) descriptions.
	 * see : https://pending.schema.org/knowsAbout
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public $knows_about;
	
	/**
	 * Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a known language. We do not distinguish skill levels or reading/writing/speaking/signing here. Use language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	 * see : https://pending.schema.org/knowsLanguage
	 * @var \Language | \Language[] | string | string[]
	 */
	public $knows_language;
	
	/**
	 * The official name of the organization, e.g. the registered company name.
	 * see : https://schema.org/legalName
	 * @var string | string[]
	 */
	public $legal_name;
	
	/**
	 * An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
	 * see : https://schema.org/leiCode
	 * @var string | string[]
	 */
	public $lei_code;
	
	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 * @var \Place | \Place[] | \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $location;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $logo;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;
	
	/**
	 * A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	 * see : https://schema.org/makesOffer
	 * @var \Offer | \Offer[]
	 */
	public $makes_offer;
	
	/**
	 * The total number of individuals that may attend an event or venue.
	 * see : https://schema.org/maximumAttendeeCapacity
	 * @var integer | integer[]
	 */
	public $maximum_attendee_capacity;
	
	/**
	 * A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	 * see : https://schema.org/member
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $member;
	
	/**
	 * An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	 * see : https://schema.org/memberOf
	 * @var \Organization | \Organization[] | \ProgramMembership | \ProgramMembership[]
	 */
	public $member_of;
	
	/**
	 * The North American Industry Classification System (NAICS) code for a particular organization or business person.
	 * see : https://schema.org/naics
	 * @var string | string[]
	 */
	public $naics;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public $name;
	
	/**
	 * The number of employees in an organization e.g. business.
	 * see : https://schema.org/numberOfEmployees
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $number_of_employees;
	
	/**
	 * The general opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.
	 * 
	 * 
	 * Days are specified using the following two-letter combinations: Mo, Tu, We, Th, Fr, Sa, Su.
	 * Times are specified using 24:00 time. For example, 3pm is specified as 15:00. 
	 * Here is an example: &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Tu,Th 16:00-20:00&quot;&gt;Tuesdays and Thursdays 4-8pm&lt;/time&gt;.
	 * If a business is open 7 days a week, then it can be specified as &lt;time itemprop=&quot;openingHours&quot; datetime=&quot;Mo-Su&quot;&gt;Monday through Sunday, all day&lt;/time&gt;.
	 * 
	 * 
	 * see : https://schema.org/openingHours
	 * @var string | string[]
	 */
	public $opening_hours;
	
	/**
	 * The opening hours of a certain place.
	 * see : https://schema.org/openingHoursSpecification
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public $opening_hours_specification;
	
	/**
	 * For an Organization (see: https://schema.org/Organization) (often but not necessarily a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a description of organizational ownership structure; funding and grants. In a news/media setting, this is with particular reference to editorial independence.   Note that the funder (see: https://schema.org/funder) is also available and can be used to make basic funder information machine-readable.
	 * see : https://pending.schema.org/ownershipFundingInfo
	 * @var \AboutPage | \AboutPage[] | \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $ownership_funding_info;
	
	/**
	 * Products owned by the organization or person.
	 * see : https://schema.org/owns
	 * @var \OwnershipInfo | \OwnershipInfo[] | \Product | \Product[]
	 */
	public $owns;
	
	/**
	 * The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
	 * see : https://schema.org/parentOrganization
	 * @var \Organization | \Organization[]
	 */
	public $parent_organization;
	
	/**
	 * Cash, Credit Card, Cryptocurrency, Local Exchange Tradings System, etc.
	 * see : https://schema.org/paymentAccepted
	 * @var string | string[]
	 */
	public $payment_accepted;
	
	/**
	 * A photograph of this place. Supersedes photos (see: https://schema.org/photos).
	 * see : https://schema.org/photo
	 * @var \ImageObject | \ImageObject[] | \Photograph | \Photograph[]
	 */
	public $photo;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public $potential_action;
	
	/**
	 * The price range of the business, for example $$$.
	 * see : https://schema.org/priceRange
	 * @var string | string[]
	 */
	public $price_range;
	
	/**
	 * A flag to signal that the Place (see: https://schema.org/Place) is open to public visitors.  If this property is omitted there is no assumed default boolean value
	 * see : https://schema.org/publicAccess
	 * @var boolean | boolean[]
	 */
	public $public_access;
	
	/**
	 * The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	 * 
	 * While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	 * see : https://schema.org/publishingPrinciples
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $publishing_principles;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review | \Review[]
	 */
	public $review;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public $same_as;
	
	/**
	 * A pointer to products or services sought by the organization or person (demand).
	 * see : https://schema.org/seeks
	 * @var \Demand | \Demand[]
	 */
	public $seeks;
	
	/**
	 * Indicates whether it is allowed to smoke in the place, e.g. in the restaurant, hotel or hotel room.
	 * see : https://schema.org/smokingAllowed
	 * @var boolean | boolean[]
	 */
	public $smoking_allowed;
	
	/**
	 * The special opening hours of a certain place.
	 * 
	 * Use this to explicitly override general opening hours brought in scope by openingHoursSpecification (see: https://schema.org/openingHoursSpecification) or openingHours (see: https://schema.org/openingHours).
	 * see : https://schema.org/specialOpeningHoursSpecification
	 * @var \OpeningHoursSpecification | \OpeningHoursSpecification[]
	 */
	public $special_opening_hours_specification;
	
	/**
	 * A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	 * see : https://schema.org/sponsor
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $sponsor;
	
	/**
	 * A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific &#39;department&#39; property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
	 * see : https://schema.org/subOrganization
	 * @var \Organization | \Organization[]
	 */
	public $sub_organization;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;
	
	/**
	 * The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	 * see : https://schema.org/taxID
	 * @var string | string[]
	 */
	public $tax_id;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string | string[]
	 */
	public $telephone;
	
	/**
	 * For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
	 * see : https://pending.schema.org/unnamedSourcesPolicy
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $unnamed_sources_policy;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public $url;
	
	/**
	 * The Value-added Tax ID of the organization or person.
	 * see : https://schema.org/vatID
	 * @var string | string[]
	 */
	public $vat_id;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'LocalBusiness'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->actionable_feedback_policy );
		if ( ! empty( $serialized ) ) {
			$out['actionableFeedbackPolicy'] = $serialized;
		}
		
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
		
		$serialized = \SchemaOrg\json_serialize( $this->alumni );
		if ( ! empty( $serialized ) ) {
			$out['alumni'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->amenity_feature );
		if ( ! empty( $serialized ) ) {
			$out['amenityFeature'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->branch_code );
		if ( ! empty( $serialized ) ) {
			$out['branchCode'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->contact_point );
		if ( ! empty( $serialized ) ) {
			$out['contactPoint'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->contained_in_place );
		if ( ! empty( $serialized ) ) {
			$out['containedInPlace'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->contains_place );
		if ( ! empty( $serialized ) ) {
			$out['containsPlace'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->corrections_policy );
		if ( ! empty( $serialized ) ) {
			$out['correctionsPolicy'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->currencies_accepted );
		if ( ! empty( $serialized ) ) {
			$out['currenciesAccepted'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->department );
		if ( ! empty( $serialized ) ) {
			$out['department'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->dissolution_date );
		if ( ! empty( $serialized ) ) {
			$out['dissolutionDate'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->diversity_policy );
		if ( ! empty( $serialized ) ) {
			$out['diversityPolicy'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->diversity_staffing_report );
		if ( ! empty( $serialized ) ) {
			$out['diversityStaffingReport'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->duns );
		if ( ! empty( $serialized ) ) {
			$out['duns'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->employee );
		if ( ! empty( $serialized ) ) {
			$out['employee'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->ethics_policy );
		if ( ! empty( $serialized ) ) {
			$out['ethicsPolicy'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->event );
		if ( ! empty( $serialized ) ) {
			$out['event'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->founder );
		if ( ! empty( $serialized ) ) {
			$out['founder'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->founding_date );
		if ( ! empty( $serialized ) ) {
			$out['foundingDate'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->founding_location );
		if ( ! empty( $serialized ) ) {
			$out['foundingLocation'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->funder );
		if ( ! empty( $serialized ) ) {
			$out['funder'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->has_pos );
		if ( ! empty( $serialized ) ) {
			$out['hasPOS'] = $serialized;
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
		
		$serialized = \SchemaOrg\json_serialize( $this->knows_about );
		if ( ! empty( $serialized ) ) {
			$out['knowsAbout'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->knows_language );
		if ( ! empty( $serialized ) ) {
			$out['knowsLanguage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->legal_name );
		if ( ! empty( $serialized ) ) {
			$out['legalName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->lei_code );
		if ( ! empty( $serialized ) ) {
			$out['leiCode'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->makes_offer );
		if ( ! empty( $serialized ) ) {
			$out['makesOffer'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->maximum_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['maximumAttendeeCapacity'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->member );
		if ( ! empty( $serialized ) ) {
			$out['member'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->member_of );
		if ( ! empty( $serialized ) ) {
			$out['memberOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->naics );
		if ( ! empty( $serialized ) ) {
			$out['naics'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->number_of_employees );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEmployees'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->opening_hours );
		if ( ! empty( $serialized ) ) {
			$out['openingHours'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['openingHoursSpecification'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->ownership_funding_info );
		if ( ! empty( $serialized ) ) {
			$out['ownershipFundingInfo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->owns );
		if ( ! empty( $serialized ) ) {
			$out['owns'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->parent_organization );
		if ( ! empty( $serialized ) ) {
			$out['parentOrganization'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->payment_accepted );
		if ( ! empty( $serialized ) ) {
			$out['paymentAccepted'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->photo );
		if ( ! empty( $serialized ) ) {
			$out['photo'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->price_range );
		if ( ! empty( $serialized ) ) {
			$out['priceRange'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->public_access );
		if ( ! empty( $serialized ) ) {
			$out['publicAccess'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->publishing_principles );
		if ( ! empty( $serialized ) ) {
			$out['publishingPrinciples'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->seeks );
		if ( ! empty( $serialized ) ) {
			$out['seeks'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->smoking_allowed );
		if ( ! empty( $serialized ) ) {
			$out['smokingAllowed'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->special_opening_hours_specification );
		if ( ! empty( $serialized ) ) {
			$out['specialOpeningHoursSpecification'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->sponsor );
		if ( ! empty( $serialized ) ) {
			$out['sponsor'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->sub_organization );
		if ( ! empty( $serialized ) ) {
			$out['subOrganization'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->tax_id );
		if ( ! empty( $serialized ) ) {
			$out['taxID'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->unnamed_sources_policy );
		if ( ! empty( $serialized ) ) {
			$out['unnamedSourcesPolicy'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->vat_id );
		if ( ! empty( $serialized ) ) {
			$out['vatID'] = $serialized;
		}
		
		return $out;
	}
}
