package core

import (
	rawCtx "context"
	"testing"

	"xins/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_session_SetID(t *testing.T) {
	sess := NewSession(nil, nil)
	_, ok := sess.ID().(string)
	assert.True(t, ok)
	sess.SetID(123)
	assert.Equal(t, sess.ID(), 123)

}

func TestSessionConn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conn := mock.NewMockConn(ctrl)
	s := NewSession(conn, nil)
	assert.Equal(t, s.Conn(), conn)
}

// TODO 调试失败
func TestSessionSend(t *testing.T) {
	t.Run("session closed", func(t *testing.T) {
		message := NewMessage(1, []byte("test"))
		sess := NewSession(nil, nil)
		sess.Close() // close session
		ctx := sess.NewContext()
		assert.False(t, ctx.SetMessage(message).Send())
	})

	t.Run("ctx is done", func(t *testing.T) {
		sess := NewSession(nil, nil)

		ctx, cancel := rawCtx.WithCancel(rawCtx.Background())

		c := sess.NewContext().WithContext(ctx)
		done := make(chan struct{})
		go func() {
			assert.False(t, c.Send())
			close(done)
		}()

		cancel()
		<-done
	})
}
