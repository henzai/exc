package exc

import "fmt"

type Option interface {
	apply(*Error)
}

type kindOption struct {
	kind Kind
}

func (o kindOption) apply(e *Error) {
	e.kind = o.kind
}

func WithKind(k Kind) Option {
	return kindOption{
		kind: k,
	}
}

type msgOption string

func (o msgOption) apply(e *Error) {
	e.msg = string(o)
}

func WithMessage(msg string) Option {
	return msgOption(msg)
}

func WithMessagef(msg string, args ...interface{}) Option {
	return msgOption(fmt.Sprintf(msg, args...))
}
