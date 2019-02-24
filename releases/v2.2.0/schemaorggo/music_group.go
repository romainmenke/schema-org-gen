package schemaorggo

import "encoding/json"

// MusicGroup see : https://schema.org/MusicGroup
type MusicGroup struct {

	// With properties from PerformingGroup see : https://schema.org/PerformingGroup
	//

	// With properties from Organization see : https://schema.org/Organization
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

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

	// Album see : https://schema.org/album
	// A music album.
	// types : MusicAlbum
	Album []*MusicAlbum `json:"album,omitempty"`

	// Albums see : https://schema.org/albums
	// A collection of music albums.
	// types : MusicAlbum
	Albums []*MusicAlbum `json:"albums,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided.
	// types : Place AdministrativeArea GeoShape Text
	AreaServed []interface{} `json:"areaServed,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// Awards see : https://schema.org/awards
	// Awards won by or for this item.
	// types : Text
	Awards []string `json:"awards,omitempty"`

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

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The &lt;a href=&quot;http://www.gs1.org/gln&quot;&gt;Global Location Number&lt;/a&gt; (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	// types : Text
	GlobalLocationNumber []string `json:"globalLocationNumber,omitempty"`

	// HasOfferCatalog see : https://schema.org/hasOfferCatalog
	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	// types : OfferCatalog
	HasOfferCatalog []*OfferCatalog `json:"hasOfferCatalog,omitempty"`

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
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
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

	// MusicGroupMember see : https://schema.org/musicGroupMember
	// A member of a music group&amp;#x2014;for example, John, Paul, George, or Ringo.
	// types : Person
	MusicGroupMember []*Person `json:"musicGroupMember,omitempty"`

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

	// OfferedBy see : https://schema.org/offeredBy
	// A pointer to the organization or person making the offer.
	// types : Person Offer
	OfferedBy []interface{} `json:"offeredBy,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns []interface{} `json:"owns,omitempty"`

	// ParentOrganization see : https://schema.org/parentOrganization
	// The larger organization that this organization is a branch of, if any.
	// types : Organization
	ParentOrganization []*Organization `json:"parentOrganization,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

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

	// ServiceArea see : https://schema.org/serviceArea
	// The geographic area where the service is provided.
	// types : Place AdministrativeArea GeoShape
	ServiceArea []interface{} `json:"serviceArea,omitempty"`

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

	// Track see : https://schema.org/track
	// A music recording (track)&amp;#x2014;usually a single song. If an ItemList is given, the list should contain items of type MusicRecording.
	// types : ItemList MusicRecording
	Track []interface{} `json:"track,omitempty"`

	// Tracks see : https://schema.org/tracks
	// A music recording (track)&amp;#x2014;usually a single song.
	// types : MusicRecording
	Tracks []*MusicRecording `json:"tracks,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// VatID see : https://schema.org/vatID
	// The Value-added Tax ID of the organization or person.
	// types : Text
	VatID []string `json:"vatID,omitempty"`
}

func (v MusicGroup) MarshalJSON() ([]byte, error) {
	type Alias MusicGroup

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"MusicGroup\","), b[1:]...), nil
}
