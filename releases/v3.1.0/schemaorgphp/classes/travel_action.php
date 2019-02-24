<?php
namespace SchemaOrg;

// TravelAction see : https://schema.org/TravelAction
class TravelAction implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'TravelAction';

	/**
	 * With properties from MoveAction see : https://schema.org/MoveAction
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
	 * A description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 *
	 * @var string | string[]
	 */
	public $disambiguating_description;

	/**
	 * The distance travelled, e.g. exercising or travelling.
	 * see : https://schema.org/distance
	 *
	 * @var \Distance | \Distance[]
	 */
	public $distance;

	/**
	 * The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to *December*.\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
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
	 * A sub property of location. The original location of the object or the agent before the action.
	 * see : https://schema.org/fromLocation
	 *
	 * @var \Place | \Place[]
	 */
	public $from_location;

	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * The object that helped the agent perform the action. e.g. John wrote a book with *a pen*.
	 * see : https://schema.org/instrument
	 *
	 * @var \Thing | \Thing[]
	 */
	public $instrument;

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
	 * The result produced in the action. e.g. John wrote *a book*.
	 * see : https://schema.org/result
	 *
	 * @var \Thing | \Thing[]
	 */
	public $result;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from *January* to December.\n\nNote that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
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
	 * A sub property of location. The final location of the object or the agent after the action.
	 * see : https://schema.org/toLocation
	 *
	 * @var \Place | \Place[]
	 */
	public $to_location;

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
			'@type'    => 'TravelAction',
		);

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

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->distance );
		if ( ! empty( $serialized ) ) {
			$out['distance'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->end_time );
		if ( ! empty( $serialized ) ) {
			$out['endTime'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->error );
		if ( ! empty( $serialized ) ) {
			$out['error'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->from_location );
		if ( ! empty( $serialized ) ) {
			$out['fromLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->instrument );
		if ( ! empty( $serialized ) ) {
			$out['instrument'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->result );
		if ( ! empty( $serialized ) ) {
			$out['result'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->to_location );
		if ( ! empty( $serialized ) ) {
			$out['toLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		return $out;
	}
}
