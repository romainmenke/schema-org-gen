package schemaorggo

import "encoding/json"

// VideoGameSeries see : https://schema.org/VideoGameSeries
type VideoGameSeries struct {
	CreativeWorkSeries

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	Actor *Person `json:"actor,omitempty"` // types : Person

	// CharacterAttribute see : https://schema.org/characterAttribute
	// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
	CharacterAttribute *Thing `json:"characterAttribute,omitempty"` // types : Thing

	// CheatCode see : https://schema.org/cheatCode
	// Cheat codes to the game.
	CheatCode *CreativeWork `json:"cheatCode,omitempty"` // types : CreativeWork

	// ContainsSeason see : https://schema.org/containsSeason
	// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
	ContainsSeason *CreativeWorkSeason `json:"containsSeason,omitempty"` // types : CreativeWorkSeason

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	Director *Person `json:"director,omitempty"` // types : Person

	// Episode see : https://schema.org/episode
	// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
	Episode *Episode `json:"episode,omitempty"` // types : Episode

	// GameItem see : https://schema.org/gameItem
	// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
	GameItem *Thing `json:"gameItem,omitempty"` // types : Thing

	// GameLocation see : https://schema.org/gameLocation
	// Real or fictional location of the game (or part of game).
	GameLocation interface{} `json:"gameLocation,omitempty"` // types : Place PostalAddress URL

	// GamePlatform see : https://schema.org/gamePlatform
	// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	GamePlatform interface{} `json:"gamePlatform,omitempty"` // types : Text Thing URL

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	MusicBy interface{} `json:"musicBy,omitempty"` // types : MusicGroup Person

	// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
	// The number of episodes in this season or series.
	NumberOfEpisodes float64 `json:"numberOfEpisodes,omitempty"` // types : Integer

	// NumberOfPlayers see : https://schema.org/numberOfPlayers
	// Indicate how many people can play this game (minimum, maximum, or range).
	NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers,omitempty"` // types : QuantitativeValue

	// NumberOfSeasons see : https://schema.org/numberOfSeasons
	// The number of seasons in this series.
	NumberOfSeasons float64 `json:"numberOfSeasons,omitempty"` // types : Integer

	// PlayMode see : https://schema.org/playMode
	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	PlayMode *GamePlayMode `json:"playMode,omitempty"` // types : GamePlayMode

	// ProductionCompany see : https://schema.org/productionCompany
	// The production company or studio responsible for the item e.g. series, video game, episode etc.
	ProductionCompany *Organization `json:"productionCompany,omitempty"` // types : Organization

	// Quest see : https://schema.org/quest
	// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
	Quest *Thing `json:"quest,omitempty"` // types : Thing

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	Trailer *VideoObject `json:"trailer,omitempty"` // types : VideoObject

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
