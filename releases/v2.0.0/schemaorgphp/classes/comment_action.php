<?php
namespace SchemaOrg;

// CommentAction see : https://schema.org/CommentAction
class CommentAction implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'CommentAction';

	/**
	 * With properties from CommunicateAction see : https://schema.org/CommunicateAction
	 */

	/**
	 * With properties from InteractAction see : https://schema.org/InteractAction
	 */

	/**
	 * With properties from Action see : https://schema.org/Action
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Action see : https://schema.org/Action
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
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
	 * Indicates the current disposition of the Action.
	 * see : https://schema.org/actionStatus
	 *
	 * @var \ActionStatusType | \ActionStatusType[]
	 */
	public $action_status;

	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * The direct performer or driver of the action (animate or inanimate). e.g. *John* wrote a book.
	 * see : https://schema.org/agent
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $agent;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to *December*.
	 *
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/endTime
	 *
	 * @var string | string[]
	 */
	public $end_time;

	/**
	 * For failed actions, more information on the cause of the failure.
	 * see : https://schema.org/error
	 *
	 * @var \Thing | \Thing[]
	 */
	public $error;

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
	 * The object that helped the agent perform the action. e.g. John wrote a book with *a pen*.
	 * see : https://schema.org/instrument
	 *
	 * @var \Thing | \Thing[]
	 */
	public $instrument;

	/**
	 * A sub property of instrument. The language used on this action.
	 * see : https://schema.org/language
	 *
	 * @var \Language | \Language[]
	 */
	public $language;

	/**
	 * The location of the event, organization or action.
	 * see : https://schema.org/location
	 *
	 * @var \Place | \Place[] | \PostalAddress | \PostalAddress[]
	 */
	public $location;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       Many (but not all) pages have a fairly clear primary topic, some entity or thing that the page describes. For
	 *       example a restaurant&#39;s home page might be primarily about that Restaurant, or an event listing page might
	 *       represent a single event. The mainEntity and mainEntityOfPage properties allow you to explicitly express the relationship
	 *       between the page and the primary entity.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       Related properties include sameAs, about, and url.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       The sameAs and url properties are both similar to mainEntityOfPage. The url property should be reserved to refer to more
	 *       official or authoritative web pages, such as the item’s official website. The sameAs property also relates a thing
	 *       to a page that indirectly identifies it. Whereas sameAs emphasises well known pages, the mainEntityOfPage property
	 *       serves more to clarify which of several entities is the main one for that page.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       mainEntityOfPage can be used for any page, including those not recognized as authoritative for that entity. For example,
	 *       for a product, sameAs might refer to a page on the manufacturer’s official site with specs for the product, while
	 *       mainEntityOfPage might be used on pages within various retailers’ sites giving details for the same product.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *
	 *       about is similar to mainEntity, with two key differences. First, about can refer to multiple entities/topics,
	 *       while mainEntity should be used for only the primary one. Second, some pages have a primary entity that itself
	 *       describes some other entity. For example, one web page may display a news article about a particular person.
	 *       Another page may display a product review for a particular product. In these cases, mainEntity for the pages
	 *       should refer to the news article or review, respectively, while about would more properly refer to the person or product.
	 *
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * The object upon the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn&#39;t). e.g. John read *a book*.
	 * see : https://schema.org/object
	 *
	 * @var \Thing | \Thing[]
	 */
	public $object;

	/**
	 * Other co-agents that participated in the action indirectly. e.g. John wrote a book with *Steve*.
	 * see : https://schema.org/participant
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $participant;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * A sub property of participant. The participant who is at the receiving end of the action.
	 * see : https://schema.org/recipient
	 *
	 * @var \Audience | \Audience[] | \Organization | \Organization[] | \Person | \Person[]
	 */
	public $recipient;

	/**
	 * The result produced in the action. e.g. John wrote *a book*.
	 * see : https://schema.org/result
	 *
	 * @var \Thing | \Thing[]
	 */
	public $result;

	/**
	 * A sub property of result. The Comment created or sent as a result of this action.
	 * see : https://schema.org/resultComment
	 *
	 * @var \Comment | \Comment[]
	 */
	public $result_comment;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from *January* to December.
	 *
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/startTime
	 *
	 * @var string | string[]
	 */
	public $start_time;

	/**
	 * Indicates a target EntryPoint for an Action.
	 * see : https://schema.org/target
	 *
	 * @var \EntryPoint | \EntryPoint[]
	 */
	public $target;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'CommentAction',
		);

		$serialized = \SchemaOrg\json_serialize( $this->about );
		if ( ! empty( $serialized ) ) {
			$out['about'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->action_status );
		if ( ! empty( $serialized ) ) {
			$out['actionStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->agent );
		if ( ! empty( $serialized ) ) {
			$out['agent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->end_time );
		if ( ! empty( $serialized ) ) {
			$out['endTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->error );
		if ( ! empty( $serialized ) ) {
			$out['error'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->in_language );
		if ( ! empty( $serialized ) ) {
			$out['inLanguage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->instrument );
		if ( ! empty( $serialized ) ) {
			$out['instrument'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->language );
		if ( ! empty( $serialized ) ) {
			$out['language'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->object );
		if ( ! empty( $serialized ) ) {
			$out['object'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->participant );
		if ( ! empty( $serialized ) ) {
			$out['participant'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recipient );
		if ( ! empty( $serialized ) ) {
			$out['recipient'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->result );
		if ( ! empty( $serialized ) ) {
			$out['result'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->result_comment );
		if ( ! empty( $serialized ) ) {
			$out['resultComment'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->start_time );
		if ( ! empty( $serialized ) ) {
			$out['startTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->target );
		if ( ! empty( $serialized ) ) {
			$out['target'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
