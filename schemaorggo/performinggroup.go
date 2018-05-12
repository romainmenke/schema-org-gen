package schemaorggo

import "encoding/json"

// PerformingGroup see : https://schema.org/PerformingGroup
type PerformingGroup struct {
	Organization

	typeContext

	// ActionableFeedbackPolicy see : http://pending.schema.org/actionableFeedbackPolicy
	// For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
	// types : CreativeWork URL
	ActionableFeedbackPolicy interface{} `json:"actionableFeedbackPolicy,omitempty"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress Text
	Address interface{} `json:"address,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating *AggregateRating `json:"aggregateRating,omitempty"`

	// Alumni see : https://schema.org/alumni
	// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	// types : Person
	Alumni *Person `json:"alumni,omitempty"`

	// AreaServed see : https://schema.org/areaServed
	// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	// types : AdministrativeArea GeoShape Place Text
	AreaServed interface{} `json:"areaServed,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award string `json:"award,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand interface{} `json:"brand,omitempty"`

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	// types : ContactPoint
	ContactPoint *ContactPoint `json:"contactPoint,omitempty"`

	// CorrectionsPolicy see : http://pending.schema.org/correctionsPolicy
	// For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
	// types : CreativeWork URL
	CorrectionsPolicy interface{} `json:"correctionsPolicy,omitempty"`

	// Department see : https://schema.org/department
	// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	// types : Organization
	Department *Organization `json:"department,omitempty"`

	// DissolutionDate see : https://schema.org/dissolutionDate
	// The date that this organization was dissolved.
	// types : Date
	DissolutionDate Date `json:"dissolutionDate,omitempty"`

	// DiversityPolicy see : http://pending.schema.org/diversityPolicy
	// Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
	// types : CreativeWork URL
	DiversityPolicy interface{} `json:"diversityPolicy,omitempty"`

	// Duns see : https://schema.org/duns
	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	// types : Text
	Duns string `json:"duns,omitempty"`

	// Email see : https://schema.org/email
	// Email address.
	// types : Text
	Email string `json:"email,omitempty"`

	// Employee see : https://schema.org/employee
	// Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
	// types : Person
	Employee *Person `json:"employee,omitempty"`

	// EthicsPolicy see : http://pending.schema.org/ethicsPolicy
	// Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
	// types : CreativeWork URL
	EthicsPolicy interface{} `json:"ethicsPolicy,omitempty"`

	// Event see : https://schema.org/event
	// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	// types : Event
	Event *Event `json:"event,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber string `json:"faxNumber,omitempty"`

	// Founder see : https://schema.org/founder
	// A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
	// types : Person
	Founder *Person `json:"founder,omitempty"`

	// FoundingDate see : https://schema.org/foundingDate
	// The date that this organization was founded.
	// types : Date
	FoundingDate Date `json:"foundingDate,omitempty"`

	// FoundingLocation see : https://schema.org/foundingLocation
	// The place where the Organization was founded.
	// types : Place
	FoundingLocation *Place `json:"foundingLocation,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder interface{} `json:"funder,omitempty"`

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	// types : Text
	GlobalLocationNumber string `json:"globalLocationNumber,omitempty"`

	// HasOfferCatalog see : https://schema.org/hasOfferCatalog
	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	// types : OfferCatalog
	HasOfferCatalog *OfferCatalog `json:"hasOfferCatalog,omitempty"`

	// HasPOS see : https://schema.org/hasPOS
	// Points-of-Sales operated by the organization or person.
	// types : Place
	HasPOS *Place `json:"hasPOS,omitempty"`

	// IsicV4 see : https://schema.org/isicV4
	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	// types : Text
	IsicV4 string `json:"isicV4,omitempty"`

	// LegalName see : https://schema.org/legalName
	// The official name of the organization, e.g. the registered company name.
	// types : Text
	LegalName string `json:"legalName,omitempty"`

	// LeiCode see : https://schema.org/leiCode
	// An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
	// types : Text
	LeiCode string `json:"leiCode,omitempty"`

	// Location see : https://schema.org/location
	// The location of for example where the event is happening, an organization is located, or where an action takes place.
	// types : Place PostalAddress Text
	Location interface{} `json:"location,omitempty"`

	// Logo see : https://schema.org/logo
	// An associated logo.
	// types : ImageObject URL
	Logo interface{} `json:"logo,omitempty"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	// types : Offer
	MakesOffer *Offer `json:"makesOffer,omitempty"`

	// Member see : https://schema.org/member
	// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	// types : Organization Person
	Member interface{} `json:"member,omitempty"`

	// MemberOf see : https://schema.org/memberOf
	// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	// types : Organization ProgramMembership
	MemberOf interface{} `json:"memberOf,omitempty"`

	// Naics see : https://schema.org/naics
	// The North American Industry Classification System (NAICS) code for a particular organization or business person.
	// types : Text
	Naics string `json:"naics,omitempty"`

	// NumberOfEmployees see : https://schema.org/numberOfEmployees
	// The number of employees in an organization e.g. business.
	// types : QuantitativeValue
	NumberOfEmployees *QuantitativeValue `json:"numberOfEmployees,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns interface{} `json:"owns,omitempty"`

	// ParentOrganization see : https://schema.org/parentOrganization
	// The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
	// types : Organization
	ParentOrganization *Organization `json:"parentOrganization,omitempty"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	// types : CreativeWork URL
	PublishingPrinciples interface{} `json:"publishingPrinciples,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review *Review `json:"review,omitempty"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	// types : Demand
	Seeks *Demand `json:"seeks,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor interface{} `json:"sponsor,omitempty"`

	// SubOrganization see : https://schema.org/subOrganization
	// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific &#39;department&#39; property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
	// types : Organization
	SubOrganization *Organization `json:"subOrganization,omitempty"`

	// TaxID see : https://schema.org/taxID
	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	// types : Text
	TaxID string `json:"taxID,omitempty"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone string `json:"telephone,omitempty"`

	// UnnamedSourcesPolicy see : http://pending.schema.org/unnamedSourcesPolicy
	// For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
	// types : CreativeWork URL
	UnnamedSourcesPolicy interface{} `json:"unnamedSourcesPolicy,omitempty"`

	// VatID see : https://schema.org/vatID
	// The Value-added Tax ID of the organization or person.
	// types : Text
	VatID string `json:"vatID,omitempty"`
}

func (v PerformingGroup) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "PerformingGroup"

	return json.Marshal(v)
}

func (v *PerformingGroup) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "PerformingGroup"

	return json.Marshal(*v)
}
