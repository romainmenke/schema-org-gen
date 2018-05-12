package schemaorggo

import "encoding/json"

// ExerciseAction see : https://schema.org/ExerciseAction
type ExerciseAction struct {
	PlayAction

	typeContext

	// Diet see : http://health-lifesci.schema.org/diet
	// A sub property of instrument. The diet used in this action.
	// types : Diet
	Diet interface{} `json:"diet,omitempty"`

	// Distance see : https://schema.org/distance
	// The distance travelled, e.g. exercising or travelling.
	// types : Distance
	Distance *Distance `json:"distance,omitempty"`

	// ExerciseCourse see : https://schema.org/exerciseCourse
	// A sub property of location. The course where this action was taken. Supersedes course (see: https://schema.org/course).
	// types : Place
	ExerciseCourse *Place `json:"exerciseCourse,omitempty"`

	// ExercisePlan see : http://health-lifesci.schema.org/exercisePlan
	// A sub property of instrument. The exercise plan used on this action.
	// types : ExercisePlan
	ExercisePlan interface{} `json:"exercisePlan,omitempty"`

	// ExerciseRelatedDiet see : http://health-lifesci.schema.org/exerciseRelatedDiet
	// A sub property of instrument. The diet used in this action.
	// types : Diet
	ExerciseRelatedDiet interface{} `json:"exerciseRelatedDiet,omitempty"`

	// ExerciseType see : http://health-lifesci.schema.org/exerciseType
	// Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
	// types : Text
	ExerciseType string `json:"exerciseType,omitempty"`

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	// types : Place
	FromLocation *Place `json:"fromLocation,omitempty"`

	// Opponent see : https://schema.org/opponent
	// A sub property of participant. The opponent on this action.
	// types : Person
	Opponent *Person `json:"opponent,omitempty"`

	// SportsActivityLocation see : https://schema.org/sportsActivityLocation
	// A sub property of location. The sports activity location where this action occurred.
	// types : SportsActivityLocation
	SportsActivityLocation *SportsActivityLocation `json:"sportsActivityLocation,omitempty"`

	// SportsEvent see : https://schema.org/sportsEvent
	// A sub property of location. The sports event where this action occurred.
	// types : SportsEvent
	SportsEvent *SportsEvent `json:"sportsEvent,omitempty"`

	// SportsTeam see : https://schema.org/sportsTeam
	// A sub property of participant. The sports team that participated on this action.
	// types : SportsTeam
	SportsTeam *SportsTeam `json:"sportsTeam,omitempty"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	// types : Place
	ToLocation *Place `json:"toLocation,omitempty"`
}

func (v ExerciseAction) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "ExerciseAction"

	return json.Marshal(v)
}

func (v *ExerciseAction) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "ExerciseAction"

	return json.Marshal(*v)
}
