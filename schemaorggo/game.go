package schemaorggo

import "encoding/json"

// Game see : https://schema.org/Game
type Game struct {
	CreativeWork

	typeContext

	// CharacterAttribute see : https://schema.org/characterAttribute
	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	// types : Thing
	CharacterAttribute []*Thing `json:"characterAttribute,omitempty"`

	// GameItem see : https://schema.org/gameItem
	// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	// types : Thing
	GameItem []*Thing `json:"gameItem,omitempty"`

	// GameLocation see : https://schema.org/gameLocation
	// Real or fictional location of the game (or part of game).
	// types : Place PostalAddress URL
	GameLocation []interface{} `json:"gameLocation,omitempty"`

	// NumberOfPlayers see : https://schema.org/numberOfPlayers
	// Indicate how many people can play this game (minimum, maximum, or range).
	// types : QuantitativeValue
	NumberOfPlayers []*QuantitativeValue `json:"numberOfPlayers,omitempty"`

	// Quest see : https://schema.org/quest
	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	// types : Thing
	Quest []*Thing `json:"quest,omitempty"`
}

func (v Game) IntoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWork.IntoMap(intop)

	into := *intop

	{
		var value interface{} = v.CharacterAttribute
		if len(v.CharacterAttribute) == 1 {
			value = v.CharacterAttribute[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["characterAttribute"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GameItem
		if len(v.GameItem) == 1 {
			value = v.GameItem[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gameItem"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GameLocation
		if len(v.GameLocation) == 1 {
			value = v.GameLocation[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gameLocation"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfPlayers
		if len(v.NumberOfPlayers) == 1 {
			value = v.NumberOfPlayers[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfPlayers"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Quest
		if len(v.Quest) == 1 {
			value = v.Quest[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["quest"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v Game) AsMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.IntoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "Game"

	return data, nil
}

func (v Game) MarshalJSON() ([]byte, error) {
	data, err := v.AsMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
