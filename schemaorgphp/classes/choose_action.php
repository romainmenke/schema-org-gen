<?php

class ChooseAction extends AssessAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ChooseAction';
	
	/**
	 * A sub property of object. The options subject to this action. Supersedes option (see: https://schema.org/option).
	 * see : https://schema.org/actionOption
	 * @var string|string[]|\Thing|\Thing[]
	 */
	public var $action_option;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ChooseAction'
		);
		
		$serialized = so_json_serialize( $this->action_option );
		if ( ! empty( $serialized ) ) {
			$out['actionOption'] = $serialized;
		}
		
		return $out;
	}
}
