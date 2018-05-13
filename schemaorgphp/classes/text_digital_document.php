<?php

class TextDigitalDocument extends DigitalDocument implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'TextDigitalDocument';
	
	/**
	 * A permission related to the access to this document (e.g. permission to read or write an electronic document). For a public document, specify a grantee with an Audience with audienceType equal to &quot;public&quot;.
	 * see : https://schema.org/hasDigitalDocumentPermission
	 * @var \DigitalDocumentPermission|\DigitalDocumentPermission[]
	 */
	public var $has_digital_document_permission;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'TextDigitalDocument'
		);
		
		$serialized = so_json_serialize( $this->has_digital_document_permission );
		if ( ! empty( $serialized ) ) {
			$out['hasDigitalDocumentPermission'] = $serialized;
		}
		
		return $out;
	}
}
