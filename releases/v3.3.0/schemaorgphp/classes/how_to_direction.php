<?php
namespace SchemaOrg;

// HowToDirection see : https://schema.org/HowToDirection
class HowToDirection implements \JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'HowToDirection';
	
	/**
	 * With properties from ListItem see : https://schema.org/ListItem
	 */
	
	/**
	 * With properties from Intangible see : https://schema.org/Intangible
	 */
	
	/**
	 * With properties from Thing see : https://schema.org/Thing
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
	 * A media object representing the circumstances after performing this direction.
	 * see : https://schema.org/afterMedia
	 * @var \MediaObject | \MediaObject[]
	 */
	public $after_media;
	
	/**
	 * An alias for the item.
	 * see : https://schema.org/alternateName
	 * @var string | string[]
	 */
	public $alternate_name;
	
	/**
	 * A media object representing the circumstances before performing this direction.
	 * see : https://schema.org/beforeMedia
	 * @var \MediaObject | \MediaObject[]
	 */
	public $before_media;
	
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
	 * A media object representing the circumstances while performing this direction.
	 * see : https://schema.org/duringMedia
	 * @var \MediaObject | \MediaObject[]
	 */
	public $during_media;
	
	/**
	 * The identifier property represents any kind of identifier for any kind of [[Thing]], such as ISBNs, GTIN codes, UUIDs etc. Schema.org provides dedicated properties for representing many of these, either as textual strings or as URL (URI) links. See [background notes](/docs/datamodel.html#identifierBg) for more details.
	 *         
	 * see : https://schema.org/identifier
	 * @var string | string[] | \PropertyValue | \PropertyValue[]
	 */
	public $identifier;
	
	/**
	 * An image of the item. This can be a [[URL]] or a fully described [[ImageObject]].
	 * see : https://schema.org/image
	 * @var string | string[] | \ImageObject | \ImageObject[]
	 */
	public $image;
	
	/**
	 * An entity represented by an entry in a list or data feed (e.g. an &#39;artist&#39; in a list of &#39;artists&#39;)’.
	 * see : https://schema.org/item
	 * @var \Thing | \Thing[]
	 */
	public $item;
	
	/**
	 * Indicates a page (or other CreativeWork) for which this thing is the main entity being described. See [background notes](/docs/datamodel.html#mainEntityBackground) for details.
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
	 * A link to the ListItem that follows the current one.
	 * see : https://schema.org/nextItem
	 * @var \ListItem | \ListItem[]
	 */
	public $next_item;
	
	/**
	 * The length of time it takes to perform instructions or a direction (not including time to prepare the supplies), in [ISO 8601 duration format](http://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/performTime
	 * @var \Duration | \Duration[]
	 */
	public $perform_time;
	
	/**
	 * The position of an item in a series or sequence of items.
	 * see : https://schema.org/position
	 * @var string | string[] | integer | integer[]
	 */
	public $position;
	
	/**
	 * Indicates a potential Action, which describes an idealized action in which this thing would play an &#39;object&#39; role.
	 * see : https://schema.org/potentialAction
	 * @var \Action | \Action[]
	 */
	public $potential_action;
	
	/**
	 * The length of time it takes to prepare the items to be used in instructions or a direction, in [ISO 8601 duration format](http://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/prepTime
	 * @var \Duration | \Duration[]
	 */
	public $prep_time;
	
	/**
	 * A link to the ListItem that preceeds the current one.
	 * see : https://schema.org/previousItem
	 * @var \ListItem | \ListItem[]
	 */
	public $previous_item;
	
	/**
	 * URL of a reference Web page that unambiguously indicates the item&#39;s identity. E.g. the URL of the item&#39;s Wikipedia page, Wikidata entry, or official website.
	 * see : https://schema.org/sameAs
	 * @var string | string[]
	 */
	public $same_as;
	
	/**
	 * A sub-property of instrument. A supply consumed when performing instructions or a direction.
	 * see : https://schema.org/supply
	 * @var string | string[] | \HowToSupply | \HowToSupply[]
	 */
	public $supply;
	
	/**
	 * A sub property of instrument. An object used (but not consumed) when performing instructions or a direction.
	 * see : https://schema.org/tool
	 * @var string | string[] | \HowToTool | \HowToTool[]
	 */
	public $tool;
	
	/**
	 * The total time required to perform instructions or a direction (including time to prepare the supplies), in [ISO 8601 duration format](http://en.wikipedia.org/wiki/ISO_8601).
	 * see : https://schema.org/totalTime
	 * @var \Duration | \Duration[]
	 */
	public $total_time;
	
	/**
	 * URL of the item.
	 * see : https://schema.org/url
	 * @var string | string[]
	 */
	public $url;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'HowToDirection'
		);
		
		$serialized = \SchemaOrg\json_serialize( $this->additional_type );
		if ( ! empty( $serialized ) ) {
			$out['additionalType'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->after_media );
		if ( ! empty( $serialized ) ) {
			$out['afterMedia'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->alternate_name );
		if ( ! empty( $serialized ) ) {
			$out['alternateName'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->before_media );
		if ( ! empty( $serialized ) ) {
			$out['beforeMedia'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->description );
		if ( ! empty( $serialized ) ) {
			$out['description'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->disambiguating_description );
		if ( ! empty( $serialized ) ) {
			$out['disambiguatingDescription'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->during_media );
		if ( ! empty( $serialized ) ) {
			$out['duringMedia'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->identifier );
		if ( ! empty( $serialized ) ) {
			$out['identifier'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->image );
		if ( ! empty( $serialized ) ) {
			$out['image'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->item );
		if ( ! empty( $serialized ) ) {
			$out['item'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->main_entity_of_page );
		if ( ! empty( $serialized ) ) {
			$out['mainEntityOfPage'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->name );
		if ( ! empty( $serialized ) ) {
			$out['name'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->next_item );
		if ( ! empty( $serialized ) ) {
			$out['nextItem'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->perform_time );
		if ( ! empty( $serialized ) ) {
			$out['performTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->position );
		if ( ! empty( $serialized ) ) {
			$out['position'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->potential_action );
		if ( ! empty( $serialized ) ) {
			$out['potentialAction'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->prep_time );
		if ( ! empty( $serialized ) ) {
			$out['prepTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->previous_item );
		if ( ! empty( $serialized ) ) {
			$out['previousItem'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->same_as );
		if ( ! empty( $serialized ) ) {
			$out['sameAs'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->supply );
		if ( ! empty( $serialized ) ) {
			$out['supply'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->tool );
		if ( ! empty( $serialized ) ) {
			$out['tool'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->total_time );
		if ( ! empty( $serialized ) ) {
			$out['totalTime'] = $serialized;
		}
		
		$serialized = \SchemaOrg\json_serialize( $this->url );
		if ( ! empty( $serialized ) ) {
			$out['url'] = $serialized;
		}
		
		return $out;
	}
}
