package event

// PlayerMessage is sent when a player receives a message. This happens for any message type, including when
// the player itself sends a message.
type PlayerMessage struct {
	// Message is the message sent by the sender.
	Message string
	// Type is the type of the message, for example 'chat'.
	Type string
	// Sender is the name of the sender of the message. This may be the name of the receiver if the receiver
	// sent a message.
	Sender string
	// Receiver is the receiver of the text message.
	Receiver string `json:",omitempty"`
}
