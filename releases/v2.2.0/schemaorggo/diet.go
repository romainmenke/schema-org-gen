package schemaorggo

import "encoding/json"

// Diet see : https://schema.org/Diet
type Diet struct {

	// With properties from CreativeWork see : https://schema.org/CreativeWork
	//

	// With properties from LifestyleModification see : https://schema.org/LifestyleModification
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from MedicalTherapy see : https://schema.org/MedicalTherapy
	//

	// With properties from MedicalEntity see : https://schema.org/MedicalEntity
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from Thing see : https://schema.org/Thing
	//

	// With properties from MedicalEntity see : https://schema.org/MedicalEntity
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

	// AccessibilityAPI see : https://schema.org/accessibilityAPI
	// Indicates that the resource is compatible with the referenced accessibility API (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	// types : Text
	AccessibilityAPI []string `json:"accessibilityAPI,omitempty"`

	// AccessibilityControl see : https://schema.org/accessibilityControl
	// Identifies input methods that are sufficient to fully control the described resource (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	// types : Text
	AccessibilityControl []string `json:"accessibilityControl,omitempty"`

	// AccessibilityFeature see : https://schema.org/accessibilityFeature
	// Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	// types : Text
	AccessibilityFeature []string `json:"accessibilityFeature,omitempty"`

	// AccessibilityHazard see : https://schema.org/accessibilityHazard
	// A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	// types : Text
	AccessibilityHazard []string `json:"accessibilityHazard,omitempty"`

	// AccountablePerson see : https://schema.org/accountablePerson
	// Specifies the Person that is legally accountable for the CreativeWork.
	// types : Person
	AccountablePerson []*Person `json:"accountablePerson,omitempty"`

	// AdditionalType see : https://schema.org/additionalType
	// An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	// types : URL
	AdditionalType []string `json:"additionalType,omitempty"`

	// AdverseOutcome see : https://schema.org/adverseOutcome
	// A possible complication and/or side effect of this therapy. If it is known that an adverse outcome is serious (resulting in death, disability, or permanent damage; requiring hospitalization; or is otherwise life-threatening or requires immediate medical attention), tag it as a seriouseAdverseOutcome instead.
	// types : MedicalEntity
	AdverseOutcome []*MedicalEntity `json:"adverseOutcome,omitempty"`

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
	// The author of this content. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
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

	// Code see : https://schema.org/code
	// A medical code for the entity, taken from a controlled vocabulary or ontology such as ICD-9, DiseasesDB, MeSH, SNOMED-CT, RxNorm, etc.
	// types : MedicalCode
	Code []*MedicalCode `json:"code,omitempty"`

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
	// Official rating of a piece of content&amp;#x2014;for example,&#39;MPAA PG-13&#39;.
	// types : Text
	ContentRating []string `json:"contentRating,omitempty"`

	// Contraindication see : https://schema.org/contraindication
	// A contraindication for this therapy.
	// types : MedicalContraindication
	Contraindication []*MedicalContraindication `json:"contraindication,omitempty"`

	// Contributor see : https://schema.org/contributor
	// A secondary contributor to the CreativeWork.
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
	// A short description of the item.
	// types : Text
	Description []string `json:"description,omitempty"`

	// DietFeatures see : https://schema.org/dietFeatures
	// Nutritional information specific to the dietary plan. May include dietary recommendations on what foods to avoid, what foods to consume, and specific alterations/deviations from the USDA or other regulatory body&#39;s approved dietary guidelines.
	// types : Text
	DietFeatures []string `json:"dietFeatures,omitempty"`

	// DiscussionUrl see : https://schema.org/discussionUrl
	// A link to the page containing the comments of the CreativeWork.
	// types : URL
	DiscussionUrl []string `json:"discussionUrl,omitempty"`

	// DuplicateTherapy see : https://schema.org/duplicateTherapy
	// A therapy that duplicates or overlaps this one.
	// types : MedicalTherapy
	DuplicateTherapy []*MedicalTherapy `json:"duplicateTherapy,omitempty"`

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

	// Encodings see : https://schema.org/encodings
	// A media object that encodes this CreativeWork.
	// types : MediaObject
	Encodings []*MediaObject `json:"encodings,omitempty"`

	// Endorsers see : https://schema.org/endorsers
	// People or organizations that endorse the plan.
	// types : Organization Person
	Endorsers []interface{} `json:"endorsers,omitempty"`

	// ExampleOfWork see : https://schema.org/exampleOfWork
	// A creative work that this work is an example/instance/realization/derivation of.
	// types : CreativeWork
	ExampleOfWork []*CreativeWork `json:"exampleOfWork,omitempty"`

	// ExpertConsiderations see : https://schema.org/expertConsiderations
	// Medical expert advice related to the plan.
	// types : Text
	ExpertConsiderations []string `json:"expertConsiderations,omitempty"`

	// FileFormat see : https://schema.org/fileFormat
	// Media type (aka MIME format, see &lt;a href=&quot;http://www.iana.org/assignments/media-types/media-types.xhtml&quot;&gt;IANA site&lt;/a&gt;) of the content e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, &#39;encoding&#39; can be used to indicate each MediaObject alongside particular fileFormat information.
	// types : Text
	FileFormat []string `json:"fileFormat,omitempty"`

	// Genre see : https://schema.org/genre
	// Genre of the creative work or group.
	// types : Text URL
	Genre []string `json:"genre,omitempty"`

	// Guideline see : https://schema.org/guideline
	// A medical guideline related to this entity.
	// types : MedicalGuideline
	Guideline []*MedicalGuideline `json:"guideline,omitempty"`

	// HasPart see : https://schema.org/hasPart
	// Indicates a CreativeWork that is (in some sense) a part of this CreativeWork.
	// types : CreativeWork
	HasPart []*CreativeWork `json:"hasPart,omitempty"`

	// Headline see : https://schema.org/headline
	// Headline of the article.
	// types : Text
	Headline []string `json:"headline,omitempty"`

	// Image see : https://schema.org/image
	// An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	// types : URL ImageObject
	Image []interface{} `json:"image,omitempty"`

	// InLanguage see : https://schema.org/inLanguage
	// The language of the content or performance or used in an action. Please use one of the language codes from the &lt;a href=&#39;http://tools.ietf.org/html/bcp47&#39;&gt;IETF BCP 47 standard&lt;/a&gt;.
	// types : Text Language
	InLanguage []interface{} `json:"inLanguage,omitempty"`

	// Indication see : https://schema.org/indication
	// A factor that indicates use of this therapy for treatment and/or prevention of a condition, symptom, etc. For therapies such as drugs, indications can include both officially-approved indications as well as off-label uses. These can be distinguished by using the ApprovedIndication subtype of MedicalIndication.
	// types : MedicalIndication
	Indication []*MedicalIndication `json:"indication,omitempty"`

	// InteractionStatistic see : https://schema.org/interactionStatistic
	// The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used.
	// types : InteractionCounter
	InteractionStatistic []*InteractionCounter `json:"interactionStatistic,omitempty"`

	// InteractivityType see : https://schema.org/interactivityType
	// The predominant mode of learning supported by the learning resource. Acceptable values are &#39;active&#39;, &#39;expositive&#39;, or &#39;mixed&#39;.
	// types : Text
	InteractivityType []string `json:"interactivityType,omitempty"`

	// IsBasedOnUrl see : https://schema.org/isBasedOnUrl
	// A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
	// types : URL
	IsBasedOnUrl []string `json:"isBasedOnUrl,omitempty"`

	// IsFamilyFriendly see : https://schema.org/isFamilyFriendly
	// Indicates whether this content is family friendly.
	// types : Boolean
	IsFamilyFriendly []bool `json:"isFamilyFriendly,omitempty"`

	// IsPartOf see : https://schema.org/isPartOf
	// Indicates a CreativeWork that this CreativeWork is (in some sense) part of.
	// types : CreativeWork
	IsPartOf []*CreativeWork `json:"isPartOf,omitempty"`

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
	// Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	//       &lt;br /&gt;&lt;br /&gt;
	//       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	//
	// types : CreativeWork URL
	MainEntityOfPage []interface{} `json:"mainEntityOfPage,omitempty"`

	// MedicineSystem see : https://schema.org/medicineSystem
	// The system of medicine that includes this MedicalEntity, for example &#39;evidence-based&#39;, &#39;homeopathic&#39;, &#39;chiropractic&#39;, etc.
	// types : MedicineSystem
	MedicineSystem []*MedicineSystem `json:"medicineSystem,omitempty"`

	// Mentions see : https://schema.org/mentions
	// Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept.
	// types : Thing
	Mentions []*Thing `json:"mentions,omitempty"`

	// Name see : https://schema.org/name
	// The name of the item.
	// types : Text
	Name []string `json:"name,omitempty"`

	// Offers see : https://schema.org/offers
	// An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	// types : Offer
	Offers []*Offer `json:"offers,omitempty"`

	// Overview see : https://schema.org/overview
	// Descriptive information establishing the overarching theory/philosophy of the plan. May include the rationale for the name, the population where the plan first came to prominence, etc.
	// types : Text
	Overview []string `json:"overview,omitempty"`

	// PhysiologicalBenefits see : https://schema.org/physiologicalBenefits
	// Specific physiologic benefits associated to the plan.
	// types : Text
	PhysiologicalBenefits []string `json:"physiologicalBenefits,omitempty"`

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

	// ProprietaryName see : https://schema.org/proprietaryName
	// Proprietary name given to the diet plan, typically by its originator or creator.
	// types : Text
	ProprietaryName []string `json:"proprietaryName,omitempty"`

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
	// Link to page describing the editorial principles of the organization primarily responsible for the creation of the CreativeWork.
	// types : URL
	PublishingPrinciples []string `json:"publishingPrinciples,omitempty"`

	// RecognizingAuthority see : https://schema.org/recognizingAuthority
	// If applicable, the organization that officially recognizes this entity as part of its endorsed system of medicine.
	// types : Organization
	RecognizingAuthority []*Organization `json:"recognizingAuthority,omitempty"`

	// RecordedAt see : https://schema.org/recordedAt
	// The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event.
	// types : Event
	RecordedAt []*Event `json:"recordedAt,omitempty"`

	// ReleasedEvent see : https://schema.org/releasedEvent
	// The place and time the release was issued, expressed as a PublicationEvent.
	// types : PublicationEvent
	ReleasedEvent []*PublicationEvent `json:"releasedEvent,omitempty"`

	// RelevantSpecialty see : https://schema.org/relevantSpecialty
	// If applicable, a medical specialty in which this entity is relevant.
	// types : MedicalSpecialty
	RelevantSpecialty []*MedicalSpecialty `json:"relevantSpecialty,omitempty"`

	// Review see : https://schema.org/review
	// A review of the item.
	// types : Review
	Review []*Review `json:"review,omitempty"`

	// Reviews see : https://schema.org/reviews
	// Review of the item.
	// types : Review
	Reviews []*Review `json:"reviews,omitempty"`

	// Risks see : https://schema.org/risks
	// Specific physiologic risks associated to the plan.
	// types : Text
	Risks []string `json:"risks,omitempty"`

	// SameAs see : https://schema.org/sameAs
	// URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	// types : URL
	SameAs []string `json:"sameAs,omitempty"`

	// SchemaVersion see : https://schema.org/schemaVersion
	// Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using a URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
	// types : URL Text
	SchemaVersion []string `json:"schemaVersion,omitempty"`

	// SeriousAdverseOutcome see : https://schema.org/seriousAdverseOutcome
	// A possible serious complication and/or serious side effect of this therapy. Serious adverse outcomes include those that are life-threatening; result in death, disability, or permanent damage; require hospitalization or prolong existing hospitalization; cause congenital anomalies or birth defects; or jeopardize the patient and may require medical or surgical intervention to prevent one of the outcomes in this definition.
	// types : MedicalEntity
	SeriousAdverseOutcome []*MedicalEntity `json:"seriousAdverseOutcome,omitempty"`

	// SourceOrganization see : https://schema.org/sourceOrganization
	// The Organization on whose behalf the creator was working.
	// types : Organization
	SourceOrganization []*Organization `json:"sourceOrganization,omitempty"`

	// Study see : https://schema.org/study
	// A medical study or trial related to this entity.
	// types : MedicalStudy
	Study []*MedicalStudy `json:"study,omitempty"`

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

	// Translator see : https://schema.org/translator
	// Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market.
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
	// types : Number
	Version []float64 `json:"version,omitempty"`

	// Video see : https://schema.org/video
	// An embedded video object.
	// types : VideoObject
	Video []*VideoObject `json:"video,omitempty"`

	// WorkExample see : https://schema.org/workExample
	// Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook.
	// types : CreativeWork
	WorkExample []*CreativeWork `json:"workExample,omitempty"`
}

func (v Diet) MarshalJSON() ([]byte, error) {
	type Alias Diet

	b, err := json.Marshal((Alias)(v))
	if err != nil {
		return nil, err
	}

	return append([]byte("{\"@context\":\"http://schema.org\",\"@type\":\"Diet\","), b[1:]...), nil
}