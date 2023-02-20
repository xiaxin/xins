package xins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionSetID(t *testing.T) {
	session := NewSession(nil, NewDefaultProtocol())
	session.SetID("session id")
	assert.Equal(t, session.ID(), "session id")
}
