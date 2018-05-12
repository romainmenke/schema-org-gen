package schemaorg

import "encoding/json"

// ExerciseAction see : https://schema.org/ExerciseAction
type ExerciseAction struct {

PlayAction

typeContext

// Diet see : http://health-lifesci.schema.org/diet
// A sub property of instrument. The diet used in this action.
Diet interface{} `json:"diet"`

// Distance see : https://schema.org/distance
// The distance travelled, e.g. exercising or travelling.
Distance *Distance `json:"distance"`

// ExerciseCourse see : https://schema.org/exerciseCourse
// A sub property of location. The course where this action was taken. Supersedes course (see: https://schema.org/course).
ExerciseCourse *Place `json:"exerciseCourse"`

// ExercisePlan see : http://health-lifesci.schema.org/exercisePlan
// A sub property of instrument. The exercise plan used on this action.
ExercisePlan interface{} `json:"exercisePlan"`

// ExerciseRelatedDiet see : http://health-lifesci.schema.org/exerciseRelatedDiet
// A sub property of instrument. The diet used in this action.
ExerciseRelatedDiet interface{} `json:"exerciseRelatedDiet"`

// ExerciseType see : http://health-lifesci.schema.org/exerciseType
// Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
ExerciseType string `json:"exerciseType"`

// FromLocation see : https://schema.org/fromLocation
// A sub property of location. The original location of the object or the agent before the action.
FromLocation *Place `json:"fromLocation"`

// Opponent see : https://schema.org/opponent
// A sub property of participant. The opponent on this action.
Opponent *Person `json:"opponent"`

// SportsActivityLocation see : https://schema.org/sportsActivityLocation
// A sub property of location. The sports activity location where this action occurred.
SportsActivityLocation *SportsActivityLocation `json:"sportsActivityLocation"`

// SportsEvent see : https://schema.org/sportsEvent
// A sub property of location. The sports event where this action occurred.
SportsEvent *SportsEvent `json:"sportsEvent"`

// SportsTeam see : https://schema.org/sportsTeam
// A sub property of participant. The sports team that participated on this action.
SportsTeam *SportsTeam `json:"sportsTeam"`

// ToLocation see : https://schema.org/toLocation
// A sub property of location. The final location of the object or the agent after the action.
ToLocation *Place `json:"toLocation"`

}

func (v *ExerciseAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ExerciseAction"

	return json.Marshal(v)
}
