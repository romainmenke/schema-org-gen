package schemaorggo

import "encoding/json"

// Person see : https://schema.org/Person
type Person struct {
	Thing

	typeContext

	// AdditionalName see : https://schema.org/additionalName
	// An additional name for a Person, can be used for a middle name.
	// types : Text
	AdditionalName string `json:"additionalName,omitempty"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress Text
	Address interface{} `json:"address,omitempty"`

	// Affiliation see : https://schema.org/affiliation
	// An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	// types : Organization
	Affiliation *Organization `json:"affiliation,omitempty"`

	// AlumniOf see : https://schema.org/alumniOf
	// An organization that the person is an alumni of. Inverse property: alumni (see: https://schema.org/alumni).
	// types : EducationalOrganization Organization
	AlumniOf interface{} `json:"alumniOf,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award string `json:"award,omitempty"`

	// BirthDate see : https://schema.org/birthDate
	// Date of birth.
	// types : Date
	BirthDate Date `json:"birthDate,omitempty"`

	// BirthPlace see : https://schema.org/birthPlace
	// The place where the person was born.
	// types : Place
	BirthPlace *Place `json:"birthPlace,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand interface{} `json:"brand,omitempty"`

	// Children see : https://schema.org/children
	// A child of the person.
	// types : Person
	Children *Person `json:"children,omitempty"`

	// Colleague see : https://schema.org/colleague
	// A colleague of the person. Supersedes colleagues (see: https://schema.org/colleagues).
	// types : Person URL
	Colleague interface{} `json:"colleague,omitempty"`

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	// types : ContactPoint
	ContactPoint *ContactPoint `json:"contactPoint,omitempty"`

	// DeathDate see : https://schema.org/deathDate
	// Date of death.
	// types : Date
	DeathDate Date `json:"deathDate,omitempty"`

	// DeathPlace see : https://schema.org/deathPlace
	// The place where the person died.
	// types : Place
	DeathPlace *Place `json:"deathPlace,omitempty"`

	// Duns see : https://schema.org/duns
	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	// types : Text
	Duns string `json:"duns,omitempty"`

	// Email see : https://schema.org/email
	// Email address.
	// types : Text
	Email string `json:"email,omitempty"`

	// FamilyName see : https://schema.org/familyName
	// Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
	// types : Text
	FamilyName string `json:"familyName,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber string `json:"faxNumber,omitempty"`

	// Follows see : https://schema.org/follows
	// The most generic uni-directional social relation.
	// types : Person
	Follows *Person `json:"follows,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder interface{} `json:"funder,omitempty"`

	// Gender see : https://schema.org/gender
	// Gender of the person. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender.
	// types : GenderType Text
	Gender interface{} `json:"gender,omitempty"`

	// GivenName see : https://schema.org/givenName
	// Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	// types : Text
	GivenName string `json:"givenName,omitempty"`

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

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height interface{} `json:"height,omitempty"`

	// HomeLocation see : https://schema.org/homeLocation
	// A contact location for a person&#39;s residence.
	// types : ContactPoint Place
	HomeLocation interface{} `json:"homeLocation,omitempty"`

	// HonorificPrefix see : https://schema.org/honorificPrefix
	// An honorific prefix preceding a Person&#39;s name such as Dr/Mrs/Mr.
	// types : Text
	HonorificPrefix string `json:"honorificPrefix,omitempty"`

	// HonorificSuffix see : https://schema.org/honorificSuffix
	// An honorific suffix preceding a Person&#39;s name such as M.D. /PhD/MSCSW.
	// types : Text
	HonorificSuffix string `json:"honorificSuffix,omitempty"`

	// IsicV4 see : https://schema.org/isicV4
	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	// types : Text
	IsicV4 string `json:"isicV4,omitempty"`

	// JobTitle see : https://schema.org/jobTitle
	// The job title of the person (for example, Financial Manager).
	// types : Text
	JobTitle string `json:"jobTitle,omitempty"`

	// Knows see : https://schema.org/knows
	// The most generic bi-directional social/work relation.
	// types : Person
	Knows *Person `json:"knows,omitempty"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	// types : Offer
	MakesOffer *Offer `json:"makesOffer,omitempty"`

	// MemberOf see : https://schema.org/memberOf
	// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	// types : Organization ProgramMembership
	MemberOf interface{} `json:"memberOf,omitempty"`

	// Naics see : https://schema.org/naics
	// The North American Industry Classification System (NAICS) code for a particular organization or business person.
	// types : Text
	Naics string `json:"naics,omitempty"`

	// Nationality see : https://schema.org/nationality
	// Nationality of the person.
	// types : Country
	Nationality *Country `json:"nationality,omitempty"`

	// NetWorth see : https://schema.org/netWorth
	// The total financial value of the person as calculated by subtracting assets from liabilities.
	// types : MonetaryAmount PriceSpecification
	NetWorth interface{} `json:"netWorth,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns interface{} `json:"owns,omitempty"`

	// Parent see : https://schema.org/parent
	// A parent of this person. Supersedes parents (see: https://schema.org/parents).
	// types : Person
	Parent *Person `json:"parent,omitempty"`

	// PerformerIn see : https://schema.org/performerIn
	// Event that this person is a performer or participant in.
	// types : Event
	PerformerIn *Event `json:"performerIn,omitempty"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	// types : CreativeWork URL
	PublishingPrinciples interface{} `json:"publishingPrinciples,omitempty"`

	// RelatedTo see : https://schema.org/relatedTo
	// The most generic familial relation.
	// types : Person
	RelatedTo *Person `json:"relatedTo,omitempty"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	// types : Demand
	Seeks *Demand `json:"seeks,omitempty"`

	// Sibling see : https://schema.org/sibling
	// A sibling of the person. Supersedes siblings (see: https://schema.org/siblings).
	// types : Person
	Sibling *Person `json:"sibling,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor interface{} `json:"sponsor,omitempty"`

	// Spouse see : https://schema.org/spouse
	// The person&#39;s spouse.
	// types : Person
	Spouse *Person `json:"spouse,omitempty"`

	// TaxID see : https://schema.org/taxID
	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	// types : Text
	TaxID string `json:"taxID,omitempty"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	// types : Text
	Telephone string `json:"telephone,omitempty"`

	// VatID see : https://schema.org/vatID
	// The Value-added Tax ID of the organization or person.
	// types : Text
	VatID string `json:"vatID,omitempty"`

	// Weight see : https://schema.org/weight
	// The weight of the product or person.
	// types : QuantitativeValue
	Weight *QuantitativeValue `json:"weight,omitempty"`

	// WorkLocation see : https://schema.org/workLocation
	// A contact location for a person&#39;s place of work.
	// types : ContactPoint Place
	WorkLocation interface{} `json:"workLocation,omitempty"`

	// WorksFor see : https://schema.org/worksFor
	// Organizations that the person works for.
	// types : Organization
	WorksFor *Organization `json:"worksFor,omitempty"`
}

func (v Person) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Person"

	return json.Marshal(v)
}

func (v *Person) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Person"

	return json.Marshal(*v)
}
