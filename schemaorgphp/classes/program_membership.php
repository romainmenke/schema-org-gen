<?php

class ProgramMembership extends Intangible implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ProgramMembership';
	
	/**
	 * The organization (airline, travelers&#39; club, etc.) the membership is made with.
	 * see : https://schema.org/hostingOrganization
	 * @var \Organization|\Organization[]
	 */
	public var $hosting_organization;
	
	/**
	 * A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	 * see : https://schema.org/member
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $member;
	
	/**
	 * A unique identifier for the membership.
	 * see : https://schema.org/membershipNumber
	 * @var string|string[]
	 */
	public var $membership_number;
	
	/**
	 * The program providing the membership.
	 * see : https://schema.org/programName
	 * @var string|string[]
	 */
	public var $program_name;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ProgramMembership'
		);
		
		$serialized = so_json_serialize( $this->hosting_organization );
		if ( ! empty( $serialized ) ) {
			$out['hostingOrganization'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->member );
		if ( ! empty( $serialized ) ) {
			$out['member'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->membership_number );
		if ( ! empty( $serialized ) ) {
			$out['membershipNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->program_name );
		if ( ! empty( $serialized ) ) {
			$out['programName'] = $serialized;
		}
		
		return $out;
	}
}
