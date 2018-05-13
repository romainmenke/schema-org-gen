<?php

class RsvpAction extends InformAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'RsvpAction';
	
	/**
	 * If responding yes, the number of guests who will attend in addition to the invitee.
	 * see : https://schema.org/additionalNumberOfGuests
	 * @var float|float[]
	 */
	public var $additional_number_of_guests;
	
	/**
	 * Comments, typically from users.
	 * see : https://schema.org/comment
	 * @var \Comment|\Comment[]
	 */
	public var $comment;
	
	/**
	 * The response (yes, no, maybe) to the RSVP.
	 * see : https://schema.org/rsvpResponse
	 * @var \RsvpResponseType|\RsvpResponseType[]
	 */
	public var $rsvp_response;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'RsvpAction'
		);
		
		$serialized = so_json_serialize( $this->additional_number_of_guests );
		if ( ! empty( $serialized ) ) {
			$out['additionalNumberOfGuests'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->comment );
		if ( ! empty( $serialized ) ) {
			$out['comment'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->rsvp_response );
		if ( ! empty( $serialized ) ) {
			$out['rsvpResponse'] = $serialized;
		}
		
		return $out;
	}
}
