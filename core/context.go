package core

import (
	"sync"
	"time"

	rawCtx "context"
)

/**
 * TODO context 的使用
 */

type Context interface {
	rawCtx.Context
	WithContext(ctx rawCtx.Context) Context

	Message() interface{}
	SetMessage(message any) Context

	Session() Session
	SetSession(sess Session) Context

	Get(key string) (value interface{}, exists bool)
	Set(key string, value interface{})
	Remove(key string)

	Send() bool
}

type context struct {
	rawCtx rawCtx.Context

	session Session
	message interface{}

	m      sync.RWMutex
	keyval map[string]interface{}
}

func NewContext() *context {
	return &context{
		rawCtx: rawCtx.Background(),
		keyval: make(map[string]any),
	}
}

func (c *context) Deadline() (time.Time, bool) {
	return c.rawCtx.Deadline()
}

func (c *context) Done() <-chan struct{} {
	return c.rawCtx.Done()
}
func (c *context) Err() error {
	return c.rawCtx.Err()
}

func (c *context) Message() any {
	return c.message
}

func (c *context) SetMessage(message any) Context {
	c.message = message
	return c
}

func (c *context) Session() Session {
	return c.session
}

func (c *context) SetSession(sess Session) Context {
	c.session = sess
	return c
}
func (c *context) Get(key string) (value interface{}, exists bool) {
	c.m.RLock()
	value, exists = c.keyval[key]
	c.m.RUnlock()
	return
}

func (c *context) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val, _ := c.Get(keyAsString)
		return val
	}
	return nil
}

func (c *context) WithContext(ctx rawCtx.Context) Context {
	c.rawCtx = ctx
	return c
}

func (c *context) Set(key string, value interface{}) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.keyval == nil {
		c.keyval = make(map[string]interface{})
	}
	c.keyval[key] = value

}

func (c *context) Remove(key string) {
	c.m.Lock()
	defer c.m.Unlock()

	delete(c.keyval, key)
}

func (c *context) reset() {
	c.session = nil
	c.message = nil
	c.keyval = nil
}

func (c *context) Send() bool {
	return c.session.Send(c)
}
