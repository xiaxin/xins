package core

import (
	"math"
	"sync"
	"time"

	rawCtx "context"
)

const abortIndex int = math.MaxInt >> 1

var _ Context = &context{}

/**
 * TODO context 的使用
 */

type Context interface {
	rawCtx.Context
	WithContext(ctx rawCtx.Context) Context

	Message() any
	SetMessage(message any) Context

	Session() Session
	SetSession(sess Session) Context

	Set(key string, value any)
	// GET
	Get(key string) (value any, exists bool)
	GetString(key string) (s string)
	GetBool(key string) (s bool)
	GetInt(key string) (i int)
	GetInt64(key string) (i64 int64)
	GetUint(key string) (ui uint)
	GetUint64(key string) (ui64 uint64)
	GetFloat64(key string) (f64 float64)
	GetTime(key string) (t time.Time)
	GetDuration(key string) (d time.Duration)
	GetStringSlice(key string) (ss []string)
	GetStringMap(key string) (sm map[string]any)
	GetStringMapString(key string) (sms map[string]string)
	GetStringMapStringSlice(key string) (smss map[string][]string)

	Remove(key string)

	Send() bool
	Next() // TODO 添加说明

	SetHandles(handles []RouteFunc)

	IsAborted() bool
	Abort()
}

type context struct {
	rawCtx rawCtx.Context

	session Session
	message interface{}

	m      sync.RWMutex
	keyval map[string]interface{}

	index   int
	handles []RouteFunc
}

func NewContext() *context {
	return &context{
		rawCtx: rawCtx.Background(),
		keyval: make(map[string]any),

		index: -1,
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
func (c *context) Get(key string) (value any, exists bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	value, exists = c.keyval[key]
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

func (c *context) Set(key string, val any) {
	c.m.Lock()
	defer c.m.Unlock()

	if c.keyval == nil {
		c.keyval = make(map[string]any)
	}
	c.keyval[key] = val

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
	c.index = -1
}

func (c *context) Send() bool {
	return c.session.Send(c)
}

func (c *context) SetHandles(handles []RouteFunc) {
	c.handles = handles
}

func (c *context) Next() {
	c.index++
	for c.index < len(c.handles) {
		c.handles[c.index](c)
		c.index++
	}
}

func (c *context) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *context) Abort() {
	c.index = abortIndex
}

// GET ...

func (c *context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

func (c *context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

func (c *context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

func (c *context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

func (c *context) GetUint(key string) (ui uint) {
	if val, ok := c.Get(key); ok && val != nil {
		ui, _ = val.(uint)
	}
	return
}

func (c *context) GetUint64(key string) (ui64 uint64) {
	if val, ok := c.Get(key); ok && val != nil {
		ui64, _ = val.(uint64)
	}
	return
}

func (c *context) GetFloat64(key string) (f64 float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

func (c *context) GetTime(key string) (t time.Time) {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

func (c *context) GetDuration(key string) (d time.Duration) {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

func (c *context) GetStringSlice(key string) (ss []string) {
	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

func (c *context) GetStringMap(key string) (sm map[string]any) {
	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]any)
	}
	return
}

func (c *context) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

func (c *context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}
