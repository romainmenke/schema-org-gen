<?php

class EntryPoint extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'EntryPoint';
	
	/**
	 * An application that can complete the request. Supersedes application (see: https://schema.org/application).
	 * see : https://schema.org/actionApplication
	 * @var \SoftwareApplication|\SoftwareApplication[]
	 */
	public var $action_application;
	
	/**
	 * The high level platform(s) where the Action can be performed for the given URL. To specify a specific application or operating system instance, use actionApplication.
	 * see : https://schema.org/actionPlatform
	 * @var string|string[]|string|string[]
	 */
	public var $action_platform;
	
	/**
	 * The supported content type(s) for an EntryPoint response.
	 * see : https://schema.org/contentType
	 * @var string|string[]
	 */
	public var $content_type;
	
	/**
	 * The supported encoding type(s) for an EntryPoint request.
	 * see : https://schema.org/encodingType
	 * @var string|string[]
	 */
	public var $encoding_type;
	
	/**
	 * An HTTP method that specifies the appropriate HTTP method for a request to an HTTP EntryPoint. Values are capitalized strings as used in HTTP.
	 * see : https://schema.org/httpMethod
	 * @var string|string[]
	 */
	public var $http_method;
	
	/**
	 * An url template (RFC6570) that will be used to construct the target of the execution of the action.
	 * see : https://schema.org/urlTemplate
	 * @var string|string[]
	 */
	public var $url_template;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'EntryPoint'
		);
		
		$serialized = so_json_serialize( $this->action_application );
		if ( ! empty( $serialized ) ) {
			$out['actionApplication'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->action_platform );
		if ( ! empty( $serialized ) ) {
			$out['actionPlatform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->content_type );
		if ( ! empty( $serialized ) ) {
			$out['contentType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->encoding_type );
		if ( ! empty( $serialized ) ) {
			$out['encodingType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->http_method );
		if ( ! empty( $serialized ) ) {
			$out['httpMethod'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->url_template );
		if ( ! empty( $serialized ) ) {
			$out['urlTemplate'] = $serialized;
		}
		
		return $out;
	}
}
