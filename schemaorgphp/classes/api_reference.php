<?php

class APIReference extends TechArticle implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'APIReference';
	
	/**
	 * Associated product/technology version. e.g., .NET Framework 4.5.
	 * see : https://schema.org/assemblyVersion
	 * @var string|string[]
	 */
	public var $assembly_version;
	
	/**
	 * Library file name e.g., mscorlib.dll, system.web.dll. Supersedes assembly (see: https://schema.org/assembly).
	 * see : https://schema.org/executableLibraryName
	 * @var string|string[]
	 */
	public var $executable_library_name;
	
	/**
	 * Indicates whether API is managed or unmanaged.
	 * see : https://schema.org/programmingModel
	 * @var string|string[]
	 */
	public var $programming_model;
	
	/**
	 * Type of app development: phone, Metro style, desktop, XBox, etc.
	 * see : https://schema.org/targetPlatform
	 * @var string|string[]
	 */
	public var $target_platform;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'APIReference'
		);
		
		$serialized = so_json_serialize( $this->assembly_version );
		if ( ! empty( $serialized ) ) {
			$out['assemblyVersion'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->executable_library_name );
		if ( ! empty( $serialized ) ) {
			$out['executableLibraryName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->programming_model );
		if ( ! empty( $serialized ) ) {
			$out['programmingModel'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->target_platform );
		if ( ! empty( $serialized ) ) {
			$out['targetPlatform'] = $serialized;
		}
		
		return $out;
	}
}
