package token

import "time"

// this interface is to manages token types
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
