package schemaorggo

import "encoding/json"

// VideoGameSeries see : https://schema.org/VideoGameSeries
type VideoGameSeries struct {
	CreativeWorkSeries

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// CharacterAttribute see : https://schema.org/characterAttribute
	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	// types : Thing
	CharacterAttribute []*Thing `json:"characterAttribute,omitempty"`

	// CheatCode see : https://schema.org/cheatCode
	// Cheat codes to the game.
	// types : CreativeWork
	CheatCode []*CreativeWork `json:"cheatCode,omitempty"`

	// ContainsSeason see : https://schema.org/containsSeason
	// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
	// types : CreativeWorkSeason
	ContainsSeason []*CreativeWorkSeason `json:"containsSeason,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	// types : Episode
	Episode []*Episode `json:"episode,omitempty"`

	// GameItem see : https://schema.org/gameItem
	// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	// types : Thing
	GameItem []*Thing `json:"gameItem,omitempty"`

	// GameLocation see : https://schema.org/gameLocation
	// Real or fictional location of the game (or part of game).
	// types : Place PostalAddress URL
	GameLocation []interface{} `json:"gameLocation,omitempty"`

	// GamePlatform see : https://schema.org/gamePlatform
	// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	// types : Text Thing URL
	GamePlatform []interface{} `json:"gamePlatform,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy []interface{} `json:"musicBy,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes []float64 `json:"numberOfEpisodes,omitempty"`

	// NumberOfPlayers see : https://schema.org/numberOfPlayers
	// Indicate how many people can play this game (minimum, maximum, or range).
	// types : QuantitativeValue
	NumberOfPlayers []*QuantitativeValue `json:"numberOfPlayers,omitempty"`

	// NumberOfSeasons see : https://schema.org/numberOfSeasons
	// The number of seasons in this series.
	// types : Integer
	NumberOfSeasons []float64 `json:"numberOfSeasons,omitempty"`

	// PlayMode see : https://schema.org/playMode
	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	// types : GamePlayMode
	PlayMode []*GamePlayMode `json:"playMode,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany []*Organization `json:"productionCompany,omitempty"`

	// Quest see : https://schema.org/quest
	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	// types : Thing
	Quest []*Thing `json:"quest,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer []*VideoObject `json:"trailer,omitempty"`
}

func (v VideoGameSeries) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.CreativeWorkSeries.intoMap(intop)

	into := *intop

	{
		var value interface{} = v.Actor
		if len(v.Actor) == 1 {
			value = v.Actor[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["actor"] = json.RawMessage(b)
		}
	}

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
		var value interface{} = v.CheatCode
		if len(v.CheatCode) == 1 {
			value = v.CheatCode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["cheatCode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ContainsSeason
		if len(v.ContainsSeason) == 1 {
			value = v.ContainsSeason[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["containsSeason"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Director
		if len(v.Director) == 1 {
			value = v.Director[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["director"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.Episode
		if len(v.Episode) == 1 {
			value = v.Episode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["episode"] = json.RawMessage(b)
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
		var value interface{} = v.GamePlatform
		if len(v.GamePlatform) == 1 {
			value = v.GamePlatform[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gamePlatform"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.MusicBy
		if len(v.MusicBy) == 1 {
			value = v.MusicBy[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["musicBy"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.NumberOfEpisodes
		if len(v.NumberOfEpisodes) == 1 {
			value = v.NumberOfEpisodes[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfEpisodes"] = json.RawMessage(b)
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
		var value interface{} = v.NumberOfSeasons
		if len(v.NumberOfSeasons) == 1 {
			value = v.NumberOfSeasons[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["numberOfSeasons"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.PlayMode
		if len(v.PlayMode) == 1 {
			value = v.PlayMode[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["playMode"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.ProductionCompany
		if len(v.ProductionCompany) == 1 {
			value = v.ProductionCompany[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["productionCompany"] = json.RawMessage(b)
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

	{
		var value interface{} = v.Trailer
		if len(v.Trailer) == 1 {
			value = v.Trailer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["trailer"] = json.RawMessage(b)
		}
	}

	*intop = into

	return nil
}

func (v VideoGameSeries) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "VideoGameSeries"

	return data, nil
}

func (v VideoGameSeries) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
