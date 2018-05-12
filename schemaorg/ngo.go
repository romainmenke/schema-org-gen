package schemaorg

import "encoding/json"

// NGO see : https://schema.org/NGO
type NGO struct {

Organization

typeContext

// ActionableFeedbackPolicy see : http://pending.schema.org/actionableFeedbackPolicy
// For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
ActionableFeedbackPolicy interface{} `json:"actionableFeedbackPolicy"` // types : CreativeWork URL

// Address see : https://schema.org/address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AggregateRating see : https://schema.org/aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating *AggregateRating `json:"aggregateRating"`

// Alumni see : https://schema.org/alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni *Person `json:"alumni"`

// AreaServed see : https://schema.org/areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Award see : https://schema.org/award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : https://schema.org/brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// ContactPoint see : https://schema.org/contactPoint
// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
ContactPoint *ContactPoint `json:"contactPoint"`

// CorrectionsPolicy see : http://pending.schema.org/correctionsPolicy
// For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
CorrectionsPolicy interface{} `json:"correctionsPolicy"` // types : CreativeWork URL

// Department see : https://schema.org/department
// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
Department *Organization `json:"department"`

// DissolutionDate see : https://schema.org/dissolutionDate
// The date that this organization was dissolved.
DissolutionDate interface{} `json:"dissolutionDate"`

// DiversityPolicy see : http://pending.schema.org/diversityPolicy
// Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
DiversityPolicy interface{} `json:"diversityPolicy"` // types : CreativeWork URL

// Duns see : https://schema.org/duns
// The Dun & Bradstreet DUNS number for identifying an organization or business person.
Duns string `json:"duns"`

// Email see : https://schema.org/email
// Email address.
Email string `json:"email"`

// Employee see : https://schema.org/employee
// Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
Employee *Person `json:"employee"`

// EthicsPolicy see : http://pending.schema.org/ethicsPolicy
// Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
EthicsPolicy interface{} `json:"ethicsPolicy"` // types : CreativeWork URL

// Event see : https://schema.org/event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event *Event `json:"event"`

// FaxNumber see : https://schema.org/faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// Founder see : https://schema.org/founder
// A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
Founder *Person `json:"founder"`

// FoundingDate see : https://schema.org/foundingDate
// The date that this organization was founded.
FoundingDate interface{} `json:"foundingDate"`

// FoundingLocation see : https://schema.org/foundingLocation
// The place where the Organization was founded.
FoundingLocation *Place `json:"foundingLocation"`

// Funder see : https://schema.org/funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// GlobalLocationNumber see : https://schema.org/globalLocationNumber
// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
GlobalLocationNumber string `json:"globalLocationNumber"`

// HasOfferCatalog see : https://schema.org/hasOfferCatalog
// Indicates an OfferCatalog listing for this Organization, Person, or Service.
HasOfferCatalog *OfferCatalog `json:"hasOfferCatalog"`

// HasPOS see : https://schema.org/hasPOS
// Points-of-Sales operated by the organization or person.
HasPOS *Place `json:"hasPOS"`

// IsicV4 see : https://schema.org/isicV4
// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
IsicV4 string `json:"isicV4"`

// LegalName see : https://schema.org/legalName
// The official name of the organization, e.g. the registered company name.
LegalName string `json:"legalName"`

// LeiCode see : https://schema.org/leiCode
// An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
LeiCode string `json:"leiCode"`

// Location see : https://schema.org/location
// The location of for example where the event is happening, an organization is located, or where an action takes place.
Location interface{} `json:"location"` // types : Place PostalAddress Text

// Logo see : https://schema.org/logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// MakesOffer see : https://schema.org/makesOffer
// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
MakesOffer *Offer `json:"makesOffer"`

// Member see : https://schema.org/member
// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
Member interface{} `json:"member"` // types : Organization Person

// MemberOf see : https://schema.org/memberOf
// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
MemberOf interface{} `json:"memberOf"` // types : Organization ProgramMembership

// Naics see : https://schema.org/naics
// The North American Industry Classification System (NAICS) code for a particular organization or business person.
Naics string `json:"naics"`

// NumberOfEmployees see : https://schema.org/numberOfEmployees
// The number of employees in an organization e.g. business.
NumberOfEmployees *QuantitativeValue `json:"numberOfEmployees"`

// Owns see : https://schema.org/owns
// Products owned by the organization or person.
Owns interface{} `json:"owns"` // types : OwnershipInfo Product

// ParentOrganization see : https://schema.org/parentOrganization
// The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
ParentOrganization *Organization `json:"parentOrganization"`

// PublishingPrinciples see : https://schema.org/publishingPrinciples
// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
// 
// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

// Review see : https://schema.org/review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review *Review `json:"review"`

// Seeks see : https://schema.org/seeks
// A pointer to products or services sought by the organization or person (demand).
Seeks *Demand `json:"seeks"`

// Sponsor see : https://schema.org/sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// SubOrganization see : https://schema.org/subOrganization
// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific 'department' property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
SubOrganization *Organization `json:"subOrganization"`

// TaxID see : https://schema.org/taxID
// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
TaxID string `json:"taxID"`

// Telephone see : https://schema.org/telephone
// The telephone number.
Telephone string `json:"telephone"`

// UnnamedSourcesPolicy see : http://pending.schema.org/unnamedSourcesPolicy
// For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
UnnamedSourcesPolicy interface{} `json:"unnamedSourcesPolicy"` // types : CreativeWork URL

// VatID see : https://schema.org/vatID
// The Value-added Tax ID of the organization or person.
VatID string `json:"vatID"`

}

func (v *NGO) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "NGO"

	return json.Marshal(v)
}
