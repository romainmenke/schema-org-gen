package schemaorggo

import "encoding/json"

// VideoGame see : https://schema.org/VideoGame
type VideoGame struct {
	SoftwareApplication

	typeContext

	// Actor see : https://schema.org/actor
	// An actor, e.g. in tv, radio, movie, video games etc., or in an event. Actors can be associated with individual items or with a series, episode, clip. Supersedes actors (see: https://schema.org/actors).
	// types : Person
	Actor []*Person `json:"actor,omitempty"`

	// CheatCode see : https://schema.org/cheatCode
	// Cheat codes to the game.
	// types : CreativeWork
	CheatCode []*CreativeWork `json:"cheatCode,omitempty"`

	// Director see : https://schema.org/director
	// A director of e.g. tv, radio, movie, video gaming etc. content, or of an event. Directors can be associated with individual items or with a series, episode, clip. Supersedes directors (see: https://schema.org/directors).
	// types : Person
	Director []*Person `json:"director,omitempty"`

	// GamePlatform see : https://schema.org/gamePlatform
	// The electronic systems used to play video games (see: https://schema.orghttp://en.wikipedia.org/wiki/Category:Video_game_platforms).
	// types : Text Thing URL
	GamePlatform []interface{} `json:"gamePlatform,omitempty"`

	// GameServer see : https://schema.org/gameServer
	// The server on which  it is possible to play the game. Inverse property: game (see: https://schema.org/game).
	// types : GameServer
	GameServer []*GameServer `json:"gameServer,omitempty"`

	// GameTip see : https://schema.org/gameTip
	// Links to tips, tactics, etc.
	// types : CreativeWork
	GameTip []*CreativeWork `json:"gameTip,omitempty"`

	// MusicBy see : https://schema.org/musicBy
	// The composer of the soundtrack.
	// types : MusicGroup Person
	MusicBy []interface{} `json:"musicBy,omitempty"`

	// PlayMode see : https://schema.org/playMode
	// Indicates whether this game is multi-player, co-op or single-player.  The game can be marked as multi-player, co-op and single-player at the same time.
	// types : GamePlayMode
	PlayMode []*GamePlayMode `json:"playMode,omitempty"`

	// Trailer see : https://schema.org/trailer
	// The trailer of a movie or tv/radio series, season, episode, etc.
	// types : VideoObject
	Trailer []*VideoObject `json:"trailer,omitempty"`
}

func (v VideoGame) intoMap(intop *map[string]interface{}) error {
	if intop == nil {
		return nil
	}

	v.SoftwareApplication.intoMap(intop)

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
		var value interface{} = v.GameServer
		if len(v.GameServer) == 1 {
			value = v.GameServer[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gameServer"] = json.RawMessage(b)
		}
	}

	{
		var value interface{} = v.GameTip
		if len(v.GameTip) == 1 {
			value = v.GameTip[0]
		}

		b, err := json.Marshal(value)
		if err != nil {
			return err
		}

		if len(b) > 0 && string(b) != "null" {
			into["gameTip"] = json.RawMessage(b)
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

func (v VideoGame) asMap() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := v.intoMap(&data)
	if err != nil {
		return nil, err
	}

	data["@context"] = "http://schema.org"
	data["@type"] = "VideoGame"

	return data, nil
}

func (v VideoGame) MarshalJSON() ([]byte, error) {
	data, err := v.asMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(data)
}
