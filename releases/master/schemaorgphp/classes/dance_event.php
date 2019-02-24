<?php
namespace SchemaOrg;

// DanceEvent see : https://schema.org/DanceEvent
class DanceEvent implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'DanceEvent';

	/**
	 * With properties from Event see : https://schema.org/Event
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
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip.
	 * see : https://schema.org/actor
	 *
	 * @var \Person | \Person[]
	 */
	public $actor;

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
	 * An intended audience, i.e. a group for whom something was created.
	 * see : https://schema.org/audience
	 *
	 * @var \Audience | \Audience[]
	 */
	public $audience;

	/**
	 * The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	 * see : https://schema.org/composer
	 *
	 * @var \Person | \Person[] | \Organization | \Organization[]
	 */
	public $composer;

	/**
	 * A secondary contributor to the CreativeWork or Event.
	 * see : https://schema.org/contributor
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $contributor;

	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip.
	 * see : https://schema.org/director
	 *
	 * @var \Person | \Person[]
	 */
	public $director;

	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 *
	 * @var string | string[]
	 */
	public $disambiguating_description;

	/**
	 * The time admission will commence.
	 * see : https://schema.org/doorTime
	 *
	 * @var string | string[]
	 */
	public $door_time;

	/**
	 * The duration of the item (movie, audio recording, event, etc.) in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/duration
	 *
	 * @var \Duration | \Duration[]
	 */
	public $duration;

	/**
	 * The end date and time of the item (in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601)).
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
	 * A person or organization that supports (sponsors) something through some kind of financial contribution.
	 * see : https://schema.org/funder
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $funder;

	/**
	 * The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	 *
	 * see : https://schema.org/identifier
	 *
	 * @var string | string[] | \PropertyValue | \PropertyValue[]
	 */
	public $identifier;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * The language of the content or performance or used in an action. Please use one of the language codes from the [IETF BCP 47 standard](http://tools.ietf.org/html/bcp47). See also [[availableLanguage]].
	 * see : https://schema.org/inLanguage
	 *
	 * @var string | string[] | \Language | \Language[]
	 */
	public $in_language;

	/**
	 * A flag to signal that the item, event, or place is accessible for free.
	 * see : https://schema.org/isAccessibleForFree
	 *
	 * @var boolean | boolean[]
	 */
	public $is_accessible_for_free;

	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 *
	 * @var \Place | \Place[] | \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $location;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * The total number of individuals that may attend an event or venue.
	 * see : https://schema.org/maximumAttendeeCapacity
	 *
	 * @var integer | integer[]
	 */
	public $maximum_attendee_capacity;

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
	 * The number of attendee places for an event that remain unallocated.
	 * see : https://schema.org/remainingAttendeeCapacity
	 *
	 * @var integer | integer[]
	 */
	public $remaining_attendee_capacity;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	 * see : https://schema.org/sponsor
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $sponsor;

	/**
	 * The start date and time of the item (in [ISO 8601 date format](http://en.wikipedia.org/wiki/ISO_8601)).
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
	 * A CreativeWork or Event about this Thing..
	 * see : https://schema.org/subjectOf
	 *
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;

	/**
	 * An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent.
	 * see : https://schema.org/superEvent
	 *
	 * @var \Event | \Event[]
	 */
	public $super_event;

	/**
	 * Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
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
	 * A work featured in some event, e.g. exhibited in an ExhibitionEvent.
	 *        Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
	 * see : https://schema.org/workFeatured
	 *
	 * @var \CreativeWork | \CreativeWork[]
	 */
	public $work_featured;

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
			'@type'    => 'DanceEvent',
		);

		$serialized = \SchemaOrg\json_serialize( $this->about );
		if ( ! empty( $serialized ) ) {
			$out['about'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->attendee );
		if ( ! empty( $serialized ) ) {
			$out['attendee'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->attendees );
		if ( ! empty( $serialized ) ) {
			$out['attendees'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->composer );
		if ( ! empty( $serialized ) ) {
			$out['composer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contributor );
		if ( ! empty( $serialized ) ) {
			$out['contributor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->funder );
		if ( ! empty( $serialized ) ) {
			$out['funder'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->in_language );
		if ( ! empty( $serialized ) ) {
			$out['inLanguage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->is_accessible_for_free );
		if ( ! empty( $serialized ) ) {
			$out['isAccessibleForFree'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->maximum_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['maximumAttendeeCapacity'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->remaining_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['remainingAttendeeCapacity'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sponsor );
		if ( ! empty( $serialized ) ) {
			$out['sponsor'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->super_event );
		if ( ! empty( $serialized ) ) {
			$out['superEvent'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->work_featured );
		if ( ! empty( $serialized ) ) {
			$out['workFeatured'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->work_performed );
		if ( ! empty( $serialized ) ) {
			$out['workPerformed'] = $serialized;
		}

		return $out;
	}
}
