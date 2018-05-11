package schemaorg

import "encoding/json"

// VideoGameSeries see : https://schema.org/VideoGameSeries
type VideoGameSeries struct {

typeContext

CreativeWorkSeries

// Actor see : https://schema.org/actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// CharacterAttribute see : https://schema.org/characterAttribute
// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
CharacterAttribute *Thing `json:"characterAttribute"`

// CheatCode see : https://schema.org/cheatCode
// Cheat codes to the game.
CheatCode *CreativeWork `json:"cheatCode"`

// ContainsSeason see : https://schema.org/containsSeason
// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
ContainsSeason *CreativeWorkSeason `json:"containsSeason"`

// Director see : https://schema.org/director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// Episode see : https://schema.org/episode
// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
Episode *Episode `json:"episode"`

// GameItem see : https://schema.org/gameItem
// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
GameItem *Thing `json:"gameItem"`

// GameLocation see : https://schema.org/gameLocation
// Real or fictional location of the game (or part of game).
GameLocation interface{} `json:"gameLocation"` // types : Place PostalAddress URL

// GamePlatform see : https://schema.org/gamePlatform
// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
GamePlatform interface{} `json:"gamePlatform"` // types : Text Thing URL

// MusicBy see : https://schema.org/musicBy
// The composer of the soundtrack.
MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

// NumberOfEpisodes see : https://schema.org/numberOfEpisodes
// The number of episodes in this season or series.
NumberOfEpisodes int `json:"numberOfEpisodes"`

// NumberOfPlayers see : https://schema.org/numberOfPlayers
// Indicate how many people can play this game (minimum, maximum, or range).
NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers"`

// NumberOfSeasons see : https://schema.org/numberOfSeasons
// The number of seasons in this series.
NumberOfSeasons int `json:"numberOfSeasons"`

// PlayMode see : https://schema.org/playMode
// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
PlayMode *GamePlayMode `json:"playMode"`

// ProductionCompany see : https://schema.org/productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc.
ProductionCompany *Organization `json:"productionCompany"`

// Quest see : https://schema.org/quest
// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
Quest *Thing `json:"quest"`

// Trailer see : https://schema.org/trailer
// The trailer of a movie or tv/radio series, season, episode, etc.
Trailer *VideoObject `json:"trailer"`

}

func (v *VideoGameSeries) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoGameSeries"

	return json.Marshal(v)
}
