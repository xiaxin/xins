package xins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionSetID(t *testing.T) {
	session := NewSession(nil, NewDefaultPacker())
	session.SetID("session id")
	assert.Equal(t, session.ID(), "session id")
}
