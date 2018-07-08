<?php

class WearAction extends UseAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'WearAction';
	
	/**
	 * A set of requirements that a must be fulfilled in order to perform an Action. If more than one value is specied, fulfilling one set of requirements will allow the Action to be performed.
	 * see : https://pending.schema.org/actionAccessibilityRequirement
	 * @var \ActionAccessSpecification|\ActionAccessSpecification[]
	 */
	public var $action_accessibility_requirement;
	
	/**
	 * An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	 * see : https://pending.schema.org/expectsAcceptanceOf
	 * @var \Offer|\Offer[]
	 */
	public var $expects_acceptance_of;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'WearAction'
		);
		
		$serialized = so_json_serialize( $this->action_accessibility_requirement );
		if ( ! empty( $serialized ) ) {
			$out['actionAccessibilityRequirement'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->expects_acceptance_of );
		if ( ! empty( $serialized ) ) {
			$out['expectsAcceptanceOf'] = $serialized;
		}
		
		return $out;
	}
}
