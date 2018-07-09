package schemaorggo

import "encoding/json"

// Person see : https://schema.org/Person
type Person struct {
	typeContext

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
	// types : PostalAddress Text
	Address []interface{} `json:"address,omitempty"`

	// Affiliation see : https://schema.org/affiliation
	// An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	// types : Organization
	Affiliation []*Organization `json:"affiliation,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AlumniOf see : https://schema.org/alumniOf
	// An organization that the person is an alumni of. Inverse property: alumni (see: https://schema.org/alumni).
	// types : EducationalOrganization Organization
	AlumniOf []interface{} `json:"alumniOf,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award []string `json:"award,omitempty"`

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
	// A colleague of the person. Supersedes colleagues (see: https://schema.org/colleagues).
	// types : Person URL
	Colleague []interface{} `json:"colleague,omitempty"`

	// ContactPoint see : https://schema.org/contactPoint
	// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	// types : ContactPoint
	ContactPoint []*ContactPoint `json:"contactPoint,omitempty"`

	// DeathDate see : https://schema.org/deathDate
	// Date of death.
	// types : Date
	DeathDate []Date `json:"deathDate,omitempty"`

	// DeathPlace see : https://schema.org/deathPlace
	// The place where the person died.
	// types : Place
	DeathPlace []*Place `json:"deathPlace,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

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

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder []interface{} `json:"funder,omitempty"`

	// Gender see : https://schema.org/gender
	// Gender of the person. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender.
	// types : GenderType Text
	Gender []interface{} `json:"gender,omitempty"`

	// GivenName see : https://schema.org/givenName
	// Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	// types : Text
	GivenName []string `json:"givenName,omitempty"`

	// GlobalLocationNumber see : https://schema.org/globalLocationNumber
	// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	// types : Text
	GlobalLocationNumber []string `json:"globalLocationNumber,omitempty"`

	// HasOccupation see : https://pending.schema.org/hasOccupation
	// The Person&#39;s occupation. For past professions, use Role for expressing dates.
	// types : Occupation
	HasOccupation []interface{} `json:"hasOccupation,omitempty"`

	// HasOfferCatalog see : https://schema.org/hasOfferCatalog
	// Indicates an OfferCatalog listing for this Organization, Person, or Service.
	// types : OfferCatalog
	HasOfferCatalog []*OfferCatalog `json:"hasOfferCatalog,omitempty"`

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

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	// types : PropertyValue Text URL
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	// types : ImageObject URL
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

	// KnowsAbout see : https://pending.schema.org/knowsAbout
	// Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a topic that is known about - suggesting possible expertise but not implying it. We do not distinguish skill levels here, or yet relate this to educational content, events, objectives or JobPosting (see: https://schema.org/JobPosting) descriptions.
	// types : Text Thing URL
	KnowsAbout []interface{} `json:"knowsAbout,omitempty"`

	// KnowsLanguage see : https://pending.schema.org/knowsLanguage
	// Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a known language. We do not distinguish skill levels or reading/writing/speaking/signing here. Use language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	// types : Language Text
	KnowsLanguage []interface{} `json:"knowsLanguage,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MakesOffer see : https://schema.org/makesOffer
	// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	// types : Offer
	MakesOffer []*Offer `json:"makesOffer,omitempty"`

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

	// Nationality see : https://schema.org/nationality
	// Nationality of the person.
	// types : Country
	Nationality []*Country `json:"nationality,omitempty"`

	// NetWorth see : https://schema.org/netWorth
	// The total financial value of the person as calculated by subtracting assets from liabilities.
	// types : MonetaryAmount PriceSpecification
	NetWorth []interface{} `json:"netWorth,omitempty"`

	// Owns see : https://schema.org/owns
	// Products owned by the organization or person.
	// types : OwnershipInfo Product
	Owns []interface{} `json:"owns,omitempty"`

	// Parent see : https://schema.org/parent
	// A parent of this person. Supersedes parents (see: https://schema.org/parents).
	// types : Person
	Parent []*Person `json:"parent,omitempty"`

	// PerformerIn see : https://schema.org/performerIn
	// Event that this person is a performer or participant in.
	// types : Event
	PerformerIn []*Event `json:"performerIn,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	// types : CreativeWork URL
	PublishingPrinciples []interface{} `json:"publishingPrinciples,omitempty"`

	// RelatedTo see : https://schema.org/relatedTo
	// The most generic familial relation.
	// types : Person
	RelatedTo []*Person `json:"relatedTo,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// Seeks see : https://schema.org/seeks
	// A pointer to products or services sought by the organization or person (demand).
	// types : Demand
	Seeks []*Demand `json:"seeks,omitempty"`

	// Sibling see : https://schema.org/sibling
	// A sibling of the person. Supersedes siblings (see: https://schema.org/siblings).
	// types : Person
	Sibling []*Person `json:"sibling,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor []interface{} `json:"sponsor,omitempty"`

	// Spouse see : https://schema.org/spouse
	// The person&#39;s spouse.
	// types : Person
	Spouse []*Person `json:"spouse,omitempty"`

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

func (v Person) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.AdditionalName
		if len(v.AdditionalName) == 1 {
			value = v.AdditionalName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["additionalName"] = json.RawMessage(b)
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
		var value interface{} = v.Affiliation
		if len(v.Affiliation) == 1 {
			value = v.Affiliation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["affiliation"] = json.RawMessage(b)
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
		var value interface{} = v.AlumniOf
		if len(v.AlumniOf) == 1 {
			value = v.AlumniOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alumniOf"] = json.RawMessage(b)
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
		var value interface{} = v.BirthDate
		if len(v.BirthDate) == 1 {
			value = v.BirthDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["birthDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.BirthPlace
		if len(v.BirthPlace) == 1 {
			value = v.BirthPlace[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["birthPlace"] = json.RawMessage(b)
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
		var value interface{} = v.Children
		if len(v.Children) == 1 {
			value = v.Children[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["children"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Colleague
		if len(v.Colleague) == 1 {
			value = v.Colleague[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["colleague"] = json.RawMessage(b)
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
		var value interface{} = v.DeathDate
		if len(v.DeathDate) == 1 {
			value = v.DeathDate[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["deathDate"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DeathPlace
		if len(v.DeathPlace) == 1 {
			value = v.DeathPlace[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["deathPlace"] = json.RawMessage(b)
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
		var value interface{} = v.FamilyName
		if len(v.FamilyName) == 1 {
			value = v.FamilyName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["familyName"] = json.RawMessage(b)
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
		var value interface{} = v.Follows
		if len(v.Follows) == 1 {
			value = v.Follows[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["follows"] = json.RawMessage(b)
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
		var value interface{} = v.Gender
		if len(v.Gender) == 1 {
			value = v.Gender[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gender"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GivenName
		if len(v.GivenName) == 1 {
			value = v.GivenName[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["givenName"] = json.RawMessage(b)
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
		var value interface{} = v.HasOccupation
		if len(v.HasOccupation) == 1 {
			value = v.HasOccupation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasOccupation"] = json.RawMessage(b)
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
		var value interface{} = v.Height
		if len(v.Height) == 1 {
			value = v.Height[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["height"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HomeLocation
		if len(v.HomeLocation) == 1 {
			value = v.HomeLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["homeLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HonorificPrefix
		if len(v.HonorificPrefix) == 1 {
			value = v.HonorificPrefix[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["honorificPrefix"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HonorificSuffix
		if len(v.HonorificSuffix) == 1 {
			value = v.HonorificSuffix[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["honorificSuffix"] = json.RawMessage(b)
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
		var value interface{} = v.JobTitle
		if len(v.JobTitle) == 1 {
			value = v.JobTitle[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["jobTitle"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Knows
		if len(v.Knows) == 1 {
			value = v.Knows[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["knows"] = json.RawMessage(b)
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
		var value interface{} = v.Nationality
		if len(v.Nationality) == 1 {
			value = v.Nationality[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["nationality"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NetWorth
		if len(v.NetWorth) == 1 {
			value = v.NetWorth[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["netWorth"] = json.RawMessage(b)
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
		var value interface{} = v.Parent
		if len(v.Parent) == 1 {
			value = v.Parent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["parent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PerformerIn
		if len(v.PerformerIn) == 1 {
			value = v.PerformerIn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["performerIn"] = json.RawMessage(b)
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
		var value interface{} = v.RelatedTo
		if len(v.RelatedTo) == 1 {
			value = v.RelatedTo[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["relatedTo"] = json.RawMessage(b)
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
		var value interface{} = v.Sibling
		if len(v.Sibling) == 1 {
			value = v.Sibling[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sibling"] = json.RawMessage(b)
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
		var value interface{} = v.Spouse
		if len(v.Spouse) == 1 {
			value = v.Spouse[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["spouse"] = json.RawMessage(b)
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

	{
		var value interface{} = v.Weight
		if len(v.Weight) == 1 {
			value = v.Weight[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["weight"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkLocation
		if len(v.WorkLocation) == 1 {
			value = v.WorkLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorksFor
		if len(v.WorksFor) == 1 {
			value = v.WorksFor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["worksFor"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Person) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Person"

	return data, nil
}

func (v Person) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
