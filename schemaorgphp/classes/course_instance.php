<?php

class CourseInstance extends Event implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'CourseInstance';
	
	/**
	 * The medium or means of delivery of the course instance or the mode of study, either as a text label (e.g. &quot;online&quot;, &quot;onsite&quot; or &quot;blended&quot;; &quot;synchronous&quot; or &quot;asynchronous&quot;; &quot;full-time&quot; or &quot;part-time&quot;) or as a URL reference to a term from a controlled vocabulary (e.g. https://ceds.ed.gov/element/001311#Asynchronous ).
	 * see : https://schema.org/courseMode
	 * @var string|string[]|string|string[]
	 */
	public var $course_mode;
	
	/**
	 * A person assigned to instruct or provide instructional assistance for the CourseInstance (see: https://schema.org/CourseInstance).
	 * see : https://schema.org/instructor
	 * @var \Person|\Person[]
	 */
	public var $instructor;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'CourseInstance'
		);
		
		$serialized = so_json_serialize( $this->course_mode );
		if ( ! empty( $serialized ) ) {
			$out['courseMode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->instructor );
		if ( ! empty( $serialized ) ) {
			$out['instructor'] = $serialized;
		}
		
		return $out;
	}
}
