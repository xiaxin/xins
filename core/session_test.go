package core

import (
	rawCtx "context"
	"net"
	"sync"
	"testing"
	"time"

	"xins/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSessionNew(t *testing.T) {
	t.Run("default new", func(t *testing.T) {
		sess := NewSession(nil, nil, nil)

		assert.Nil(t, sess.ID())
		assert.Nil(t, sess.Conn())
		assert.Nil(t, sess.Protocol())
	})

	t.Run("uuid new", func(t *testing.T) {
		sess := NewSessionByUUID(nil, nil)

		assert.NotNil(t, sess.ID())
		assert.Nil(t, sess.Conn())
		assert.Nil(t, sess.Protocol())
	})
}

func TestSessionClose(t *testing.T) {
	sess := NewSessionByUUID(nil, nil)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sess.Close() // goroutine safe
		}()
	}
	wg.Wait()
	_, ok := <-sess.closed
	assert.False(t, ok)
}

func TestSessionRead(t *testing.T) {
	t.Run("connection set read timeout failed", func(t *testing.T) {
		p1, _ := net.Pipe()
		_ = p1.Close()
		sess := NewSessionByUUID(p1, nil)
		go sess.read(time.Millisecond)
		<-sess.closed
	})

	t.Run("connection read timeout", func(t *testing.T) {
		p1, _ := net.Pipe()
		sess := NewSessionByUUID(p1, NewProtocol())
		done := make(chan struct{})
		go func() {
			sess.read(time.Millisecond * 10)
			close(done)
		}()
		<-done
	})
}

func TestSessionSetID(t *testing.T) {
	sess := NewSessionByUUID(nil, nil)
	_, ok := sess.ID().(string)
	assert.True(t, ok)
	sess.SetID(123)
	assert.Equal(t, sess.ID(), 123)

}

func TestSessionConn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conn := mock.NewMockConn(ctrl)
	s := NewSessionByUUID(conn, nil)
	assert.Equal(t, s.Conn(), conn)
}

func TestSessionSend(t *testing.T) {
	t.Run("session closed", func(t *testing.T) {
		message := NewMessage(1, []byte("test"))
		sess := NewSessionByUUID(nil, nil)
		sess.Close() // close session

		ctx := sess.NewContext()
		assert.False(t, ctx.SetMessage(message).Send())
	})

	t.Run("ctx is done", func(t *testing.T) {
		sess := NewSessionByUUID(nil, nil)

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

	t.Run("when send succeed", func(t *testing.T) {
		sess := NewSessionByUUID(nil, nil)
		sess.ctxQueue = make(chan Context) // no buffer

		go func() { <-sess.ctxQueue }()

		assert.True(t, sess.NewContext().SetMessage(NewMessage(1, []byte("test"))).Send())
		sess.Close()
	})
}
