package protocol

// EncryptResponse is sent by the client on connection to deliver the public key.
type EncryptResponse struct {
	// PublicKey is the client's public key.
	PublicKey string `json:"publicKey"`
}
