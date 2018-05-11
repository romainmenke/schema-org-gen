package schemaorg

// Organization see : https://schema.org/Organization
type Organization struct {

Thing// ActionableFeedbackPolicy see : https://pending.schema.org/actionableFeedbackPolicy
// For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
ActionableFeedbackPolicy interface{} `json:"actionableFeedbackPolicy"` // types : CreativeWork URL

// Address see : /address
// Physical address of the item.
Address interface{} `json:"address"` // types : PostalAddress Text

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// Alumni see : /alumni
// Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
Alumni string `json:"alumni"`

// AreaServed see : /areaServed
// The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
AreaServed interface{} `json:"areaServed"` // types : AdministrativeArea GeoShape Place Text

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
Brand interface{} `json:"brand"` // types : Brand Organization

// ContactPoint see : /contactPoint
// A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
ContactPoint string `json:"contactPoint"`

// CorrectionsPolicy see : https://pending.schema.org/correctionsPolicy
// For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
CorrectionsPolicy interface{} `json:"correctionsPolicy"` // types : CreativeWork URL

// Department see : /department
// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
Department string `json:"department"`

// DissolutionDate see : /dissolutionDate
// The date that this organization was dissolved.
DissolutionDate string `json:"dissolutionDate"`

// DiversityPolicy see : https://pending.schema.org/diversityPolicy
// Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
DiversityPolicy interface{} `json:"diversityPolicy"` // types : CreativeWork URL

// Duns see : /duns
// The Dun & Bradstreet DUNS number for identifying an organization or business person.
Duns string `json:"duns"`

// Email see : /email
// Email address.
Email string `json:"email"`

// Employee see : /employee
// Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
Employee string `json:"employee"`

// EthicsPolicy see : https://pending.schema.org/ethicsPolicy
// Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
EthicsPolicy interface{} `json:"ethicsPolicy"` // types : CreativeWork URL

// Event see : /event
// Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
Event string `json:"event"`

// FaxNumber see : /faxNumber
// The fax number.
FaxNumber string `json:"faxNumber"`

// Founder see : /founder
// A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
Founder string `json:"founder"`

// FoundingDate see : /foundingDate
// The date that this organization was founded.
FoundingDate string `json:"foundingDate"`

// FoundingLocation see : /foundingLocation
// The place where the Organization was founded.
FoundingLocation string `json:"foundingLocation"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// GlobalLocationNumber see : /globalLocationNumber
// The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
GlobalLocationNumber string `json:"globalLocationNumber"`

// HasOfferCatalog see : /hasOfferCatalog
// Indicates an OfferCatalog listing for this Organization, Person, or Service.
HasOfferCatalog string `json:"hasOfferCatalog"`

// HasPOS see : /hasPOS
// Points-of-Sales operated by the organization or person.
HasPOS string `json:"hasPOS"`

// IsicV4 see : /isicV4
// The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
IsicV4 string `json:"isicV4"`

// LegalName see : /legalName
// The official name of the organization, e.g. the registered company name.
LegalName string `json:"legalName"`

// LeiCode see : /leiCode
// An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
LeiCode string `json:"leiCode"`

// Location see : /location
// The location of for example where the event is happening, an organization is located, or where an action takes place.
Location interface{} `json:"location"` // types : Place PostalAddress Text

// Logo see : /logo
// An associated logo.
Logo interface{} `json:"logo"` // types : ImageObject URL

// MakesOffer see : /makesOffer
// A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
MakesOffer string `json:"makesOffer"`

// Member see : /member
// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
Member interface{} `json:"member"` // types : Organization Person

// MemberOf see : /memberOf
// An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
MemberOf interface{} `json:"memberOf"` // types : Organization ProgramMembership

// Naics see : /naics
// The North American Industry Classification System (NAICS) code for a particular organization or business person.
Naics string `json:"naics"`

// NumberOfEmployees see : /numberOfEmployees
// The number of employees in an organization e.g. business.
NumberOfEmployees string `json:"numberOfEmployees"`

// Owns see : /owns
// Products owned by the organization or person.
Owns interface{} `json:"owns"` // types : OwnershipInfo Product

// ParentOrganization see : /parentOrganization
// The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
ParentOrganization string `json:"parentOrganization"`

// PublishingPrinciples see : /publishingPrinciples
// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
// 
// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// Seeks see : /seeks
// A pointer to products or services sought by the organization or person (demand).
Seeks string `json:"seeks"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// SubOrganization see : /subOrganization
// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific 'department' property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
SubOrganization string `json:"subOrganization"`

// TaxID see : /taxID
// The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
TaxID string `json:"taxID"`

// Telephone see : /telephone
// The telephone number.
Telephone string `json:"telephone"`

// UnnamedSourcesPolicy see : https://pending.schema.org/unnamedSourcesPolicy
// For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
UnnamedSourcesPolicy interface{} `json:"unnamedSourcesPolicy"` // types : CreativeWork URL

// VatID see : /vatID
// The Value-added Tax ID of the organization or person.
VatID string `json:"vatID"`

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

// AcquiredFrom see : /acquiredFrom
// The organization or person from which the product was acquired. 
AcquiredFrom string `json:"acquiredFrom"`

// Affiliation see : /affiliation
// An organization that this person is affiliated with. For example, a school/university, a club, or a team. 
Affiliation string `json:"affiliation"`

// Agent see : /agent
// The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book. 
Agent string `json:"agent"`

// AlumniOf see : /alumniOf
// An organization that the person is an alumni of.  inverse property: alumni (see: https://schema.org/alumni).
AlumniOf string `json:"alumniOf"`

// Attendee see : /attendee
// A person or organization attending the event.  Supersedes attendees (see: https://schema.org/attendees).
Attendee string `json:"attendee"`

// Author see : /author
// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably. 
Author interface{} `json:"author"` // types : CreativeWork Rating

// BccRecipient see : /bccRecipient
// A sub property of recipient. The recipient blind copied on a message. 
BccRecipient string `json:"bccRecipient"`

// Brand see : /brand
// The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person. 
Brand interface{} `json:"brand"` // types : Organization Person Product Service

// BroadcastAffiliateOf see : /broadcastAffiliateOf
// The media network(s) whose content is broadcast on this station. 
BroadcastAffiliateOf string `json:"broadcastAffiliateOf"`

// Broadcaster see : /broadcaster
// The organization owning or operating the broadcast service. 
Broadcaster string `json:"broadcaster"`

// Broker see : /broker
// An entity that arranges for an exchange between a buyer and a seller.  In most cases a broker never acquires or releases ownership of a product or service involved in an exchange.  If it is not clear whether an entity is a broker, seller, or buyer, the latter two terms are preferred.  Supersedes bookingAgent (see: https://schema.org/bookingAgent).
Broker interface{} `json:"broker"` // types : Invoice Order Reservation Service

// CcRecipient see : /ccRecipient
// A sub property of recipient. The recipient copied on a message. 
CcRecipient string `json:"ccRecipient"`

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

// Department see : /department
// A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe. 
Department string `json:"department"`

// Endorsee see : /endorsee
// A sub property of participant. The person/organization being supported. 
Endorsee string `json:"endorsee"`

// Followee see : /followee
// A sub property of object. The person or organization being followed. 
Followee string `json:"followee"`

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution. 
Funder interface{} `json:"funder"` // types : CreativeWork Event Organization Person

// Grantee see : /grantee
// The person, organization, contact point, or audience that has been granted this permission. 
Grantee string `json:"grantee"`

// HiringOrganization see : /hiringOrganization
// Organization offering the job position. 
HiringOrganization string `json:"hiringOrganization"`

// HostingOrganization see : /hostingOrganization
// The organization (airline, travelers' club, etc.) the membership is made with. 
HostingOrganization string `json:"hostingOrganization"`

// IssuedBy see : /issuedBy
// The organization issuing the ticket or permit. 
IssuedBy interface{} `json:"issuedBy"` // types : Permit Ticket

// Landlord see : /landlord
// A sub property of participant. The owner of the real estate property. 
Landlord string `json:"landlord"`

// Lender see : /lender
// A sub property of participant. The person that lends the object being borrowed. 
Lender string `json:"lender"`

// Manufacturer see : /manufacturer
// The manufacturer of the product. 
Manufacturer interface{} `json:"manufacturer"` // types : DietarySupplement Drug Product

// Member see : /member
// A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals.  Supersedes members (see: https://schema.org/members). inverse property: memberOf (see: https://schema.org/memberOf).
Member interface{} `json:"member"` // types : Organization ProgramMembership

// MemberOf see : /memberOf
// An Organization (or ProgramMembership) to which this Person or Organization belongs.  inverse property: member (see: https://schema.org/member).
MemberOf interface{} `json:"memberOf"` // types : Organization Person

// OfferedBy see : /offeredBy
// A pointer to the organization or person making the offer.  inverse property: makesOffer (see: https://schema.org/makesOffer).
OfferedBy string `json:"offeredBy"`

// Organizer see : /organizer
// An organizer of an Event. 
Organizer string `json:"organizer"`

// ParentOrganization see : /parentOrganization
// The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any.  Supersedes branchOf (see: https://schema.org/branchOf). inverse property: subOrganization (see: https://schema.org/subOrganization).
ParentOrganization string `json:"parentOrganization"`

// Participant see : /participant
// Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve. 
Participant string `json:"participant"`

// Performer see : /performer
// A performer at the event—for example, a presenter, musician, musical group or actor.  Supersedes performers (see: https://schema.org/performers).
Performer string `json:"performer"`

// Producer see : /producer
// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.). 
Producer string `json:"producer"`

// ProductionCompany see : /productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc. 
ProductionCompany interface{} `json:"productionCompany"` // types : CreativeWorkSeason Episode MediaObject Movie MovieSeries RadioSeries TVSeries VideoGameSeries

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.  Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : BusTrip CreativeWork Flight Invoice ParcelDelivery Reservation Service TrainTrip

// Publisher see : /publisher
// The publisher of the creative work. 
Publisher string `json:"publisher"`

// Recipient see : /recipient
// A sub property of participant. The participant who is at the receiving end of the action. 
Recipient interface{} `json:"recipient"` // types : AuthorizeAction CommunicateAction DonateAction GiveAction Message PayAction ReturnAction SendAction TipAction

// RecordLabel see : /recordLabel
// The label that issued the release. 
RecordLabel string `json:"recordLabel"`

// ReviewedBy see : /reviewedBy
// People or organizations that have reviewed the content on this web page for accuracy and/or completeness. 
ReviewedBy string `json:"reviewedBy"`

// Seller see : /seller
// An entity which offers (sells / leases / lends / loans) the services / goods.  A seller may also be a provider.  Supersedes vendor (see: https://schema.org/vendor).
Seller interface{} `json:"seller"` // types : BuyAction Demand Flight Offer Order

// Sender see : /sender
// A sub property of participant. The participant who is at the sending end of the action. 
Sender interface{} `json:"sender"` // types : Message ReceiveAction

// ServiceOperator see : /serviceOperator
// The operating organization, if different from the provider.  This enables the representation of services that are provided by an organization, but operated by another organization like a subcontractor. 
ServiceOperator string `json:"serviceOperator"`

// SourceOrganization see : /sourceOrganization
// The Organization on whose behalf the creator was working. 
SourceOrganization string `json:"sourceOrganization"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event. 
Sponsor interface{} `json:"sponsor"` // types : CreativeWork Event MedicalStudy Organization Person

// SubOrganization see : /subOrganization
// A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific 'department' property.  inverse property: parentOrganization (see: https://schema.org/parentOrganization).
SubOrganization string `json:"subOrganization"`

// ToRecipient see : /toRecipient
// A sub property of recipient. The recipient who was directly sent the message. 
ToRecipient string `json:"toRecipient"`

// Translator see : /translator
// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event. 
Translator interface{} `json:"translator"` // types : CreativeWork Event

// UnderName see : /underName
// The person or organization the reservation or ticket is for. 
UnderName interface{} `json:"underName"` // types : Reservation Ticket

// WorksFor see : /worksFor
// Organizations that the person works for. 
WorksFor string `json:"worksFor"`

}

