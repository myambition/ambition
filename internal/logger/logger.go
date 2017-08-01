package logger

import (
	"fmt"
	"os"
	"time"

	stdlog "log"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/go-stack/stack"
	//. "github.com/y0ssar1an/q"
)

var (
	globalLogger levelLogger
)

func init() {
	l := log.NewJSONLogger(os.Stdout)
	lctx := log.With(l,
		"ts", log.DefaultTimestampUTC,
		"caller", caller(3),
		"function", function(3),
	)

	stdlog.SetFlags(0)
	stdlog.SetOutput(log.NewStdlibAdapter(lctx))

	globalLogger = levelLogger{lctx}
}

// Init adds values to the globalLogger, having them be on all logs
func Init(keyvals ...interface{}) {
	globalLogger = levelLogger{log.With(globalLogger.Logger, keyvals...)}
}

func Info() log.Logger {
	return globalLogger.Info()
}
func Debug() log.Logger {
	return globalLogger.Debug()
}
func Error() log.Logger {
	return globalLogger.Error()
}
func With(keyvals ...interface{}) Leveler {
	return globalLogger.With(keyvals...)
}
func WithCustomDepth(depth int, keyvals ...interface{}) Leveler {
	return globalLogger.WithCustomDepth(depth, keyvals...)
}
func LogError(err error, keyvals ...interface{}) error {
	return globalLogger.LogError(err, keyvals...)
}

func (l levelLogger) Info() log.Logger {
	return level.Info(l.Logger)
}
func (l levelLogger) Debug() log.Logger {
	return level.Debug(l.Logger)
}
func (l levelLogger) Error() log.Logger {
	return level.Error(l.Logger)
}
func (l levelLogger) With(keyvals ...interface{}) Leveler {
	lctx := log.With(l.Logger, keyvals...)
	return levelLogger{lctx}
}
func (l levelLogger) WithCustomDepth(depth int, keyvals ...interface{}) Leveler {
	lctx := log.With(l.Logger,
		"caller", caller(depth),
		"function", function(depth),
	)
	lctx = log.With(lctx, keyvals...)
	return levelLogger{lctx}
}
func (l levelLogger) LogError(err error, keyvals ...interface{}) error {
	return l.WithCustomDepth(5, keyvals...).Error().Log("err", err)
}

// Leveler forces a levelLogger to chose a level before being able to log
type Leveler interface {
	With(keyval ...interface{}) Leveler
	Info() log.Logger
	Debug() log.Logger
	Error() log.Logger
}

type levelLogger struct {
	log.Logger
}

// Helpers

func caller(depth int) log.Valuer {
	return func() interface{} {
		return fmt.Sprintf("%+v", stack.Caller(depth))
	}
}

func function(depth int) log.Valuer {
	return func() interface{} {
		return fmt.Sprintf("%n", stack.Caller(depth))
	}
}

func defaultTSNano() log.Valuer {
	return func() interface{} {
		return time.Now().UTC().Format(time.RFC3339Nano)
	}
}
