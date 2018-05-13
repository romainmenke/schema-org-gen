<?php

class ReplaceAction extends UpdateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ReplaceAction';
	
	/**
	 * A sub property of object. The object that is being replaced.
	 * see : https://schema.org/replacee
	 * @var \Thing|\Thing[]
	 */
	public var $replacee;
	
	/**
	 * A sub property of object. The object that replaces.
	 * see : https://schema.org/replacer
	 * @var \Thing|\Thing[]
	 */
	public var $replacer;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ReplaceAction'
		);
		
		$serialized = so_json_serialize( $this->replacee );
		if ( ! empty( $serialized ) ) {
			$out['replacee'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->replacer );
		if ( ! empty( $serialized ) ) {
			$out['replacer'] = $serialized;
		}
		
		return $out;
	}
}
