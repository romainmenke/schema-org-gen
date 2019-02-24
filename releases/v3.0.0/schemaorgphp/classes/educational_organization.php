<?php
namespace SchemaOrg;

// EducationalOrganization see : https://schema.org/EducationalOrganization
class EducationalOrganization implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'EducationalOrganization';

	/**
	 * With properties from Organization see : https://schema.org/Organization
	 */

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 *
	 * @var string | string[]
	 */
	public $additional_type;

	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 *
	 * @var \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $address;

	/**
	 * The overall rating, based on a collection of reviews or ratings, of the item.
	 * see : https://schema.org/aggregateRating
	 *
	 * @var \AggregateRating | \AggregateRating[]
	 */
	public $aggregate_rating;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * Alumni of an organization.
	 * see : https://schema.org/alumni
	 *
	 * @var \Person | \Person[]
	 */
	public $alumni;

	/**
	 * The geographic area where a service or offered item is provided.
	 * see : https://schema.org/areaServed
	 *
	 * @var \Place | \Place[] | \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[] | string | string[]
	 */
	public $area_served;

	/**
	 * An award won by or for this item.
	 * see : https://schema.org/award
	 *
	 * @var string | string[]
	 */
	public $award;

	/**
	 * Awards won by or for this item.
	 * see : https://schema.org/awards
	 *
	 * @var string | string[]
	 */
	public $awards;

	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 *
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public $brand;

	/**
	 * A contact point for a person or organization.
	 * see : https://schema.org/contactPoint
	 *
	 * @var \ContactPoint | \ContactPoint[]
	 */
	public $contact_point;

	/**
	 * A contact point for a person or organization.
	 * see : https://schema.org/contactPoints
	 *
	 * @var \ContactPoint | \ContactPoint[]
	 */
	public $contact_points;

	/**
	 * A relationship between an organization and a department of that organization, also described as an organization (allowing different urls, logos, opening hours). For example: a store with a pharmacy, or a bakery with a cafe.
	 * see : https://schema.org/department
	 *
	 * @var \Organization | \Organization[]
	 */
	public $department;

	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 *
	 * @var string | string[]
	 */
	public $disambiguating_description;

	/**
	 * The date that this organization was dissolved.
	 * see : https://schema.org/dissolutionDate
	 *
	 * @var string | string[]
	 */
	public $dissolution_date;

	/**
	 * The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	 * see : https://schema.org/duns
	 *
	 * @var string | string[]
	 */
	public $duns;

	/**
	 * Email address.
	 * see : https://schema.org/email
	 *
	 * @var string | string[]
	 */
	public $email;

	/**
	 * Someone working for this organization.
	 * see : https://schema.org/employee
	 *
	 * @var \Person | \Person[]
	 */
	public $employee;

	/**
	 * People working for this organization.
	 * see : https://schema.org/employees
	 *
	 * @var \Person | \Person[]
	 */
	public $employees;

	/**
	 * Upcoming or past event associated with this place, organization, or action.
	 * see : https://schema.org/event
	 *
	 * @var \Event | \Event[]
	 */
	public $event;

	/**
	 * Upcoming or past events associated with this place or organization.
	 * see : https://schema.org/events
	 *
	 * @var \Event | \Event[]
	 */
	public $events;

	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 *
	 * @var string | string[]
	 */
	public $fax_number;

	/**
	 * A person who founded this organization.
	 * see : https://schema.org/founder
	 *
	 * @var \Person | \Person[]
	 */
	public $founder;

	/**
	 * A person who founded this organization.
	 * see : https://schema.org/founders
	 *
	 * @var \Person | \Person[]
	 */
	public $founders;

	/**
	 * The date that this organization was founded.
	 * see : https://schema.org/foundingDate
	 *
	 * @var string | string[]
	 */
	public $founding_date;

	/**
	 * The place where the Organization was founded.
	 * see : https://schema.org/foundingLocation
	 *
	 * @var \Place | \Place[]
	 */
	public $founding_location;

	/**
	 * The &lt;a href=&quot;http://www.gs1.org/gln&quot;&gt;Global Location Number&lt;/a&gt; (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 *
	 * @var string | string[]
	 */
	public $global_location_number;

	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 *
	 * @var \OfferCatalog | \OfferCatalog[]
	 */
	public $has_offer_catalog;

	/**
	 * Points-of-Sales operated by the organization or person.
	 * see : https://schema.org/hasPOS
	 *
	 * @var \Place | \Place[]
	 */
	public $has_pos;

	/**
	 * An image of the item. This can be a &lt;a href=&quot;http://schema.org/URL&quot;&gt;URL&lt;/a&gt; or a fully described &lt;a href=&quot;http://schema.org/ImageObject&quot;&gt;ImageObject&lt;/a&gt;.
	 * see : https://schema.org/image
	 *
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;

	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 *
	 * @var string | string[]
	 */
	public $isic_v_4;

	/**
	 * The official name of the organization, e.g. the registered company name.
	 * see : https://schema.org/legalName
	 *
	 * @var string | string[]
	 */
	public $legal_name;

	/**
	 * An organization identifier that uniquely identifies a legal entity as defined in ISO 17442.
	 * see : https://schema.org/leiCode
	 *
	 * @var string | string[]
	 */
	public $lei_code;

	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 *
	 * @var \Place | \Place[] | \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public $location;

	/**
	 * An associated logo.
	 * see : https://schema.org/logo
	 *
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $logo;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	 * see : https://schema.org/mainEntityOfPage
	 *
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;

	/**
	 * A pointer to products or services offered by the organization or person.
	 * see : https://schema.org/makesOffer
	 *
	 * @var \Offer | \Offer[]
	 */
	public $makes_offer;

	/**
	 * A member of an Organization or a ProgramMembership. Organizations can be members of organizations; ProgramMembership is typically for individuals.
	 * see : https://schema.org/member
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $member;

	/**
	 * An Organization (or ProgramMembership) to which this Person or Organization belongs.
	 * see : https://schema.org/memberOf
	 *
	 * @var \Organization | \Organization[] | \ProgramMembership | \ProgramMembership[]
	 */
	public $member_of;

	/**
	 * A member of this organization.
	 * see : https://schema.org/members
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $members;

	/**
	 * The North American Industry Classification System (NAICS) code for a particular organization or business person.
	 * see : https://schema.org/naics
	 *
	 * @var string | string[]
	 */
	public $naics;

	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 *
	 * @var string | string[]
	 */
	public $name;

	/**
	 * The number of employees in an organization e.g. business.
	 * see : https://schema.org/numberOfEmployees
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $number_of_employees;

	/**
	 * A pointer to the organization or person making the offer.
	 * see : https://schema.org/offeredBy
	 *
	 * @var \Person | \Person[] | \Offer | \Offer[]
	 */
	public $offered_by;

	/**
	 * Products owned by the organization or person.
	 * see : https://schema.org/owns
	 *
	 * @var \OwnershipInfo | \OwnershipInfo[] | \Product | \Product[]
	 */
	public $owns;

	/**
	 * The larger organization that this local business is a branch of, if any.
	 * see : https://schema.org/parentOrganization
	 *
	 * @var \Organization | \Organization[]
	 */
	public $parent_organization;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * A review of the item.
	 * see : https://schema.org/review
	 *
	 * @var \Review | \Review[]
	 */
	public $review;

	/**
	 * Review of the item.
	 * see : https://schema.org/reviews
	 *
	 * @var \Review | \Review[]
	 */
	public $reviews;

	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Freebase page, or official website.
	 * see : https://schema.org/sameAs
	 *
	 * @var string | string[]
	 */
	public $same_as;

	/**
	 * A pointer to products or services sought by the organization or person (demand).
	 * see : https://schema.org/seeks
	 *
	 * @var \Demand | \Demand[]
	 */
	public $seeks;

	/**
	 * The geographic area where the service is provided.
	 * see : https://schema.org/serviceArea
	 *
	 * @var \Place | \Place[] | \AdministrativeArea | \AdministrativeArea[] | \GeoShape | \GeoShape[]
	 */
	public $service_area;

	/**
	 * A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	 * see : https://schema.org/sponsor
	 *
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public $sponsor;

	/**
	 * A relationship between two organizations where the first includes the second, e.g., as a subsidiary. See also: the more specific &#39;department&#39; property.
	 * see : https://schema.org/subOrganization
	 *
	 * @var \Organization | \Organization[]
	 */
	public $sub_organization;

	/**
	 * The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	 * see : https://schema.org/taxID
	 *
	 * @var string | string[]
	 */
	public $tax_id;

	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 *
	 * @var string | string[]
	 */
	public $telephone;

	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 *
	 * @var string | string[]
	 */
	public $url;

	/**
	 * The Value-added Tax ID of the organization or person.
	 * see : https://schema.org/vatID
	 *
	 * @var string | string[]
	 */
	public $vat_id;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'EducationalOrganization',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->aggregate_rating );
		if ( ! empty( $serialized ) ) {
			$out['aggregateRating'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alumni );
		if ( ! empty( $serialized ) ) {
			$out['alumni'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->area_served );
		if ( ! empty( $serialized ) ) {
			$out['areaServed'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->awards );
		if ( ! empty( $serialized ) ) {
			$out['awards'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contact_point );
		if ( ! empty( $serialized ) ) {
			$out['contactPoint'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contact_points );
		if ( ! empty( $serialized ) ) {
			$out['contactPoints'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->department );
		if ( ! empty( $serialized ) ) {
			$out['department'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->dissolution_date );
		if ( ! empty( $serialized ) ) {
			$out['dissolutionDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->duns );
		if ( ! empty( $serialized ) ) {
			$out['duns'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->employee );
		if ( ! empty( $serialized ) ) {
			$out['employee'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->employees );
		if ( ! empty( $serialized ) ) {
			$out['employees'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->event );
		if ( ! empty( $serialized ) ) {
			$out['event'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->events );
		if ( ! empty( $serialized ) ) {
			$out['events'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->founder );
		if ( ! empty( $serialized ) ) {
			$out['founder'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->founders );
		if ( ! empty( $serialized ) ) {
			$out['founders'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->founding_date );
		if ( ! empty( $serialized ) ) {
			$out['foundingDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->founding_location );
		if ( ! empty( $serialized ) ) {
			$out['foundingLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->global_location_number );
		if ( ! empty( $serialized ) ) {
			$out['globalLocationNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->has_pos );
		if ( ! empty( $serialized ) ) {
			$out['hasPOS'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->isic_v_4 );
		if ( ! empty( $serialized ) ) {
			$out['isicV4'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->legal_name );
		if ( ! empty( $serialized ) ) {
			$out['legalName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->lei_code );
		if ( ! empty( $serialized ) ) {
			$out['leiCode'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->logo );
		if ( ! empty( $serialized ) ) {
			$out['logo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->makes_offer );
		if ( ! empty( $serialized ) ) {
			$out['makesOffer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->member );
		if ( ! empty( $serialized ) ) {
			$out['member'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->member_of );
		if ( ! empty( $serialized ) ) {
			$out['memberOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->members );
		if ( ! empty( $serialized ) ) {
			$out['members'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->naics );
		if ( ! empty( $serialized ) ) {
			$out['naics'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->number_of_employees );
		if ( ! empty( $serialized ) ) {
			$out['numberOfEmployees'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->offered_by );
		if ( ! empty( $serialized ) ) {
			$out['offeredBy'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->owns );
		if ( ! empty( $serialized ) ) {
			$out['owns'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->parent_organization );
		if ( ! empty( $serialized ) ) {
			$out['parentOrganization'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->review );
		if ( ! empty( $serialized ) ) {
			$out['review'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->reviews );
		if ( ! empty( $serialized ) ) {
			$out['reviews'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->seeks );
		if ( ! empty( $serialized ) ) {
			$out['seeks'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->service_area );
		if ( ! empty( $serialized ) ) {
			$out['serviceArea'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sponsor );
		if ( ! empty( $serialized ) ) {
			$out['sponsor'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sub_organization );
		if ( ! empty( $serialized ) ) {
			$out['subOrganization'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->tax_id );
		if ( ! empty( $serialized ) ) {
			$out['taxID'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->vat_id );
		if ( ! empty( $serialized ) ) {
			$out['vatID'] = $serialized;
		}

		return $out;
	}
}
