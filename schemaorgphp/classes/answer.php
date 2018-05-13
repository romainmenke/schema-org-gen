<?php

class Answer extends Comment implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Answer';
	
	/**
	 * The number of downvotes this question, answer or comment has received from the community.
	 * see : https://schema.org/downvoteCount
	 * @var integer|integer[]
	 */
	public var $downvote_count;
	
	/**
	 * The parent of a question, answer or item in general.
	 * see : https://schema.org/parentItem
	 * @var \Question|\Question[]
	 */
	public var $parent_item;
	
	/**
	 * The number of upvotes this question, answer or comment has received from the community.
	 * see : https://schema.org/upvoteCount
	 * @var integer|integer[]
	 */
	public var $upvote_count;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Answer'
		);
		
		$serialized = so_json_serialize( $this->downvote_count );
		if ( ! empty( $serialized ) ) {
			$out['downvoteCount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->parent_item );
		if ( ! empty( $serialized ) ) {
			$out['parentItem'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->upvote_count );
		if ( ! empty( $serialized ) ) {
			$out['upvoteCount'] = $serialized;
		}
		
		return $out;
	}
}
