package schemaorg

import "encoding/json"

// TravelAction see : https://schema.org/TravelAction
type TravelAction struct {

MoveAction

typeContext

// Distance see : https://schema.org/distance
// The distance travelled, e.g. exercising or travelling.
Distance *Distance `json:"distance"`

}

func (v *TravelAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "TravelAction"

	return json.Marshal(v)
}
