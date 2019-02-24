<?php
namespace SchemaOrg;

// Person see : https://schema.org/Person
class Person implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type    = 'Person';

	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */


	/**
	 * An additional name for a Person, can be used for a middle name.
	 * see : https://schema.org/additionalName
	 *
	 * @var string | string[]
	 */
	public $additional_name;

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
	 * An organization that this person is affiliated with. For example, a school/university, a club, or a team.
	 * see : https://schema.org/affiliation
	 *
	 * @var \Organization | \Organization[]
	 */
	public $affiliation;

	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 *
	 * @var string | string[]
	 */
	public $alternate_name;

	/**
	 * An organization that the person is an alumni of.
	 * see : https://schema.org/alumniOf
	 *
	 * @var \EducationalOrganization | \EducationalOrganization[]
	 */
	public $alumni_of;

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
	 * Date of birth.
	 * see : https://schema.org/birthDate
	 *
	 * @var string | string[]
	 */
	public $birth_date;

	/**
	 * The place where the person was born.
	 * see : https://schema.org/birthPlace
	 *
	 * @var \Place | \Place[]
	 */
	public $birth_place;

	/**
	 * The brand(s) associated with a product or service, or the brand(s) maintained by an organization or business person.
	 * see : https://schema.org/brand
	 *
	 * @var \Brand | \Brand[] | \Organization | \Organization[]
	 */
	public $brand;

	/**
	 * A child of the person.
	 * see : https://schema.org/children
	 *
	 * @var \Person | \Person[]
	 */
	public $children;

	/**
	 * A colleague of the person.
	 * see : https://schema.org/colleague
	 *
	 * @var \Person | \Person[]
	 */
	public $colleague;

	/**
	 * A colleague of the person.
	 * see : https://schema.org/colleagues
	 *
	 * @var \Person | \Person[]
	 */
	public $colleagues;

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
	 * Date of death.
	 * see : https://schema.org/deathDate
	 *
	 * @var string | string[]
	 */
	public $death_date;

	/**
	 * The place where the person died.
	 * see : https://schema.org/deathPlace
	 *
	 * @var \Place | \Place[]
	 */
	public $death_place;

	/**
	 * A short description of the item.
	 * see : https://schema.org/description
	 *
	 * @var string | string[]
	 */
	public $description;

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
	 * Family name. In the U.S., the last name of an Person. This can be used along with givenName instead of the name property.
	 * see : https://schema.org/familyName
	 *
	 * @var string | string[]
	 */
	public $family_name;

	/**
	 * The fax number.
	 * see : https://schema.org/faxNumber
	 *
	 * @var string | string[]
	 */
	public $fax_number;

	/**
	 * The most generic uni-directional social relation.
	 * see : https://schema.org/follows
	 *
	 * @var \Person | \Person[]
	 */
	public $follows;

	/**
	 * Gender of the person.
	 * see : https://schema.org/gender
	 *
	 * @var string | string[]
	 */
	public $gender;

	/**
	 * Given name. In the U.S., the first name of a Person. This can be used along with familyName instead of the name property.
	 * see : https://schema.org/givenName
	 *
	 * @var string | string[]
	 */
	public $given_name;

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
	 * The height of the item.
	 * see : https://schema.org/height
	 *
	 * @var \Distance | \Distance[] | \QuantitativeValue | \QuantitativeValue[]
	 */
	public $height;

	/**
	 * A contact location for a person&#39;s residence.
	 * see : https://schema.org/homeLocation
	 *
	 * @var \ContactPoint | \ContactPoint[] | \Place | \Place[]
	 */
	public $home_location;

	/**
	 * An honorific prefix preceding a Person&#39;s name such as Dr/Mrs/Mr.
	 * see : https://schema.org/honorificPrefix
	 *
	 * @var string | string[]
	 */
	public $honorific_prefix;

	/**
	 * An honorific suffix preceding a Person&#39;s name such as M.D. /PhD/MSCSW.
	 * see : https://schema.org/honorificSuffix
	 *
	 * @var string | string[]
	 */
	public $honorific_suffix;

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
	 * The job title of the person (for example, Financial Manager).
	 * see : https://schema.org/jobTitle
	 *
	 * @var string | string[]
	 */
	public $job_title;

	/**
	 * The most generic bi-directional social/work relation.
	 * see : https://schema.org/knows
	 *
	 * @var \Person | \Person[]
	 */
	public $knows;

	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described.
	 *       &lt;br /&gt;&lt;br /&gt;
	 *       See &lt;a href=&quot;/docs/datamodel.html#mainEntityBackground&quot;&gt;background notes&lt;/a&gt; for details.
	 *
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
	 * An Organization (or ProgramMembership) to which this Person or Organization belongs.
	 * see : https://schema.org/memberOf
	 *
	 * @var \Organization | \Organization[] | \ProgramMembership | \ProgramMembership[]
	 */
	public $member_of;

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
	 * Nationality of the person.
	 * see : https://schema.org/nationality
	 *
	 * @var \Country | \Country[]
	 */
	public $nationality;

	/**
	 * The total financial value of the person as calculated by subtracting assets from liabilities.
	 * see : https://schema.org/netWorth
	 *
	 * @var \PriceSpecification | \PriceSpecification[]
	 */
	public $net_worth;

	/**
	 * Products owned by the organization or person.
	 * see : https://schema.org/owns
	 *
	 * @var \OwnershipInfo | \OwnershipInfo[] | \Product | \Product[]
	 */
	public $owns;

	/**
	 * A parent of this person.
	 * see : https://schema.org/parent
	 *
	 * @var \Person | \Person[]
	 */
	public $parent;

	/**
	 * A parents of the person.
	 * see : https://schema.org/parents
	 *
	 * @var \Person | \Person[]
	 */
	public $parents;

	/**
	 * Event that this person is a performer or participant in.
	 * see : https://schema.org/performerIn
	 *
	 * @var \Event | \Event[]
	 */
	public $performer_in;

	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 *
	 * @var \Action | \Action[]
	 */
	public $potential_action;

	/**
	 * The most generic familial relation.
	 * see : https://schema.org/relatedTo
	 *
	 * @var \Person | \Person[]
	 */
	public $related_to;

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
	 * A sibling of the person.
	 * see : https://schema.org/sibling
	 *
	 * @var \Person | \Person[]
	 */
	public $sibling;

	/**
	 * A sibling of the person.
	 * see : https://schema.org/siblings
	 *
	 * @var \Person | \Person[]
	 */
	public $siblings;

	/**
	 * The person&#39;s spouse.
	 * see : https://schema.org/spouse
	 *
	 * @var \Person | \Person[]
	 */
	public $spouse;

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

	/**
	 * The weight of the product or person.
	 * see : https://schema.org/weight
	 *
	 * @var \QuantitativeValue | \QuantitativeValue[]
	 */
	public $weight;

	/**
	 * A contact location for a person&#39;s place of work.
	 * see : https://schema.org/workLocation
	 *
	 * @var \ContactPoint | \ContactPoint[] | \Place | \Place[]
	 */
	public $work_location;

	/**
	 * Organizations that the person works for.
	 * see : https://schema.org/worksFor
	 *
	 * @var \Organization | \Organization[]
	 */
	public $works_for;

	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type'    => 'Person',
		);

		$serialized = \SchemaOrg\json_serialize( $this->additional_name );
		if ( ! empty( $serialized ) ) {
			$out['additionalName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->address );
		if ( ! empty( $serialized ) ) {
			$out['address'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->affiliation );
		if ( ! empty( $serialized ) ) {
			$out['affiliation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->alumni_of );
		if ( ! empty( $serialized ) ) {
			$out['alumniOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->award );
		if ( ! empty( $serialized ) ) {
			$out['award'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->awards );
		if ( ! empty( $serialized ) ) {
			$out['awards'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->birth_date );
		if ( ! empty( $serialized ) ) {
			$out['birthDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->birth_place );
		if ( ! empty( $serialized ) ) {
			$out['birthPlace'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->brand );
		if ( ! empty( $serialized ) ) {
			$out['brand'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->children );
		if ( ! empty( $serialized ) ) {
			$out['children'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->colleague );
		if ( ! empty( $serialized ) ) {
			$out['colleague'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->colleagues );
		if ( ! empty( $serialized ) ) {
			$out['colleagues'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contact_point );
		if ( ! empty( $serialized ) ) {
			$out['contactPoint'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->contact_points );
		if ( ! empty( $serialized ) ) {
			$out['contactPoints'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->death_date );
		if ( ! empty( $serialized ) ) {
			$out['deathDate'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->death_place );
		if ( ! empty( $serialized ) ) {
			$out['deathPlace'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->duns );
		if ( ! empty( $serialized ) ) {
			$out['duns'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->email );
		if ( ! empty( $serialized ) ) {
			$out['email'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->family_name );
		if ( ! empty( $serialized ) ) {
			$out['familyName'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->fax_number );
		if ( ! empty( $serialized ) ) {
			$out['faxNumber'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->follows );
		if ( ! empty( $serialized ) ) {
			$out['follows'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->gender );
		if ( ! empty( $serialized ) ) {
			$out['gender'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->given_name );
		if ( ! empty( $serialized ) ) {
			$out['givenName'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->height );
		if ( ! empty( $serialized ) ) {
			$out['height'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->home_location );
		if ( ! empty( $serialized ) ) {
			$out['homeLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->honorific_prefix );
		if ( ! empty( $serialized ) ) {
			$out['honorificPrefix'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->honorific_suffix );
		if ( ! empty( $serialized ) ) {
			$out['honorificSuffix'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->isic_v_4 );
		if ( ! empty( $serialized ) ) {
			$out['isicV4'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->job_title );
		if ( ! empty( $serialized ) ) {
			$out['jobTitle'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->knows );
		if ( ! empty( $serialized ) ) {
			$out['knows'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->makes_offer );
		if ( ! empty( $serialized ) ) {
			$out['makesOffer'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->member_of );
		if ( ! empty( $serialized ) ) {
			$out['memberOf'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->naics );
		if ( ! empty( $serialized ) ) {
			$out['naics'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->nationality );
		if ( ! empty( $serialized ) ) {
			$out['nationality'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->net_worth );
		if ( ! empty( $serialized ) ) {
			$out['netWorth'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->owns );
		if ( ! empty( $serialized ) ) {
			$out['owns'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->parent );
		if ( ! empty( $serialized ) ) {
			$out['parent'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->parents );
		if ( ! empty( $serialized ) ) {
			$out['parents'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->performer_in );
		if ( ! empty( $serialized ) ) {
			$out['performerIn'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->related_to );
		if ( ! empty( $serialized ) ) {
			$out['relatedTo'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->seeks );
		if ( ! empty( $serialized ) ) {
			$out['seeks'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->sibling );
		if ( ! empty( $serialized ) ) {
			$out['sibling'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->siblings );
		if ( ! empty( $serialized ) ) {
			$out['siblings'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->spouse );
		if ( ! empty( $serialized ) ) {
			$out['spouse'] = $serialized;
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

		$serialized = \SchemaOrg\json_serialize( $this->weight );
		if ( ! empty( $serialized ) ) {
			$out['weight'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->work_location );
		if ( ! empty( $serialized ) ) {
			$out['workLocation'] = $serialized;
		}

		$serialized = \SchemaOrg\json_serialize( $this->works_for );
		if ( ! empty( $serialized ) ) {
			$out['worksFor'] = $serialized;
		}

		return $out;
	}
}
