package safe

import "fmt"

type sopt struct {
	logger   func(error)
	callback func()
	final    func()
}

func newSopt(opts ...Option) *sopt {
	so := new(sopt)
	for _, fn := range opts {
		fn(so)
	}
	return so
}

type Option func(*sopt)

func OptLog(fn func(error)) Option {
	return func(s *sopt) {
		s.logger = fn
	}
}

func OptCallback(fn func()) Option {
	return func(s *sopt) {
		s.callback = fn
	}
}

func OptFinal(fn func()) Option {
	return func(s *sopt) {
		s.final = fn
	}
}

func Go(fn func(), opts ...Option) {
	go Run(fn, opts...)
}

func Run(fn func(), opts ...Option) {
	so := newSopt(opts...)
	defer func() {
		if so.final != nil {
			safeRun(so.final, so.logger)
		}
	}()

	safeRun(func() {
		fn()
		if so.callback != nil {
			so.callback()
		}
	}, so.logger)
}

func safeRun(fn func(), logger func(error)) {
	defer func() {
		if err := recover(); err != nil && logger != nil {
			safeRun(func() {
				logger(fmt.Errorf("%+v", err))
			}, nil)
		}
	}()
	fn()
}
