package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)

	//VerifyToken checks if the token is vaild or not
	VerifyToken(token string) (*Payload, error)
}
