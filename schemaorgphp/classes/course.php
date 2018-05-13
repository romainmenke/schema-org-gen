<?php

class Course extends CreativeWork implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'Course';
	
	/**
	 * The identifier for the Course (see: https://schema.org/Course) used by the course provider (see: https://schema.org/provider) (e.g. CS101 or 6.001).
	 * see : https://schema.org/courseCode
	 * @var string|string[]
	 */
	public var $course_code;
	
	/**
	 * Requirements for taking the Course. May be completion of another Course (see: https://schema.org/Course) or a textual description like &quot;permission of instructor&quot;. Requirements may be a pre-requisite competency, referenced using AlignmentObject (see: https://schema.org/AlignmentObject).
	 * see : https://schema.org/coursePrerequisites
	 * @var \AlignmentObject|\AlignmentObject[]|\Course|\Course[]|string|string[]
	 */
	public var $course_prerequisites;
	
	/**
	 * A description of the qualification, award, certificate, diploma or other educational credential awarded as a consequence of successful completion of this course.
	 * see : http://pending.schema.org/educationalCredentialAwarded
	 * @var string|string[]|string|string[]
	 */
	public var $educational_credential_awarded;
	
	/**
	 * An offering of the course at a specific time and place or through specific media or mode of study or to a specific section of students.
	 * see : https://schema.org/hasCourseInstance
	 * @var \CourseInstance|\CourseInstance[]
	 */
	public var $has_course_instance;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'Course'
		);
		
		$serialized = so_json_serialize( $this->course_code );
		if ( ! empty( $serialized ) ) {
			$out['courseCode'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->course_prerequisites );
		if ( ! empty( $serialized ) ) {
			$out['coursePrerequisites'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->educational_credential_awarded );
		if ( ! empty( $serialized ) ) {
			$out['educationalCredentialAwarded'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->has_course_instance );
		if ( ! empty( $serialized ) ) {
			$out['hasCourseInstance'] = $serialized;
		}
		
		return $out;
	}
}
