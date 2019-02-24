<?php
namespace SchemaOrg;

// UserPlays see : https://schema.org/UserPlays
class UserPlays implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'UserPlays';

	/**
	 * With properties from UserInteraction see : https://schema.org/UserInteraction
	 */

	/**
	 * With properties from Event see : https://schema.org/Event
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


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
	 * A person or organization attending the event.
	 * see : https://schema.org/attendee
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $attendee;

	/**
	 * A person attending the event.
	 * see : https://schema.org/attendees
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $attendees;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * The time admission will commence.
	 * see : https://schema.org/doorTime
	 *
	 * @var string | string[]
	 */
	public $door_time;

	/**
	 * The duration of the item (movie, audio recording, event, etc.) in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 date format&lt;/a&gt;.
	 * see : https://schema.org/duration
	 *
	 * @var \Duration | \Duration[]
	 */
	public $duration;

	/**
	 * The end date and time of the item (in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 date format&lt;/a&gt;).
	 * see : https://schema.org/endDate
	 *
	 * @var string | string[]
	 */
	public $end_date;

	/**
	 * An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
	 * see : https://schema.org/eventStatus
	 *
	 * @var \EventStatusType | \EventStatusType[]
	 */
	public $event_status;

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
	 * The location of the event, organization or action.
	 * see : https://schema.org/location
	 *
	 * @var \Place | \Place[] | \PostalAddress | \PostalAddress[]
	 */
	public $location;

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
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * An offer to provide this item&amp;#x2014;for example, an offer to sell a product, rent the DVD of a movie, or give away tickets to an event.
	 * see : https://schema.org/offers
	 *
	 * @var \Offer | \Offer[]
	 */
	public $offers;

	/**
	 * An organizer of an Event.
	 * see : https://schema.org/organizer
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $organizer;

	/**
	 * A performer at the event&amp;#x2014;for example, a presenter, musician, musical group or actor.
	 * see : https://schema.org/performer
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $performer;

	/**
	 * The main performer or performers of the event&amp;#x2014;for example, a presenter, musician, or actor.
	 * see : https://schema.org/performers
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $performers;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
	 * see : https://schema.org/previousStartDate
	 *
	 * @var string | string[]
	 */
	public $previous_start_date;

	/**
	 * The CreativeWork that captured all or part of this Event.
	 * see : https://schema.org/recordedIn
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $recorded_in;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The start date and time of the item (in &lt;a href=&#39;http://en.wikipedia.org/wiki/ISO_8601&#39;&gt;ISO 8601 date format&lt;/a&gt;).
	 * see : https://schema.org/startDate
	 *
	 * @var string | string[]
	 */
	public $start_date;

	/**
	 * An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference.
	 * see : https://schema.org/subEvent
	 *
	 * @var \Event | \Event[]
	 */
	public $sub_event;

	/**
	 * Events that are a part of this event. For example, a conference event includes many presentations, each subEvents of the conference.
	 * see : https://schema.org/subEvents
	 *
	 * @var \Event | \Event[]
	 */
	public $sub_events;

	/**
	 * An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent.
	 * see : https://schema.org/superEvent
	 *
	 * @var \Event | \Event[]
	 */
	public $super_event;

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
	 * A work performed in some event, for example a play performed in a TheaterEvent.
	 * see : https://schema.org/workPerformed
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $work_performed;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'UserPlays',
		);

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

		$serialized = \SchemaOrg\json_serialize( $this->attendee );
		if ( ! empty( $serialized ) ) {
			$out['attendee'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->attendees );
		if ( ! empty( $serialized ) ) {
			$out['attendees'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->door_time );
		if ( ! empty( $serialized ) ) {
			$out['doorTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->duration );
		if ( ! empty( $serialized ) ) {
			$out['duration'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->end_date );
		if ( ! empty( $serialized ) ) {
			$out['endDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->event_status );
		if ( ! empty( $serialized ) ) {
			$out['eventStatus'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->in_language );
		if ( ! empty( $serialized ) ) {
			$out['inLanguage'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->organizer );
		if ( ! empty( $serialized ) ) {
			$out['organizer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->performer );
		if ( ! empty( $serialized ) ) {
			$out['performer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->performers );
		if ( ! empty( $serialized ) ) {
			$out['performers'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->previous_start_date );
		if ( ! empty( $serialized ) ) {
			$out['previousStartDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->recorded_in );
		if ( ! empty( $serialized ) ) {
			$out['recordedIn'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->start_date );
		if ( ! empty( $serialized ) ) {
			$out['startDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sub_event );
		if ( ! empty( $serialized ) ) {
			$out['subEvent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sub_events );
		if ( ! empty( $serialized ) ) {
			$out['subEvents'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->super_event );
		if ( ! empty( $serialized ) ) {
			$out['superEvent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->typical_age_range );
		if ( ! empty( $serialized ) ) {
			$out['typicalAgeRange'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->work_performed );
		if ( ! empty( $serialized ) ) {
			$out['workPerformed'] = $serialized;
		}

		return $out;
	}
}
