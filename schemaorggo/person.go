package schemaorggo

import "encoding/json"

// Person see : https://schema.org/Person
type Person struct {
	Thing

	typeContext

	// AdditionalName see : https://schema.org/additionalName
	// An additional name for a Person, can be used for a middle name.
	AdditionalName string `json:"additionalName"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	Address interface{} `json:"address"` // types : PostalAddress Text

	// Affiliation see : https://schema.org/affiliation
	// An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	Affiliation *Organization `json:"affiliation"`

	// AlumniOf see : https://schema.org/alumniOf
	// An organization that the person is an alumni of. Inverse property: alumni (see: https://schema.org/alumni).
	AlumniOf interface{} `json:"alumniOf"` // types : EducationalOrganization Organization

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	Award string `json:"award"`

	// BirthDate see : https://schema.org/birthDate
	// Date of birth.
	BirthDate Date `json:"birthDate"`

	// BirthPlace see : https://schema.org/birthPlace
	// The place where the person was born.
	BirthPlace *Place `json:"birthPlace"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	Brand interface{} `json:"brand"` // types : Brand Organization

	// Children see : https://schema.org/children
	// A child of the person.
	Children *Person `json:"children"`

	// Colleague see : https://schema.org/colleague
	// A colleague of the person. Supersedes colleagues (see: https://schema.org/colleagues).
	Colleague interface{} `json:"colleague"` // types : Person URL

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	ContactPoint *ContactPoint `json:"contactPoint"`

	// DeathDate see : https://schema.org/deathDate
	// Date of death.
	DeathDate Date `json:"deathDate"`

	// DeathPlace see : https://schema.org/deathPlace
	// The place where the person died.
	DeathPlace *Place `json:"deathPlace"`

	// Duns see : https://schema.org/duns
	// The Dun & Bradstreet DUNS number for identifying an organization or business person.
	Duns string `json:"duns"`

	// Email see : https://schema.org/email
	// Email address.
	Email string `json:"email"`

	// FamilyName see : https://schema.org/familyName
	// Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
	FamilyName string `json:"familyName"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	FaxNumber string `json:"faxNumber"`

	// Follows see : https://schema.org/follows
	// The most generic uni-directional social relation.
	Follows *Person `json:"follows"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	Funder interface{} `json:"funder"` // types : Organization Person

	// Gender see : https://schema.org/gender
	// Gender of the person. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender.
	Gender interface{} `json:"gender"` // types : GenderType Text

	// GivenName see : https://schema.org/givenName
	// Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	GivenName string `json:"givenName"`

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	GlobalLocationNumber string `json:"globalLocationNumber"`

	// HasOfferCatalog see : https://schema.org/hasOfferCatalog
	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	HasOfferCatalog *OfferCatalog `json:"hasOfferCatalog"`

	// HasPOS see : https://schema.org/hasPOS
	// Points-of-Sales operated by the organization or person.
	HasPOS *Place `json:"hasPOS"`

	// Height see : https://schema.org/height
	// The height of the item.
	Height interface{} `json:"height"` // types : Distance QuantitativeValue

	// HomeLocation see : https://schema.org/homeLocation
	// A contact location for a person's residence.
	HomeLocation interface{} `json:"homeLocation"` // types : ContactPoint Place

	// HonorificPrefix see : https://schema.org/honorificPrefix
	// An honorific prefix preceding a Person's name such as Dr/Mrs/Mr.
	HonorificPrefix string `json:"honorificPrefix"`

	// HonorificSuffix see : https://schema.org/honorificSuffix
	// An honorific suffix preceding a Person's name such as M.D. /PhD/MSCSW.
	HonorificSuffix string `json:"honorificSuffix"`

	// IsicV4 see : https://schema.org/isicV4
	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	IsicV4 string `json:"isicV4"`

	// JobTitle see : https://schema.org/jobTitle
	// The job title of the person (for example, Financial Manager).
	JobTitle string `json:"jobTitle"`

	// Knows see : https://schema.org/knows
	// The most generic bi-directional social/work relation.
	Knows *Person `json:"knows"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	MakesOffer *Offer `json:"makesOffer"`

	// MemberOf see : https://schema.org/memberOf
	// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	MemberOf interface{} `json:"memberOf"` // types : Organization ProgramMembership

	// Naics see : https://schema.org/naics
	// The North American Industry Classification System (NAICS) code for a particular organization or business person.
	Naics string `json:"naics"`

	// Nationality see : https://schema.org/nationality
	// Nationality of the person.
	Nationality *Country `json:"nationality"`

	// NetWorth see : https://schema.org/netWorth
	// The total financial value of the person as calculated by subtracting assets from liabilities.
	NetWorth interface{} `json:"netWorth"` // types : MonetaryAmount PriceSpecification

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	Owns interface{} `json:"owns"` // types : OwnershipInfo Product

	// Parent see : https://schema.org/parent
	// A parent of this person. Supersedes parents (see: https://schema.org/parents).
	Parent *Person `json:"parent"`

	// PerformerIn see : https://schema.org/performerIn
	// Event that this person is a performer or participant in.
	PerformerIn *Event `json:"performerIn"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

	// RelatedTo see : https://schema.org/relatedTo
	// The most generic familial relation.
	RelatedTo *Person `json:"relatedTo"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	Seeks *Demand `json:"seeks"`

	// Sibling see : https://schema.org/sibling
	// A sibling of the person. Supersedes siblings (see: https://schema.org/siblings).
	Sibling *Person `json:"sibling"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	Sponsor interface{} `json:"sponsor"` // types : Organization Person

	// Spouse see : https://schema.org/spouse
	// The person's spouse.
	Spouse *Person `json:"spouse"`

	// TaxID see : https://schema.org/taxID
	// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	TaxID string `json:"taxID"`

	// Telephone see : https://schema.org/telephone
	// The telephone number.
	Telephone string `json:"telephone"`

	// VatID see : https://schema.org/vatID
	// The Value-added Tax ID of the organization or person.
	VatID string `json:"vatID"`

	// Weight see : https://schema.org/weight
	// The weight of the product or person.
	Weight *QuantitativeValue `json:"weight"`

	// WorkLocation see : https://schema.org/workLocation
	// A contact location for a person's place of work.
	WorkLocation interface{} `json:"workLocation"` // types : ContactPoint Place

	// WorksFor see : https://schema.org/worksFor
	// Organizations that the person works for.
	WorksFor *Organization `json:"worksFor"`
}

func (v *Person) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Person"

	return json.Marshal(v)
}
