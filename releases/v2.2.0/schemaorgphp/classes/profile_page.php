<?php
namespace SchemaOrg;

// ProfilePage see : https://schema.org/ProfilePage
class ProfilePage implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'ProfilePage';

	/**
	 * With properties from WebPage see : https://schema.org/WebPage
	 */

	/**
	 * With properties from CreativeWork see : https://schema.org/CreativeWork
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * The subject matter of the content.
	 * see : https://schema.org/about
	 *
	 * @var \Thing | \Thing[]
	 */
	public $about;

	/**
	 * Indicates that the resource is compatible with the referenced accessibility API (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	 * see : https://schema.org/accessibilityAPI
	 *
	 * @var string | string[]
	 */
	public $accessibility_api;

	/**
	 * Identifies input methods that are sufficient to fully control the described resource (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	 * see : https://schema.org/accessibilityControl
	 *
	 * @var string | string[]
	 */
	public $accessibility_control;

	/**
	 * Content features of the resource, such as accessible media, alternatives and supported enhancements for accessibility (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	 * see : https://schema.org/accessibilityFeature
	 *
	 * @var string | string[]
	 */
	public $accessibility_feature;

	/**
	 * A characteristic of the described resource that is physiologically dangerous to some users. Related to WCAG 2.0 guideline 2.3 (&lt;a href=&quot;http://www.w3.org/wiki/WebSchemas/Accessibility&quot;&gt;WebSchemas wiki lists possible values&lt;/a&gt;).
	 * see : https://schema.org/accessibilityHazard
	 *
	 * @var string | string[]
	 */
	public $accessibility_hazard;

	/**
	 * Specifies the Person that is legally accountable for the CreativeWork.
	 * see : https://schema.org/accountablePerson
	 *
	 * @var \Person | \Person[]
	 */
	public $accountable_person;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 *
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public $aggregate_rating;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * A secondary title of the CreativeWork.
	 * see : https://schema.org/alternativeHeadline
	 *
	 * @var string | string[]
	 */
	public $alternative_headline;

	/**
	 * A media object that encodes this CreativeWork. This property is a synonym for encoding.
	 * see : https://schema.org/associatedMedia
	 *
	 * @var \MediaObject | \MediaObject[]
	 */
	public $associated_media;

	/**
	 * An intended audience, i.e. a group for whom something was created.
	 * see : https://schema.org/audience
	 *
	 * @var \Audience | \Audience[]
	 */
	public $audience;

	/**
	 * An embedded audio object.
	 * see : https://schema.org/audio
	 *
	 * @var \AudioObject | \AudioObject[]
	 */
	public $audio;

	/**
	 * The author of this content. Please note that author is special in that HTML 5 provides a special mechanism for indicating authorship via the rel tag. That is equivalent to this and may be used interchangeably.
	 * see : https://schema.org/author
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $author;

	/**
	 * An award won by or for this item.
	 * see : https://schema.org/award
	 *
	 * @var string | string[]
	 */
	public $award;

	/**
	 * Awards won by or for this item.
	 * see : https://schema.org/awards
	 *
	 * @var string | string[]
	 */
	public $awards;

	/**
	 * A set of links that can help a user understand and navigate a website hierarchy.
	 * see : https://schema.org/breadcrumb
	 *
	 * @var string | string[] | \BreadcrumbList | \BreadcrumbList[]
	 */
	public $breadcrumb;

	/**
	 * Fictional person connected with a creative work.
	 * see : https://schema.org/character
	 *
	 * @var \Person | \Person[]
	 */
	public $character;

	/**
	 * A citation or reference to another creative work, such as another publication, web page, scholarly article, etc.
	 * see : https://schema.org/citation
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $citation;

	/**
	 * Comments, typically from users.
	 * see : https://schema.org/comment
	 *
	 * @var \Comment | \Comment[]
	 */
	public $comment;

	/**
	 * The number of comments this CreativeWork (e.g. Article, Question or Answer) has received. This is most applicable to works published in Web sites with commenting system; additional comments may exist elsewhere.
	 * see : https://schema.org/commentCount
	 *
	 * @var integer | integer[]
	 */
	public $comment_count;

	/**
	 * The location depicted or described in the content. For example, the location in a photograph or painting.
	 * see : https://schema.org/contentLocation
	 *
	 * @var \Place | \Place[]
	 */
	public $content_location;

	/**
	 * Official rating of a piece of content&amp;#x2014;for example,&#39;MPAA PG-13&#39;.
	 * see : https://schema.org/contentRating
	 *
	 * @var string | string[]
	 */
	public $content_rating;

	/**
	 * A secondary contributor to the CreativeWork.
	 * see : https://schema.org/contributor
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $contributor;

	/**
	 * The party holding the legal copyright to the CreativeWork.
	 * see : https://schema.org/copyrightHolder
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $copyright_holder;

	/**
	 * The year during which the claimed copyright for the CreativeWork was first asserted.
	 * see : https://schema.org/copyrightYear
	 *
	 * @var float | float[]
	 */
	public $copyright_year;

	/**
	 * The creator/author of this CreativeWork. This is the same as the Author property for CreativeWork.
	 * see : https://schema.org/creator
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $creator;

	/**
	 * The date on which the CreativeWork was created or the item was added to a DataFeed.
	 * see : https://schema.org/dateCreated
	 *
	 * @var string | string[]
	 */
	public $date_created;

	/**
	 * The date on which the CreativeWork was most recently modified or when the item&#39;s entry was modified within a DataFeed.
	 * see : https://schema.org/dateModified
	 *
	 * @var string | string[]
	 */
	public $date_modified;

	/**
	 * Date of first broadcast/publication.
	 * see : https://schema.org/datePublished
	 *
	 * @var string | string[]
	 */
	public $date_published;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A link to the page containing the comments of the CreativeWork.
	 * see : https://schema.org/discussionUrl
	 *
	 * @var string | string[]
	 */
	public $discussion_url;

	/**
	 * Specifies the Person who edited the CreativeWork.
	 * see : https://schema.org/editor
	 *
	 * @var \Person | \Person[]
	 */
	public $editor;

	/**
	 * An alignment to an established educational framework.
	 * see : https://schema.org/educationalAlignment
	 *
	 * @var \AlignmentObject | \AlignmentObject[]
	 */
	public $educational_alignment;

	/**
	 * The purpose of a work in the context of education; for example, &#39;assignment&#39;, &#39;group work&#39;.
	 * see : https://schema.org/educationalUse
	 *
	 * @var string | string[]
	 */
	public $educational_use;

	/**
	 * A media object that encodes this CreativeWork. This property is a synonym for associatedMedia.
	 * see : https://schema.org/encoding
	 *
	 * @var \MediaObject | \MediaObject[]
	 */
	public $encoding;

	/**
	 * A media object that encodes this CreativeWork.
	 * see : https://schema.org/encodings
	 *
	 * @var \MediaObject | \MediaObject[]
	 */
	public $encodings;

	/**
	 * A creative work that this work is an example/instance/realization/derivation of.
	 * see : https://schema.org/exampleOfWork
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $example_of_work;

	/**
	 * Media type (aka MIME format, see &lt;a href=&quot;http://www.iana.org/assignments/media-types/media-types.xhtml&quot;&gt;IANA site&lt;/a&gt;) of the content e.g. application/zip of a SoftwareApplication binary. In cases where a CreativeWork has several media type representations, &#39;encoding&#39; can be used to indicate each MediaObject alongside particular fileFormat information.
	 * see : https://schema.org/fileFormat
	 *
	 * @var string | string[]
	 */
	public $file_format;

	/**
	 * Genre of the creative work or group.
	 * see : https://schema.org/genre
	 *
	 * @var string | string[]
	 */
	public $genre;

	/**
	 * Indicates a CreativeWork that is (in some sense) a part of this CreativeWork.
	 * see : https://schema.org/hasPart
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $has_part;

	/**
	 * Headline of the article.
	 * see : https://schema.org/headline
	 *
	 * @var string | string[]
	 */
	public $headline;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * The language of the content or performance or used in an action. Please use one of the language codes from the &lt;a href=&#39;http://tools.ietf.org/html/bcp47&#39;&gt;IETF BCP 47 standard&lt;/a&gt;.
	 * see : https://schema.org/inLanguage
	 *
	 * @var string | string[] | \Language | \Language[]
	 */
	public $in_language;

	/**
	 * The number of interactions for the CreativeWork using the WebSite or SoftwareApplication. The most specific child type of InteractionCounter should be used.
	 * see : https://schema.org/interactionStatistic
	 *
	 * @var \InteractionCounter | \InteractionCounter[]
	 */
	public $interaction_statistic;

	/**
	 * The predominant mode of learning supported by the learning resource. Acceptable values are &#39;active&#39;, &#39;expositive&#39;, or &#39;mixed&#39;.
	 * see : https://schema.org/interactivityType
	 *
	 * @var string | string[]
	 */
	public $interactivity_type;

	/**
	 * A resource that was used in the creation of this resource. This term can be repeated for multiple sources. For example, http://example.com/great-multiplication-intro.html.
	 * see : https://schema.org/isBasedOnUrl
	 *
	 * @var string | string[]
	 */
	public $is_based_on_url;

	/**
	 * Indicates whether this content is family friendly.
	 * see : https://schema.org/isFamilyFriendly
	 *
	 * @var boolean | boolean[]
	 */
	public $is_family_friendly;

	/**
	 * Indicates a CreativeWork that this CreativeWork is (in some sense) part of.
	 * see : https://schema.org/isPartOf
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $is_part_of;

	/**
	 * Keywords or tags used to describe this content. Multiple entries in a keywords list are typically delimited by commas.
	 * see : https://schema.org/keywords
	 *
	 * @var string | string[]
	 */
	public $keywords;

	/**
	 * Date on which the content on this web page was last reviewed for accuracy and/or completeness.
	 * see : https://schema.org/lastReviewed
	 *
	 * @var string | string[]
	 */
	public $last_reviewed;

	/**
	 * The predominant type or kind characterizing the learning resource. For example, &#39;presentation&#39;, &#39;handout&#39;.
	 * see : https://schema.org/learningResourceType
	 *
	 * @var string | string[]
	 */
	public $learning_resource_type;

	/**
	 * A license document that applies to this content, typically indicated by URL.
	 * see : https://schema.org/license
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $license;

	/**
	 * The location where the CreativeWork was created, which may not be the same as the location depicted in the CreativeWork.
	 * see : https://schema.org/locationCreated
	 *
	 * @var \Place | \Place[]
	 */
	public $location_created;

	/**
	 * Indicates if this web page element is the main subject of the page.
	 * see : https://schema.org/mainContentOfPage
	 *
	 * @var \WebPageElement | \WebPageElement[]
	 */
	public $main_content_of_page;

	/**
	 * Indicates the primary entity described in some page or other CreativeWork.
	 * see : https://schema.org/mainEntity
	 *
	 * @var \Thing | \Thing[]
	 */
	public $main_entity;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * Indicates that the CreativeWork contains a reference to, but is not necessarily about a concept.
	 * see : https://schema.org/mentions
	 *
	 * @var \Thing | \Thing[]
	 */
	public $mentions;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 *
	 * @var \Offer | \Offer[]
	 */
	public $offers;

	/**
	 * The position of an item in a series or sequence of items.
	 * see : https://schema.org/position
	 *
	 * @var string | string[] | integer | integer[]
	 */
	public $position;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * Indicates the main image on the page.
	 * see : https://schema.org/primaryImageOfPage
	 *
	 * @var \ImageObject | \ImageObject[]
	 */
	public $primary_image_of_page;

	/**
	 * The person or organization who produced the work (e.g. music album, movie, tv/radio series etc.).
	 * see : https://schema.org/producer
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $producer;

	/**
	 * The service provider, service operator, or service performer; the goods producer. Another party (a seller) may offer those services or goods on behalf of the provider. A provider may also serve as the seller.
	 * see : https://schema.org/provider
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $provider;

	/**
	 * A publication event associated with the item.
	 * see : https://schema.org/publication
	 *
	 * @var \PublicationEvent | \PublicationEvent[]
	 */
	public $publication;

	/**
	 * The publisher of the creative work.
	 * see : https://schema.org/publisher
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $publisher;

	/**
	 * Link to page describing the editorial principles of the organization primarily responsible for the creation of the CreativeWork.
	 * see : https://schema.org/publishingPrinciples
	 *
	 * @var string | string[]
	 */
	public $publishing_principles;

	/**
	 * The Event where the CreativeWork was recorded. The CreativeWork may capture all or part of the event.
	 * see : https://schema.org/recordedAt
	 *
	 * @var \Event | \Event[]
	 */
	public $recorded_at;

	/**
	 * A link related to this web page, for example to other related web pages.
	 * see : https://schema.org/relatedLink
	 *
	 * @var string | string[]
	 */
	public $related_link;

	/**
	 * The place and time the release was issued, expressed as a PublicationEvent.
	 * see : https://schema.org/releasedEvent
	 *
	 * @var \PublicationEvent | \PublicationEvent[]
	 */
	public $released_event;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * People or organizations that have reviewed the content on this web page for accuracy and/or completeness.
	 * see : https://schema.org/reviewedBy
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $reviewed_by;

	/**
	 * Review of the item.
	 * see : https://schema.org/reviews
	 *
	 * @var \Review | \Review[]
	 */
	public $reviews;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * Indicates (by URL or string) a particular version of a schema used in some CreativeWork. For example, a document could declare a schemaVersion using a URL such as http://schema.org/version/2.0/ if precise indication of schema version was required by some application.
	 * see : https://schema.org/schemaVersion
	 *
	 * @var string | string[]
	 */
	public $schema_version;

	/**
	 * One of the more significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most.
	 * see : https://schema.org/significantLink
	 *
	 * @var string | string[]
	 */
	public $significant_link;

	/**
	 * The most significant URLs on the page. Typically, these are the non-navigation links that are clicked on the most.
	 * see : https://schema.org/significantLinks
	 *
	 * @var string | string[]
	 */
	public $significant_links;

	/**
	 * The Organization on whose behalf the creator was working.
	 * see : https://schema.org/sourceOrganization
	 *
	 * @var \Organization | \Organization[]
	 */
	public $source_organization;

	/**
	 * One of the domain specialities to which this web page&#39;s content applies.
	 * see : https://schema.org/specialty
	 *
	 * @var \Specialty | \Specialty[]
	 */
	public $specialty;

	/**
	 * The textual content of this CreativeWork.
	 * see : https://schema.org/text
	 *
	 * @var string | string[]
	 */
	public $text;

	/**
	 * A thumbnail image relevant to the Thing.
	 * see : https://schema.org/thumbnailUrl
	 *
	 * @var string | string[]
	 */
	public $thumbnail_url;

	/**
	 * Approximate or typical time it takes to work with or through this learning resource for the typical intended target audience, e.g. &#39;P30M&#39;, &#39;P1H25M&#39;.
	 * see : https://schema.org/timeRequired
	 *
	 * @var \Duration | \Duration[]
	 */
	public $time_required;

	/**
	 * Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market.
	 * see : https://schema.org/translator
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $translator;

	/**
	 * The typical expected age range, e.g. &#39;7-9&#39;, &#39;11-&#39;.
	 * see : https://schema.org/typicalAgeRange
	 *
	 * @var string | string[]
	 */
	public $typical_age_range;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The version of the CreativeWork embodied by a specified resource.
	 * see : https://schema.org/version
	 *
	 * @var float | float[]
	 */
	public $version;

	/**
	 * An embedded video object.
	 * see : https://schema.org/video
	 *
	 * @var \VideoObject | \VideoObject[]
	 */
	public $video;

	/**
	 * Example/instance/realization/derivation of the concept of this creative work. eg. The paperback edition, first edition, or eBook.
	 * see : https://schema.org/workExample
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $work_example;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'ProfilePage',
		);

		$serialized = \SchemaOrg\json_serialize( $this->about );
		if ( ! empty( $serialized ) ) {
			$out['about'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->accessibility_api );
		if ( ! empty( $serialized ) ) {
			$out['accessibilityAPI'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->accessibility_control );
		if ( ! empty( $serialized ) ) {
			$out['accessibilityControl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->accessibility_feature );
		if ( ! empty( $serialized ) ) {
			$out['accessibilityFeature'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->accessibility_hazard );
		if ( ! empty( $serialized ) ) {
			$out['accessibilityHazard'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->accountable_person );
		if ( ! empty( $serialized ) ) {
			$out['accountablePerson'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternative_headline );
		if ( ! empty( $serialized ) ) {
			$out['alternativeHeadline'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->associated_media );
		if ( ! empty( $serialized ) ) {
			$out['associatedMedia'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->audio );
		if ( ! empty( $serialized ) ) {
			$out['audio'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->author );
		if ( ! empty( $serialized ) ) {
			$out['author'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->awards );
		if ( ! empty( $serialized ) ) {
			$out['awards'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->breadcrumb );
		if ( ! empty( $serialized ) ) {
			$out['breadcrumb'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->character );
		if ( ! empty( $serialized ) ) {
			$out['character'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->citation );
		if ( ! empty( $serialized ) ) {
			$out['citation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->comment );
		if ( ! empty( $serialized ) ) {
			$out['comment'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->comment_count );
		if ( ! empty( $serialized ) ) {
			$out['commentCount'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->content_location );
		if ( ! empty( $serialized ) ) {
			$out['contentLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->content_rating );
		if ( ! empty( $serialized ) ) {
			$out['contentRating'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contributor );
		if ( ! empty( $serialized ) ) {
			$out['contributor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->copyright_holder );
		if ( ! empty( $serialized ) ) {
			$out['copyrightHolder'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->copyright_year );
		if ( ! empty( $serialized ) ) {
			$out['copyrightYear'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->creator );
		if ( ! empty( $serialized ) ) {
			$out['creator'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->date_created );
		if ( ! empty( $serialized ) ) {
			$out['dateCreated'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->date_modified );
		if ( ! empty( $serialized ) ) {
			$out['dateModified'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->date_published );
		if ( ! empty( $serialized ) ) {
			$out['datePublished'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->discussion_url );
		if ( ! empty( $serialized ) ) {
			$out['discussionUrl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->editor );
		if ( ! empty( $serialized ) ) {
			$out['editor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->educational_alignment );
		if ( ! empty( $serialized ) ) {
			$out['educationalAlignment'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->educational_use );
		if ( ! empty( $serialized ) ) {
			$out['educationalUse'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->encoding );
		if ( ! empty( $serialized ) ) {
			$out['encoding'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->encodings );
		if ( ! empty( $serialized ) ) {
			$out['encodings'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->example_of_work );
		if ( ! empty( $serialized ) ) {
			$out['exampleOfWork'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->file_format );
		if ( ! empty( $serialized ) ) {
			$out['fileFormat'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->genre );
		if ( ! empty( $serialized ) ) {
			$out['genre'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->has_part );
		if ( ! empty( $serialized ) ) {
			$out['hasPart'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->headline );
		if ( ! empty( $serialized ) ) {
			$out['headline'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->in_language );
		if ( ! empty( $serialized ) ) {
			$out['inLanguage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->interaction_statistic );
		if ( ! empty( $serialized ) ) {
			$out['interactionStatistic'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->interactivity_type );
		if ( ! empty( $serialized ) ) {
			$out['interactivityType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_based_on_url );
		if ( ! empty( $serialized ) ) {
			$out['isBasedOnUrl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_family_friendly );
		if ( ! empty( $serialized ) ) {
			$out['isFamilyFriendly'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_part_of );
		if ( ! empty( $serialized ) ) {
			$out['isPartOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->keywords );
		if ( ! empty( $serialized ) ) {
			$out['keywords'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->last_reviewed );
		if ( ! empty( $serialized ) ) {
			$out['lastReviewed'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->learning_resource_type );
		if ( ! empty( $serialized ) ) {
			$out['learningResourceType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->license );
		if ( ! empty( $serialized ) ) {
			$out['license'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->location_created );
		if ( ! empty( $serialized ) ) {
			$out['locationCreated'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_content_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainContentOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity );
		if ( ! empty( $serialized ) ) {
			$out['mainEntity'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->mentions );
		if ( ! empty( $serialized ) ) {
			$out['mentions'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->position );
		if ( ! empty( $serialized ) ) {
			$out['position'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->primary_image_of_page );
		if ( ! empty( $serialized ) ) {
			$out['primaryImageOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->producer );
		if ( ! empty( $serialized ) ) {
			$out['producer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->provider );
		if ( ! empty( $serialized ) ) {
			$out['provider'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->publication );
		if ( ! empty( $serialized ) ) {
			$out['publication'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->publisher );
		if ( ! empty( $serialized ) ) {
			$out['publisher'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->publishing_principles );
		if ( ! empty( $serialized ) ) {
			$out['publishingPrinciples'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recorded_at );
		if ( ! empty( $serialized ) ) {
			$out['recordedAt'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->related_link );
		if ( ! empty( $serialized ) ) {
			$out['relatedLink'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->released_event );
		if ( ! empty( $serialized ) ) {
			$out['releasedEvent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->reviewed_by );
		if ( ! empty( $serialized ) ) {
			$out['reviewedBy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->reviews );
		if ( ! empty( $serialized ) ) {
			$out['reviews'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->schema_version );
		if ( ! empty( $serialized ) ) {
			$out['schemaVersion'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->significant_link );
		if ( ! empty( $serialized ) ) {
			$out['significantLink'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->significant_links );
		if ( ! empty( $serialized ) ) {
			$out['significantLinks'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->source_organization );
		if ( ! empty( $serialized ) ) {
			$out['sourceOrganization'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->specialty );
		if ( ! empty( $serialized ) ) {
			$out['specialty'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->text );
		if ( ! empty( $serialized ) ) {
			$out['text'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->thumbnail_url );
		if ( ! empty( $serialized ) ) {
			$out['thumbnailUrl'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->time_required );
		if ( ! empty( $serialized ) ) {
			$out['timeRequired'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->translator );
		if ( ! empty( $serialized ) ) {
			$out['translator'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->typical_age_range );
		if ( ! empty( $serialized ) ) {
			$out['typicalAgeRange'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->version );
		if ( ! empty( $serialized ) ) {
			$out['version'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->video );
		if ( ! empty( $serialized ) ) {
			$out['video'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->work_example );
		if ( ! empty( $serialized ) ) {
			$out['workExample'] = $serialized;
		}

		return $out;
	}
}
