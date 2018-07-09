<?php

// Person see : https://schema.org/Person
class Person implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Person';
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional name for a Person, can be used for a middle name.
	 * see : https://schema.org/additionalName
	 * @var string | string[]
	 */
	public var $additional_name;
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * Physical address of the item.
	 * see : https://schema.org/address
	 * @var \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public var $address;
	
	/**
	 * An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	 * see : https://schema.org/affiliation
	 * @var \Organization | \Organization[]
	 */
	public var $affiliation;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
	/**
	 * An organization that the person is an alumni of. Inverse property: alumni (see: https://schema.org/alumni).
	 * see : https://schema.org/alumniOf
	 * @var \EducationalOrganization | \EducationalOrganization[] | \Organization | \Organization[]
	 */
	public var $alumni_of;
	
	/**
	 * An award won by or for this item. Supersedes awards (see: https://schema.org/awards).
	 * see : https://schema.org/award
	 * @var string | string[]
	 */
	public var $award;
	
	/**
	 * Date of birth.
	 * see : https://schema.org/birthDate
	 * @var string | string[]
	 */
	public var $birth_date;
	
	/**
	 * The place where the person was born.
	 * see : https://schema.org/birthPlace
	 * @var \Place | \Place[]
	 */
	public var $birth_place;
	
	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public var $brand;
	
	/**
	 * A child of the person.
	 * see : https://schema.org/children
	 * @var \Person | \Person[]
	 */
	public var $children;
	
	/**
	 * A colleague of the person. Supersedes colleagues (see: https://schema.org/colleagues).
	 * see : https://schema.org/colleague
	 * @var \Person | \Person[] | string | string[]
	 */
	public var $colleague;
	
	/**
	 * A contact point for a person or organization. Supersedes contactPoints (see: https://schema.org/contactPoints).
	 * see : https://schema.org/contactPoint
	 * @var \ContactPoint | \ContactPoint[]
	 */
	public var $contact_point;
	
	/**
	 * Date of death.
	 * see : https://schema.org/deathDate
	 * @var string | string[]
	 */
	public var $death_date;
	
	/**
	 * The place where the person died.
	 * see : https://schema.org/deathPlace
	 * @var \Place | \Place[]
	 */
	public var $death_place;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public var $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public var $disambiguating_description;
	
	/**
	 * The Dun &amp; Bradstreet DUNS number for identifying an organization or business person.
	 * see : https://schema.org/duns
	 * @var string | string[]
	 */
	public var $duns;
	
	/**
	 * Email address.
	 * see : https://schema.org/email
	 * @var string | string[]
	 */
	public var $email;
	
	/**
	 * Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
	 * see : https://schema.org/familyName
	 * @var string | string[]
	 */
	public var $family_name;
	
	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 * @var string | string[]
	 */
	public var $fax_number;
	
	/**
	 * The most generic uni-directional social relation.
	 * see : https://schema.org/follows
	 * @var \Person | \Person[]
	 */
	public var $follows;
	
	/**
	 * A person or organization that supports (sponsors) something through some kind of financial contribution.
	 * see : https://schema.org/funder
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $funder;
	
	/**
	 * Gender of the person. While http://schema.org/Male and http://schema.org/Female may be used, text strings are also acceptable for people who do not identify as a binary gender.
	 * see : https://schema.org/gender
	 * @var \GenderType | \GenderType[] | string | string[]
	 */
	public var $gender;
	
	/**
	 * Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	 * see : https://schema.org/givenName
	 * @var string | string[]
	 */
	public var $given_name;
	
	/**
	 * The Global Location Number (see: https://schema.orghttp://www.gs1.org/gln) (GLN, sometimes also referred to as International Location Number or ILN) of the respective organization, person, or place. The GLN is a 13-digit number used to identify parties and physical locations.
	 * see : https://schema.org/globalLocationNumber
	 * @var string | string[]
	 */
	public var $global_location_number;
	
	/**
	 * The Person&#39;s occupation. For past professions, use Role for expressing dates.
	 * see : https://pending.schema.org/hasOccupation
	 * @var \Occupation | \Occupation[]
	 */
	public var $has_occupation;
	
	/**
	 * Indicates an OfferCatalog listing for this Organization, Person, or Service.
	 * see : https://schema.org/hasOfferCatalog
	 * @var \OfferCatalog | \OfferCatalog[]
	 */
	public var $has_offer_catalog;
	
	/**
	 * Points-of-Sales operated by the organization or person.
	 * see : https://schema.org/hasPOS
	 * @var \Place | \Place[]
	 */
	public var $haspo_s;
	
	/**
	 * The height of the item.
	 * see : https://schema.org/height
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $height;
	
	/**
	 * A contact location for a person&#39;s residence.
	 * see : https://schema.org/homeLocation
	 * @var \ContactPoint | \ContactPoint[] | \Place | \Place[]
	 */
	public var $home_location;
	
	/**
	 * An honorific prefix preceding a Person&#39;s name such as Dr/Mrs/Mr.
	 * see : https://schema.org/honorificPrefix
	 * @var string | string[]
	 */
	public var $honorific_prefix;
	
	/**
	 * An honorific suffix preceding a Person&#39;s name such as M.D. /PhD/MSCSW.
	 * see : https://schema.org/honorificSuffix
	 * @var string | string[]
	 */
	public var $honorific_suffix;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public var $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public var $image;
	
	/**
	 * The International Standard of Industrial Classification of All Economic Activities (ISIC), Revision 4 code for a particular organization, business person, or place.
	 * see : https://schema.org/isicV4
	 * @var string | string[]
	 */
	public var $isic_v_4;
	
	/**
	 * The job title of the person (for example, Financial Manager).
	 * see : https://schema.org/jobTitle
	 * @var string | string[]
	 */
	public var $job_title;
	
	/**
	 * The most generic bi-directional social/work relation.
	 * see : https://schema.org/knows
	 * @var \Person | \Person[]
	 */
	public var $knows;
	
	/**
	 * Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a topic that is known about - suggesting possible expertise but not implying it. We do not distinguish skill levels here, or yet relate this to educational content, events, objectives or JobPosting (see: https://schema.org/JobPosting) descriptions.
	 * see : https://pending.schema.org/knowsAbout
	 * @var string | string[] | \Thing | \Thing[]
	 */
	public var $knows_about;
	
	/**
	 * Of a Person (see: https://schema.org/Person), and less typically of an Organization (see: https://schema.org/Organization), to indicate a known language. We do not distinguish skill levels or reading/writing/speaking/signing here. Use language codes from the IETF BCP 47 standard (see: https://schema.orghttp://tools.ietf.org/html/bcp47).
	 * see : https://pending.schema.org/knowsLanguage
	 * @var \Language | \Language[] | string | string[]
	 */
	public var $knows_language;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * A pointer to products or services offered by the organization or person. Inverse property: offeredBy (see: https://schema.org/offeredBy).
	 * see : https://schema.org/makesOffer
	 * @var \Offer | \Offer[]
	 */
	public var $makes_offer;
	
	/**
	 * An Organization (or ProgramMembership) to which this Person or Organization belongs. Inverse property: member (see: https://schema.org/member).
	 * see : https://schema.org/memberOf
	 * @var \Organization | \Organization[] | \ProgramMembership | \ProgramMembership[]
	 */
	public var $member_of;
	
	/**
	 * The North American Industry Classification System (NAICS) code for a particular organization or business person.
	 * see : https://schema.org/naics
	 * @var string | string[]
	 */
	public var $naics;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * Nationality of the person.
	 * see : https://schema.org/nationality
	 * @var \Country | \Country[]
	 */
	public var $nationality;
	
	/**
	 * The total financial value of the person as calculated by subtracting assets from liabilities.
	 * see : https://schema.org/netWorth
	 * @var \MonetaryAmount | \MonetaryAmount[] | \PriceSpecification | \PriceSpecification[]
	 */
	public var $net_worth;
	
	/**
	 * Products owned by the organization or person.
	 * see : https://schema.org/owns
	 * @var \OwnershipInfo | \OwnershipInfo[] | \Product | \Product[]
	 */
	public var $owns;
	
	/**
	 * A parent of this person. Supersedes parents (see: https://schema.org/parents).
	 * see : https://schema.org/parent
	 * @var \Person | \Person[]
	 */
	public var $parent;
	
	/**
	 * Event that this person is a performer or participant in.
	 * see : https://schema.org/performerIn
	 * @var \Event | \Event[]
	 */
	public var $performer_in;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The publishingPrinciples property indicates (typically via URL (see: https://schema.org/URL)) a document describing the editorial principles of an Organization (see: https://schema.org/Organization) (or individual e.g. a Person (see: https://schema.org/Person) writing a blog) that relate to their activities as a publisher, e.g. ethics or diversity policies. When applied to a CreativeWork (see: https://schema.org/CreativeWork) (e.g. NewsArticle (see: https://schema.org/NewsArticle)) the principles are those of the party primarily responsible for the creation of the CreativeWork (see: https://schema.org/CreativeWork).
	 * 
	 * While such policies are most typically expressed in natural language, sometimes related information (e.g. indicating a funder (see: https://schema.org/funder)) can be expressed using schema.org terminology.
	 * see : https://schema.org/publishingPrinciples
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $publishing_principles;
	
	/**
	 * The most generic familial relation.
	 * see : https://schema.org/relatedTo
	 * @var \Person | \Person[]
	 */
	public var $related_to;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * A pointer to products or services sought by the organization or person (demand).
	 * see : https://schema.org/seeks
	 * @var \Demand | \Demand[]
	 */
	public var $seeks;
	
	/**
	 * A sibling of the person. Supersedes siblings (see: https://schema.org/siblings).
	 * see : https://schema.org/sibling
	 * @var \Person | \Person[]
	 */
	public var $sibling;
	
	/**
	 * A person or organization that supports a thing through a pledge, promise, or financial contribution. e.g. a sponsor of a Medical Study or a corporate sponsor of an event.
	 * see : https://schema.org/sponsor
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $sponsor;
	
	/**
	 * The person&#39;s spouse.
	 * see : https://schema.org/spouse
	 * @var \Person | \Person[]
	 */
	public var $spouse;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * The Tax / Fiscal ID of the organization or person, e.g. the TIN in the US or the CIF/NIF in Spain.
	 * see : https://schema.org/taxID
	 * @var string | string[]
	 */
	public var $taxi_d;
	
	/**
	 * The telephone number.
	 * see : https://schema.org/telephone
	 * @var string | string[]
	 */
	public var $telephone;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	/**
	 * The Value-added Tax ID of the organization or person.
	 * see : https://schema.org/vatID
	 * @var string | string[]
	 */
	public var $vati_d;
	
	/**
	 * The weight of the product or person.
	 * see : https://schema.org/weight
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public var $weight;
	
	/**
	 * A contact location for a person&#39;s place of work.
	 * see : https://schema.org/workLocation
	 * @var \ContactPoint | \ContactPoint[] | \Place | \Place[]
	 */
	public var $work_location;
	
	/**
	 * Organizations that the person works for.
	 * see : https://schema.org/worksFor
	 * @var \Organization | \Organization[]
	 */
	public var $works_for;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Person'
		);
		
		$serialized = so_json_serialize( $this->additional_name );
		if ( ! empty( $serialized ) ) {
			$out['additionalName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->affiliation );
		if ( ! empty( $serialized ) ) {
			$out['affiliation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->alumni_of );
		if ( ! empty( $serialized ) ) {
			$out['alumniOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->birth_date );
		if ( ! empty( $serialized ) ) {
			$out['birthDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->birth_place );
		if ( ! empty( $serialized ) ) {
			$out['birthPlace'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->children );
		if ( ! empty( $serialized ) ) {
			$out['children'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->colleague );
		if ( ! empty( $serialized ) ) {
			$out['colleague'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->contact_point );
		if ( ! empty( $serialized ) ) {
			$out['contactPoint'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->death_date );
		if ( ! empty( $serialized ) ) {
			$out['deathDate'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->death_place );
		if ( ! empty( $serialized ) ) {
			$out['deathPlace'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->duns );
		if ( ! empty( $serialized ) ) {
			$out['duns'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->family_name );
		if ( ! empty( $serialized ) ) {
			$out['familyName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->follows );
		if ( ! empty( $serialized ) ) {
			$out['follows'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->funder );
		if ( ! empty( $serialized ) ) {
			$out['funder'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->gender );
		if ( ! empty( $serialized ) ) {
			$out['gender'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->given_name );
		if ( ! empty( $serialized ) ) {
			$out['givenName'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->global_location_number );
		if ( ! empty( $serialized ) ) {
			$out['globalLocationNumber'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_occupation );
		if ( ! empty( $serialized ) ) {
			$out['hasOccupation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_offer_catalog );
		if ( ! empty( $serialized ) ) {
			$out['hasOfferCatalog'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->haspo_s );
		if ( ! empty( $serialized ) ) {
			$out['hasPOS'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->home_location );
		if ( ! empty( $serialized ) ) {
			$out['homeLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->honorific_prefix );
		if ( ! empty( $serialized ) ) {
			$out['honorificPrefix'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->honorific_suffix );
		if ( ! empty( $serialized ) ) {
			$out['honorificSuffix'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->isic_v_4 );
		if ( ! empty( $serialized ) ) {
			$out['isicV4'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->job_title );
		if ( ! empty( $serialized ) ) {
			$out['jobTitle'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->knows );
		if ( ! empty( $serialized ) ) {
			$out['knows'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->knows_about );
		if ( ! empty( $serialized ) ) {
			$out['knowsAbout'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->knows_language );
		if ( ! empty( $serialized ) ) {
			$out['knowsLanguage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->makes_offer );
		if ( ! empty( $serialized ) ) {
			$out['makesOffer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->member_of );
		if ( ! empty( $serialized ) ) {
			$out['memberOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->naics );
		if ( ! empty( $serialized ) ) {
			$out['naics'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->nationality );
		if ( ! empty( $serialized ) ) {
			$out['nationality'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->net_worth );
		if ( ! empty( $serialized ) ) {
			$out['netWorth'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->owns );
		if ( ! empty( $serialized ) ) {
			$out['owns'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->parent );
		if ( ! empty( $serialized ) ) {
			$out['parent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->performer_in );
		if ( ! empty( $serialized ) ) {
			$out['performerIn'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->publishing_principles );
		if ( ! empty( $serialized ) ) {
			$out['publishingPrinciples'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->related_to );
		if ( ! empty( $serialized ) ) {
			$out['relatedTo'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->seeks );
		if ( ! empty( $serialized ) ) {
			$out['seeks'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sibling );
		if ( ! empty( $serialized ) ) {
			$out['sibling'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sponsor );
		if ( ! empty( $serialized ) ) {
			$out['sponsor'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->spouse );
		if ( ! empty( $serialized ) ) {
			$out['spouse'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->taxi_d );
		if ( ! empty( $serialized ) ) {
			$out['taxID'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->telephone );
		if ( ! empty( $serialized ) ) {
			$out['telephone'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->vati_d );
		if ( ! empty( $serialized ) ) {
			$out['vatID'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->weight );
		if ( ! empty( $serialized ) ) {
			$out['weight'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->work_location );
		if ( ! empty( $serialized ) ) {
			$out['workLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->works_for );
		if ( ! empty( $serialized ) ) {
			$out['worksFor'] = $serialized;
		}
		
		return $out;
	}
}
