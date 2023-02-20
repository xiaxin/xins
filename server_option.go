package xins

type Options struct {
	protocol Protocol
}

type Option func(*Options)

func ServerProtocol(p Protocol) Option {
	return func(o *Options) {
		o.protocol = p
	}
}

func newOptions(opt ...Option) *Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}
	if opts.protocol == nil {
		// TODO 报错
	}

	return &opts
}

func (o *Options) Protocol() Protocol {
	return o.protocol
}
