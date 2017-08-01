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
	globalLogger myLogger
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

	globalLogger = myLogger{lctx}
}

func Init(keyvals ...interface{}) {
	globalLogger = myLogger{log.With(globalLogger.Logger, keyvals...)}
}

type Leveler interface {
	With(keyval ...interface{}) Leveler
	Info() log.Logger
	Debug() log.Logger
	Error() log.Logger
}

type myLogger struct {
	log.Logger
}

func Info() log.Logger {
	return globalLogger.Info()
}

func (l myLogger) Info() log.Logger {
	return level.Info(l.Logger)
}

func Debug() log.Logger {
	return globalLogger.Debug()
}

func (l myLogger) Debug() log.Logger {
	return level.Debug(l.Logger)
}

func Error() log.Logger {
	return globalLogger.Error()
}

func (l myLogger) Error() log.Logger {
	return level.Error(l.Logger)
}

func With(keyvals ...interface{}) Leveler {
	lctx := log.With(globalLogger.Logger, keyvals...)
	return myLogger{lctx}
}

func (l myLogger) With(keyvals ...interface{}) Leveler {
	lctx := log.With(l.Logger, keyvals...)
	return myLogger{lctx}
}

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
