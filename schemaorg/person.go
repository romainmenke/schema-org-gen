package schemaorg

// Person see : https://schema.org/Person
type Person struct {

Thing// AdditionalName see : /additionalName
// An additional name for a Person, can be used for a middle name.
AdditionalName string `json:"additionalName"`

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// Affiliation see : /affiliation
// An organization that this person is affiliated with. For example, a school/university, a club, or a team.
Affiliation string `json:"affiliation"`

// AlumniOf see : /alumniOf
// An organization that the person is an alumni of. Inverse property: alumni (see: https://schema.org/alumni).
AlumniOf interface{} `json:"alumniOf"` // types : EducationalOrganization Organization

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// BirthDate see : /birthDate
// Date of birth.
BirthDate string `json:"birthDate"`

// BirthPlace see : /birthPlace
// The place where the person was born.
BirthPlace string `json:"birthPlace"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// Children see : /children
// A child of the person.
Children string `json:"children"`

// Colleague see : /colleague
// A colleague of the person. Supersedes colleagues (see: https://schema.org/colleagues).
Colleague interface{} `json:"colleague"` // types : Person URL

// ContactPoint see : /contactPoint
// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
ContactPoint string `json:"contactPoint"`

// DeathDate see : /deathDate
// Date of death.
DeathDate string `json:"deathDate"`

// DeathPlace see : /deathPlace
// The place where the person died.
DeathPlace string `json:"deathPlace"`

// Duns see : /duns
// The Dun & Bradstreet DUNS number for identifying an organization or business person.
Duns string `json:"duns"`

// Email see : /email
// Email address.
Email string `json:"email"`

// FamilyName see : /familyName
// Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
FamilyName string `json:"familyName"`

// FaxNumber see : /faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// Follows see : /follows
// The most generic uni-directional social relation.
Follows string `json:"follows"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// Gender see : /gender
// Gender of the person. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender.
Gender interface{} `json:"gender"` // types : GenderType Text

// GivenName see : /givenName
// Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
GivenName string `json:"givenName"`

// GlobalLocationNumber see : /globalLocationNumber
// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
GlobalLocationNumber string `json:"globalLocationNumber"`

// HasOfferCatalog see : /hasOfferCatalog
// Indicates an OfferCatalog listing for this Organization, Person, or Service.
HasOfferCatalog string `json:"hasOfferCatalog"`

// HasPOS see : /hasPOS
// Points-of-Sales operated by the organization or person.
HasPOS string `json:"hasPOS"`

// Height see : /height
// The height of the item.
Height interface{} `json:"height"` // types : Distance QuantitativeValue

// HomeLocation see : /homeLocation
// A contact location for a person's residence.
HomeLocation interface{} `json:"homeLocation"` // types : ContactPoint Place

// HonorificPrefix see : /honorificPrefix
// An honorific prefix preceding a Person's name such as Dr/Mrs/Mr.
HonorificPrefix string `json:"honorificPrefix"`

// HonorificSuffix see : /honorificSuffix
// An honorific suffix preceding a Person's name such as M.D. /PhD/MSCSW.
HonorificSuffix string `json:"honorificSuffix"`

// IsicV4 see : /isicV4
// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
IsicV4 string `json:"isicV4"`

// JobTitle see : /jobTitle
// The job title of the person (for example, Financial Manager).
JobTitle string `json:"jobTitle"`

// Knows see : /knows
// The most generic bi-directional social/work relation.
Knows string `json:"knows"`

// MakesOffer see : /makesOffer
// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
MakesOffer string `json:"makesOffer"`

// MemberOf see : /memberOf
// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
MemberOf interface{} `json:"memberOf"` // types : Organization ProgramMembership

// Naics see : /naics
// The North American Industry Classification System (NAICS) code for a particular organization or business person.
Naics string `json:"naics"`

// Nationality see : /nationality
// Nationality of the person.
Nationality string `json:"nationality"`

// NetWorth see : /netWorth
// The total financial value of the person as calculated by subtracting assets from liabilities.
NetWorth interface{} `json:"netWorth"` // types : MonetaryAmount PriceSpecification

// Owns see : /owns
// Products owned by the organization or person.
Owns interface{} `json:"owns"` // types : OwnershipInfo Product

// Parent see : /parent
// A parent of this person. Supersedes parents (see: https://schema.org/parents).
Parent string `json:"parent"`

// PerformerIn see : /performerIn
// Event that this person is a performer or participant in.
PerformerIn string `json:"performerIn"`

// PublishingPrinciples see : /publishingPrinciples
// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
// 
// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

// RelatedTo see : /relatedTo
// The most generic familial relation.
RelatedTo string `json:"relatedTo"`

// Seeks see : /seeks
// A pointer to products or services sought by the organization or person (demand).
Seeks string `json:"seeks"`

// Sibling see : /sibling
// A sibling of the person. Supersedes siblings (see: https://schema.org/siblings).
Sibling string `json:"sibling"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// Spouse see : /spouse
// The person's spouse.
Spouse string `json:"spouse"`

// TaxID see : /taxID
// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
TaxID string `json:"taxID"`

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

// VatID see : /vatID
// The Value-added Tax ID of the organization or person.
VatID string `json:"vatID"`

// Weight see : /weight
// The weight of the product or person.
Weight string `json:"weight"`

// WorkLocation see : /workLocation
// A contact location for a person's place of work.
WorkLocation interface{} `json:"workLocation"` // types : ContactPoint Place

// WorksFor see : /worksFor
// Organizations that the person works for.
WorksFor string `json:"worksFor"`

// AdditionalType see : /additionalType
// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the 'typeof' attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
AdditionalType string `json:"additionalType"`

// AlternateName see : /alternateName
// An alias for the item.
AlternateName string `json:"alternateName"`

// Description see : /description
// A description of the item.
Description string `json:"description"`

// DisambiguatingDescription see : /disambiguatingDescription
// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
DisambiguatingDescription string `json:"disambiguatingDescription"`

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

// AccountablePerson see : /accountablePerson
// Specifies the Person that is legally accountable for the CreativeWork. 
AccountablePerson string `json:"accountablePerson"`

// AcquiredFrom see : /acquiredFrom
// The organization or person from which the product was acquired. 
AcquiredFrom string `json:"acquiredFrom"`

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip.  Supersedes actors (see: https://schema.org/actors).
Actor interface{} `json:"actor"` // types : Clip CreativeWorkSeason Episode Event Movie MovieSeries RadioSeries TVSeries VideoGame VideoGameSeries VideoObject

// Agent see : /agent
// The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book. 
Agent string `json:"agent"`

// Alumni see : /alumni
// Alumni of an organization.  inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni interface{} `json:"alumni"` // types : EducationalOrganization Organization

// Athlete see : /athlete
// A person that acts as performing member of a sports team; a player as opposed to a coach. 
Athlete string `json:"athlete"`

// Attendee see : /attendee
// A person or organization attending the event.  Supersedes attendees (see: https://schema.org/attendees).
Attendee string `json:"attendee"`

// Author see : /author
// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably. 
Author interface{} `json:"author"` // types : CreativeWork Rating

// AwayTeam see : /awayTeam
// The away team in a sports event. 
AwayTeam string `json:"awayTeam"`

// BccRecipient see : /bccRecipient
// A sub property of recipient. The recipient blind copied on a message. 
BccRecipient string `json:"bccRecipient"`

// Borrower see : /borrower
// A sub property of participant. The person that borrows the object being lent. 
Borrower string `json:"borrower"`

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.  Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Invoice Order Reservation Service

// Buyer see : /buyer
// A sub property of participant. The participant/person/organization that bought the object. 
Buyer string `json:"buyer"`

// Candidate see : /candidate
// A sub property of object. The candidate subject of this action. 
Candidate string `json:"candidate"`

// CcRecipient see : /ccRecipient
// A sub property of recipient. The recipient copied on a message. 
CcRecipient string `json:"ccRecipient"`

// Character see : /character
// Fictional person connected with a creative work. 
Character string `json:"character"`

// Children see : /children
// A child of the person. 
Children string `json:"children"`

// Coach see : /coach
// A person that acts in a coaching role for a sports team. 
Coach string `json:"coach"`

// Colleague see : /colleague
// A colleague of the person.  Supersedes colleagues (see: https://schema.org/colleagues).
Colleague string `json:"colleague"`

// Competitor see : /competitor
// A competitor in a sports event. 
Competitor string `json:"competitor"`

// Composer see : /composer
// The person or organization who wrote a composition, or who is the composer of a work performed at some event. 
Composer interface{} `json:"composer"` // types : Event MusicComposition

// Contributor see : /contributor
// A secondary contributor to the CreativeWork or Event. 
Contributor interface{} `json:"contributor"` // types : CreativeWork Event

// CopyrightHolder see : /copyrightHolder
// The party holding the legal copyright to the CreativeWork. 
CopyrightHolder string `json:"copyrightHolder"`

// Creator see : /creator
// The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork. 
Creator interface{} `json:"creator"` // types : CreativeWork UserComments

// CreditedTo see : /creditedTo
// The group the release is credited to if different than the byArtist. For example, Red and Blue is credited to "Stefani Germanotta Band", but by Lady Gaga. 
CreditedTo string `json:"creditedTo"`

// Customer see : /customer
// Party placing the order or paying the invoice. 
Customer interface{} `json:"customer"` // types : Invoice Order

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip.  Supersedes directors (see: https://schema.org/directors).
Director interface{} `json:"director"` // types : Clip CreativeWorkSeason Episode Event Movie MovieSeries RadioSeries TVSeries VideoGame VideoGameSeries VideoObject

// Editor see : /editor
// Specifies the Person who edited the CreativeWork. 
Editor string `json:"editor"`

// Employee see : /employee
// Someone working for this organization.  Supersedes employees (see: https://schema.org/employees).
Employee string `json:"employee"`

// Endorsee see : /endorsee
// A sub property of participant. The person/organization being supported. 
Endorsee string `json:"endorsee"`

// Followee see : /followee
// A sub property of object. The person or organization being followed. 
Followee string `json:"followee"`

// Follows see : /follows
// The most generic uni-directional social relation. 
Follows string `json:"follows"`

// Founder see : /founder
// A person who founded this organization.  Supersedes founders (see: https://schema.org/founders).
Founder string `json:"founder"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution. 
Funder interface{} `json:"funder"` // types : CreativeWork Event Organization Person

// Grantee see : /grantee
// The person, organization, contact point, or audience that has been granted this permission. 
Grantee string `json:"grantee"`

// HomeTeam see : /homeTeam
// The home team in a sports event. 
HomeTeam string `json:"homeTeam"`

// Illustrator see : /illustrator
// The illustrator of the book. 
Illustrator string `json:"illustrator"`

// Instructor see : /instructor
// A person assigned to instruct or provide instructional assistance for the CourseInstance (see: https://schema.org/CourseInstance). 
Instructor string `json:"instructor"`

// Knows see : /knows
// The most generic bi-directional social/work relation. 
Knows string `json:"knows"`

// Landlord see : /landlord
// A sub property of participant. The owner of the real estate property. 
Landlord string `json:"landlord"`

// Lender see : /lender
// A sub property of participant. The person that lends the object being borrowed. 
Lender string `json:"lender"`

// Loser see : /loser
// A sub property of participant. The loser of the action. 
Loser string `json:"loser"`

// Lyricist see : /lyricist
// The person who wrote the words. 
Lyricist string `json:"lyricist"`

// Member see : /member
// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals.  Supersedes members (see: https://schema.org/members). inverse property: memberOf (see: https://schema.org/memberOf).
Member interface{} `json:"member"` // types : Organization ProgramMembership

// MusicBy see : /musicBy
// The composer of the soundtrack. 
MusicBy interface{} `json:"musicBy"` // types : Clip Episode Movie MovieSeries RadioSeries TVSeries VideoGame VideoGameSeries VideoObject

// OfferedBy see : /offeredBy
// A pointer to the organization or person making the offer.  inverse property: makesOffer (see: https://schema.org/makesOffer).
OfferedBy string `json:"offeredBy"`

// Opponent see : /opponent
// A sub property of participant. The opponent on this action. 
Opponent string `json:"opponent"`

// Organizer see : /organizer
// An organizer of an Event. 
Organizer string `json:"organizer"`

// Parent see : /parent
// A parent of this person.  Supersedes parents (see: https://schema.org/parents).
Parent string `json:"parent"`

// Participant see : /participant
// Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve. 
Participant string `json:"participant"`

// Performer see : /performer
// A performer at the event—for example, a presenter, musician, musical group or actor.  Supersedes performers (see: https://schema.org/performers).
Performer string `json:"performer"`

// Producer see : /producer
// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.). 
Producer string `json:"producer"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.  Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : BusTrip CreativeWork Flight Invoice ParcelDelivery Reservation Service TrainTrip

// Publisher see : /publisher
// The publisher of the creative work. 
Publisher string `json:"publisher"`

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action. 
Recipient interface{} `json:"recipient"` // types : AuthorizeAction CommunicateAction DonateAction GiveAction Message PayAction ReturnAction SendAction TipAction

// RelatedTo see : /relatedTo
// The most generic familial relation. 
RelatedTo string `json:"relatedTo"`

// ReviewedBy see : /reviewedBy
// People or organizations that have reviewed the content on this web page for accuracy and/or completeness. 
ReviewedBy string `json:"reviewedBy"`

// Seller see : /seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.  Supersedes merchant (see: https://schema.org/merchant).
Seller interface{} `json:"seller"` // types : BuyAction Demand Flight Offer Order

// Sender see : /sender
// A sub property of participant. The participant who is at the sending end of the action. 
Sender interface{} `json:"sender"` // types : Message ReceiveAction

// Sibling see : /sibling
// A sibling of the person.  Supersedes siblings (see: https://schema.org/siblings).
Sibling string `json:"sibling"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event. 
Sponsor interface{} `json:"sponsor"` // types : CreativeWork Event MedicalStudy Organization Person

// Spouse see : /spouse
// The person's spouse. 
Spouse string `json:"spouse"`

// ToRecipient see : /toRecipient
// A sub property of recipient. The recipient who was directly sent the message. 
ToRecipient string `json:"toRecipient"`

// Translator see : /translator
// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event. 
Translator interface{} `json:"translator"` // types : CreativeWork Event

// UnderName see : /underName
// The person or organization the reservation or ticket is for. 
UnderName interface{} `json:"underName"` // types : Reservation Ticket

// Winner see : /winner
// A sub property of participant. The winner of the action. 
Winner string `json:"winner"`

}

