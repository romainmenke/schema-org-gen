<?php

class ExerciseAction extends PlayAction implements JsonSerializable {

	public static $context = 'http://schema.org';
	public static $type = 'ExerciseAction';
	
	/**
	 * A sub property of instrument. The diet used in this action.
	 * see : https://health-lifesci.schema.org/diet
	 * @var \Diet|\Diet[]
	 */
	public var $diet;
	
	/**
	 * The distance travelled, e.g. exercising or travelling.
	 * see : https://schema.org/distance
	 * @var \Distance|\Distance[]
	 */
	public var $distance;
	
	/**
	 * A sub property of location. The course where this action was taken. Supersedes course (see: https://schema.org/course).
	 * see : https://schema.org/exerciseCourse
	 * @var \Place|\Place[]
	 */
	public var $exercise_course;
	
	/**
	 * A sub property of instrument. The exercise plan used on this action.
	 * see : https://health-lifesci.schema.org/exercisePlan
	 * @var \ExercisePlan|\ExercisePlan[]
	 */
	public var $exercise_plan;
	
	/**
	 * A sub property of instrument. The diet used in this action.
	 * see : https://health-lifesci.schema.org/exerciseRelatedDiet
	 * @var \Diet|\Diet[]
	 */
	public var $exercise_related_diet;
	
	/**
	 * Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
	 * see : https://health-lifesci.schema.org/exerciseType
	 * @var string|string[]
	 */
	public var $exercise_type;
	
	/**
	 * A sub property of location. The original location of the object or the agent before the action.
	 * see : https://schema.org/fromLocation
	 * @var \Place|\Place[]
	 */
	public var $from_location;
	
	/**
	 * A sub property of participant. The opponent on this action.
	 * see : https://schema.org/opponent
	 * @var \Person|\Person[]
	 */
	public var $opponent;
	
	/**
	 * A sub property of location. The sports activity location where this action occurred.
	 * see : https://schema.org/sportsActivityLocation
	 * @var \SportsActivityLocation|\SportsActivityLocation[]
	 */
	public var $sports_activity_location;
	
	/**
	 * A sub property of location. The sports event where this action occurred.
	 * see : https://schema.org/sportsEvent
	 * @var \SportsEvent|\SportsEvent[]
	 */
	public var $sports_event;
	
	/**
	 * A sub property of participant. The sports team that participated on this action.
	 * see : https://schema.org/sportsTeam
	 * @var \SportsTeam|\SportsTeam[]
	 */
	public var $sports_team;
	
	/**
	 * A sub property of location. The final location of the object or the agent after the action.
	 * see : https://schema.org/toLocation
	 * @var \Place|\Place[]
	 */
	public var $to_location;
	
	public function jsonSerialize() {
		$out = array(
			'@context' => 'http://schema.org',
			'@type' => 'ExerciseAction'
		);
		
		$serialized = so_json_serialize( $this->diet );
		if ( ! empty( $serialized ) ) {
			$out['diet'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->distance );
		if ( ! empty( $serialized ) ) {
			$out['distance'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->exercise_course );
		if ( ! empty( $serialized ) ) {
			$out['exerciseCourse'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->exercise_plan );
		if ( ! empty( $serialized ) ) {
			$out['exercisePlan'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->exercise_related_diet );
		if ( ! empty( $serialized ) ) {
			$out['exerciseRelatedDiet'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->exercise_type );
		if ( ! empty( $serialized ) ) {
			$out['exerciseType'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->from_location );
		if ( ! empty( $serialized ) ) {
			$out['fromLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->opponent );
		if ( ! empty( $serialized ) ) {
			$out['opponent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sports_activity_location );
		if ( ! empty( $serialized ) ) {
			$out['sportsActivityLocation'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sports_event );
		if ( ! empty( $serialized ) ) {
			$out['sportsEvent'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->sports_team );
		if ( ! empty( $serialized ) ) {
			$out['sportsTeam'] = $serialized;
		}
		
		$serialized = so_json_serialize( $this->to_location );
		if ( ! empty( $serialized ) ) {
			$out['toLocation'] = $serialized;
		}
		
		return $out;
	}
}
