<?php

class ClaimReview extends Review implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ClaimReview';
	
	/**
	 * A short summary of the specific claims reviewed in a ClaimReview.
	 * see : https://schema.org/claimReviewed
	 * @var string|string[]
	 */
	public var $claim_reviewed;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ClaimReview'
		);
		
		$serialized = so_json_serialize( $this->claim_reviewed );
		if ( ! empty( $serialized ) ) {
			$out['claimReviewed'] = $serialized;
		}
		
		return $out;
	}
}
