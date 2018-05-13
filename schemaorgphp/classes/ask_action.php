<?php

class AskAction extends CommunicateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'AskAction';
	
	/**
	 * A sub property of object. A question.
	 * see : https://schema.org/question
	 * @var \Question|\Question[]
	 */
	public var $question;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'AskAction'
		);
		
		$serialized = so_json_serialize( $this->question );
		if ( ! empty( $serialized ) ) {
			$out['question'] = $serialized;
		}
		
		return $out;
	}
}
