<?php

class Question extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Question';
	
	/**
	 * The answer that has been accepted as best, typically on a Question/Answer site. Sites vary in their selection mechanisms, e.g. drawing on community opinion and/or the view of the Question author.
	 * see : https://schema.org/acceptedAnswer
	 * @var \Answer|\Answer[]
	 */
	public var $accepted_answer;
	
	/**
	 * The number of answers this question has received.
	 * see : https://schema.org/answerCount
	 * @var integer|integer[]
	 */
	public var $answer_count;
	
	/**
	 * The number of downvotes this question, answer or comment has received from the community.
	 * see : https://schema.org/downvoteCount
	 * @var integer|integer[]
	 */
	public var $downvote_count;
	
	/**
	 * An answer (possibly one of several, possibly incorrect) to a Question, e.g. on a Question/Answer site.
	 * see : https://schema.org/suggestedAnswer
	 * @var \Answer|\Answer[]
	 */
	public var $suggested_answer;
	
	/**
	 * The number of upvotes this question, answer or comment has received from the community.
	 * see : https://schema.org/upvoteCount
	 * @var integer|integer[]
	 */
	public var $upvote_count;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Question'
		);
		
		$serialized = so_json_serialize( $this->accepted_answer );
		if ( ! empty( $serialized ) ) {
			$out['acceptedAnswer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->answer_count );
		if ( ! empty( $serialized ) ) {
			$out['answerCount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->downvote_count );
		if ( ! empty( $serialized ) ) {
			$out['downvoteCount'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->suggested_answer );
		if ( ! empty( $serialized ) ) {
			$out['suggestedAnswer'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->upvote_count );
		if ( ! empty( $serialized ) ) {
			$out['upvoteCount'] = $serialized;
		}
		
		return $out;
	}
}
