<?php

class ViewAction extends ConsumeAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ViewAction';
	
	/**
	 * An Offer which must be accepted before the user can perform the Action. For example, the user may need to buy a movie before being able to watch it.
	 * see : https://schema.org/expectsAcceptanceOf
	 * @var \Offer|\Offer[]
	 */
	public var $expects_acceptance_of;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ViewAction'
		);
		
		$serialized = so_json_serialize( $this->expects_acceptance_of );
		if ( ! empty( $serialized ) ) {
			$out['expectsAcceptanceOf'] = $serialized;
		}
		
		return $out;
	}
}
