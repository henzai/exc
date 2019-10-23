package exc

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

type Error struct {
	msg   string
	kind  Kind
	err   error
	frame xerrors.Frame
}

type Kind string

func (e *Error) LastKind() Kind {
	if e.kind != "" {
		return e.kind
	}
	var baseErr *Error
	if ok := errors.As(e.err, &baseErr); ok {
		return baseErr.LastKind()
	}
	return ""
}

func (e *Error) Error() string { return e.msg }

func (e *Error) Unwrap() error { return e.err }

func (e *Error) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *Error) FormatError(p xerrors.Printer) error {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

func New(opts ...Option) error {
	err := &Error{
		err:   nil,
		frame: xerrors.Caller(1),
	}
	for _, opt := range opts {
		opt.apply(err)
	}
	return err
}

func Wrap(err error, opts ...Option) error {
	if err == nil {
		return nil
	}
	e := &Error{
		err:   err,
		frame: xerrors.Caller(1),
	}
	for _, opt := range opts {
		opt.apply(e)
	}
	return e
}

func Is(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.kind != "" {
		return e.kind == kind
	}
	if e.err != nil {
		return Is(kind, e.err)
	}
	return false
}
