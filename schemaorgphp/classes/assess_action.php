<?php

class AssessAction extends Action implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AssessAction';
	
	/**
	 * Indicates the current disposition of the Action.
	 * see : https://schema.org/actionStatus
	 * @var \ActionStatusType|\ActionStatusType[]
	 */
	public var $action_status;
	
	/**
	 * The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book.
	 * see : https://schema.org/agent
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $agent;
	
	/**
	 * The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	 * 
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/endTime
	 * @var string|string[]
	 */
	public var $end_time;
	
	/**
	 * For failed actions, more information on the cause of the failure.
	 * see : https://schema.org/error
	 * @var \Thing|\Thing[]
	 */
	public var $error;
	
	/**
	 * The object that helped the agent perform the action. e.g. John wrote a book with a pen.
	 * see : https://schema.org/instrument
	 * @var \Thing|\Thing[]
	 */
	public var $instrument;
	
	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 * @var \Place|\Place[]|\PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $location;
	
	/**
	 * The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn&#39;t). e.g. John read a book.
	 * see : https://schema.org/object
	 * @var \Thing|\Thing[]
	 */
	public var $object;
	
	/**
	 * Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve.
	 * see : https://schema.org/participant
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $participant;
	
	/**
	 * The result produced in the action. e.g. John wrote a book.
	 * see : https://schema.org/result
	 * @var \Thing|\Thing[]
	 */
	public var $result;
	
	/**
	 * The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	 * 
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/startTime
	 * @var string|string[]
	 */
	public var $start_time;
	
	/**
	 * Indicates a target EntryPoint for an Action.
	 * see : https://schema.org/target
	 * @var \EntryPoint|\EntryPoint[]
	 */
	public var $target;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AssessAction'
		);
		
		$serialized = so_json_serialize( $this->action_status );
		if ( ! empty( $serialized ) ) {
			$out['actionStatus'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->agent );
		if ( ! empty( $serialized ) ) {
			$out['agent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->end_time );
		if ( ! empty( $serialized ) ) {
			$out['endTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->error );
		if ( ! empty( $serialized ) ) {
			$out['error'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->instrument );
		if ( ! empty( $serialized ) ) {
			$out['instrument'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->object );
		if ( ! empty( $serialized ) ) {
			$out['object'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->participant );
		if ( ! empty( $serialized ) ) {
			$out['participant'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->result );
		if ( ! empty( $serialized ) ) {
			$out['result'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->start_time );
		if ( ! empty( $serialized ) ) {
			$out['startTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->target );
		if ( ! empty( $serialized ) ) {
			$out['target'] = $serialized;
		}
		
		return $out;
	}
}
