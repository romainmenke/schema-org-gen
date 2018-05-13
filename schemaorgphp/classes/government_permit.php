<?php

class GovernmentPermit extends Permit implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'GovernmentPermit';
	
	/**
	 * The organization issuing the ticket or permit.
	 * see : https://schema.org/issuedBy
	 * @var \Organization|\Organization[]
	 */
	public var $issued_by;
	
	/**
	 * The service through with the permit was granted.
	 * see : https://schema.org/issuedThrough
	 * @var \Service|\Service[]
	 */
	public var $issued_through;
	
	/**
	 * The target audience for this permit.
	 * see : https://schema.org/permitAudience
	 * @var \Audience|\Audience[]
	 */
	public var $permit_audience;
	
	/**
	 * The time validity of the permit.
	 * see : https://schema.org/validFor
	 * @var \Duration|\Duration[]
	 */
	public var $valid_for;
	
	/**
	 * The date when the item becomes valid.
	 * see : https://schema.org/validFrom
	 * @var string|string[]
	 */
	public var $valid_from;
	
	/**
	 * The geographic area where the permit is valid.
	 * see : https://schema.org/validIn
	 * @var \AdministrativeArea|\AdministrativeArea[]
	 */
	public var $valid_in;
	
	/**
	 * The date when the item is no longer valid.
	 * see : https://schema.org/validUntil
	 * @var string|string[]
	 */
	public var $valid_until;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'GovernmentPermit'
		);
		
		$serialized = so_json_serialize( $this->issued_by );
		if ( ! empty( $serialized ) ) {
			$out['issuedBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->issued_through );
		if ( ! empty( $serialized ) ) {
			$out['issuedThrough'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->permit_audience );
		if ( ! empty( $serialized ) ) {
			$out['permitAudience'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_for );
		if ( ! empty( $serialized ) ) {
			$out['validFor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_from );
		if ( ! empty( $serialized ) ) {
			$out['validFrom'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_in );
		if ( ! empty( $serialized ) ) {
			$out['validIn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->valid_until );
		if ( ! empty( $serialized ) ) {
			$out['validUntil'] = $serialized;
		}
		
		return $out;
	}
}
