package graceful

import (
	"context"
	"github.com/realHoangHai/authenticator/pkg/log"
	"github.com/sirupsen/logrus"
)

// Option interface for configuration.
type Option interface {
	Apply(*Options)
}

// OptionFunc is a function that configures a graceful shutdown.
type OptionFunc func(*Options)

// Apply calls f(option)
func (f OptionFunc) Apply(option *Options) {
	f(option)
}

// Options for graceful shutdown
type Options struct {
	ctx    context.Context
	logger *logrus.Logger
}

// WithContext custom context
func WithContext(ctx context.Context) Option {
	return OptionFunc(func(o *Options) {
		o.ctx = ctx
	})
}

// WithLogger custom logger
func WithLogger(logger *logrus.Logger) Option {
	return OptionFunc(func(o *Options) {
		o.logger = logger
	})
}

func newOptions(opts ...Option) Options {
	defaultOpts := Options{
		ctx:    context.Background(),
		logger: log.NewLogger(log.DebugLevel, "graceful"),
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt.Apply(&defaultOpts)
	}

	return defaultOpts
}
