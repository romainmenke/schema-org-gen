package schemaorggo

import "encoding/json"

// ExerciseAction see : https://schema.org/ExerciseAction
type ExerciseAction struct {
	PlayAction

	typeContext

	// Diet see : https://health-lifesci.schema.org/diet
	// A sub property of instrument. The diet used in this action.
	// types : Diet
	Diet []interface{} `json:"diet,omitempty"`

	// Distance see : https://schema.org/distance
	// The distance travelled, e.g. exercising or travelling.
	// types : Distance
	Distance []*Distance `json:"distance,omitempty"`

	// ExerciseCourse see : https://schema.org/exerciseCourse
	// A sub property of location. The course where this action was taken. Supersedes course (see: https://schema.org/course).
	// types : Place
	ExerciseCourse []*Place `json:"exerciseCourse,omitempty"`

	// ExercisePlan see : https://health-lifesci.schema.org/exercisePlan
	// A sub property of instrument. The exercise plan used on this action.
	// types : ExercisePlan
	ExercisePlan []interface{} `json:"exercisePlan,omitempty"`

	// ExerciseRelatedDiet see : https://health-lifesci.schema.org/exerciseRelatedDiet
	// A sub property of instrument. The diet used in this action.
	// types : Diet
	ExerciseRelatedDiet []interface{} `json:"exerciseRelatedDiet,omitempty"`

	// ExerciseType see : https://health-lifesci.schema.org/exerciseType
	// Type(s) of exercise or activity, such as strength training, flexibility training, aerobics, cardiac rehabilitation, etc.
	// types : Text
	ExerciseType []string `json:"exerciseType,omitempty"`

	// FromLocation see : https://schema.org/fromLocation
	// A sub property of location. The original location of the object or the agent before the action.
	// types : Place
	FromLocation []*Place `json:"fromLocation,omitempty"`

	// Opponent see : https://schema.org/opponent
	// A sub property of participant. The opponent on this action.
	// types : Person
	Opponent []*Person `json:"opponent,omitempty"`

	// SportsActivityLocation see : https://schema.org/sportsActivityLocation
	// A sub property of location. The sports activity location where this action occurred.
	// types : SportsActivityLocation
	SportsActivityLocation []*SportsActivityLocation `json:"sportsActivityLocation,omitempty"`

	// SportsEvent see : https://schema.org/sportsEvent
	// A sub property of location. The sports event where this action occurred.
	// types : SportsEvent
	SportsEvent []*SportsEvent `json:"sportsEvent,omitempty"`

	// SportsTeam see : https://schema.org/sportsTeam
	// A sub property of participant. The sports team that participated on this action.
	// types : SportsTeam
	SportsTeam []*SportsTeam `json:"sportsTeam,omitempty"`

	// ToLocation see : https://schema.org/toLocation
	// A sub property of location. The final location of the object or the agent after the action.
	// types : Place
	ToLocation []*Place `json:"toLocation,omitempty"`
}

func (v ExerciseAction) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.PlayAction.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Diet
		if len(v.Diet) == 1 {
			value = v.Diet[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["diet"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Distance
		if len(v.Distance) == 1 {
			value = v.Distance[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["distance"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExerciseCourse
		if len(v.ExerciseCourse) == 1 {
			value = v.ExerciseCourse[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["exerciseCourse"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExercisePlan
		if len(v.ExercisePlan) == 1 {
			value = v.ExercisePlan[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["exercisePlan"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExerciseRelatedDiet
		if len(v.ExerciseRelatedDiet) == 1 {
			value = v.ExerciseRelatedDiet[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["exerciseRelatedDiet"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ExerciseType
		if len(v.ExerciseType) == 1 {
			value = v.ExerciseType[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["exerciseType"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.FromLocation
		if len(v.FromLocation) == 1 {
			value = v.FromLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["fromLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Opponent
		if len(v.Opponent) == 1 {
			value = v.Opponent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["opponent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SportsActivityLocation
		if len(v.SportsActivityLocation) == 1 {
			value = v.SportsActivityLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sportsActivityLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SportsEvent
		if len(v.SportsEvent) == 1 {
			value = v.SportsEvent[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sportsEvent"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.SportsTeam
		if len(v.SportsTeam) == 1 {
			value = v.SportsTeam[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["sportsTeam"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ToLocation
		if len(v.ToLocation) == 1 {
			value = v.ToLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["toLocation"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v ExerciseAction) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "ExerciseAction"

	return data, nil
}

func (v ExerciseAction) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
