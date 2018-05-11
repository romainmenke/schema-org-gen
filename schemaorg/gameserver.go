package schemaorg

import "encoding/json"

// GameServer see : https://schema.org/GameServer
type GameServer struct {

typeContext

Intangible

// Game see : https://schema.org/game
// Video game which is played on this server. Inverse property: gameServer (see: https://schema.org/gameServer).
Game *VideoGame `json:"game"`

// PlayersOnline see : https://schema.org/playersOnline
// Number of players on the server.
PlayersOnline int `json:"playersOnline"`

// ServerStatus see : https://schema.org/serverStatus
// Status of a game server.
ServerStatus *GameServerStatus `json:"serverStatus"`

}

func (v *GameServer) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "GameServer"

	return json.Marshal(v)
}
