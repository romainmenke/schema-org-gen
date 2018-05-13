<?php

class EducationEvent extends Event implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EducationEvent';
	
	/**
	 * The subject matter of the content.
	 * see : https://schema.org/about
	 * @var \Thing|\Thing[]
	 */
	public var $about;
	
	/**
	 * An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	 * see : https://schema.org/actor
	 * @var \Person|\Person[]
	 */
	public var $actor;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating|\AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * A person or organization attending the event. Supersedes attendees (see: https://schema.org/attendees).
	 * see : https://schema.org/attendee
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $attendee;
	
	/**
	 * An intended audience, i.e. a group for whom something was created. Supersedes serviceAudience (see: https://schema.org/serviceAudience).
	 * see : https://schema.org/audience
	 * @var \Audience|\Audience[]
	 */
	public var $audience;
	
	/**
	 * The person or organization who wrote a composition, or who is the composer of a work performed at some event.
	 * see : https://schema.org/composer
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $composer;
	
	/**
	 * A secondary contributor to the CreativeWork or Event.
	 * see : https://schema.org/contributor
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $contributor;
	
	/**
	 * A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	 * see : https://schema.org/director
	 * @var \Person|\Person[]
	 */
	public var $director;
	
	/**
	 * The time admission will commence.
	 * see : https://schema.org/doorTime
	 * @var string|string[]
	 */
	public var $door_time;
	
	/**
	 * The duration of the item (movie, audio recording, event, etc.) in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/duration
	 * @var \Duration|\Duration[]
	 */
	public var $duration;
	
	/**
	 * The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/endDate
	 * @var string|string[]|string|string[]
	 */
	public var $end_date;
	
	/**
	 * An eventStatus of an event represents its status; particularly useful when an event is cancelled or rescheduled.
	 * see : https://schema.org/eventStatus
	 * @var \EventStatusType|\EventStatusType[]
	 */
	public var $event_status;
	
	/**
	 * A person or organization that supports (sponsors) something through some kind of financial contribution.
	 * see : https://schema.org/funder
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $funder;
	
	/**
	 * The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	 * see : https://schema.org/inLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $in_language;
	
	/**
	 * A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	 * see : https://schema.org/isAccessibleForFree
	 * @var boolean|boolean[]
	 */
	public var $is_accessible_for_free;
	
	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 * @var \Place|\Place[]|\PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $location;
	
	/**
	 * The total number of individuals that may attend an event or venue.
	 * see : https://schema.org/maximumAttendeeCapacity
	 * @var integer|integer[]
	 */
	public var $maximum_attendee_capacity;
	
	/**
	 * An offer to provide this item—for example, an offer to sell a product, rent the DVD of a movie, perform a service, or give away tickets to an event.
	 * see : https://schema.org/offers
	 * @var \Offer|\Offer[]
	 */
	public var $offers;
	
	/**
	 * An organizer of an Event.
	 * see : https://schema.org/organizer
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $organizer;
	
	/**
	 * A performer at the event—for example, a presenter, musician, musical group or actor. Supersedes performers (see: https://schema.org/performers).
	 * see : https://schema.org/performer
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $performer;
	
	/**
	 * Used in conjunction with eventStatus for rescheduled or cancelled events. This property contains the previously scheduled start date. For rescheduled events, the startDate property should be used for the newly scheduled start date. In the (rare) case of an event that has been postponed and rescheduled multiple times, this field may be repeated.
	 * see : https://schema.org/previousStartDate
	 * @var string|string[]
	 */
	public var $previous_start_date;
	
	/**
	 * The CreativeWork that captured all or part of this Event. Inverse property: recordedAt (see: https://schema.org/recordedAt).
	 * see : https://schema.org/recordedIn
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $recorded_in;
	
	/**
	 * The number of attendee places for an event that remain unallocated.
	 * see : https://schema.org/remainingAttendeeCapacity
	 * @var integer|integer[]
	 */
	public var $remaining_attendee_capacity;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review|\Review[]
	 */
	public var $review;
	
	/**
	 * A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	 * see : https://schema.org/sponsor
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $sponsor;
	
	/**
	 * The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/startDate
	 * @var string|string[]|string|string[]
	 */
	public var $start_date;
	
	/**
	 * An Event that is part of this event. For example, a conference event includes many presentations, each of which is a subEvent of the conference. Supersedes subEvents (see: https://schema.org/subEvents). Inverse property: superEvent (see: https://schema.org/superEvent).
	 * see : https://schema.org/subEvent
	 * @var \Event|\Event[]
	 */
	public var $sub_event;
	
	/**
	 * An event that this event is a part of. For example, a collection of individual music performances might each have a music festival as their superEvent. Inverse property: subEvent (see: https://schema.org/subEvent).
	 * see : https://schema.org/superEvent
	 * @var \Event|\Event[]
	 */
	public var $super_event;
	
	/**
	 * Organization or person who adapts a creative work to different languages, regional differences and technical requirements of a target market, or that translates during some event.
	 * see : https://schema.org/translator
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $translator;
	
	/**
	 * The typical expected age range, e.g. &#39;7-9&#39;, &#39;11-&#39;.
	 * see : https://schema.org/typicalAgeRange
	 * @var string|string[]
	 */
	public var $typical_age_range;
	
	/**
	 * A work featured in some event, e.g. exhibited in an ExhibitionEvent.
	 *        Specific subproperties are available for workPerformed (e.g. a play), or a workPresented (a Movie at a ScreeningEvent).
	 * see : https://schema.org/workFeatured
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $work_featured;
	
	/**
	 * A work performed in some event, for example a play performed in a TheaterEvent.
	 * see : https://schema.org/workPerformed
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $work_performed;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EducationEvent'
		);
		
		$serialized = so_json_serialize( $this->about );
		if ( ! empty( $serialized ) ) {
			$out['about'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->actor );
		if ( ! empty( $serialized ) ) {
			$out['actor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->attendee );
		if ( ! empty( $serialized ) ) {
			$out['attendee'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->audience );
		if ( ! empty( $serialized ) ) {
			$out['audience'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->composer );
		if ( ! empty( $serialized ) ) {
			$out['composer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contributor );
		if ( ! empty( $serialized ) ) {
			$out['contributor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->director );
		if ( ! empty( $serialized ) ) {
			$out['director'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->door_time );
		if ( ! empty( $serialized ) ) {
			$out['doorTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->duration );
		if ( ! empty( $serialized ) ) {
			$out['duration'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->end_date );
		if ( ! empty( $serialized ) ) {
			$out['endDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->event_status );
		if ( ! empty( $serialized ) ) {
			$out['eventStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->funder );
		if ( ! empty( $serialized ) ) {
			$out['funder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->in_language );
		if ( ! empty( $serialized ) ) {
			$out['inLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->is_accessible_for_free );
		if ( ! empty( $serialized ) ) {
			$out['isAccessibleForFree'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->maximum_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['maximumAttendeeCapacity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->offers );
		if ( ! empty( $serialized ) ) {
			$out['offers'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->organizer );
		if ( ! empty( $serialized ) ) {
			$out['organizer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->performer );
		if ( ! empty( $serialized ) ) {
			$out['performer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->previous_start_date );
		if ( ! empty( $serialized ) ) {
			$out['previousStartDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recorded_in );
		if ( ! empty( $serialized ) ) {
			$out['recordedIn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->remaining_attendee_capacity );
		if ( ! empty( $serialized ) ) {
			$out['remainingAttendeeCapacity'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sponsor );
		if ( ! empty( $serialized ) ) {
			$out['sponsor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->start_date );
		if ( ! empty( $serialized ) ) {
			$out['startDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sub_event );
		if ( ! empty( $serialized ) ) {
			$out['subEvent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->super_event );
		if ( ! empty( $serialized ) ) {
			$out['superEvent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->translator );
		if ( ! empty( $serialized ) ) {
			$out['translator'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->typical_age_range );
		if ( ! empty( $serialized ) ) {
			$out['typicalAgeRange'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->work_featured );
		if ( ! empty( $serialized ) ) {
			$out['workFeatured'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->work_performed );
		if ( ! empty( $serialized ) ) {
			$out['workPerformed'] = $serialized;
		}
		
		return $out;
	}
}
