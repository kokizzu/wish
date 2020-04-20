package wish

import (
	"encoding/base64"
	"fmt"

	"github.com/gliderlabs/ssh"
)

type Session struct {
	ssh.Session
}

type SessionHandler func(s Session)

func (s *Session) KeyText() (string, error) {
	if s.Session.PublicKey() == nil {
		return "", fmt.Errorf("Session doesn't have public key")
	}
	return base64.StdEncoding.EncodeToString(s.Session.PublicKey().Marshal()), nil
}