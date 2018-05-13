<?php

class EmailMessage extends Message implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EmailMessage';
	
	/**
	 * A sub property of recipient. The recipient blind copied on a message.
	 * see : https://schema.org/bccRecipient
	 * @var \ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $bcc_recipient;
	
	/**
	 * A sub property of recipient. The recipient copied on a message.
	 * see : https://schema.org/ccRecipient
	 * @var \ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $cc_recipient;
	
	/**
	 * The date/time at which the message has been read by the recipient if a single recipient exists.
	 * see : https://schema.org/dateRead
	 * @var string|string[]
	 */
	public var $date_read;
	
	/**
	 * The date/time the message was received if a single recipient exists.
	 * see : https://schema.org/dateReceived
	 * @var string|string[]
	 */
	public var $date_received;
	
	/**
	 * The date/time at which the message was sent.
	 * see : https://schema.org/dateSent
	 * @var string|string[]
	 */
	public var $date_sent;
	
	/**
	 * A CreativeWork attached to the message.
	 * see : https://schema.org/messageAttachment
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $message_attachment;
	
	/**
	 * A sub property of participant. The participant who is at the receiving end of the action.
	 * see : https://schema.org/recipient
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $recipient;
	
	/**
	 * A sub property of participant. The participant who is at the sending end of the action.
	 * see : https://schema.org/sender
	 * @var \Audience|\Audience[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $sender;
	
	/**
	 * A sub property of recipient. The recipient who was directly sent the message.
	 * see : https://schema.org/toRecipient
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $to_recipient;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EmailMessage'
		);
		
		$serialized = so_json_serialize( $this->bcc_recipient );
		if ( ! empty( $serialized ) ) {
			$out['bccRecipient'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->cc_recipient );
		if ( ! empty( $serialized ) ) {
			$out['ccRecipient'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_read );
		if ( ! empty( $serialized ) ) {
			$out['dateRead'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_received );
		if ( ! empty( $serialized ) ) {
			$out['dateReceived'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->date_sent );
		if ( ! empty( $serialized ) ) {
			$out['dateSent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->message_attachment );
		if ( ! empty( $serialized ) ) {
			$out['messageAttachment'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipient );
		if ( ! empty( $serialized ) ) {
			$out['recipient'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sender );
		if ( ! empty( $serialized ) ) {
			$out['sender'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->to_recipient );
		if ( ! empty( $serialized ) ) {
			$out['toRecipient'] = $serialized;
		}
		
		return $out;
	}
}
