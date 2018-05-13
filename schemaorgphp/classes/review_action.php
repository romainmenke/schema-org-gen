<?php

class ReviewAction extends AssessAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ReviewAction';
	
	/**
	 * A sub property of result. The review that resulted in the performing of the action.
	 * see : https://schema.org/resultReview
	 * @var \Review|\Review[]
	 */
	public var $result_review;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ReviewAction'
		);
		
		$serialized = so_json_serialize( $this->result_review );
		if ( ! empty( $serialized ) ) {
			$out['resultReview'] = $serialized;
		}
		
		return $out;
	}
}
