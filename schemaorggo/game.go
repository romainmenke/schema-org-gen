package schemaorggo

import "encoding/json"

// Game see : https://schema.org/Game
type Game struct {
	CreativeWork

	typeContext

	// CharacterAttribute see : https://schema.org/characterAttribute
	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	// types : Thing
	CharacterAttribute *Thing `json:"characterAttribute,omitempty"`

	// GameItem see : https://schema.org/gameItem
	// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	// types : Thing
	GameItem *Thing `json:"gameItem,omitempty"`

	// GameLocation see : https://schema.org/gameLocation
	// Real or fictional location of the game (or part of game).
	// types : Place PostalAddress URL
	GameLocation interface{} `json:"gameLocation,omitempty"`

	// NumberOfPlayers see : https://schema.org/numberOfPlayers
	// Indicate how many people can play this game (minimum, maximum, or range).
	// types : QuantitativeValue
	NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers,omitempty"`

	// Quest see : https://schema.org/quest
	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	// types : Thing
	Quest *Thing `json:"quest,omitempty"`
}

func (v Game) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Game"

	return json.Marshal(v)
}

func (v *Game) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Game"

	return json.Marshal(*v)
}
