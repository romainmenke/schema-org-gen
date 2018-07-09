<?php

namespace SchemaOrg;

// IgnoreAction see : https://schema.org/IgnoreAction
class IgnoreAction implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'IgnoreAction';
	
	/**
	 * With properties from Action see : https://schema.org/Action
	 */
	
	/**
	 * With properties from AssessAction see : https://schema.org/AssessAction
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
	 */
	
	
	/**
	 * Indicates the current disposition of the Action.
	 * see : https://schema.org/actionStatus
	 * @var \ActionStatusType | \ActionStatusType[]
	 */
	public var $action_status;
	
	/**
	 * An additional type for the item, typically used for adding more specific types from external vocabularies in microdata syntax. This is a relationship between something and a class that the thing is in. In RDFa syntax, it is better to use the native RDFa syntax - the &#39;typeof&#39; attribute - for multiple types. Schema.org tools may have only weaker understanding of extra types, in particular those defined externally.
	 * see : https://schema.org/additionalType
	 * @var string | string[]
	 */
	public var $additional_type;
	
	/**
	 * The direct performer or driver of the action (animate or inanimate). e.g. John wrote a book.
	 * see : https://schema.org/agent
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $agent;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public var $alternate_name;
	
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
	 * The endTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to end. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	 * 
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/endTime
	 * @var string | string[]
	 */
	public var $end_time;
	
	/**
	 * For failed actions, more information on the cause of the failure.
	 * see : https://schema.org/error
	 * @var \Thing | \Thing[]
	 */
	public var $error;
	
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
	 * The object that helped the agent perform the action. e.g. John wrote a book with a pen.
	 * see : https://schema.org/instrument
	 * @var \Thing | \Thing[]
	 */
	public var $instrument;
	
	/**
	 * The location of for example where the event is happening, an organization is located, or where an action takes place.
	 * see : https://schema.org/location
	 * @var \Place | \Place[] | \PostalAddress | \PostalAddress[] | string | string[]
	 */
	public var $location;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See background notes (see: https://schema.org/docs/datamodel.html#mainEntityBackground) for details. Inverse property: mainEntity (see: https://schema.org/mainEntity).
	 * see : https://schema.org/mainEntityOfPage
	 * @var \CreativeWork | \CreativeWork[] | string | string[]
	 */
	public var $main_entity_of_page;
	
	/**
	 * The name of the item.
	 * see : https://schema.org/name
	 * @var string | string[]
	 */
	public var $name;
	
	/**
	 * The object upon which the action is carried out, whose state is kept intact or changed. Also known as the semantic roles patient, affected or undergoer (which change their state) or theme (which doesn&#39;t). e.g. John read a book.
	 * see : https://schema.org/object
	 * @var \Thing | \Thing[]
	 */
	public var $object;
	
	/**
	 * Other co-agents that participated in the action indirectly. e.g. John wrote a book with Steve.
	 * see : https://schema.org/participant
	 * @var \Organization | \Organization[] | \Person | \Person[]
	 */
	public var $participant;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public var $potential_action;
	
	/**
	 * The result produced in the action. e.g. John wrote a book.
	 * see : https://schema.org/result
	 * @var \Thing | \Thing[]
	 */
	public var $result;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public var $same_as;
	
	/**
	 * The startTime of something. For a reserved event or service (e.g. FoodEstablishmentReservation), the time that it is expected to start. For actions that span a period of time, when the action was performed. e.g. John wrote a book from January to December.
	 * 
	 * Note that Event uses startDate/endDate instead of startTime/endTime, even when describing dates with times. This situation may be clarified in future revisions.
	 * see : https://schema.org/startTime
	 * @var string | string[]
	 */
	public var $start_time;
	
	/**
	 * A CreativeWork or Event about this Thing.. Inverse property: about (see: https://schema.org/about).
	 * see : https://pending.schema.org/subjectOf
	 * @var \CreativeWork | \CreativeWork[] | \Event | \Event[]
	 */
	public var $subject_of;
	
	/**
	 * Indicates a target EntryPoint for an Action.
	 * see : https://schema.org/target
	 * @var \EntryPoint | \EntryPoint[]
	 */
	public var $target;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public var $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'IgnoreAction'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->action_status );
		if ( ! empty( $serialized ) ) {
			$out['actionStatus'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->agent );
		if ( ! empty( $serialized ) ) {
			$out['agent'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->end_time );
		if ( ! empty( $serialized ) ) {
			$out['endTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->error );
		if ( ! empty( $serialized ) ) {
			$out['error'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->instrument );
		if ( ! empty( $serialized ) ) {
			$out['instrument'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->location );
		if ( ! empty( $serialized ) ) {
			$out['location'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->object );
		if ( ! empty( $serialized ) ) {
			$out['object'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->participant );
		if ( ! empty( $serialized ) ) {
			$out['participant'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->result );
		if ( ! empty( $serialized ) ) {
			$out['result'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->start_time );
		if ( ! empty( $serialized ) ) {
			$out['startTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->subject_of );
		if ( ! empty( $serialized ) ) {
			$out['subjectOf'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->target );
		if ( ! empty( $serialized ) ) {
			$out['target'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
