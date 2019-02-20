package command

// LocalPlayerName is sent by the server to find out the player name of the player connected to the websocket.
// This command is executed automatically by the server and stored in a field of the Player struct.
type LocalPlayerName struct {
	// LocalPlayerName is the name of the player connected.
	LocalPlayerName string `json:"localplayername"`
	// StatusCode is the status code of the command. This is generally 0 for this command.
	StatusCode int `json:"statusCode"`
	// StatusMessage is the same as LocalPlayerName for this command.
	StatusMessage string `json:"statusMessage"`
}