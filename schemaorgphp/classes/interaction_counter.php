<?php

class InteractionCounter extends StructuredValue implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'InteractionCounter';
	
	/**
	 * The WebSite or SoftwareApplication where the interactions took place.
	 * see : https://schema.org/interactionService
	 * @var \SoftwareApplication|\SoftwareApplication[]|\WebSite|\WebSite[]
	 */
	public var $interaction_service;
	
	/**
	 * The Action representing the type of interaction. For up votes, +1s, etc. use LikeAction (see: https://schema.org/LikeAction). For down votes use DislikeAction (see: https://schema.org/DislikeAction). Otherwise, use the most specific Action.
	 * see : https://schema.org/interactionType
	 * @var \Action|\Action[]
	 */
	public var $interaction_type;
	
	/**
	 * The number of interactions for the CreativeWork using the WebSite or SoftwareApplication.
	 * see : https://schema.org/userInteractionCount
	 * @var integer|integer[]
	 */
	public var $user_interaction_count;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'InteractionCounter'
		);
		
		$serialized = so_json_serialize( $this->interaction_service );
		if ( ! empty( $serialized ) ) {
			$out['interactionService'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->interaction_type );
		if ( ! empty( $serialized ) ) {
			$out['interactionType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->user_interaction_count );
		if ( ! empty( $serialized ) ) {
			$out['userInteractionCount'] = $serialized;
		}
		
		return $out;
	}
}
