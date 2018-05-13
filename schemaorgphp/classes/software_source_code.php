<?php

class SoftwareSourceCode extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SoftwareSourceCode';
	
	/**
	 * Link to the repository where the un-compiled, human readable code and related code is located (SVN, github, CodePlex).
	 * see : https://schema.org/codeRepository
	 * @var string|string[]
	 */
	public var $code_repository;
	
	/**
	 * What type of code sample: full (compile ready) solution, code snippet, inline code, scripts, template. Supersedes sampleType (see: https://schema.org/sampleType).
	 * see : https://schema.org/codeSampleType
	 * @var string|string[]
	 */
	public var $code_sample_type;
	
	/**
	 * The computer programming language.
	 * see : https://schema.org/programmingLanguage
	 * @var \ComputerLanguage|\ComputerLanguage[]|string|string[]
	 */
	public var $programming_language;
	
	/**
	 * Runtime platform or script interpreter dependencies (Example - Java v1, Python2.3, .Net Framework 3.0). Supersedes runtime (see: https://schema.org/runtime).
	 * see : https://schema.org/runtimePlatform
	 * @var string|string[]
	 */
	public var $runtime_platform;
	
	/**
	 * Target Operating System / Product to which the code applies.  If applies to several versions, just the product name can be used.
	 * see : https://schema.org/targetProduct
	 * @var \SoftwareApplication|\SoftwareApplication[]
	 */
	public var $target_product;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SoftwareSourceCode'
		);
		
		$serialized = so_json_serialize( $this->code_repository );
		if ( ! empty( $serialized ) ) {
			$out['codeRepository'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->code_sample_type );
		if ( ! empty( $serialized ) ) {
			$out['codeSampleType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->programming_language );
		if ( ! empty( $serialized ) ) {
			$out['programmingLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->runtime_platform );
		if ( ! empty( $serialized ) ) {
			$out['runtimePlatform'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->target_product );
		if ( ! empty( $serialized ) ) {
			$out['targetProduct'] = $serialized;
		}
		
		return $out;
	}
}
