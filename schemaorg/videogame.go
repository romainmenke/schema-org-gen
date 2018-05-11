package schemaorg

import "encoding/json"

// VideoGame see : https://schema.org/VideoGame
type VideoGame struct {

typeContext

SoftwareApplication

// Actor see : /actor
// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
Actor *Person `json:"actor"`

// CheatCode see : /cheatCode
// Cheat codes to the game.
CheatCode *CreativeWork `json:"cheatCode"`

// Director see : /director
// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
Director *Person `json:"director"`

// GamePlatform see : /gamePlatform
// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
GamePlatform interface{} `json:"gamePlatform"` // types : Text Thing URL

// GameServer see : /gameServer
// The server on which  it is possible to play the game. Inverse property: game (see: https://schema.org/game).
GameServer *GameServer `json:"gameServer"`

// GameTip see : /gameTip
// Links to tips, tactics, etc.
GameTip *CreativeWork `json:"gameTip"`

// MusicBy see : /musicBy
// The composer of the soundtrack.
MusicBy interface{} `json:"musicBy"` // types : MusicGroup Person

// PlayMode see : /playMode
// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
PlayMode *GamePlayMode `json:"playMode"`

// Trailer see : /trailer
// The trailer of a movie or tv/radio series, season, episode, etc.
Trailer *VideoObject `json:"trailer"`

}

func (v *VideoGame) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoGame"

	return json.Marshal(v)
}
