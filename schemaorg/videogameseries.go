package schemaorg

import "encoding/json"

// VideoGameSeries see : https://schema.org/VideoGameSeries
type VideoGameSeries struct {

typeContext

CreativeWorkSeries

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// CharacterAttribute see : /characterAttribute
// A piece of data that represents a particular aspect of a fictional character (skill, power, character points, advantage, disadvantage).
CharacterAttribute *Thing `json:"characterAttribute"`

// CheatCode see : /cheatCode
// Cheat codes to the game.
CheatCode *CreativeWork `json:"cheatCode"`

// ContainsSeason see : /containsSeason
// A season that is part of the media series. Supersedes season (see: https://schema.org/season).
ContainsSeason *CreativeWorkSeason `json:"containsSeason"`

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// Episode see : /episode
// An episode of a tv, radio or game media within a series or season. Supersedes episodes (see: https://schema.org/episodes).
Episode *Episode `json:"episode"`

// GameItem see : /gameItem
// An item is an object within the game world that can be collected by a player or, occasionally, a non-player character.
GameItem *Thing `json:"gameItem"`

// GameLocation see : /gameLocation
// Real or fictional location of the game (or part of game).
GameLocation interface{} `json:"gameLocation"` // types : Place PostalAddress URL

// GamePlatform see : /gamePlatform
// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
GamePlatform interface{} `json:"gamePlatform"` // types : Text Thing URL

// MusicBy see : /musicBy
// The composer of the soundtrack.
MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

// NumberOfEpisodes see : /numberOfEpisodes
// The number of episodes in this season or series.
NumberOfEpisodes int `json:"numberOfEpisodes"`

// NumberOfPlayers see : /numberOfPlayers
// Indicate how many people can play this game (minimum, maximum, or range).
NumberOfPlayers *QuantitativeValue `json:"numberOfPlayers"`

// NumberOfSeasons see : /numberOfSeasons
// The number of seasons in this series.
NumberOfSeasons int `json:"numberOfSeasons"`

// PlayMode see : /playMode
// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
PlayMode *GamePlayMode `json:"playMode"`

// ProductionCompany see : /productionCompany
// The production company or studio responsible for the item e.g. series, video game, episode etc.
ProductionCompany *Organization `json:"productionCompany"`

// Quest see : /quest
// The task that a player-controlled character, or group of characters may complete in order to gain a reward.
Quest *Thing `json:"quest"`

// Trailer see : /trailer
// The trailer of a movie or tv/radio series, season, episode, etc.
Trailer *VideoObject `json:"trailer"`

}

func (v *VideoGameSeries) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoGameSeries"

	return json.Marshal(v)
}
