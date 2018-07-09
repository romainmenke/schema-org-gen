package schemaorggo

import "encoding/json"

// Clip see : https://schema.org/Clip
type Clip struct {
	typeContext

	// With properties from CreativeWork see : https://schema.org/CreativeWork
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// About see : https://schema.org/about
	// The subject matter of the content. Inverse property: subjectOf (see: https://schema.orghttps://pending.schema.org/subjectOf).
	// types : Thing
	About []*Thing `json:"about,omitempty"`

	// AccessMode see : https://schema.org/accessMode
	// The human sensory perceptual system or cognitive faculty through which a person may process or perceive information. Expected values include: auditory, tactile, textual, visual, colorDependent, chartOnVisual, chemOnVisual, diagramOnVisual, mathOnVisual, musicOnVisual, textOnVisual.
	// types : Text
	AccessMode []string `json:"accessMode,omitempty"`

	// AccessModeSufficient see : https://schema.org/accessModeSufficient
	// A list of single or combined accessModes that are sufficient to understand all the intellectual content of a resource. Expected values include:  auditory, tactile, textual, visual.
	// types : Text
	AccessModeSufficient []string `json:"accessModeSufficient,omitempty"`

	// AccessibilityAPI see : https://schema.org/accessibilityAPI
	// Indicates that the resource is compatible with the referenced accessibility API (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityAPI []string `json:"accessibilityAPI,omitempty"`

	// AccessibilityControl see : https://schema.org/accessibilityControl
	// Identifies input methods that are sufficient to fully control the described resource (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityControl []string `json:"accessibilityControl,omitempty"`

	// AccessibilityFeature see : https://schema.org/accessibilityFeature
	// Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	// types : Text
	AccessibilityFeature []string `json:"accessibilityFeature,omitempty"`

	// AccessibilityHazard see : https://schema.org/accessibilityHazard
	// A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
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
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

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
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
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
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	// types : Text
	Award []string `json:"award,omitempty"`

	// Character see : https://schema.org/character
	// Fictional person connected with a creative work.
	// types : Person
	Character []*Person `json:"character,omitempty"`

	// Citation see : https://schema.org/citation
	// A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
	// types : CreativeWork Text
	Citation []interface{} `json:"citation,omitempty"`

	// ClipNumber see : https://schema.org/clipNumber
	// Position of the clip within an ordered group of clips.
	// types : Integer Text
	ClipNumber []interface{} `json:"clipNumber,omitempty"`

	// Comment see : https://schema.org/comment
	// Comments, typically from users.
	// types : Comment
	Comment []*Comment `json:"comment,omitempty"`

	// CommentCount see : https://schema.org/commentCount
	// The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
	// types : Integer
	CommentCount []float64 `json:"commentCount,omitempty"`

	// ContentLocation see : https://schema.org/contentLocation
	// The location depicted or described in the content. For example, the location in a photograph or painting.
	// types : Place
	ContentLocation []*Place `json:"contentLocation,omitempty"`

	// ContentRating see : https://schema.org/contentRating
	// Official rating of a piece of content—for example,&#39;MPAA PG-13&#39;.
	// types : Rating Text
	ContentRating []interface{} `json:"contentRating,omitempty"`

	// ContentReferenceTime see : https://pending.schema.org/contentReferenceTime
	// The specific time described by a creative work, for works (e.g. articles, video objects etc.) that emphasise a particular moment within an Event.
	// types : DateTime
	ContentReferenceTime []DateTime `json:"contentReferenceTime,omitempty"`

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

	// Correction see : https://pending.schema.org/correction
	// Indicates a correction to a CreativeWork (see: https://schema.org/CreativeWork), either via a CorrectionComment (see: https://schema.org/CorrectionComment), textually or in another document.
	// types : CorrectionComment Text URL
	Correction []interface{} `json:"correction,omitempty"`

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
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

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
	// A media object that encodes this CreativeWork. This property is a synonym for associatedMedia. Supersedes encodings (see: https://schema.org/encodings).
	// types : MediaObject
	Encoding []*MediaObject `json:"encoding,omitempty"`

	// EncodingFormat see : https://schema.org/encodingFormat
	// Media type typically expressed using a MIME format (see IANA site (see: https://schema.orghttp://www.iana.org/assignments/media-types/media-types.xhtml) and MDN reference (see: https://schema.orghttps://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types)) e.g. application/zip for a SoftwareApplication binary, audio/mpeg for .mp3 etc.).
	//
	// In cases where a CreativeWork (see: https://schema.org/CreativeWork) has several media type representations, encoding (see: https://schema.org/encoding) can be used to indicate each MediaObject (see: https://schema.org/MediaObject) alongside particular encodingFormat (see: https://schema.org/encodingFormat) information.
	//
	// Unregistered or niche encoding and file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia/Wikidata entry. Supersedes fileFormat (see: https://schema.org/fileFormat).
	// types : Text URL
	EncodingFormat []string `json:"encodingFormat,omitempty"`

	// ExampleOfWork see : https://schema.org/exampleOfWork
	// A creative work that this work is an example/instance/realization/derivation of. Inverse property: workExample (see: https://schema.org/workExample).
	// types : CreativeWork
	ExampleOfWork []*CreativeWork `json:"exampleOfWork,omitempty"`

	// Expires see : https://schema.org/expires
	// Date the content expires and is no longer useful or available. For example a VideoObject (see: https://schema.org/VideoObject) or NewsArticle (see: https://schema.org/NewsArticle) whose availability or relevance is time-limited, or a ClaimReview (see: https://schema.org/ClaimReview) fact check whose publisher wants to indicate that it may no longer be relevant (or helpful to highlight) after some date.
	// types : Date
	Expires []Date `json:"expires,omitempty"`

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	// types : Organization Person
	Funder []interface{} `json:"funder,omitempty"`

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	// types : Text URL
	Genre []string `json:"genre,omitempty"`

	// HasPart see : https://schema.org/hasPart
	// Indicates an item or CreativeWork that is part of this item, or CreativeWork (in some sense). Inverse property: isPartOf (see: https://schema.org/isPartOf).
	// types : CreativeWork Trip
	HasPart []interface{} `json:"hasPart,omitempty"`

	// Headline see : https://schema.org/headline
	// Headline of the article.
	// types : Text
	Headline []string `json:"headline,omitempty"`

	// Identifier see : https://schema.org/identifier
	// The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	// types : PropertyValue Text URL
	Identifier []interface{} `json:"identifier,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	// types : ImageObject URL
	Image []interface{} `json:"image,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	// types : Language Text
	InLanguage []interface{} `json:"inLanguage,omitempty"`

	// InteractionStatistic see : https://schema.org/interactionStatistic
	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used. Supersedes interactionCount (see: https://schema.org/interactionCount).
	// types : InteractionCounter
	InteractionStatistic []*InteractionCounter `json:"interactionStatistic,omitempty"`

	// InteractivityType see : https://schema.org/interactivityType
	// The predominant mode of learning supported by the learning resource. Acceptable values are &#39;active&#39;, &#39;expositive&#39;, or &#39;mixed&#39;.
	// types : Text
	InteractivityType []string `json:"interactivityType,omitempty"`

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	// types : Boolean
	IsAccessibleForFree []bool `json:"isAccessibleForFree,omitempty"`

	// IsBasedOn see : https://schema.org/isBasedOn
	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html. Supersedes isBasedOnUrl (see: https://schema.org/isBasedOnUrl).
	// types : CreativeWork Product URL
	IsBasedOn []interface{} `json:"isBasedOn,omitempty"`

	// IsFamilyFriendly see : https://schema.org/isFamilyFriendly
	// Indicates whether this content is family friendly.
	// types : Boolean
	IsFamilyFriendly []bool `json:"isFamilyFriendly,omitempty"`

	// IsPartOf see : https://schema.org/isPartOf
	// Indicates an item or CreativeWork that this item, or CreativeWork (in some sense), is part of. Inverse property: hasPart (see: https://schema.org/hasPart).
	// types : CreativeWork Trip
	IsPartOf []interface{} `json:"isPartOf,omitempty"`

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
	// Indicates the primary entity described in some page or other CreativeWork. Inverse property: mainEntityOfPage (see: https://schema.org/mainEntityOfPage).
	// types : Thing
	MainEntity []*Thing `json:"mainEntity,omitempty"`

	// MainEntityOfPage see : https://schema.org/mainEntityOfPage
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// Material see : https://schema.org/material
	// A material that something is made from, e.g. leather, wool, cotton, paper.
	// types : Product Text URL
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

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// PartOfEpisode see : https://schema.org/partOfEpisode
	// The episode to which this clip belongs.
	// types : Episode
	PartOfEpisode []*Episode `json:"partOfEpisode,omitempty"`

	// PartOfSeason see : https://schema.org/partOfSeason
	// The season to which this episode belongs.
	// types : CreativeWorkSeason
	PartOfSeason []*CreativeWorkSeason `json:"partOfSeason,omitempty"`

	// PartOfSeries see : https://schema.org/partOfSeries
	// The series to which this episode or season belongs. Supersedes partOfTVSeries (see: https://schema.org/partOfTVSeries).
	// types : CreativeWorkSeries
	PartOfSeries []*CreativeWorkSeries `json:"partOfSeries,omitempty"`

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	// types : Integer Text
	Position []interface{} `json:"position,omitempty"`

	// PotentialAction see : https://schema.org/potentialAction
	// Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	// types : Action
	PotentialAction []*Action `json:"potentialAction,omitempty"`

	// Producer see : https://schema.org/producer
	// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.).
	// types : Organization Person
	Producer []interface{} `json:"producer,omitempty"`

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	// types : Organization Person
	Provider []interface{} `json:"provider,omitempty"`

	// Publication see : https://schema.org/publication
	// A publication event associated with the item.
	// types : PublicationEvent
	Publication []*PublicationEvent `json:"publication,omitempty"`

	// Publisher see : https://schema.org/publisher
	// The publisher of the creative work.
	// types : Organization Person
	Publisher []interface{} `json:"publisher,omitempty"`

	// PublisherImprint see : https://bib.schema.org/publisherImprint
	// The publishing division which published the comic.
	// types : Organization
	PublisherImprint []*Organization `json:"publisherImprint,omitempty"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	// types : CreativeWork URL
	PublishingPrinciples []interface{} `json:"publishingPrinciples,omitempty"`

	// RecordedAt see : https://schema.org/recordedAt
	// The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event. Inverse property: recordedIn (see: https://schema.org/recordedIn).
	// types : Event
	RecordedAt []*Event `json:"recordedAt,omitempty"`

	// ReleasedEvent see : https://schema.org/releasedEvent
	// The place and time the release was issued, expressed as a PublicationEvent.
	// types : PublicationEvent
	ReleasedEvent []*PublicationEvent `json:"releasedEvent,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SchemaVersion see : https://schema.org/schemaVersion
	// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using an URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
	// types : Text URL
	SchemaVersion []string `json:"schemaVersion,omitempty"`

	// SdDatePublished see : https://pending.schema.org/sdDatePublished
	// Indicates the date on which the current structured data was generated / published. Typically used alongside sdPublisher (see: https://schema.org/sdPublisher)
	// types : Date
	SdDatePublished []Date `json:"sdDatePublished,omitempty"`

	// SdLicense see : https://pending.schema.org/sdLicense
	// A license document that applies to this structured data, typically indicated by URL.
	// types : CreativeWork URL
	SdLicense []interface{} `json:"sdLicense,omitempty"`

	// SdPublisher see : https://pending.schema.org/sdPublisher
	// Indicates the party responsible for generating and publishing the current structured data markup, typically in cases where the structured data is derived automatically from existing published content but published on a different site. For example, student projects and open data initiatives often re-publish existing content with more explicitly structured metadata. The
	// sdPublisher (see: https://schema.org/sdPublisher) property helps make such practices more explicit.
	// types : Organization Person
	SdPublisher []interface{} `json:"sdPublisher,omitempty"`

	// SourceOrganization see : https://schema.org/sourceOrganization
	// The Organization on whose behalf the creator was working.
	// types : Organization
	SourceOrganization []*Organization `json:"sourceOrganization,omitempty"`

	// SpatialCoverage see : https://schema.org/spatialCoverage
	// The spatialCoverage of a CreativeWork indicates the place(s) which are the focus of the content. It is a subproperty of
	//       contentLocation intended primarily for more technical and detailed materials. For example with a Dataset, it indicates
	//       areas that the dataset describes: a dataset of New York weather would have spatialCoverage which was the place: the state of New York. Supersedes spatial (see: https://schema.org/spatial).
	// types : Place
	SpatialCoverage []*Place `json:"spatialCoverage,omitempty"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	// types : Organization Person
	Sponsor []interface{} `json:"sponsor,omitempty"`

	// SubjectOf see : https://pending.schema.org/subjectOf
	// A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	// types : CreativeWork Event
	SubjectOf []interface{} `json:"subjectOf,omitempty"`

	// TemporalCoverage see : https://schema.org/temporalCoverage
	// The temporalCoverage of a CreativeWork indicates the period that the content applies to, i.e. that it describes, either as a DateTime or as a textual string indicating a time period in ISO 8601 time interval format (see: https://schema.orghttps://en.wikipedia.org/wiki/ISO_8601#Time_intervals). In
	//       the case of a Dataset it will typically indicate the relevant time period in a precise notation (e.g. for a 2011 census dataset, the year 2011 would be written &quot;2011/2012&quot;). Other forms of content e.g. ScholarlyArticle, Book, TVSeries or TVEpisode may indicate their temporalCoverage in broader terms - textually or via well-known URL.
	//       Written works such as books may sometimes have precise temporal coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601 interval format format via &quot;1939/1945&quot;. Supersedes datasetTimeInterval (see: https://schema.org/datasetTimeInterval), temporal (see: https://schema.org/temporal).
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

	// TranslationOfWork see : https://bib.schema.org/translationOfWork
	// The work that this work has been translated from. e.g. 物种起源 is a translationOf “On the Origin of Species” Inverse property: workTranslation (see: https://schema.orghttps://bib.schema.org/workTranslation).
	// types : CreativeWork
	TranslationOfWork []*CreativeWork `json:"translationOfWork,omitempty"`

	// Translator see : https://schema.org/translator
	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	// types : Organization Person
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
	// Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook. Inverse property: exampleOfWork (see: https://schema.org/exampleOfWork).
	// types : CreativeWork
	WorkExample []*CreativeWork `json:"workExample,omitempty"`

	// WorkTranslation see : https://bib.schema.org/workTranslation
	// A work that is a translation of the content of this work. e.g. 西遊記 has an English workTranslation “Journey to the West”,a German workTranslation “Monkeys Pilgerfahrt” and a Vietnamese  translation Tây du ký bình khảo. Inverse property: translationOfWork (see: https://schema.orghttps://bib.schema.org/translationOfWork).
	// types : CreativeWork
	WorkTranslation []*CreativeWork `json:"workTranslation,omitempty"`
}

func (v Clip) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	into := *intop

	{
		var value interface{} = v.About
		if len(v.About) == 1 {
			value = v.About[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["about"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessMode
		if len(v.AccessMode) == 1 {
			value = v.AccessMode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessMode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessModeSufficient
		if len(v.AccessModeSufficient) == 1 {
			value = v.AccessModeSufficient[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessModeSufficient"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessibilityAPI
		if len(v.AccessibilityAPI) == 1 {
			value = v.AccessibilityAPI[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessibilityAPI"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessibilityControl
		if len(v.AccessibilityControl) == 1 {
			value = v.AccessibilityControl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessibilityControl"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessibilityFeature
		if len(v.AccessibilityFeature) == 1 {
			value = v.AccessibilityFeature[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessibilityFeature"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessibilityHazard
		if len(v.AccessibilityHazard) == 1 {
			value = v.AccessibilityHazard[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessibilityHazard"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccessibilitySummary
		if len(v.AccessibilitySummary) == 1 {
			value = v.AccessibilitySummary[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accessibilitySummary"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AccountablePerson
		if len(v.AccountablePerson) == 1 {
			value = v.AccountablePerson[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["accountablePerson"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Actor
		if len(v.Actor) == 1 {
			value = v.Actor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actor"] = json.RawMessage(b)
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
		var value interface{} = v.AggregateRating
		if len(v.AggregateRating) == 1 {
			value = v.AggregateRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["aggregateRating"] = json.RawMessage(b)
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
		var value interface{} = v.AlternativeHeadline
		if len(v.AlternativeHeadline) == 1 {
			value = v.AlternativeHeadline[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["alternativeHeadline"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.AssociatedMedia
		if len(v.AssociatedMedia) == 1 {
			value = v.AssociatedMedia[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["associatedMedia"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Audience
		if len(v.Audience) == 1 {
			value = v.Audience[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audience"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Audio
		if len(v.Audio) == 1 {
			value = v.Audio[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["audio"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Author
		if len(v.Author) == 1 {
			value = v.Author[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["author"] = json.RawMessage(b)
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
		var value interface{} = v.Character
		if len(v.Character) == 1 {
			value = v.Character[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["character"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Citation
		if len(v.Citation) == 1 {
			value = v.Citation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["citation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ClipNumber
		if len(v.ClipNumber) == 1 {
			value = v.ClipNumber[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["clipNumber"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Comment
		if len(v.Comment) == 1 {
			value = v.Comment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["comment"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CommentCount
		if len(v.CommentCount) == 1 {
			value = v.CommentCount[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["commentCount"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContentLocation
		if len(v.ContentLocation) == 1 {
			value = v.ContentLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contentLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContentRating
		if len(v.ContentRating) == 1 {
			value = v.ContentRating[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contentRating"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContentReferenceTime
		if len(v.ContentReferenceTime) == 1 {
			value = v.ContentReferenceTime[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contentReferenceTime"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Contributor
		if len(v.Contributor) == 1 {
			value = v.Contributor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["contributor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CopyrightHolder
		if len(v.CopyrightHolder) == 1 {
			value = v.CopyrightHolder[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["copyrightHolder"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.CopyrightYear
		if len(v.CopyrightYear) == 1 {
			value = v.CopyrightYear[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["copyrightYear"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Correction
		if len(v.Correction) == 1 {
			value = v.Correction[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["correction"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Creator
		if len(v.Creator) == 1 {
			value = v.Creator[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["creator"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateCreated
		if len(v.DateCreated) == 1 {
			value = v.DateCreated[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateCreated"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DateModified
		if len(v.DateModified) == 1 {
			value = v.DateModified[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["dateModified"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.DatePublished
		if len(v.DatePublished) == 1 {
			value = v.DatePublished[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["datePublished"] = json.RawMessage(b)
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
		var value interface{} = v.Director
		if len(v.Director) == 1 {
			value = v.Director[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["director"] = json.RawMessage(b)
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
		var value interface{} = v.DiscussionUrl
		if len(v.DiscussionUrl) == 1 {
			value = v.DiscussionUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["discussionUrl"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Editor
		if len(v.Editor) == 1 {
			value = v.Editor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["editor"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EducationalAlignment
		if len(v.EducationalAlignment) == 1 {
			value = v.EducationalAlignment[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["educationalAlignment"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EducationalUse
		if len(v.EducationalUse) == 1 {
			value = v.EducationalUse[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["educationalUse"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Encoding
		if len(v.Encoding) == 1 {
			value = v.Encoding[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["encoding"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.EncodingFormat
		if len(v.EncodingFormat) == 1 {
			value = v.EncodingFormat[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["encodingFormat"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExampleOfWork
		if len(v.ExampleOfWork) == 1 {
			value = v.ExampleOfWork[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["exampleOfWork"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Expires
		if len(v.Expires) == 1 {
			value = v.Expires[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["expires"] = json.RawMessage(b)
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
		var value interface{} = v.Genre
		if len(v.Genre) == 1 {
			value = v.Genre[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["genre"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.HasPart
		if len(v.HasPart) == 1 {
			value = v.HasPart[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["hasPart"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Headline
		if len(v.Headline) == 1 {
			value = v.Headline[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["headline"] = json.RawMessage(b)
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
		var value interface{} = v.InLanguage
		if len(v.InLanguage) == 1 {
			value = v.InLanguage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["inLanguage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InteractionStatistic
		if len(v.InteractionStatistic) == 1 {
			value = v.InteractionStatistic[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["interactionStatistic"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.InteractivityType
		if len(v.InteractivityType) == 1 {
			value = v.InteractivityType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["interactivityType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsAccessibleForFree
		if len(v.IsAccessibleForFree) == 1 {
			value = v.IsAccessibleForFree[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isAccessibleForFree"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsBasedOn
		if len(v.IsBasedOn) == 1 {
			value = v.IsBasedOn[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isBasedOn"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsFamilyFriendly
		if len(v.IsFamilyFriendly) == 1 {
			value = v.IsFamilyFriendly[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isFamilyFriendly"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.IsPartOf
		if len(v.IsPartOf) == 1 {
			value = v.IsPartOf[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["isPartOf"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Keywords
		if len(v.Keywords) == 1 {
			value = v.Keywords[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["keywords"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LearningResourceType
		if len(v.LearningResourceType) == 1 {
			value = v.LearningResourceType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["learningResourceType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.License
		if len(v.License) == 1 {
			value = v.License[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["license"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.LocationCreated
		if len(v.LocationCreated) == 1 {
			value = v.LocationCreated[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["locationCreated"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MainEntity
		if len(v.MainEntity) == 1 {
			value = v.MainEntity[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mainEntity"] = json.RawMessage(b)
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
		var value interface{} = v.Material
		if len(v.Material) == 1 {
			value = v.Material[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["material"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Mentions
		if len(v.Mentions) == 1 {
			value = v.Mentions[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["mentions"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicBy
		if len(v.MusicBy) == 1 {
			value = v.MusicBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicBy"] = json.RawMessage(b)
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
		var value interface{} = v.Offers
		if len(v.Offers) == 1 {
			value = v.Offers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["offers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfEpisode
		if len(v.PartOfEpisode) == 1 {
			value = v.PartOfEpisode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfEpisode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfSeason
		if len(v.PartOfSeason) == 1 {
			value = v.PartOfSeason[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfSeason"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PartOfSeries
		if len(v.PartOfSeries) == 1 {
			value = v.PartOfSeries[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["partOfSeries"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Position
		if len(v.Position) == 1 {
			value = v.Position[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["position"] = json.RawMessage(b)
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
		var value interface{} = v.Producer
		if len(v.Producer) == 1 {
			value = v.Producer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["producer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Provider
		if len(v.Provider) == 1 {
			value = v.Provider[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["provider"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Publication
		if len(v.Publication) == 1 {
			value = v.Publication[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publication"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Publisher
		if len(v.Publisher) == 1 {
			value = v.Publisher[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publisher"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PublisherImprint
		if len(v.PublisherImprint) == 1 {
			value = v.PublisherImprint[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["publisherImprint"] = json.RawMessage(b)
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
		var value interface{} = v.RecordedAt
		if len(v.RecordedAt) == 1 {
			value = v.RecordedAt[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["recordedAt"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ReleasedEvent
		if len(v.ReleasedEvent) == 1 {
			value = v.ReleasedEvent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["releasedEvent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Review
		if len(v.Review) == 1 {
			value = v.Review[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["review"] = json.RawMessage(b)
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
		var value interface{} = v.SchemaVersion
		if len(v.SchemaVersion) == 1 {
			value = v.SchemaVersion[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["schemaVersion"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SdDatePublished
		if len(v.SdDatePublished) == 1 {
			value = v.SdDatePublished[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sdDatePublished"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SdLicense
		if len(v.SdLicense) == 1 {
			value = v.SdLicense[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sdLicense"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SdPublisher
		if len(v.SdPublisher) == 1 {
			value = v.SdPublisher[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sdPublisher"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SourceOrganization
		if len(v.SourceOrganization) == 1 {
			value = v.SourceOrganization[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sourceOrganization"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SpatialCoverage
		if len(v.SpatialCoverage) == 1 {
			value = v.SpatialCoverage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["spatialCoverage"] = json.RawMessage(b)
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
		var value interface{} = v.TemporalCoverage
		if len(v.TemporalCoverage) == 1 {
			value = v.TemporalCoverage[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["temporalCoverage"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Text
		if len(v.Text) == 1 {
			value = v.Text[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["text"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ThumbnailUrl
		if len(v.ThumbnailUrl) == 1 {
			value = v.ThumbnailUrl[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["thumbnailUrl"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TimeRequired
		if len(v.TimeRequired) == 1 {
			value = v.TimeRequired[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["timeRequired"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TranslationOfWork
		if len(v.TranslationOfWork) == 1 {
			value = v.TranslationOfWork[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["translationOfWork"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Translator
		if len(v.Translator) == 1 {
			value = v.Translator[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["translator"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.TypicalAgeRange
		if len(v.TypicalAgeRange) == 1 {
			value = v.TypicalAgeRange[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["typicalAgeRange"] = json.RawMessage(b)
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
		var value interface{} = v.Version
		if len(v.Version) == 1 {
			value = v.Version[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["version"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Video
		if len(v.Video) == 1 {
			value = v.Video[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["video"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkExample
		if len(v.WorkExample) == 1 {
			value = v.WorkExample[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workExample"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.WorkTranslation
		if len(v.WorkTranslation) == 1 {
			value = v.WorkTranslation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["workTranslation"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Clip) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Clip"

	return data, nil
}

func (v Clip) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
