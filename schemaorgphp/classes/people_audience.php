<?php

namespace SchemaOrg;

// PeopleAudience see : https://schema.org/PeopleAudience
class PeopleAudience implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'PeopleAudience';
	
	/**
	 * With properties from Audience see : https://schema.org/Audience
	 */
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public $additional_type;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public $alternate_name;
	
	/**
	 * The target group associated with a given audience (e.g. veterans, car owners, musicians, etc.).
	 * see : https://schema.org/audienceType
	 * @var string | string[]
	 */
	public $audience_type;
	
	/**
	 * A description of the item.
	 * see : https://schema.org/description
	 * @var string | string[]
	 */
	public $description;
	
	/**
	 * A sub property of description. A short description of the item used to disambiguate from other, similar items. Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	 * see : https://schema.org/disambiguatingDescription
	 * @var string | string[]
	 */
	public $disambiguating_description;
	
	/**
	 * The geographic area associated with the audience.
	 * see : https://schema.org/geographicArea
	 * @var \AdministrativeArea | \AdministrativeArea[]
	 */
	public $geographic_area;
	
	/**
	 * Specifying the health condition(s) of a patient, medical study, or other target audience.
	 * see : https://health-lifesci.schema.org/healthCondition
	 * @var \MedicalCondition | \MedicalCondition[]
	 */
	public $health_condition;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of Thing (see: https://schema.org/Thing), such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See background notes (see: https://schema.org/docs/datamodel.html#identifierBg) for more details.
	 * see : https://schema.org/identifier
	 * @var \PropertyValue | \PropertyValue[] | string | string[]
	 */
	public $identifier;
	
	/**
	 * An image of the item. This can be a URL (see: https://schema.org/URL) or a fully described ImageObject (see: https://schema.org/ImageObject).
	 * see : https://schema.org/image
	 * @var \ImageObject | \ImageObject[] | string | string[]
	 */
	public $image;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public $main_entity_of_page;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public $name;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public $potential_action;
	
	/**
	 * Audiences defined by a person&#39;s gender.
	 * see : https://schema.org/requiredGender
	 * @var string | string[]
	 */
	public $required_gender;
	
	/**
	 * Audiences defined by a person&#39;s maximum age.
	 * see : https://schema.org/requiredMaxAge
	 * @var integer | integer[]
	 */
	public $required_max_age;
	
	/**
	 * Audiences defined by a person&#39;s minimum age.
	 * see : https://schema.org/requiredMinAge
	 * @var integer | integer[]
	 */
	public $required_min_age;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public $same_as;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public $subject_of;
	
	/**
	 * The gender of the person or audience.
	 * see : https://schema.org/suggestedGender
	 * @var string | string[]
	 */
	public $suggested_gender;
	
	/**
	 * Maximal age recommended for viewing content.
	 * see : https://schema.org/suggestedMaxAge
	 * @var float | float[]
	 */
	public $suggested_max_age;
	
	/**
	 * Minimal age recommended for viewing content.
	 * see : https://schema.org/suggestedMinAge
	 * @var float | float[]
	 */
	public $suggested_min_age;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'PeopleAudience'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->audience_type );
		if ( ! empty( $serialized ) ) {
			$out['audienceType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->geographic_area );
		if ( ! empty( $serialized ) ) {
			$out['geographicArea'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->health_condition );
		if ( ! empty( $serialized ) ) {
			$out['healthCondition'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->required_gender );
		if ( ! empty( $serialized ) ) {
			$out['requiredGender'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->required_max_age );
		if ( ! empty( $serialized ) ) {
			$out['requiredMaxAge'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->required_min_age );
		if ( ! empty( $serialized ) ) {
			$out['requiredMinAge'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->suggested_gender );
		if ( ! empty( $serialized ) ) {
			$out['suggestedGender'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->suggested_max_age );
		if ( ! empty( $serialized ) ) {
			$out['suggestedMaxAge'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->suggested_min_age );
		if ( ! empty( $serialized ) ) {
			$out['suggestedMinAge'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
