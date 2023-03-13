package xins

type Options struct {
	protocol Protocol // 协议

	onSessionStart func(session *Session) // session 开始回调
	onSessionStop  func(session *Session) // session 结束回调
}

type Option func(*Options)

func ServerProtocol(p Protocol) Option {
	return func(o *Options) {
		o.protocol = p
	}
}

func SessionOnStart(f func(session *Session)) Option {
	return func(o *Options) {
		o.onSessionStart = f
	}
}

func SessionOnStop(f func(session *Session)) Option {
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

func (o *Options) Protocol() Protocol {
	return o.protocol
}

func (o *Options) OnSessionStart() func(session *Session) {
	return o.onSessionStart
}

func (o *Options) OnSessionStop() func(session *Session) {
	return o.onSessionStop
}
