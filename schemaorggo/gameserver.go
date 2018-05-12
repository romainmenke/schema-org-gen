package schemaorggo

import "encoding/json"

// GameServer see : https://schema.org/GameServer
type GameServer struct {
	Intangible

	typeContext

	// Game see : https://schema.org/game
	// Video game which is played on this server. Inverse property: gameServer (see: https://schema.org/gameServer).
	Game *VideoGame `json:"game,omitempty"` // types : VideoGame

	// PlayersOnline see : https://schema.org/playersOnline
	// Number of players on the server.
	PlayersOnline float64 `json:"playersOnline,omitempty"` // types : Integer

	// ServerStatus see : https://schema.org/serverStatus
	// Status of a game server.
	ServerStatus *GameServerStatus `json:"serverStatus,omitempty"` // types : GameServerStatus

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
