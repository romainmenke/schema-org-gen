package schemaorggo

import "encoding/json"

// GameServer see : https://schema.org/GameServer
type GameServer struct {
	Intangible

	typeContext

	// Game see : https://schema.org/game
	// Video game which is played on this server. Inverse property: gameServer (see: https://schema.org/gameServer).
	// types : VideoGame
	Game *VideoGame `json:"game,omitempty"`

	// PlayersOnline see : https://schema.org/playersOnline
	// Number of players on the server.
	// types : Integer
	PlayersOnline float64 `json:"playersOnline,omitempty"`

	// ServerStatus see : https://schema.org/serverStatus
	// Status of a game server.
	// types : GameServerStatus
	ServerStatus *GameServerStatus `json:"serverStatus,omitempty"`
}

func (v GameServer) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GameServer"

	return json.Marshal(v)
}

func (v *GameServer) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "GameServer"

	return json.Marshal(*v)
}
