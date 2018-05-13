<?php

class DigitalDocumentPermission extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DigitalDocumentPermission';
	
	/**
	 * The person, organization, contact point, or audience that has been granted this permission.
	 * see : https://schema.org/grantee
	 * @var \Audience|\Audience[]|\ContactPoint|\ContactPoint[]|\Organization|\Organization[]|\Person|\Person[]
	 */
	public var $grantee;
	
	/**
	 * The type of permission granted the person, organization, or audience.
	 * see : https://schema.org/permissionType
	 * @var \DigitalDocumentPermissionType|\DigitalDocumentPermissionType[]
	 */
	public var $permission_type;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DigitalDocumentPermission'
		);
		
		$serialized = so_json_serialize( $this->grantee );
		if ( ! empty( $serialized ) ) {
			$out['grantee'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->permission_type );
		if ( ! empty( $serialized ) ) {
			$out['permissionType'] = $serialized;
		}
		
		return $out;
	}
}
