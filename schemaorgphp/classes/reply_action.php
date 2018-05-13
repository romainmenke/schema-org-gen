<?php

class ReplyAction extends CommunicateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ReplyAction';
	
	/**
	 * A sub property of result. The Comment created or sent as a result of this action.
	 * see : https://schema.org/resultComment
	 * @var \Comment|\Comment[]
	 */
	public var $result_comment;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ReplyAction'
		);
		
		$serialized = so_json_serialize( $this->result_comment );
		if ( ! empty( $serialized ) ) {
			$out['resultComment'] = $serialized;
		}
		
		return $out;
	}
}
