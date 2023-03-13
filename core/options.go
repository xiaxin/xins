package core

type Option func(*Options)

type Options struct {
	protocol Protocol // 协议

	beforeSession func(session Session) // session 开始回调
	afterSession  func(session Session) // session 结束回调
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

func (o *Options) BeforeSession() func(session Session) {
	return o.BeforeSession()
}

func (o *Options) AfterSession() func(session Session) {
	return o.afterSession
}
