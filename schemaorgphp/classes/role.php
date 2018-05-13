<?php

class Role extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Role';
	
	/**
	 * The end date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/endDate
	 * @var string|string[]|string|string[]
	 */
	public var $end_date;
	
	/**
	 * A role played, performed or filled by a person or organization. For example, the team of creators for a comic book might fill the roles named &#39;inker&#39;, &#39;penciller&#39;, and &#39;letterer&#39;; or an athlete in a SportsTeam might play in the position named &#39;Quarterback&#39;. Supersedes namedPosition (see: https://schema.org/namedPosition).
	 * see : https://schema.org/roleName
	 * @var string|string[]|string|string[]
	 */
	public var $role_name;
	
	/**
	 * The start date and time of the item (in ISO 8601 date format (see: https://schema.orghttp://en.wikipedia.org/wiki/ISO_8601)).
	 * see : https://schema.org/startDate
	 * @var string|string[]|string|string[]
	 */
	public var $start_date;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Role'
		);
		
		$serialized = so_json_serialize( $this->end_date );
		if ( ! empty( $serialized ) ) {
			$out['endDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->role_name );
		if ( ! empty( $serialized ) ) {
			$out['roleName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->start_date );
		if ( ! empty( $serialized ) ) {
			$out['startDate'] = $serialized;
		}
		
		return $out;
	}
}
