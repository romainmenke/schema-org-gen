package schemaorggo

import "encoding/json"

// RadioSeries see : https://schema.org/RadioSeries
type RadioSeries struct {

	// With properties from CreativeWorkSeries see : https://schema.org/CreativeWorkSeries
	//

	// With properties from CreativeWork see : https://schema.org/CreativeWork
	//

	// With properties from Series see : https://schema.org/Series
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Intangible see : https://schema.org/Intangible
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// About see : https://schema.org/about
	// The subject matter of the content.
	// types : Thing
	About []*Thing `json:"about,omitempty"`

	// AccessMode see : https://schema.org/accessMode
	// The human sensory perceptual system or cognitive faculty through which a person may process or perceive information. Expected values include: auditory, tactile, textual, visual, colorDependent, chartOnVisual, chemOnVisual, diagramOnVisual, mathOnVisual, musicOnVisual, textOnVisual.
	//
	// types : Text
	AccessMode []string `json:"accessMode,omitempty"`

	// AccessModeSufficient see : https://schema.org/accessModeSufficient
	// A list of single or combined accessModes that are sufficient to understand all the intellectual content of a resource. Expected values include:  auditory, tactile, textual, visual.
	//
	// types : Text
	AccessModeSufficient []string `json:"accessModeSufficient,omitempty"`

	// AccessibilityAPI see : https://schema.org/accessibilityAPI
	// Indicates that the resource is compatible with the referenced accessibility API ([WebSchemas wiki lists possible values](http://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityAPI []string `json:"accessibilityAPI,omitempty"`

	// AccessibilityControl see : https://schema.org/accessibilityControl
	// Identifies input methods that are sufficient to fully control the described resource ([WebSchemas wiki lists possible values](http://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityControl []string `json:"accessibilityControl,omitempty"`

	// AccessibilityFeature see : https://schema.org/accessibilityFeature
	// Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility ([WebSchemas wiki lists possible values](http://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityFeature []string `json:"accessibilityFeature,omitempty"`

	// AccessibilityHazard see : https://schema.org/accessibilityHazard
	// A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 ([WebSchemas wiki lists possible values](http://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityHazard []string `json:"accessibilityHazard,omitempty"`

	// AccessibilitySummary see : https://schema.org/accessibilitySummary
	// A human-readable summary of specific accessibility features or deficiencies, consistent with the other accessibility metadata but expressing subtleties such as &quot;short descriptions are present but long descriptions will be needed for non-visual users&quot; or &quot;short descriptions are present and no long descriptions are needed.&quot;
	// types : Text
	AccessibilitySummary []string `json:"accessibilitySummary,omitempty"`

	// AccountablePerson see : https://schema.org/accountablePerson
	// Specifies the Person that is legally accountable for the CreativeWork.
	// types : Person
	AccountablePerson []*Person `json:"accountablePerson,omitempty"`

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip.
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// Actors see : https://schema.org/actors
	// An actor, e.g. in tv, radio, movie, video games etc. Actors can be associated with individual items or with a series, episode, clip.
	// types : Person
	Actors []*Person `json:"actors,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	// types : AggregateRating
	AggregateRating []*AggregateRating `json:"aggregateRating,omitempty"`

	// AlternateName see : https://schema.org/alternateName
	// An alias for the item.
	// types : Text
	AlternateName []string `json:"alternateName,omitempty"`

	// AlternativeHeadline see : https://schema.org/alternativeHeadline
	// A secondary title of the CreativeWork.
	// types : Text
	AlternativeHeadline []string `json:"alternativeHeadline,omitempty"`

	// AssociatedMedia see : https://schema.org/associatedMedia
	// A media object that encodes this CreativeWork. This property is a synonym for encoding.
	// types : MediaObject
	AssociatedMedia []*MediaObject `json:"associatedMedia,omitempty"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created.
	// types : Audience
	Audience []*Audience `json:"audience,omitempty"`

	// Audio see : https://schema.org/audio
	// An embedded audio object.
	// types : AudioObject
	Audio []*AudioObject `json:"audio,omitempty"`

	// Author see : https://schema.org/author
	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	// types : Organization Person
	Author []interface{} `json:"author,omitempty"`

	// Award see : https://schema.org/award
	// An award won by or for this item.
	// types : Text
	Award []string `json:"award,omitempty"`

	// Awards see : https://schema.org/awards
	// Awards won by or for this item.
	// types : Text
	Awards []string `json:"awards,omitempty"`

	// Character see : https://schema.org/character
	// Fictional person connected with a creative work.
	// types : Person
	Character []*Person `json:"character,omitempty"`

	// Citation see : https://schema.org/citation
	// A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
	// types : CreativeWork Text
	Citation []interface{} `json:"citation,omitempty"`

	// Comment see : https://schema.org/comment
	// Comments, typically from users.
	// types : Comment
	Comment []*Comment `json:"comment,omitempty"`

	// CommentCount see : https://schema.org/commentCount
	// The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
	// types : Integer
	CommentCount []float64 `json:"commentCount,omitempty"`

	// ContainsSeason see : https://schema.org/containsSeason
	// A season that is part of the media series.
	// types : CreativeWorkSeason
	ContainsSeason []*CreativeWorkSeason `json:"containsSeason,omitempty"`

	// ContentLocation see : https://schema.org/contentLocation
	// The location depicted or described in the content. For example, the location in a photograph or painting.
	// types : Place
	ContentLocation []*Place `json:"contentLocation,omitempty"`

	// ContentRating see : https://schema.org/contentRating
	// Official rating of a piece of content&amp;#x2014;for example,&#39;MPAA PG-13&#39;.
	// types : Text Rating
	ContentRating []interface{} `json:"contentRating,omitempty"`

	// Contributor see : https://schema.org/contributor
	// A secondary contributor to the CreativeWork or Event.
	// types : Organization Person
	Contributor []interface{} `json:"contributor,omitempty"`

	// CopyrightHolder see : https://schema.org/copyrightHolder
	// The party holding the legal copyright to the CreativeWork.
	// types : Organization Person
	CopyrightHolder []interface{} `json:"copyrightHolder,omitempty"`

	// CopyrightYear see : https://schema.org/copyrightYear
	// The year during which the claimed copyright for the CreativeWork was first asserted.
	// types : Number
	CopyrightYear []float64 `json:"copyrightYear,omitempty"`

	// Creator see : https://schema.org/creator
	// The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork.
	// types : Organization Person
	Creator []interface{} `json:"creator,omitempty"`

	// DateCreated see : https://schema.org/dateCreated
	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	// types : Date DateTime
	DateCreated []interface{} `json:"dateCreated,omitempty"`

	// DateModified see : https://schema.org/dateModified
	// The date on which the CreativeWork was most recently modified or when the item&#39;s entry was modified within a DataFeed.
	// types : Date DateTime
	DateModified []interface{} `json:"dateModified,omitempty"`

	// DatePublished see : https://schema.org/datePublished
	// Date of first broadcast/publication.
	// types : Date
	DatePublished []Date `json:"datePublished,omitempty"`

	// Description see : https://schema.org/description
	// A description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip.
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// Directors see : https://schema.org/directors
	// A director of e.g. tv, radio, movie, video games etc. content. Directors can be associated with individual items or with a series, episode, clip.
	// types : Person
	Directors []*Person `json:"directors,omitempty"`

	// DisambiguatingDescription see : https://schema.org/disambiguatingDescription
	// A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	// types : Text
	DisambiguatingDescription []string `json:"disambiguatingDescription,omitempty"`

	// DiscussionUrl see : https://schema.org/discussionUrl
	// A link to the page containing the comments of the CreativeWork.
	// types : URL
	DiscussionUrl []string `json:"discussionUrl,omitempty"`

	// Editor see : https://schema.org/editor
	// Specifies the Person who edited the CreativeWork.
	// types : Person
	Editor []*Person `json:"editor,omitempty"`

	// EducationalAlignment see : https://schema.org/educationalAlignment
	// An alignment to an established educational framework.
	// types : AlignmentObject
	EducationalAlignment []*AlignmentObject `json:"educationalAlignment,omitempty"`

	// EducationalUse see : https://schema.org/educationalUse
	// The purpose of a work in the context of education; for example, &#39;assignment&#39;, &#39;group work&#39;.
	// types : Text
	EducationalUse []string `json:"educationalUse,omitempty"`

	// Encoding see : https://schema.org/encoding
	// A media object that encodes this CreativeWork. This property is a synonym for associatedMedia.
	// types : MediaObject
	Encoding []*MediaObject `json:"encoding,omitempty"`

	// EncodingFormat see : https://schema.org/encodingFormat
	// Media type typically expressed using a MIME format (see [IANA site](http://www.iana.org/assignments/media-types/media-types.xhtml) and [MDN reference](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)) e.g. application/zip for a SoftwareApplication binary, audio/mpeg for .mp3 etc.).
	//
	// In cases where a [[CreativeWork]] has several media type representations, [[encoding]] can be used to indicate each [[MediaObject]] alongside particular [[encodingFormat]] information.
	//
	// Unregistered or niche encoding and file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia/Wikidata entry.
	// types : Text URL
	EncodingFormat []string `json:"encodingFormat,omitempty"`

	// Encodings see : https://schema.org/encodings
	// A media object that encodes this CreativeWork.
	// types : MediaObject
	Encodings []*MediaObject `json:"encodings,omitempty"`

	// EndDate see : https://schema.org/endDate
	// The end date and time of the item (in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date
	EndDate []Date `json:"endDate,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season.
	// types : Episode
	Episode []*Episode `json:"episode,omitempty"`

	// Episodes see : https://schema.org/episodes
	// An episode of a TV/radio series or season.
	// types : Episode
	Episodes []*Episode `json:"episodes,omitempty"`

	// ExampleOfWork see : https://schema.org/exampleOfWork
	// A creative work that this work is an example/instance/realization/derivation of.
	// types : CreativeWork
	ExampleOfWork []*CreativeWork `json:"exampleOfWork,omitempty"`

	// Expires see : https://schema.org/expires
	// Date the content expires and is no longer useful or available. For example a [[VideoObject]] or [[NewsArticle]] whose availability or relevance is time-limited, or a [[ClaimReview]] fact check whose publisher wants to indicate that it may no longer be relevant (or helpful to highlight) after some date.
	// types : Date
	Expires []Date `json:"expires,omitempty"`

	// FileFormat see : https://schema.org/fileFormat
	// Media type, typically MIME format (see [IANA site](http://www.iana.org/assignments/media-types/media-types.xhtml)) of the content e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, &#39;encoding&#39; can be used to indicate each MediaObject alongside particular fileFormat information. Unregistered or niche file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia entry.
	// types : Text URL
	FileFormat []string `json:"fileFormat,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder []interface{} `json:"funder,omitempty"`

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	// types : Text URL
	Genre []string `json:"genre,omitempty"`

	// HasPart see : https://schema.org/hasPart
	// Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense).
	// types : CreativeWork
	HasPart []*CreativeWork `json:"hasPart,omitempty"`

	// Headline see : https://schema.org/headline
	// Headline of the article.
	// types : Text
	Headline []string `json:"headline,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	//
	// types : URL Text PropertyValue
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the [IETF BCP 47 standard](http://tools.ietf.org/html/bcp47). See also [[availableLanguage]].
	// types : Text Language
	InLanguage []interface{} `json:"inLanguage,omitempty"`

	// InteractionStatistic see : https://schema.org/interactionStatistic
	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used.
	// types : InteractionCounter
	InteractionStatistic []*InteractionCounter `json:"interactionStatistic,omitempty"`

	// InteractivityType see : https://schema.org/interactivityType
	// The predominant mode of learning supported by the learning resource. Acceptable values are &#39;active&#39;, &#39;expositive&#39;, or &#39;mixed&#39;.
	// types : Text
	InteractivityType []string `json:"interactivityType,omitempty"`

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free.
	// types : Boolean
	IsAccessibleForFree []bool `json:"isAccessibleForFree,omitempty"`

	// IsBasedOn see : https://schema.org/isBasedOn
	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
	// types : URL CreativeWork Product
	IsBasedOn []interface{} `json:"isBasedOn,omitempty"`

	// IsBasedOnUrl see : https://schema.org/isBasedOnUrl
	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
	// types : URL CreativeWork Product
	IsBasedOnUrl []interface{} `json:"isBasedOnUrl,omitempty"`

	// IsFamilyFriendly see : https://schema.org/isFamilyFriendly
	// Indicates whether this content is family friendly.
	// types : Boolean
	IsFamilyFriendly []bool `json:"isFamilyFriendly,omitempty"`

	// IsPartOf see : https://schema.org/isPartOf
	// Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of.
	// types : CreativeWork
	IsPartOf []*CreativeWork `json:"isPartOf,omitempty"`

	// Issn see : https://schema.org/issn
	// The International Standard Serial Number (ISSN) that identifies this serial publication. You can repeat this property to identify different formats of, or the linking ISSN (ISSN-L) for, this serial publication.
	// types : Text
	Issn []string `json:"issn,omitempty"`

	// Keywords see : https://schema.org/keywords
	// Keywords or tags used to describe this content. Multiple entries in a keywords list are typically delimited by commas.
	// types : Text
	Keywords []string `json:"keywords,omitempty"`

	// LearningResourceType see : https://schema.org/learningResourceType
	// The predominant type or kind characterizing the learning resource. For example, &#39;presentation&#39;, &#39;handout&#39;.
	// types : Text
	LearningResourceType []string `json:"learningResourceType,omitempty"`

	// License see : https://schema.org/license
	// A license document that applies to this content, typically indicated by URL.
	// types : CreativeWork URL
	License []interface{} `json:"license,omitempty"`

	// LocationCreated see : https://schema.org/locationCreated
	// The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork.
	// types : Place
	LocationCreated []*Place `json:"locationCreated,omitempty"`

	// MainEntity see : https://schema.org/mainEntity
	// Indicates the primary entity described in some page or other CreativeWork.
	// types : Thing
	MainEntity []*Thing `json:"mainEntity,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Material see : https://schema.org/material
	// A material that something is made from, e.g. leather, wool, cotton, paper.
	// types : Text URL Product
	Material []interface{} `json:"material,omitempty"`

	// Mentions see : https://schema.org/mentions
	// Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept.
	// types : Thing
	Mentions []*Thing `json:"mentions,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy []interface{} `json:"musicBy,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes []float64 `json:"numberOfEpisodes,omitempty"`

	// NumberOfSeasons see : https://schema.org/numberOfSeasons
	// The number of seasons in this series.
	// types : Integer
	NumberOfSeasons []float64 `json:"numberOfSeasons,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	// types : Text Integer
	Position []interface{} `json:"position,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Producer see : https://schema.org/producer
	// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.).
	// types : Person Organization
	Producer []interface{} `json:"producer,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany []*Organization `json:"productionCompany,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	// types : Person Organization
	Provider []interface{} `json:"provider,omitempty"`

	// Publication see : https://schema.org/publication
	// A publication event associated with the item.
	// types : PublicationEvent
	Publication []*PublicationEvent `json:"publication,omitempty"`

	// Publisher see : https://schema.org/publisher
	// The publisher of the creative work.
	// types : Organization Person
	Publisher []interface{} `json:"publisher,omitempty"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via [[URL]]) a document describing the editorial principles of an [[Organization]] (or individual e.g. a [[Person]] writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a [[CreativeWork]] (e.g. [[NewsArticle]]) the principles are those of the party primarily responsible for the creation of the [[CreativeWork]].
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a [[funder]]) can be expressed using schema.org terminology.
	//
	// types : URL CreativeWork
	PublishingPrinciples []interface{} `json:"publishingPrinciples,omitempty"`

	// RecordedAt see : https://schema.org/recordedAt
	// The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event.
	// types : Event
	RecordedAt []*Event `json:"recordedAt,omitempty"`

	// ReleasedEvent see : https://schema.org/releasedEvent
	// The place and time the release was issued, expressed as a PublicationEvent.
	// types : PublicationEvent
	ReleasedEvent []*PublicationEvent `json:"releasedEvent,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Reviews see : https://schema.org/reviews
	// Review of the item.
	// types : Review
	Reviews []*Review `json:"reviews,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SchemaVersion see : https://schema.org/schemaVersion
	// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using an URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
	// types : URL Text
	SchemaVersion []string `json:"schemaVersion,omitempty"`

	// Season see : https://schema.org/season
	// A season in a media series.
	// types : CreativeWorkSeason
	Season []*CreativeWorkSeason `json:"season,omitempty"`

	// Seasons see : https://schema.org/seasons
	// A season in a media series.
	// types : CreativeWorkSeason
	Seasons []*CreativeWorkSeason `json:"seasons,omitempty"`

	// SourceOrganization see : https://schema.org/sourceOrganization
	// The Organization on whose behalf the creator was working.
	// types : Organization
	SourceOrganization []*Organization `json:"sourceOrganization,omitempty"`

	// SpatialCoverage see : https://schema.org/spatialCoverage
	// The spatialCoverage of a CreativeWork indicates the place(s) which are the focus of the content. It is a subproperty of
	//       contentLocation intended primarily for more technical and detailed materials. For example with a Dataset, it indicates
	//       areas that the dataset describes: a dataset of New York weather would have spatialCoverage which was the place: the state of New York.
	// types : Place
	SpatialCoverage []*Place `json:"spatialCoverage,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor []interface{} `json:"sponsor,omitempty"`

	// StartDate see : https://schema.org/startDate
	// The start date and time of the item (in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601)).
	// types : Date DateTime
	StartDate []interface{} `json:"startDate,omitempty"`

	// TemporalCoverage see : https://schema.org/temporalCoverage
	// The temporalCoverage of a CreativeWork indicates the period that the content applies to, i.e. that it describes, either as a DateTime or as a textual string indicating a time period in [ISO 8601 time interval format](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals). In
	//       the case of a Dataset it will typically indicate the relevant time period in a precise notation (e.g. for a 2011 census dataset, the year 2011 would be written &quot;2011/2012&quot;). Other forms of content e.g. ScholarlyArticle, Book, TVSeries or TVEpisode may indicate their temporalCoverage in broader terms - textually or via well-known URL.
	//       Written works such as books may sometimes have precise temporal coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601 interval format format via &quot;1939/1945&quot;.
	// types : DateTime Text URL
	TemporalCoverage []interface{} `json:"temporalCoverage,omitempty"`

	// Text see : https://schema.org/text
	// The textual content of this CreativeWork.
	// types : Text
	Text []string `json:"text,omitempty"`

	// ThumbnailUrl see : https://schema.org/thumbnailUrl
	// A thumbnail image relevant to the Thing.
	// types : URL
	ThumbnailUrl []string `json:"thumbnailUrl,omitempty"`

	// TimeRequired see : https://schema.org/timeRequired
	// Approximate or typical time it takes to work with or through this learning resource for the typical intended target audience, e.g. &#39;P30M&#39;, &#39;P1H25M&#39;.
	// types : Duration
	TimeRequired []*Duration `json:"timeRequired,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer []*VideoObject `json:"trailer,omitempty"`

	// Translator see : https://schema.org/translator
	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	// types : Person Organization
	Translator []interface{} `json:"translator,omitempty"`

	// TypicalAgeRange see : https://schema.org/typicalAgeRange
	// The typical expected age range, e.g. &#39;7-9&#39;, &#39;11-&#39;.
	// types : Text
	TypicalAgeRange []string `json:"typicalAgeRange,omitempty"`

	// Url see : https://schema.org/url
	// URL of the item.
	// types : URL
	Url []string `json:"url,omitempty"`

	// Version see : https://schema.org/version
	// The version of the CreativeWork embodied by a specified resource.
	// types : Number Text
	Version []interface{} `json:"version,omitempty"`

	// Video see : https://schema.org/video
	// An embedded video object.
	// types : VideoObject
	Video []*VideoObject `json:"video,omitempty"`

	// WorkExample see : https://schema.org/workExample
	// Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook.
	// types : CreativeWork
	WorkExample []*CreativeWork `json:"workExample,omitempty"`
}

func (v RadioSeries) MarshalJSON() ([]byte, error) {
	type Alias RadioSeries

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"RadioSeries\","), b[1:]...), nil
}
