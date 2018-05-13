<?php

class PerformingGroup extends Organization implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PerformingGroup';
	
	/**
	 * For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) or other news-related Organization (see: https://schema.org/Organization), a statement about public engagement activities (for news media, the newsroom’s), including involving the public - digitally or otherwise -- in coverage decisions, reporting and activities after publication.
	 * see : http://pending.schema.org/actionableFeedbackPolicy
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $actionable_feedback_policy;
	
	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 * @var \PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $address;
	
	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 * @var \AggregateRating|\AggregateRating[]
	 */
	public var $aggregate_rating;
	
	/**
	 * Alumni of an organization. Inverse property: alumniOf (see: https://schema.org/alumniOf).
	 * see : https://schema.org/alumni
	 * @var \Person|\Person[]
	 */
	public var $alumni;
	
	/**
	 * The geographic area where a service or offered item is provided. Supersedes serviceArea (see: https://schema.org/serviceArea).
	 * see : https://schema.org/areaServed
	 * @var \AdministrativeArea|\AdministrativeArea[]|\GeoShape|\GeoShape[]|\Place|\Place[]|string|string[]
	 */
	public var $area_served;
	
	/**
	 * An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	 * see : https://schema.org/award
	 * @var string|string[]
	 */
	public var $award;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand|\Brand[]|\Organization|\Organization[]
	 */
	public var $brand;
	
	/**
	 * A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	 * see : https://schema.org/contactPoint
	 * @var \ContactPoint|\ContactPoint[]
	 */
	public var $contact_point;
	
	/**
	 * For an Organization (see: https://schema.org/Organization) (e.g. NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement describing (in news media, the newsroom’s) disclosure and correction policy for errors.
	 * see : http://pending.schema.org/correctionsPolicy
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $corrections_policy;
	
	/**
	 * A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	 * see : https://schema.org/department
	 * @var \Organization|\Organization[]
	 */
	public var $department;
	
	/**
	 * The date that this organization was dissolved.
	 * see : https://schema.org/dissolutionDate
	 * @var string|string[]
	 */
	public var $dissolution_date;
	
	/**
	 * Statement on diversity policy by an Organization (see: https://schema.org/Organization) e.g. a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization). For a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), a statement describing the newsroom’s diversity policy on both staffing and sources, typically providing staffing data.
	 * see : http://pending.schema.org/diversityPolicy
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $diversity_policy;
	
	/**
	 * The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	 * see : https://schema.org/duns
	 * @var string|string[]
	 */
	public var $duns;
	
	/**
	 * Email address.
	 * see : https://schema.org/email
	 * @var string|string[]
	 */
	public var $email;
	
	/**
	 * Someone working for this organization. Supersedes employees (see: https://schema.org/employees).
	 * see : https://schema.org/employee
	 * @var \Person|\Person[]
	 */
	public var $employee;
	
	/**
	 * Statement about ethics policy, e.g. of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization) regarding journalistic and publishing practices, or of a Restaurant (see: https://schema.org/Restaurant), a page describing food source policies. In the case of a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization), an ethicsPolicy is typically a statement describing the personal, organizational, and corporate standards of behavior expected by the organization.
	 * see : http://pending.schema.org/ethicsPolicy
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $ethics_policy;
	
	/**
	 * Upcoming or past event associated with this place, organization, or action. Supersedes events (see: https://schema.org/events).
	 * see : https://schema.org/event
	 * @var \Event|\Event[]
	 */
	public var $event;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string|string[]
	 */
	public var $fax_number;
	
	/**
	 * A person who founded this organization. Supersedes founders (see: https://schema.org/founders).
	 * see : https://schema.org/founder
	 * @var \Person|\Person[]
	 */
	public var $founder;
	
	/**
	 * The date that this organization was founded.
	 * see : https://schema.org/foundingDate
	 * @var string|string[]
	 */
	public var $founding_date;
	
	/**
	 * The place where the Organization was founded.
	 * see : https://schema.org/foundingLocation
	 * @var \Place|\Place[]
	 */
	public var $founding_location;
	
	/**
	 * A person or organization that supports (sponsors) something through some kind of financial contribution.
	 * see : https://schema.org/funder
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $funder;
	
	/**
	 * The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 * @var string|string[]
	 */
	public var $global_location_number;
	
	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 * @var \OfferCatalog|\OfferCatalog[]
	 */
	public var $has_offer_catalog;
	
	/**
	 * Points-of-Sales operated by the organization or person.
	 * see : https://schema.org/hasPOS
	 * @var \Place|\Place[]
	 */
	public var $haspo_s;
	
	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 * @var string|string[]
	 */
	public var $isic_v_4;
	
	/**
	 * The official name of the organization, e.g. the registered company name.
	 * see : https://schema.org/legalName
	 * @var string|string[]
	 */
	public var $legal_name;
	
	/**
	 * An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
	 * see : https://schema.org/leiCode
	 * @var string|string[]
	 */
	public var $lei_code;
	
	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 * @var \Place|\Place[]|\PostalAddress|\PostalAddress[]|string|string[]
	 */
	public var $location;
	
	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 * @var \ImageObject|\ImageObject[]|string|string[]
	 */
	public var $logo;
	
	/**
	 * A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	 * see : https://schema.org/makesOffer
	 * @var \Offer|\Offer[]
	 */
	public var $makes_offer;
	
	/**
	 * A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals. Supersedes members (see: https://schema.org/members), musicGroupMember (see: https://schema.org/musicGroupMember). Inverse property: memberOf (see: https://schema.org/memberOf).
	 * see : https://schema.org/member
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $member;
	
	/**
	 * An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	 * see : https://schema.org/memberOf
	 * @var \Organization|\Organization[]|\ProgramMembership|\ProgramMembership[]
	 */
	public var $member_of;
	
	/**
	 * The North American Industry Classification System (NAICS) code for a particular organization or business person.
	 * see : https://schema.org/naics
	 * @var string|string[]
	 */
	public var $naics;
	
	/**
	 * The number of employees in an organization e.g. business.
	 * see : https://schema.org/numberOfEmployees
	 * @var \QuantitativeValue|\QuantitativeValue[]
	 */
	public var $number_of_employees;
	
	/**
	 * Products owned by the organization or person.
	 * see : https://schema.org/owns
	 * @var \OwnershipInfo|\OwnershipInfo[]|\Product|\Product[]
	 */
	public var $owns;
	
	/**
	 * The larger organization that this organization is a subOrganization (see: https://schema.org/subOrganization) of, if any. Supersedes branchOf (see: https://schema.org/branchOf). Inverse property: subOrganization (see: https://schema.org/subOrganization).
	 * see : https://schema.org/parentOrganization
	 * @var \Organization|\Organization[]
	 */
	public var $parent_organization;
	
	/**
	 * The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	 * 
	 * While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	 * see : https://schema.org/publishingPrinciples
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $publishing_principles;
	
	/**
	 * A review of the item. Supersedes reviews (see: https://schema.org/reviews).
	 * see : https://schema.org/review
	 * @var \Review|\Review[]
	 */
	public var $review;
	
	/**
	 * A pointer to products or services sought by the organization or person (demand).
	 * see : https://schema.org/seeks
	 * @var \Demand|\Demand[]
	 */
	public var $seeks;
	
	/**
	 * A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	 * see : https://schema.org/sponsor
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $sponsor;
	
	/**
	 * A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific &#39;department&#39; property. Inverse property: parentOrganization (see: https://schema.org/parentOrganization).
	 * see : https://schema.org/subOrganization
	 * @var \Organization|\Organization[]
	 */
	public var $sub_organization;
	
	/**
	 * The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	 * see : https://schema.org/taxID
	 * @var string|string[]
	 */
	public var $taxi_d;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string|string[]
	 */
	public var $telephone;
	
	/**
	 * For an Organization (see: https://schema.org/Organization) (typically a NewsMediaOrganization (see: https://schema.org/NewsMediaOrganization)), a statement about policy on use of unnamed sources and the decision process required.
	 * see : http://pending.schema.org/unnamedSourcesPolicy
	 * @var \CreativeWork|\CreativeWork[]|string|string[]
	 */
	public var $unnamed_sources_policy;
	
	/**
	 * The Value-added Tax ID of the organization or person.
	 * see : https://schema.org/vatID
	 * @var string|string[]
	 */
	public var $vati_d;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PerformingGroup'
		);
		
		$serialized = so_json_serialize( $this->actionable_feedback_policy );
		if ( ! empty( $serialized ) ) {
			$out['actionableFeedbackPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->alumni );
		if ( ! empty( $serialized ) ) {
			$out['alumni'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contact_point );
		if ( ! empty( $serialized ) ) {
			$out['contactPoint'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->corrections_policy );
		if ( ! empty( $serialized ) ) {
			$out['correctionsPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->department );
		if ( ! empty( $serialized ) ) {
			$out['department'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->dissolution_date );
		if ( ! empty( $serialized ) ) {
			$out['dissolutionDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->diversity_policy );
		if ( ! empty( $serialized ) ) {
			$out['diversityPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->duns );
		if ( ! empty( $serialized ) ) {
			$out['duns'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->employee );
		if ( ! empty( $serialized ) ) {
			$out['employee'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->ethics_policy );
		if ( ! empty( $serialized ) ) {
			$out['ethicsPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->event );
		if ( ! empty( $serialized ) ) {
			$out['event'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->founder );
		if ( ! empty( $serialized ) ) {
			$out['founder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->founding_date );
		if ( ! empty( $serialized ) ) {
			$out['foundingDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->founding_location );
		if ( ! empty( $serialized ) ) {
			$out['foundingLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->funder );
		if ( ! empty( $serialized ) ) {
			$out['funder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->global_location_number );
		if ( ! empty( $serialized ) ) {
			$out['globalLocationNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->haspo_s );
		if ( ! empty( $serialized ) ) {
			$out['hasPOS'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->isic_v_4 );
		if ( ! empty( $serialized ) ) {
			$out['isicV4'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->legal_name );
		if ( ! empty( $serialized ) ) {
			$out['legalName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->lei_code );
		if ( ! empty( $serialized ) ) {
			$out['leiCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->makes_offer );
		if ( ! empty( $serialized ) ) {
			$out['makesOffer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->member );
		if ( ! empty( $serialized ) ) {
			$out['member'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->member_of );
		if ( ! empty( $serialized ) ) {
			$out['memberOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->naics );
		if ( ! empty( $serialized ) ) {
			$out['naics'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->number_of_employees );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEmployees'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->owns );
		if ( ! empty( $serialized ) ) {
			$out['owns'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->parent_organization );
		if ( ! empty( $serialized ) ) {
			$out['parentOrganization'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->publishing_principles );
		if ( ! empty( $serialized ) ) {
			$out['publishingPrinciples'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seeks );
		if ( ! empty( $serialized ) ) {
			$out['seeks'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sponsor );
		if ( ! empty( $serialized ) ) {
			$out['sponsor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sub_organization );
		if ( ! empty( $serialized ) ) {
			$out['subOrganization'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->taxi_d );
		if ( ! empty( $serialized ) ) {
			$out['taxID'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->unnamed_sources_policy );
		if ( ! empty( $serialized ) ) {
			$out['unnamedSourcesPolicy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vati_d );
		if ( ! empty( $serialized ) ) {
			$out['vatID'] = $serialized;
		}
		
		return $out;
	}
}
