<?php

class CheckInAction extends CommunicateAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CheckInAction';
	
	/**
	 * The subject matter of the content.
	 * see : https://schema.org/about
	 * @var \Thing|\Thing[]
	 */
	public var $about;
	
	/**
	 * The language of the content or performance or used in an action. Please use one of the language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47). See also availableLanguage (see: https://schema.org/availableLanguage). Supersedes language (see: https://schema.org/language).
	 * see : https://schema.org/inLanguage
	 * @var \Language|\Language[]|string|string[]
	 */
	public var $in_language;
	
	/**
	 * A sub property of participant. The participant who is at the receiving end of the action.
	 * see : https://schema.org/recipient
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $recipient;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CheckInAction'
		);
		
		$serialized = so_json_serialize( $this->about );
		if ( ! empty( $serialized ) ) {
			$out['about'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->in_language );
		if ( ! empty( $serialized ) ) {
			$out['inLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->recipient );
		if ( ! empty( $serialized ) ) {
			$out['recipient'] = $serialized;
		}
		
		return $out;
	}
}
