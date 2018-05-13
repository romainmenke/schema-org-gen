<?php

class SoftwareApplication extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'SoftwareApplication';
	
	/**
	 * Type of software application, e.g. &#39;Game, Multimedia&#39;.
	 * see : https://schema.org/applicationCategory
	 * @var string|string[]|string|string[]
	 */
	public var $application_category;
	
	/**
	 * Subcategory of the application, e.g. &#39;Arcade Game&#39;.
	 * see : https://schema.org/applicationSubCategory
	 * @var string|string[]|string|string[]
	 */
	public var $application_sub_category;
	
	/**
	 * The name of the application suite to which the application belongs (e.g. Excel belongs to Office).
	 * see : https://schema.org/applicationSuite
	 * @var string|string[]
	 */
	public var $application_suite;
	
	/**
	 * Device required to run the application. Used in cases where a specific make/model is required to run the application. Supersedes device (see: https://schema.org/device).
	 * see : https://schema.org/availableOnDevice
	 * @var string|string[]
	 */
	public var $available_on_device;
	
	/**
	 * Countries for which the application is not supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	 * see : https://schema.org/countriesNotSupported
	 * @var string|string[]
	 */
	public var $countries_not_supported;
	
	/**
	 * Countries for which the application is supported. You can also provide the two-letter ISO 3166-1 alpha-2 country code.
	 * see : https://schema.org/countriesSupported
	 * @var string|string[]
	 */
	public var $countries_supported;
	
	/**
	 * If the file can be downloaded, URL to download the binary.
	 * see : https://schema.org/downloadUrl
	 * @var string|string[]
	 */
	public var $download_url;
	
	/**
	 * Features or modules provided by this application (and possibly required by other applications).
	 * see : https://schema.org/featureList
	 * @var string|string[]|string|string[]
	 */
	public var $feature_list;
	
	/**
	 * Size of the application / package (e.g. 18MB). In the absence of a unit (MB, KB etc.), KB will be assumed.
	 * see : https://schema.org/fileSize
	 * @var string|string[]
	 */
	public var $file_size;
	
	/**
	 * URL at which the app may be installed, if different from the URL of the item.
	 * see : https://schema.org/installUrl
	 * @var string|string[]
	 */
	public var $install_url;
	
	/**
	 * Minimum memory requirements.
	 * see : https://schema.org/memoryRequirements
	 * @var string|string[]|string|string[]
	 */
	public var $memory_requirements;
	
	/**
	 * Operating systems supported (Windows 7, OSX 10.6, Android 1.6).
	 * see : https://schema.org/operatingSystem
	 * @var string|string[]
	 */
	public var $operating_system;
	
	/**
	 * Permission(s) required to run the app (for example, a mobile app may require full internet access or may run only on wifi).
	 * see : https://schema.org/permissions
	 * @var string|string[]
	 */
	public var $permissions;
	
	/**
	 * Processor architecture required to run the application (e.g. IA64).
	 * see : https://schema.org/processorRequirements
	 * @var string|string[]
	 */
	public var $processor_requirements;
	
	/**
	 * Description of what changed in this version.
	 * see : https://schema.org/releaseNotes
	 * @var string|string[]|string|string[]
	 */
	public var $release_notes;
	
	/**
	 * A link to a screenshot image of the app.
	 * see : https://schema.org/screenshot
	 * @var \ImageObject|\ImageObject[]|string|string[]
	 */
	public var $screenshot;
	
	/**
	 * Additional content for a software application.
	 * see : https://schema.org/softwareAddOn
	 * @var \SoftwareApplication|\SoftwareApplication[]
	 */
	public var $software_add_on;
	
	/**
	 * Software application help.
	 * see : https://schema.org/softwareHelp
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $software_help;
	
	/**
	 * Component dependency requirements for application. This includes runtime environments and shared libraries that are not included in the application distribution package, but required to run the application (Examples: DirectX, Java or .NET runtime). Supersedes requirements (see: https://schema.org/requirements).
	 * see : https://schema.org/softwareRequirements
	 * @var string|string[]|string|string[]
	 */
	public var $software_requirements;
	
	/**
	 * Version of the software instance.
	 * see : https://schema.org/softwareVersion
	 * @var string|string[]
	 */
	public var $software_version;
	
	/**
	 * Storage requirements (free space required).
	 * see : https://schema.org/storageRequirements
	 * @var string|string[]|string|string[]
	 */
	public var $storage_requirements;
	
	/**
	 * Supporting data for a SoftwareApplication.
	 * see : https://schema.org/supportingData
	 * @var \DataFeed|\DataFeed[]
	 */
	public var $supporting_data;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'SoftwareApplication'
		);
		
		$serialized = so_json_serialize( $this->application_category );
		if ( ! empty( $serialized ) ) {
			$out['applicationCategory'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->application_sub_category );
		if ( ! empty( $serialized ) ) {
			$out['applicationSubCategory'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->application_suite );
		if ( ! empty( $serialized ) ) {
			$out['applicationSuite'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->available_on_device );
		if ( ! empty( $serialized ) ) {
			$out['availableOnDevice'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->countries_not_supported );
		if ( ! empty( $serialized ) ) {
			$out['countriesNotSupported'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->countries_supported );
		if ( ! empty( $serialized ) ) {
			$out['countriesSupported'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->download_url );
		if ( ! empty( $serialized ) ) {
			$out['downloadUrl'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->feature_list );
		if ( ! empty( $serialized ) ) {
			$out['featureList'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->file_size );
		if ( ! empty( $serialized ) ) {
			$out['fileSize'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->install_url );
		if ( ! empty( $serialized ) ) {
			$out['installUrl'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->memory_requirements );
		if ( ! empty( $serialized ) ) {
			$out['memoryRequirements'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->operating_system );
		if ( ! empty( $serialized ) ) {
			$out['operatingSystem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->permissions );
		if ( ! empty( $serialized ) ) {
			$out['permissions'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->processor_requirements );
		if ( ! empty( $serialized ) ) {
			$out['processorRequirements'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->release_notes );
		if ( ! empty( $serialized ) ) {
			$out['releaseNotes'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->screenshot );
		if ( ! empty( $serialized ) ) {
			$out['screenshot'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->software_add_on );
		if ( ! empty( $serialized ) ) {
			$out['softwareAddOn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->software_help );
		if ( ! empty( $serialized ) ) {
			$out['softwareHelp'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->software_requirements );
		if ( ! empty( $serialized ) ) {
			$out['softwareRequirements'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->software_version );
		if ( ! empty( $serialized ) ) {
			$out['softwareVersion'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->storage_requirements );
		if ( ! empty( $serialized ) ) {
			$out['storageRequirements'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->supporting_data );
		if ( ! empty( $serialized ) ) {
			$out['supportingData'] = $serialized;
		}
		
		return $out;
	}
}
