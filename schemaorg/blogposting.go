package schemaorg

// BlogPosting see : https://schema.org/BlogPosting
type BlogPosting struct {

SocialMediaPosting// SharedContent see : /sharedContent
// A CreativeWork such as an image, video, or audio clip shared as part of this posting.
SharedContent string `json:"sharedContent"`

// ArticleBody see : /articleBody
// The actual body of the article.
ArticleBody string `json:"articleBody"`

// ArticleSection see : /articleSection
// Articles may belong to one or more 'sections' in a magazine or newspaper, such as Sports, Lifestyle, etc.
ArticleSection string `json:"articleSection"`

// PageEnd see : /pageEnd
// The page on which the work ends; for example "138" or "xvi".
PageEnd interface{} `json:"pageEnd"` // types : Integer Text

// PageStart see : /pageStart
// The page on which the work starts; for example "135" or "xiii".
PageStart interface{} `json:"pageStart"` // types : Integer Text

// Pagination see : /pagination
// Any description of pages that is not separated into pageStart and pageEnd; for example, "1-6, 9, 55" or "10-12, 46-49".
Pagination string `json:"pagination"`

// Speakable see : http://pending.schema.org/speakable
// Indicates sections of a Web page that are particularly 'speakable' in the sense of being highlighted as being especially appropriate for text-to-speech conversion. Other sections of a page may also be usefully spoken in particular circumstances; the 'speakable' property serves to indicate the parts most likely to be generally useful for speech.
// 
// The speakable property can be repeated an arbitrary number of times, with three kinds of possible 'content-locator' values:
// 
// 1.) id-value URL references - uses id-value of an element in the page being annotated. The simplest use of speakable has (potentially relative) URL values, referencing identified sections of the document concerned.
// 
// 2.) CSS Selectors - addresses content in the annotated page, eg. via class attribute. Use the cssSelector (see: https://schema.org/cssSelector) property.
// 
// 3.)  XPaths - addresses content via XPaths (assuming an XML view of the content). Use the xpath (see: https://schema.org/xpath) property.
// 
// For more sophisticated markup of speakable sections beyond simple ID references, either CSS selectors or XPath expressions to pick out document section(s) as speakable. For this
// we define a supporting type, SpeakableSpecification (see: https://schema.org/SpeakableSpecification)  which is defined to be a possible value of the speakable property.
Speakable interface{} `json:"speakable"` // types : SpeakableSpecification URL

// WordCount see : /wordCount
// The number of words in the text of the Article.
WordCount string `json:"wordCount"`

// About see : /about
// The subject matter of the content.
About string `json:"about"`

// AccessMode see : /accessMode
// The human sensory perceptual system or cognitive faculty through which a person may process or perceive information. Expected values include: auditory, tactile, textual, visual, colorDependent, chartOnVisual, chemOnVisual, diagramOnVisual, mathOnVisual, musicOnVisual, textOnVisual.
AccessMode string `json:"accessMode"`

// AccessModeSufficient see : /accessModeSufficient
// A list of single or combined accessModes that are sufficient to understand all the intellectual content of a resource. Expected values include:  auditory, tactile, textual, visual.
AccessModeSufficient string `json:"accessModeSufficient"`

// AccessibilityAPI see : /accessibilityAPI
// Indicates that the resource is compatible with the referenced accessibility API (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
AccessibilityAPI string `json:"accessibilityAPI"`

// AccessibilityControl see : /accessibilityControl
// Identifies input methods that are sufficient to fully control the described resource (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
AccessibilityControl string `json:"accessibilityControl"`

// AccessibilityFeature see : /accessibilityFeature
// Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
AccessibilityFeature string `json:"accessibilityFeature"`

// AccessibilityHazard see : /accessibilityHazard
// A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 (WebSchemas wiki lists possible values (see: https://schema.orghttp://www.w3.org/wiki/WebSchemas/Accessibility)).
AccessibilityHazard string `json:"accessibilityHazard"`

// AccessibilitySummary see : /accessibilitySummary
// A human-readable summary of specific accessibility features or deficiencies, consistent with the other accessibility metadata but expressing subtleties such as "short descriptions are present but long descriptions will be needed for non-visual users" or "short descriptions are present and no long descriptions are needed."
AccessibilitySummary string `json:"accessibilitySummary"`

// AccountablePerson see : /accountablePerson
// Specifies the Person that is legally accountable for the CreativeWork.
AccountablePerson string `json:"accountablePerson"`

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

// Author see : /author
// The author of this content or rating. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
Author interface{} `json:"author"` // types : Organization Person

// Award see : /award
// An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
Award string `json:"award"`

// Character see : /character
// Fictional person connected with a creative work.
Character string `json:"character"`

// Citation see : /citation
// A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
Citation interface{} `json:"citation"` // types : CreativeWork Text

// Comment see : /comment
// Comments, typically from users.
Comment string `json:"comment"`

// CommentCount see : /commentCount
// The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
CommentCount string `json:"commentCount"`

// ContentLocation see : /contentLocation
// The location depicted or described in the content. For example, the location in a photograph or painting.
ContentLocation string `json:"contentLocation"`

// ContentRating see : /contentRating
// Official rating of a piece of content—for example,'MPAA PG-13'.
ContentRating string `json:"contentRating"`

// ContentReferenceTime see : http://pending.schema.org/contentReferenceTime
// The specific time described by a creative work, for works (e.g. articles, video objects etc.) that emphasise a particular moment within an Event.
ContentReferenceTime string `json:"contentReferenceTime"`

// Contributor see : /contributor
// A secondary contributor to the CreativeWork or Event.
Contributor interface{} `json:"contributor"` // types : Organization Person

// CopyrightHolder see : /copyrightHolder
// The party holding the legal copyright to the CreativeWork.
CopyrightHolder interface{} `json:"copyrightHolder"` // types : Organization Person

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

// DiscussionUrl see : /discussionUrl
// A link to the page containing the comments of the CreativeWork.
DiscussionUrl string `json:"discussionUrl"`

// Editor see : /editor
// Specifies the Person who edited the CreativeWork.
Editor string `json:"editor"`

// EducationalAlignment see : /educationalAlignment
// An alignment to an established educational framework.
EducationalAlignment string `json:"educationalAlignment"`

// EducationalUse see : /educationalUse
// The purpose of a work in the context of education; for example, 'assignment', 'group work'.
EducationalUse string `json:"educationalUse"`

// Encoding see : /encoding
// A media object that encodes this CreativeWork. This property is a synonym for associatedMedia. Supersedes encodings (see: https://schema.org/encodings).
Encoding string `json:"encoding"`

// ExampleOfWork see : /exampleOfWork
// A creative work that this work is an example/instance/realization/derivation of. Inverse property: workExample (see: https://schema.org/workExample).
ExampleOfWork string `json:"exampleOfWork"`

// Expires see : /expires
// Date the content expires and is no longer useful or available. For example a VideoObject (see: https://schema.org/VideoObject) or NewsArticle (see: https://schema.org/NewsArticle) whose availability or relevance is time-limited, or a ClaimReview (see: https://schema.org/ClaimReview) fact check whose publisher wants to indicate that it may no longer be relevant (or helpful to highlight) after some date.
Expires string `json:"expires"`

// FileFormat see : /fileFormat
// Media type, typically MIME format (see IANA site (see: https://schema.orghttp://www.iana.org/assignments/media-types/media-types.xhtml)) of the content e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, 'encoding' can be used to indicate each MediaObject alongside particular fileFormat information. Unregistered or niche file formats can be indicated instead via the most appropriate URL, e.g. defining Web page or a Wikipedia entry.
FileFormat interface{} `json:"fileFormat"` // types : Text URL

// Funder see : /funder
// A person or organization that supports (sponsors) something through some kind of financial contribution.
Funder interface{} `json:"funder"` // types : Organization Person

// Genre see : /genre
// Genre of the creative work, broadcast channel or group.
Genre interface{} `json:"genre"` // types : Text URL

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

// InteractivityType see : /interactivityType
// The predominant mode of learning supported by the learning resource. Acceptable values are 'active', 'expositive', or 'mixed'.
InteractivityType string `json:"interactivityType"`

// IsAccessibleForFree see : /isAccessibleForFree
// A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
IsAccessibleForFree string `json:"isAccessibleForFree"`

// IsBasedOn see : /isBasedOn
// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html. Supersedes isBasedOnUrl (see: https://schema.org/isBasedOnUrl).
IsBasedOn interface{} `json:"isBasedOn"` // types : CreativeWork Product URL

// IsFamilyFriendly see : /isFamilyFriendly
// Indicates whether this content is family friendly.
IsFamilyFriendly string `json:"isFamilyFriendly"`

// IsPartOf see : /isPartOf
// Indicates a CreativeWork that this CreativeWork is (in some sense) part of. Inverse property: hasPart (see: https://schema.org/hasPart).
IsPartOf string `json:"isPartOf"`

// Keywords see : /keywords
// Keywords or tags used to describe this content. Multiple entries in a keywords list are typically delimited by commas.
Keywords string `json:"keywords"`

// LearningResourceType see : /learningResourceType
// The predominant type or kind characterizing the learning resource. For example, 'presentation', 'handout'.
LearningResourceType string `json:"learningResourceType"`

// License see : /license
// A license document that applies to this content, typically indicated by URL.
License interface{} `json:"license"` // types : CreativeWork URL

// LocationCreated see : /locationCreated
// The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork.
LocationCreated string `json:"locationCreated"`

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

// Position see : /position
// The position of an item in a series or sequence of items.
Position interface{} `json:"position"` // types : Integer Text

// Producer see : /producer
// The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.).
Producer interface{} `json:"producer"` // types : Organization Person

// Provider see : /provider
// The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller. Supersedes carrier (see: https://schema.org/carrier).
Provider interface{} `json:"provider"` // types : Organization Person

// Publication see : /publication
// A publication event associated with the item.
Publication string `json:"publication"`

// Publisher see : /publisher
// The publisher of the creative work.
Publisher interface{} `json:"publisher"` // types : Organization Person

// PublisherImprint see : http://bib.schema.org/publisherImprint
// The publishing division which published the comic.
PublisherImprint string `json:"publisherImprint"`

// PublishingPrinciples see : /publishingPrinciples
// The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
// 
// While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
PublishingPrinciples interface{} `json:"publishingPrinciples"` // types : CreativeWork URL

// RecordedAt see : /recordedAt
// The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event. Inverse property: recordedIn (see: https://schema.org/recordedIn).
RecordedAt string `json:"recordedAt"`

// ReleasedEvent see : /releasedEvent
// The place and time the release was issued, expressed as a PublicationEvent.
ReleasedEvent string `json:"releasedEvent"`

// Review see : /review
// A review of the item. Supersedes reviews (see: https://schema.org/reviews).
Review string `json:"review"`

// SchemaVersion see : /schemaVersion
// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using an URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
SchemaVersion interface{} `json:"schemaVersion"` // types : Text URL

// SourceOrganization see : /sourceOrganization
// The Organization on whose behalf the creator was working.
SourceOrganization string `json:"sourceOrganization"`

// SpatialCoverage see : /spatialCoverage
// The spatialCoverage of a CreativeWork indicates the place(s) which are the focus of the content. It is a subproperty of
//       contentLocation intended primarily for more technical and detailed materials. For example with a Dataset, it indicates
//       areas that the dataset describes: a dataset of New York weather would have spatialCoverage which was the place: the state of New York. Supersedes spatial (see: https://schema.org/spatial).
SpatialCoverage string `json:"spatialCoverage"`

// Sponsor see : /sponsor
// A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
Sponsor interface{} `json:"sponsor"` // types : Organization Person

// TemporalCoverage see : /temporalCoverage
// The temporalCoverage of a CreativeWork indicates the period that the content applies to, i.e. that it describes, either as a DateTime or as a textual string indicating a time period in ISO 8601 time interval format (see: https://schema.orghttps://en.wikipedia.org/wiki/ISO_8601#Time_intervals). In
//       the case of a Dataset it will typically indicate the relevant time period in a precise notation (e.g. for a 2011 census dataset, the year 2011 would be written "2011/2012"). Other forms of content e.g. ScholarlyArticle, Book, TVSeries or TVEpisode may indicate their temporalCoverage in broader terms - textually or via well-known URL.
//       Written works such as books may sometimes have precise temporal coverage too, e.g. a work set in 1939 - 1945 can be indicated in ISO 8601 interval format format via "1939/1945". Supersedes datasetTimeInterval (see: https://schema.org/datasetTimeInterval), temporal (see: https://schema.org/temporal).
TemporalCoverage interface{} `json:"temporalCoverage"` // types : DateTime Text URL

// Text see : /text
// The textual content of this CreativeWork.
Text string `json:"text"`

// ThumbnailUrl see : /thumbnailUrl
// A thumbnail image relevant to the Thing.
ThumbnailUrl string `json:"thumbnailUrl"`

// TimeRequired see : /timeRequired
// Approximate or typical time it takes to work with or through this learning resource for the typical intended target audience, e.g. 'P30M', 'P1H25M'.
TimeRequired string `json:"timeRequired"`

// TranslationOfWork see : http://bib.schema.org/translationOfWork
// The work that this work has been translated from. e.g. 物种起源 is a translationOf “On the Origin of Species” Inverse property: workTranslation (see: https://schema.orghttp://bib.schema.org/workTranslation).
TranslationOfWork string `json:"translationOfWork"`

// Translator see : /translator
// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
Translator interface{} `json:"translator"` // types : Organization Person

// TypicalAgeRange see : /typicalAgeRange
// The typical expected age range, e.g. '7-9', '11-'.
TypicalAgeRange string `json:"typicalAgeRange"`

// Version see : /version
// The version of the CreativeWork embodied by a specified resource.
Version interface{} `json:"version"` // types : Number Text

// Video see : /video
// An embedded video object.
Video string `json:"video"`

// WorkExample see : /workExample
// Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook. Inverse property: exampleOfWork (see: https://schema.org/exampleOfWork).
WorkExample string `json:"workExample"`

// WorkTranslation see : http://bib.schema.org/workTranslation
// A work that is a translation of the content of this work. e.g. 西遊記 has an English workTranslation “Journey to the West”,a German workTranslation “Monkeys Pilgerfahrt” and a Vietnamese  translation Tây du ký bình khảo. Inverse property: translationOfWork (see: https://schema.orghttp://bib.schema.org/translationOfWork).
WorkTranslation string `json:"workTranslation"`

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

// BlogPost see : /blogPost
// A posting that is part of this blog.  Supersedes blogPosts (see: https://schema.org/blogPosts).
BlogPost string `json:"blogPost"`

// LiveBlogUpdate see : /liveBlogUpdate
// An update to the LiveBlog. 
LiveBlogUpdate string `json:"liveBlogUpdate"`

}

