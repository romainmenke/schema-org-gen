package schemaorggo

import "encoding/json"

// FurnitureStore see : https://schema.org/FurnitureStore
type FurnitureStore struct {

	// With properties from Store see : https://schema.org/Store
	//

	// With properties from LocalBusiness see : https://schema.org/LocalBusiness
	//

	// With properties from Organization see : https://schema.org/Organization
	//

	// With properties from Place see : https://schema.org/Place
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Organization see : https://schema.org/Organization
	//

	// With properties from Place see : https://schema.org/Place
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalProperty see : https://schema.org/additionalProperty
	// A property-value pair representing an additional characteristics of the entitity, e.g. a product feature or another characteristic for which there is no matching property in schema.org. &lt;br /&gt;&lt;br /&gt;
	//
	// Note: Publishers should be aware that applications designed to use specific schema.org properties (e.g. http://schema.org/width, http://schema.org/color, http://schema.org/gtin13, ...) will typically expect such data to be provided using those properties, rather than using the generic property/value mechanism.
	//
	// types : PropertyValue
	AdditionalProperty []*PropertyValue `json:"additionalProperty,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress
	Address []*PostalAddress `json:"address,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// Awards see : https://schema.org/awards
	// Awards won by or for this item.
	// types : Text
	Awards []string `json:"awards,omitempty"`

	// BranchOf see : https://schema.org/branchOf
	// The larger organization that this local business is a branch of, if any.
	// types : Organization
	BranchOf []*Organization `json:"branchOf,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand []interface{} `json:"brand,omitempty"`

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization.
	// types : ContactPoint
	ContactPoint []*ContactPoint `json:"contactPoint,omitempty"`

	// ContactPoints see : https://schema.org/contactPoints
	// A contact point for a person or organization.
	// types : ContactPoint
	ContactPoints []*ContactPoint `json:"contactPoints,omitempty"`

	// ContainedIn see : https://schema.org/containedIn
	// The basic containment relation between places.
	// types : Place
	ContainedIn []*Place `json:"containedIn,omitempty"`

	// CurrenciesAccepted see : https://schema.org/currenciesAccepted
	// The currency accepted (in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_4217&#39;&gt;ISO 4217 currency format&lt;/a&gt;).
	// types : Text
	CurrenciesAccepted []string `json:"currenciesAccepted,omitempty"`

	// Department see : https://schema.org/department
	// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	// types : Organization
	Department []*Organization `json:"department,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DissolutionDate see : https://schema.org/dissolutionDate
	// The date that this organization was dissolved.
	// types : Date
	DissolutionDate []Date `json:"dissolutionDate,omitempty"`

	// Duns see : https://schema.org/duns
	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	// types : Text
	Duns []string `json:"duns,omitempty"`

	// Email see : https://schema.org/email
	// Email address.
	// types : Text
	Email []string `json:"email,omitempty"`

	// Employee see : https://schema.org/employee
	// Someone working for this organization.
	// types : Person
	Employee []*Person `json:"employee,omitempty"`

	// Employees see : https://schema.org/employees
	// People working for this organization.
	// types : Person
	Employees []*Person `json:"employees,omitempty"`

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action.
	// types : Event
	Event []*Event `json:"event,omitempty"`

	// Events see : https://schema.org/events
	// Upcoming or past events associated with this place or organization.
	// types : Event
	Events []*Event `json:"events,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber []string `json:"faxNumber,omitempty"`

	// Founder see : https://schema.org/founder
	// A person who founded this organization.
	// types : Person
	Founder []*Person `json:"founder,omitempty"`

	// Founders see : https://schema.org/founders
	// A person who founded this organization.
	// types : Person
	Founders []*Person `json:"founders,omitempty"`

	// FoundingDate see : https://schema.org/foundingDate
	// The date that this organization was founded.
	// types : Date
	FoundingDate []Date `json:"foundingDate,omitempty"`

	// FoundingLocation see : https://schema.org/foundingLocation
	// The place where the Organization was founded.
	// types : Place
	FoundingLocation []*Place `json:"foundingLocation,omitempty"`

	// Geo see : https://schema.org/geo
	// The geo coordinates of the place.
	// types : GeoCoordinates GeoShape
	Geo []interface{} `json:"geo,omitempty"`

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The &lt;a href=&quot;http://www.gs1.org/gln&quot;&gt;Global Location Number&lt;/a&gt; (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	// types : Text
	GlobalLocationNumber []string `json:"globalLocationNumber,omitempty"`

	// HasMap see : https://schema.org/hasMap
	// A URL to a map of the place.
	// types : URL Map
	HasMap []interface{} `json:"hasMap,omitempty"`

	// HasPOS see : https://schema.org/hasPOS
	// Points-of-Sales operated by the organization or person.
	// types : Place
	HasPOS []*Place `json:"hasPOS,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IsicV4 see : https://schema.org/isicV4
	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	// types : Text
	IsicV4 []string `json:"isicV4,omitempty"`

	// LegalName see : https://schema.org/legalName
	// The official name of the organization, e.g. the registered company name.
	// types : Text
	LegalName []string `json:"legalName,omitempty"`

	// Location see : https://schema.org/location
	// The location of the event, organization or action.
	// types : Place PostalAddress
	Location []interface{} `json:"location,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo []interface{} `json:"logo,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person.
	// types : Offer
	MakesOffer []*Offer `json:"makesOffer,omitempty"`

	// Map see : https://schema.org/map
	// A URL to a map of the place.
	// types : URL
	Map []string `json:"map,omitempty"`

	// Maps see : https://schema.org/maps
	// A URL to a map of the place.
	// types : URL
	Maps []string `json:"maps,omitempty"`

	// Member see : https://schema.org/member
	// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals.
	// types : Organization Person
	Member []interface{} `json:"member,omitempty"`

	// MemberOf see : https://schema.org/memberOf
	// An Organization (or ProgramMembership) to which this Person or Organization belongs.
	// types : Organization ProgramMembership
	MemberOf []interface{} `json:"memberOf,omitempty"`

	// Members see : https://schema.org/members
	// A member of this organization.
	// types : Organization Person
	Members []interface{} `json:"members,omitempty"`

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
	// The opening hours for a business. Opening hours can be specified as a weekly time range, starting with days, then times per day. Multiple days can be listed with commas &#39;,&#39; separating each day. Day or time ranges are specified using a hyphen &#39;-&#39;.&lt;br /&gt;- Days are specified using the following two-letter combinations: &lt;code&gt;Mo&lt;/code&gt;, &lt;code&gt;Tu&lt;/code&gt;, &lt;code&gt;We&lt;/code&gt;, &lt;code&gt;Th&lt;/code&gt;, &lt;code&gt;Fr&lt;/code&gt;, &lt;code&gt;Sa&lt;/code&gt;, &lt;code&gt;Su&lt;/code&gt;.&lt;br /&gt;- Times are specified using 24:00 time. For example, 3pm is specified as &lt;code&gt;15:00&lt;/code&gt;. &lt;br /&gt;- Here is an example: &lt;code&gt;&amp;lt;time itemprop=&amp;quot;openingHours&amp;quot; datetime=&amp;quot;Tu,Th 16:00-20:00&amp;quot;&amp;gt;Tuesdays and Thursdays 4-8pm&amp;lt;/time&amp;gt;&lt;/code&gt;. &lt;br /&gt;- If a business is open 7 days a week, then it can be specified as &lt;code&gt;&amp;lt;time itemprop=&amp;quot;openingHours&amp;quot; datetime=&amp;quot;Mo-Su&amp;quot;&amp;gt;Monday through Sunday, all day&amp;lt;/time&amp;gt;&lt;/code&gt;.
	// types : Text
	OpeningHours []string `json:"openingHours,omitempty"`

	// OpeningHoursSpecification see : https://schema.org/openingHoursSpecification
	// The opening hours of a certain place.
	// types : OpeningHoursSpecification
	OpeningHoursSpecification []*OpeningHoursSpecification `json:"openingHoursSpecification,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns []interface{} `json:"owns,omitempty"`

	// ParentOrganization see : https://schema.org/parentOrganization
	// The larger organization that this organization is a branch of, if any.
	// types : Organization
	ParentOrganization []*Organization `json:"parentOrganization,omitempty"`

	// PaymentAccepted see : https://schema.org/paymentAccepted
	// Cash, credit card, etc.
	// types : Text
	PaymentAccepted []string `json:"paymentAccepted,omitempty"`

	// Photo see : https://schema.org/photo
	// A photograph of this place.
	// types : ImageObject Photograph
	Photo []interface{} `json:"photo,omitempty"`

	// Photos see : https://schema.org/photos
	// Photographs of this place.
	// types : ImageObject Photograph
	Photos []interface{} `json:"photos,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PriceRange see : https://schema.org/priceRange
	// The price range of the business, for example &lt;code&gt;$$$&lt;/code&gt;.
	// types : Text
	PriceRange []string `json:"priceRange,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Reviews see : https://schema.org/reviews
	// Review of the item.
	// types : Review
	Reviews []*Review `json:"reviews,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	// types : Demand
	Seeks []*Demand `json:"seeks,omitempty"`

	// SubOrganization see : https://schema.org/subOrganization
	// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific &#39;department&#39; property.
	// types : Organization
	SubOrganization []*Organization `json:"subOrganization,omitempty"`

	// TaxID see : https://schema.org/taxID
	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	// types : Text
	TaxID []string `json:"taxID,omitempty"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone []string `json:"telephone,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// VatID see : https://schema.org/vatID
	// The Value-added Tax ID of the organization or person.
	// types : Text
	VatID []string `json:"vatID,omitempty"`
}

func (v FurnitureStore) MarshalJSON() ([]byte, error) {
	type Alias FurnitureStore

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"FurnitureStore\","), b[1:]...), nil
}
