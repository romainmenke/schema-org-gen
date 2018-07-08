<?php

class OnDemandEvent extends PublicationEvent implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'OnDemandEvent';
	
	/**
	 * A flag to signal that the item, event, or place is accessible for free. Supersedes free (see: https://schema.org/free).
	 * see : https://schema.org/isAccessibleForFree
	 * @var boolean|boolean[]
	 */
	public var $is_accessible_for_free;
	
	/**
	 * An agent associated with the publication event.
	 * see : https://bib.schema.org/publishedBy
	 * @var \Organization|\Organization[]|\Person|\Person[]
	 */
	public var $published_by;
	
	/**
	 * A broadcast service associated with the publication event.
	 * see : https://schema.org/publishedOn
	 * @var \BroadcastService|\BroadcastService[]
	 */
	public var $published_on;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'OnDemandEvent'
		);
		
		$serialized = so_json_serialize( $this->is_accessible_for_free );
		if ( ! empty( $serialized ) ) {
			$out['isAccessibleForFree'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->published_by );
		if ( ! empty( $serialized ) ) {
			$out['publishedBy'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->published_on );
		if ( ! empty( $serialized ) ) {
			$out['publishedOn'] = $serialized;
		}
		
		return $out;
	}
}
