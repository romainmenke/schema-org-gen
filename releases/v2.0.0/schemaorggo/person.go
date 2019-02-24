package schemaorggo

import "encoding/json"

// Person see : https://schema.org/Person
type Person struct {

	// With properties from Thing see : https://schema.org/Thing
	//

	// AdditionalName see : https://schema.org/additionalName
	// An additional name for a Person, can be used for a middle name.
	// types : Text
	AdditionalName []string `json:"additionalName,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// Address see : https://schema.org/address
	// Physical address of the item.
	// types : PostalAddress
	Address []*PostalAddress `json:"address,omitempty"`

	// Affiliation see : https://schema.org/affiliation
	// An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	// types : Organization
	Affiliation []*Organization `json:"affiliation,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AlumniOf see : https://schema.org/alumniOf
	// An educational organizations that the person is an alumni of.
	// types : EducationalOrganization
	AlumniOf []*EducationalOrganization `json:"alumniOf,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// Awards see : https://schema.org/awards
	// Awards won by or for this item.
	// types : Text
	Awards []string `json:"awards,omitempty"`

	// BirthDate see : https://schema.org/birthDate
	// Date of birth.
	// types : Date
	BirthDate []Date `json:"birthDate,omitempty"`

	// BirthPlace see : https://schema.org/birthPlace
	// The place where the person was born.
	// types : Place
	BirthPlace []*Place `json:"birthPlace,omitempty"`

	// Brand see : https://schema.org/brand
	// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	// types : Brand Organization
	Brand []interface{} `json:"brand,omitempty"`

	// Children see : https://schema.org/children
	// A child of the person.
	// types : Person
	Children []*Person `json:"children,omitempty"`

	// Colleague see : https://schema.org/colleague
	// A colleague of the person.
	// types : Person
	Colleague []*Person `json:"colleague,omitempty"`

	// Colleagues see : https://schema.org/colleagues
	// A colleague of the person.
	// types : Person
	Colleagues []*Person `json:"colleagues,omitempty"`

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization.
	// types : ContactPoint
	ContactPoint []*ContactPoint `json:"contactPoint,omitempty"`

	// ContactPoints see : https://schema.org/contactPoints
	// A contact point for a person or organization.
	// types : ContactPoint
	ContactPoints []*ContactPoint `json:"contactPoints,omitempty"`

	// DeathDate see : https://schema.org/deathDate
	// Date of death.
	// types : Date
	DeathDate []Date `json:"deathDate,omitempty"`

	// DeathPlace see : https://schema.org/deathPlace
	// The place where the person died.
	// types : Place
	DeathPlace []*Place `json:"deathPlace,omitempty"`

	// Description see : https://schema.org/description
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// Duns see : https://schema.org/duns
	// The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	// types : Text
	Duns []string `json:"duns,omitempty"`

	// Email see : https://schema.org/email
	// Email address.
	// types : Text
	Email []string `json:"email,omitempty"`

	// FamilyName see : https://schema.org/familyName
	// Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
	// types : Text
	FamilyName []string `json:"familyName,omitempty"`

	// FaxNumber see : https://schema.org/faxNumber
	// The fax number.
	// types : Text
	FaxNumber []string `json:"faxNumber,omitempty"`

	// Follows see : https://schema.org/follows
	// The most generic uni-directional social relation.
	// types : Person
	Follows []*Person `json:"follows,omitempty"`

	// Gender see : https://schema.org/gender
	// Gender of the person.
	// types : Text
	Gender []string `json:"gender,omitempty"`

	// GivenName see : https://schema.org/givenName
	// Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	// types : Text
	GivenName []string `json:"givenName,omitempty"`

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The &lt;a href=&quot;http://www.gs1.org/gln&quot;&gt;Global Location Number&lt;/a&gt; (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	// types : Text
	GlobalLocationNumber []string `json:"globalLocationNumber,omitempty"`

	// HasPOS see : https://schema.org/hasPOS
	// Points-of-Sales operated by the organization or person.
	// types : Place
	HasPOS []*Place `json:"hasPOS,omitempty"`

	// Height see : https://schema.org/height
	// The height of the item.
	// types : Distance QuantitativeValue
	Height []interface{} `json:"height,omitempty"`

	// HomeLocation see : https://schema.org/homeLocation
	// A contact location for a person&#39;s residence.
	// types : ContactPoint Place
	HomeLocation []interface{} `json:"homeLocation,omitempty"`

	// HonorificPrefix see : https://schema.org/honorificPrefix
	// An honorific prefix preceding a Person&#39;s name such as Dr/Mrs/Mr.
	// types : Text
	HonorificPrefix []string `json:"honorificPrefix,omitempty"`

	// HonorificSuffix see : https://schema.org/honorificSuffix
	// An honorific suffix preceding a Person&#39;s name such as M.D. /PhD/MSCSW.
	// types : Text
	HonorificSuffix []string `json:"honorificSuffix,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// IsicV4 see : https://schema.org/isicV4
	// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	// types : Text
	IsicV4 []string `json:"isicV4,omitempty"`

	// JobTitle see : https://schema.org/jobTitle
	// The job title of the person (for example, Financial Manager).
	// types : Text
	JobTitle []string `json:"jobTitle,omitempty"`

	// Knows see : https://schema.org/knows
	// The most generic bi-directional social/work relation.
	// types : Person
	Knows []*Person `json:"knows,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	//       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	//       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	//       between the page and the primary entity.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       Related properties include sameAs, about, and url.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	//       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	//       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	//       serves more to clarify which of several entities is the main one for that page.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	//       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	//       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	//       &lt;br /&gt;&lt;br /&gt;
	//
	//       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	//       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	//       describes some other entity. For example, one web page may display a news article about a particular person.
	//       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	//       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person.
	// types : Offer
	MakesOffer []*Offer `json:"makesOffer,omitempty"`

	// MemberOf see : https://schema.org/memberOf
	// An Organization (or ProgramMembership) to which this Person or Organization belongs.
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

	// Nationality see : https://schema.org/nationality
	// Nationality of the person.
	// types : Country
	Nationality []*Country `json:"nationality,omitempty"`

	// NetWorth see : https://schema.org/netWorth
	// The total financial value of the organization or person as calculated by subtracting assets from liabilities.
	// types : PriceSpecification
	NetWorth []*PriceSpecification `json:"netWorth,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns []interface{} `json:"owns,omitempty"`

	// Parent see : https://schema.org/parent
	// A parent of this person.
	// types : Person
	Parent []*Person `json:"parent,omitempty"`

	// Parents see : https://schema.org/parents
	// A parents of the person.
	// types : Person
	Parents []*Person `json:"parents,omitempty"`

	// PerformerIn see : https://schema.org/performerIn
	// Event that this person is a performer or participant in.
	// types : Event
	PerformerIn []*Event `json:"performerIn,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// RelatedTo see : https://schema.org/relatedTo
	// The most generic familial relation.
	// types : Person
	RelatedTo []*Person `json:"relatedTo,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	// types : Demand
	Seeks []*Demand `json:"seeks,omitempty"`

	// Sibling see : https://schema.org/sibling
	// A sibling of the person.
	// types : Person
	Sibling []*Person `json:"sibling,omitempty"`

	// Siblings see : https://schema.org/siblings
	// A sibling of the person.
	// types : Person
	Siblings []*Person `json:"siblings,omitempty"`

	// Spouse see : https://schema.org/spouse
	// The person&#39;s spouse.
	// types : Person
	Spouse []*Person `json:"spouse,omitempty"`

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

	// Weight see : https://schema.org/weight
	// The weight of the product or person.
	// types : QuantitativeValue
	Weight []*QuantitativeValue `json:"weight,omitempty"`

	// WorkLocation see : https://schema.org/workLocation
	// A contact location for a person&#39;s place of work.
	// types : ContactPoint Place
	WorkLocation []interface{} `json:"workLocation,omitempty"`

	// WorksFor see : https://schema.org/worksFor
	// Organizations that the person works for.
	// types : Organization
	WorksFor []*Organization `json:"worksFor,omitempty"`
}

func (v Person) MarshalJSON() ([]byte, error) {
	type Alias Person

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Person\","), b[1:]...), nil
}
