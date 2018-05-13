<?php

class DiscussionForumPosting extends SocialMediaPosting implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'DiscussionForumPosting';
	
	/**
	 * A CreativeWork such as an image, video, or audio clip shared as part of this posting.
	 * see : https://schema.org/sharedContent
	 * @var \CreativeWork|\CreativeWork[]
	 */
	public var $shared_content;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'DiscussionForumPosting'
		);
		
		$serialized = so_json_serialize( $this->shared_content );
		if ( ! empty( $serialized ) ) {
			$out['sharedContent'] = $serialized;
		}
		
		return $out;
	}
}
