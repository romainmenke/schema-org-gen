package schemaorg

// MusicRecording see : https://schema.org/MusicRecording
type MusicRecording struct {

CreativeWork// ByArtist see : /byArtist
// The artist that performed this album or recording.
ByArtist string `json:"byArtist"`

// Duration see : /duration
// The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
Duration string `json:"duration"`

// InAlbum see : /inAlbum
// The album to which this recording belongs.
InAlbum string `json:"inAlbum"`

// InPlaylist see : /inPlaylist
// The playlist to which this recording belongs.
InPlaylist string `json:"inPlaylist"`

// IsrcCode see : /isrcCode
// The International Standard Recording Code for the recording.
IsrcCode string `json:"isrcCode"`

// RecordingOf see : /recordingOf
// The composition this track is a recording of. Inverse property: recordedAs (see: https://schema.org/recordedAs).
RecordingOf string `json:"recordingOf"`

// About see : /about
// The subject matter of the content.
About string `json:"about"`

// AccessibilityAPI see : /accessibilityAPI
// Indicates that the resource is compatible with the referenced accessibility API (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
AccessibilityAPI string `json:"accessibilityAPI"`

// AggregateRating see : /aggregateRating
// The overall rating, based on a collection of reviews or ratings, of the item.
AggregateRating string `json:"aggregateRating"`

// AlternativeHeadline see : /alternativeHeadline
// A secondary title of the CreativeWork.
AlternativeHeadline string `json:"alternativeHeadline"`

// AssociatedMedia see : /associatedMedia
// A media object that encodes this CreativeWork. This property is a synonym for encoding.
AssociatedMedia string `json:"associatedMedia"`

// Audience see : /audience
// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
Audience string `json:"audience"`

// Audio see : /audio
// An embedded audio object.
Audio string `json:"audio"`

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// ContentRating see : /contentRating
// Official rating of a piece of content—for example,'MPAA PG-13'.
ContentRating string `json:"contentRating"`

// Contributor see : /contributor
// A secondary contributor to the CreativeWork or Event.
Contributor interface{} `json:"contributor"` // types : Organization Person

// CopyrightYear see : /copyrightYear
// The year during which the claimed copyright for the CreativeWork was first asserted.
CopyrightYear string `json:"copyrightYear"`

// Creator see : /creator
// The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork.
Creator interface{} `json:"creator"` // types : Organization Person

// DateCreated see : /dateCreated
// The date on which the CreativeWork was created or the item was added to a DataFeed.
DateCreated interface{} `json:"dateCreated"` // types : Date DateTime

// DateModified see : /dateModified
// The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
DateModified interface{} `json:"dateModified"` // types : Date DateTime

// DatePublished see : /datePublished
// Date of first broadcast/publication.
DatePublished string `json:"datePublished"`

// HasPart see : /hasPart
// Indicates a CreativeWork that is (in some sense) a part of this CreativeWork. Inverse property: isPartOf (see: https://schema.org/isPartOf).
HasPart string `json:"hasPart"`

// Headline see : /headline
// Headline of the article.
Headline string `json:"headline"`

// InLanguage see : /inLanguage
// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
InLanguage interface{} `json:"inLanguage"` // types : Language Text

// InteractionStatistic see : /interactionStatistic
// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used. Supersedes interactionCount (see: https://schema.org/interactionCount).
InteractionStatistic string `json:"interactionStatistic"`

// IsFamilyFriendly see : /isFamilyFriendly
// Indicates whether this content is family friendly.
IsFamilyFriendly string `json:"isFamilyFriendly"`

// IsPartOf see : /isPartOf
// Indicates a CreativeWork that this CreativeWork is (in some sense) part of. Inverse property: hasPart (see: https://schema.org/hasPart).
IsPartOf string `json:"isPartOf"`

// Keywords see : /keywords
// Keywords or tags used to describe this content. Multiple entries in a keywords list are typically delimited by commas.
Keywords string `json:"keywords"`

// MainEntity see : /mainEntity
// Indicates the primary entity described in some page or other CreativeWork. Inverse property: mainEntityOfPage (see: https://schema.org/mainEntityOfPage).
MainEntity string `json:"mainEntity"`

// Material see : /material
// A material that something is made from, e.g. leather, wool, cotton, paper.
Material interface{} `json:"material"` // types : Product Text URL

// Mentions see : /mentions
// Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept.
Mentions string `json:"mentions"`

// Offers see : /offers
// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
Offers string `json:"offers"`

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// SchemaVersion see : /schemaVersion
// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using an URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
SchemaVersion interface{} `json:"schemaVersion"` // types : Text URL

// TimeRequired see : /timeRequired
// Approximate or typical time it takes to work with or through this learning resource for the typical intended target audience, e.g. 'P30M', 'P1H25M'.
TimeRequired string `json:"timeRequired"`

// Video see : /video
// An embedded video object.
Video string `json:"video"`

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

// RecordedAs see : /recordedAs
// An audio recording of the work.  inverse property: recordingOf (see: https://schema.org/recordingOf).
RecordedAs string `json:"recordedAs"`

// Track see : /track
// A music recording (track)—usually a single song. If an ItemList is given, the list should contain items of type MusicRecording.  Supersedes tracks (see: https://schema.org/tracks).
Track interface{} `json:"track"` // types : MusicGroup MusicPlaylist

}

