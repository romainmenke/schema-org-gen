<?php

class LiveBlogPosting extends BlogPosting implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'LiveBlogPosting';
	
	/**
	 * The time when the live blog will stop covering the Event. Note that coverage may continue after the Event concludes.
	 * see : https://schema.org/coverageEndTime
	 * @var string|string[]
	 */
	public var $coverage_end_time;
	
	/**
	 * The time when the live blog will begin covering the Event. Note that coverage may begin before the Event&#39;s start time. The LiveBlogPosting may also be created before coverage begins.
	 * see : https://schema.org/coverageStartTime
	 * @var string|string[]
	 */
	public var $coverage_start_time;
	
	/**
	 * An update to the LiveBlog.
	 * see : https://schema.org/liveBlogUpdate
	 * @var \BlogPosting|\BlogPosting[]
	 */
	public var $live_blog_update;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'LiveBlogPosting'
		);
		
		$serialized = so_json_serialize( $this->coverage_end_time );
		if ( ! empty( $serialized ) ) {
			$out['coverageEndTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->coverage_start_time );
		if ( ! empty( $serialized ) ) {
			$out['coverageStartTime'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->live_blog_update );
		if ( ! empty( $serialized ) ) {
			$out['liveBlogUpdate'] = $serialized;
		}
		
		return $out;
	}
}
