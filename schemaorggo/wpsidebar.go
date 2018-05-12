package schemaorggo

import "encoding/json"

// WPSideBar see : https://schema.org/WPSideBar
type WPSideBar struct {
	WebPageElement

	typeContext

	// About see : https://schema.org/about
	// The subject matter of the content.
	About *Thing `json:"about"`

	// AccessMode see : https://schema.org/accessMode
	// The human sensory perceptual system or cognitive faculty through which a person may process or perceive information. Expected values include: auditory, tactile, textual, visual, colorDependent, chartOnVisual, chemOnVisual, diagramOnVisual, mathOnVisual, musicOnVisual, textOnVisual.
	AccessMode string `json:"accessMode"`

	// AccessModeSufficient see : https://schema.org/accessModeSufficient
	// A list of single or combined accessModes that are sufficient to understand all the intellectual content of a resource. Expected values include:  auditory, tactile, textual, visual.
	AccessModeSufficient string `json:"accessModeSufficient"`

	// AccessibilityAPI see : https://schema.org/accessibilityAPI
	// Indicates that the resource is compatible with the referenced accessibility API (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	AccessibilityAPI string `json:"accessibilityAPI"`

	// AccessibilityControl see : https://schema.org/accessibilityControl
	// Identifies input methods that are sufficient to fully control the described resource (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	AccessibilityControl string `json:"accessibilityControl"`

	// AccessibilityFeature see : https://schema.org/accessibilityFeature
	// Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	AccessibilityFeature string `json:"accessibilityFeature"`

	// AccessibilityHazard see : https://schema.org/accessibilityHazard
	// A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
	AccessibilityHazard string `json:"accessibilityHazard"`

	// AccessibilitySummary see : https://schema.org/accessibilitySummary
	// A human-readable summary of specific accessibility features or deficiencies, consistent with the other accessibility metadata but expressing subtleties such as "short descriptions are present but long descriptions will be needed for non-visual users" or "short descriptions are present and no long descriptions are needed."
	AccessibilitySummary string `json:"accessibilitySummary"`

	// AccountablePerson see : https://schema.org/accountablePerson
	// Specifies the Person that is legally accountable for the CreativeWork.
	AccountablePerson *Person `json:"accountablePerson"`

	// AggregateRating see : https://schema.org/aggregateRating
	// The overall rating, based on a collection of reviews or ratings, of the item.
	AggregateRating *AggregateRating `json:"aggregateRating"`

	// AlternativeHeadline see : https://schema.org/alternativeHeadline
	// A secondary title of the CreativeWork.
	AlternativeHeadline string `json:"alternativeHeadline"`

	// AssociatedMedia see : https://schema.org/associatedMedia
	// A media object that encodes this CreativeWork. This property is a synonym for encoding.
	AssociatedMedia *MediaObject `json:"associatedMedia"`

	// Audience see : https://schema.org/audience
	// An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	Audience *Audience `json:"audience"`

	// Audio see : https://schema.org/audio
	// An embedded audio object.
	Audio *AudioObject `json:"audio"`

	// Author see : https://schema.org/author
	// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	Author interface{} `json:"author"` // types : Organization Person

	// Award see : https://schema.org/award
	// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	Award string `json:"award"`

	// Character see : https://schema.org/character
	// Fictional person connected with a creative work.
	Character *Person `json:"character"`

	// Citation see : https://schema.org/citation
	// A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
	Citation interface{} `json:"citation"` // types : CreativeWork Text

	// Comment see : https://schema.org/comment
	// Comments, typically from users.
	Comment *Comment `json:"comment"`

	// CommentCount see : https://schema.org/commentCount
	// The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
	CommentCount int `json:"commentCount"`

	// ContentLocation see : https://schema.org/contentLocation
	// The location depicted or described in the content. For example, the location in a photograph or painting.
	ContentLocation *Place `json:"contentLocation"`

	// ContentRating see : https://schema.org/contentRating
	// Official rating of a piece of content—for example,'MPAA PG-13'.
	ContentRating string `json:"contentRating"`

	// ContentReferenceTime see : http://pending.schema.org/contentReferenceTime
	// The specific time described by a creative work, for works (e.g. articles, video objects etc.) that emphasise a particular moment within an Event.
	ContentReferenceTime interface{} `json:"contentReferenceTime"`

	// Contributor see : https://schema.org/contributor
	// A secondary contributor to the CreativeWork or Event.
	Contributor interface{} `json:"contributor"` // types : Organization Person

	// CopyrightHolder see : https://schema.org/copyrightHolder
	// The party holding the legal copyright to the CreativeWork.
	CopyrightHolder interface{} `json:"copyrightHolder"` // types : Organization Person

	// CopyrightYear see : https://schema.org/copyrightYear
	// The year during which the claimed copyright for the CreativeWork was first asserted.
	CopyrightYear float64 `json:"copyrightYear"`

	// Creator see : https://schema.org/creator
	// The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork.
	Creator interface{} `json:"creator"` // types : Organization Person

	// DateCreated see : https://schema.org/dateCreated
	// The date on which the CreativeWork was created or the item was added to a DataFeed.
	DateCreated interface{} `json:"dateCreated"` // types : Date DateTime

	// DateModified see : https://schema.org/dateModified
	// The date on which the CreativeWork was most recently modified or when the item's entry was modified within a DataFeed.
	DateModified interface{} `json:"dateModified"` // types : Date DateTime

	// DatePublished see : https://schema.org/datePublished
	// Date of first broadcast/publication.
	DatePublished interface{} `json:"datePublished"`

	// DiscussionUrl see : https://schema.org/discussionUrl
	// A link to the page containing the comments of the CreativeWork.
	DiscussionUrl string `json:"discussionUrl"`

	// Editor see : https://schema.org/editor
	// Specifies the Person who edited the CreativeWork.
	Editor *Person `json:"editor"`

	// EducationalAlignment see : https://schema.org/educationalAlignment
	// An alignment to an established educational framework.
	EducationalAlignment *AlignmentObject `json:"educationalAlignment"`

	// EducationalUse see : https://schema.org/educationalUse
	// The purpose of a work in the context of education; for example, 'assignment', 'group work'.
	EducationalUse string `json:"educationalUse"`

	// Encoding see : https://schema.org/encoding
	// A media object that encodes this CreativeWork. This property is a synonym for associatedMedia. Supersedes encodings (see: https://schema.org/encodings).
	Encoding *MediaObject `json:"encoding"`

	// ExampleOfWork see : https://schema.org/exampleOfWork
	// A creative work that this work is an example/instance/realization/derivation of. Inverse property: workExample (see: https://schema.org/workExample).
	ExampleOfWork *CreativeWork `json:"exampleOfWork"`

	// Expires see : https://schema.org/expires
	// Date the content expires and is no longer useful or available. For example a VideoObject (see: https://schema.org/VideoObject) or NewsArticle (see: https://schema.org/NewsArticle) whose availability or relevance is time-limited, or a ClaimReview (see: https://schema.org/ClaimReview) fact check whose publisher wants to indicate that it may no longer be relevant (or helpful to highlight) after some date.
	Expires interface{} `json:"expires"`

	// FileFormat see : https://schema.org/fileFormat
	// Media type, typically MIME format (see IANA site (see: https://schema.orghttp://www.iana.org/assignments/media-types/media-types.xhtml)) of the content e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, 'encoding' can be used to indicate each MediaObject alongside particular fileFormat information. Unregistered or niche file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia entry.
	FileFormat interface{} `json:"fileFormat"` // types : Text URL

	// Funder see : https://schema.org/funder
	// A person or organization that supports (sponsors) something through some kind of financial contribution.
	Funder interface{} `json:"funder"` // types : Organization Person

	// Genre see : https://schema.org/genre
	// Genre of the creative work, broadcast channel or group.
	Genre interface{} `json:"genre"` // types : Text URL

	// HasPart see : https://schema.org/hasPart
	// Indicates a CreativeWork that is (in some sense) a part of this CreativeWork. Inverse property: isPartOf (see: https://schema.org/isPartOf).
	HasPart *CreativeWork `json:"hasPart"`

	// Headline see : https://schema.org/headline
	// Headline of the article.
	Headline string `json:"headline"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	InLanguage interface{} `json:"inLanguage"` // types : Language Text

	// InteractionStatistic see : https://schema.org/interactionStatistic
	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used. Supersedes interactionCount (see: https://schema.org/interactionCount).
	InteractionStatistic *InteractionCounter `json:"interactionStatistic"`

	// InteractivityType see : https://schema.org/interactivityType
	// The predominant mode of learning supported by the learning resource. Acceptable values are 'active', 'expositive', or 'mixed'.
	InteractivityType string `json:"interactivityType"`

	// IsAccessibleForFree see : https://schema.org/isAccessibleForFree
	// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	IsAccessibleForFree bool `json:"isAccessibleForFree"`

	// IsBasedOn see : https://schema.org/isBasedOn
	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html. Supersedes isBasedOnUrl (see: https://schema.org/isBasedOnUrl).
	IsBasedOn interface{} `json:"isBasedOn"` // types : CreativeWork Product URL

	// IsFamilyFriendly see : https://schema.org/isFamilyFriendly
	// Indicates whether this content is family friendly.
	IsFamilyFriendly bool `json:"isFamilyFriendly"`

	// IsPartOf see : https://schema.org/isPartOf
	// Indicates a CreativeWork that this CreativeWork is (in some sense) part of. Inverse property: hasPart (see: https://schema.org/hasPart).
	IsPartOf *CreativeWork `json:"isPartOf"`

	// Keywords see : https://schema.org/keywords
	// Keywords or tags used to describe this content. Multiple entries in a keywords list are typically delimited by commas.
	Keywords string `json:"keywords"`

	// LearningResourceType see : https://schema.org/learningResourceType
	// The predominant type or kind characterizing the learning resource. For example, 'presentation', 'handout'.
	LearningResourceType string `json:"learningResourceType"`

	// License see : https://schema.org/license
	// A license document that applies to this content, typically indicated by URL.
	License interface{} `json:"license"` // types : CreativeWork URL

	// LocationCreated see : https://schema.org/locationCreated
	// The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork.
	LocationCreated *Place `json:"locationCreated"`

	// MainEntity see : https://schema.org/mainEntity
	// Indicates the primary entity described in some page or other CreativeWork. Inverse property: mainEntityOfPage (see: https://schema.org/mainEntityOfPage).
	MainEntity *Thing `json:"mainEntity"`

	// Material see : https://schema.org/material
	// A material that something is made from, e.g. leather, wool, cotton, paper.
	Material interface{} `json:"material"` // types : Product Text URL

	// Mentions see : https://schema.org/mentions
	// Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept.
	Mentions *Thing `json:"mentions"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	Offers *Offer `json:"offers"`

	// Position see : https://schema.org/position
	// The position of an item in a series or sequence of items.
	Position interface{} `json:"position"` // types : Integer Text

	// Producer see : https://schema.org/producer
	// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.).
	Producer interface{} `json:"producer"` // types : Organization Person

	// Provider see : https://schema.org/provider
	// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
	Provider interface{} `json:"provider"` // types : Organization Person

	// Publication see : https://schema.org/publication
	// A publication event associated with the item.
	Publication *PublicationEvent `json:"publication"`

	// Publisher see : https://schema.org/publisher
	// The publisher of the creative work.
	Publisher interface{} `json:"publisher"` // types : Organization Person

	// PublisherImprint see : http://bib.schema.org/publisherImprint
	// The publishing division which published the comic.
	PublisherImprint *Organization `json:"publisherImprint"`

	// PublishingPrinciples see : https://schema.org/publishingPrinciples
	// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	//
	// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

	// RecordedAt see : https://schema.org/recordedAt
	// The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event. Inverse property: recordedIn (see: https://schema.org/recordedIn).
	RecordedAt *Event `json:"recordedAt"`

	// ReleasedEvent see : https://schema.org/releasedEvent
	// The place and time the release was issued, expressed as a PublicationEvent.
	ReleasedEvent *PublicationEvent `json:"releasedEvent"`

	// Review see : https://schema.org/review
	// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	Review *Review `json:"review"`

	// SchemaVersion see : https://schema.org/schemaVersion
	// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using an URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
	SchemaVersion interface{} `json:"schemaVersion"` // types : Text URL

	// SourceOrganization see : https://schema.org/sourceOrganization
	// The Organization on whose behalf the creator was working.
	SourceOrganization *Organization `json:"sourceOrganization"`

	// SpatialCoverage see : https://schema.org/spatialCoverage
	// The spatialCoverage of a CreativeWork indicates the place(s) which are the focus of the content. It is a subproperty of
	//       contentLocation intended primarily for more technical and detailed materials. For example with a Dataset, it indicates
	//       areas that the dataset describes: a dataset of New York weather would have spatialCoverage which was the place: the state of New York. Supersedes spatial (see: https://schema.org/spatial).
	SpatialCoverage *Place `json:"spatialCoverage"`

	// Sponsor see : https://schema.org/sponsor
	// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	Sponsor interface{} `json:"sponsor"` // types : Organization Person

	// TemporalCoverage see : https://schema.org/temporalCoverage
	// The temporalCoverage of a CreativeWork indicates the period that the content applies to, i.e. that it describes, either as a DateTime or as a textual string indicating a time period in ISO 8601 time interval format (see: https://schema.orghttps://en.wikipedia.org/wiki/ISO_8601#Time_intervals). In
	//       the case of a Dataset it will typically indicate the relevant time period in a precise notation (e.g. for a 2011 census dataset, the year 2011 would be written "2011/2012"). Other forms of content e.g. ScholarlyArticle, Book, TVSeries or TVEpisode may indicate their temporalCoverage in broader terms - textually or via well-known URL.
	//       Written works such as books may sometimes have precise temporal coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601 interval format format via "1939/1945". Supersedes datasetTimeInterval (see: https://schema.org/datasetTimeInterval), temporal (see: https://schema.org/temporal).
	TemporalCoverage interface{} `json:"temporalCoverage"` // types : DateTime Text URL

	// Text see : https://schema.org/text
	// The textual content of this CreativeWork.
	Text string `json:"text"`

	// ThumbnailUrl see : https://schema.org/thumbnailUrl
	// A thumbnail image relevant to the Thing.
	ThumbnailUrl string `json:"thumbnailUrl"`

	// TimeRequired see : https://schema.org/timeRequired
	// Approximate or typical time it takes to work with or through this learning resource for the typical intended target audience, e.g. 'P30M', 'P1H25M'.
	TimeRequired *Duration `json:"timeRequired"`

	// TranslationOfWork see : http://bib.schema.org/translationOfWork
	// The work that this work has been translated from. e.g. 物种起源 is a translationOf “On the Origin of Species” Inverse property: workTranslation (see: https://schema.orghttp://bib.schema.org/workTranslation).
	TranslationOfWork *CreativeWork `json:"translationOfWork"`

	// Translator see : https://schema.org/translator
	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	Translator interface{} `json:"translator"` // types : Organization Person

	// TypicalAgeRange see : https://schema.org/typicalAgeRange
	// The typical expected age range, e.g. '7-9', '11-'.
	TypicalAgeRange string `json:"typicalAgeRange"`

	// Version see : https://schema.org/version
	// The version of the CreativeWork embodied by a specified resource.
	Version interface{} `json:"version"` // types : Number Text

	// Video see : https://schema.org/video
	// An embedded video object.
	Video *VideoObject `json:"video"`

	// WorkExample see : https://schema.org/workExample
	// Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook. Inverse property: exampleOfWork (see: https://schema.org/exampleOfWork).
	WorkExample *CreativeWork `json:"workExample"`

	// WorkTranslation see : http://bib.schema.org/workTranslation
	// A work that is a translation of the content of this work. e.g. 西遊記 has an English workTranslation “Journey to the West”,a German workTranslation “Monkeys Pilgerfahrt” and a Vietnamese  translation Tây du ký bình khảo. Inverse property: translationOfWork (see: https://schema.orghttp://bib.schema.org/translationOfWork).
	WorkTranslation *CreativeWork `json:"workTranslation"`
}

func (v *WPSideBar) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "WPSideBar"

	return json.Marshal(v)
}
