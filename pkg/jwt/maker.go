package jwttoken

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and Duration
	CreateToken(info map[string]interface{}) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
