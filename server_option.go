package xins

import "xins/core"

type Options struct {
	protocol core.Protocol // 协议

	onSessionStart func(session core.Session) // session 开始回调
	onSessionStop  func(session core.Session) // session 结束回调
}

type Option func(*Options)

func ServerProtocol(p core.Protocol) Option {
	return func(o *Options) {
		o.protocol = p
	}
}

func SessionOnStart(f func(session core.Session)) Option {
	return func(o *Options) {
		o.onSessionStart = f
	}
}

func SessionOnStop(f func(session core.Session)) Option {
	return func(o *Options) {
		o.onSessionStop = f
	}
}

func newOptions(opt ...Option) *Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	return &opts
}

func (o *Options) Protocol() core.Protocol {
	return o.protocol
}

func (o *Options) OnSessionStart() func(session core.Session) {
	return o.onSessionStart
}

func (o *Options) OnSessionStop() func(session core.Session) {
	return o.onSessionStop
}
