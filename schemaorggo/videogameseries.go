package schemaorggo

import "encoding/json"

// VideoGameSeries see : https://schema.org/VideoGameSeries
type VideoGameSeries struct {
	CreativeWorkSeries

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// CharacterAttribute see : https://schema.org/characterAttribute
	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	// types : Thing
	CharacterAttribute *Thing `json:"characterAttribute,omitempty"`

	// CheatCode see : https://schema.org/cheatCode
	// Cheat codes to the game.
	// types : CreativeWork
	CheatCode *CreativeWork `json:"cheatCode,omitempty"`

	// ContainsSeason see : https://schema.org/containsSeason
	// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
	// types : CreativeWorkSeason
	ContainsSeason *CreativeWorkSeason `json:"containsSeason,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	// types : Episode
	Episode *Episode `json:"episode,omitempty"`

	// GameItem see : https://schema.org/gameItem
	// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	// types : Thing
	GameItem *Thing `json:"gameItem,omitempty"`

	// GameLocation see : https://schema.org/gameLocation
	// Real or fictional location of the game (or part of game).
	// types : Place PostalAddress URL
	GameLocation interface{} `json:"gameLocation,omitempty"`

	// GamePlatform see : https://schema.org/gamePlatform
	// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	// types : Text Thing URL
	GamePlatform interface{} `json:"gamePlatform,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy interface{} `json:"musicBy,omitempty"`

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	// types : Integer
	NumberOfEpisodes float64 `json:"numberOfEpisodes,omitempty"`

	// NumberOfPlayers see : https://schema.org/numberOfPlayers
	// Indicate how many people can play this game (minimum, maximum, or range).
	// types : QuantitativeValue
	NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers,omitempty"`

	// NumberOfSeasons see : https://schema.org/numberOfSeasons
	// The number of seasons in this series.
	// types : Integer
	NumberOfSeasons float64 `json:"numberOfSeasons,omitempty"`

	// PlayMode see : https://schema.org/playMode
	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	// types : GamePlayMode
	PlayMode *GamePlayMode `json:"playMode,omitempty"`

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	// types : Organization
	ProductionCompany *Organization `json:"productionCompany,omitempty"`

	// Quest see : https://schema.org/quest
	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	// types : Thing
	Quest *Thing `json:"quest,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer *VideoObject `json:"trailer,omitempty"`
}

func (v VideoGameSeries) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoGameSeries"

	return json.Marshal(v)
}

func (v *VideoGameSeries) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "VideoGameSeries"

	return json.Marshal(*v)
}
