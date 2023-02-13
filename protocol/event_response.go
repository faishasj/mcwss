package protocol

import "encoding/json"

// EventResponse is sent by the client. It holds information about a particular event listened on by the
// sever.
type EventResponse json.RawMessage
