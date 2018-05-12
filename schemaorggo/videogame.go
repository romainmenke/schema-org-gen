package schemaorggo

import "encoding/json"

// VideoGame see : https://schema.org/VideoGame
type VideoGame struct {
	SoftwareApplication

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor *Person `json:"actor,omitempty"`

	// CheatCode see : https://schema.org/cheatCode
	// Cheat codes to the game.
	// types : CreativeWork
	CheatCode *CreativeWork `json:"cheatCode,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director *Person `json:"director,omitempty"`

	// GamePlatform see : https://schema.org/gamePlatform
	// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	// types : Text Thing URL
	GamePlatform interface{} `json:"gamePlatform,omitempty"`

	// GameServer see : https://schema.org/gameServer
	// The server on which  it is possible to play the game. Inverse property: game (see: https://schema.org/game).
	// types : GameServer
	GameServer *GameServer `json:"gameServer,omitempty"`

	// GameTip see : https://schema.org/gameTip
	// Links to tips, tactics, etc.
	// types : CreativeWork
	GameTip *CreativeWork `json:"gameTip,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy interface{} `json:"musicBy,omitempty"`

	// PlayMode see : https://schema.org/playMode
	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	// types : GamePlayMode
	PlayMode *GamePlayMode `json:"playMode,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer *VideoObject `json:"trailer,omitempty"`
}

func (v VideoGame) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "VideoGame"

	return json.Marshal(v)
}

func (v *VideoGame) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "VideoGame"

	return json.Marshal(*v)
}
