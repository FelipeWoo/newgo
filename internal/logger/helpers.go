package logger

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Trace(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Trace().Str("module", getModule()).Msg(TraceSym + " " + msg)
}

func Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Debug().Str("module", getModule()).Msg(DebugSym + " " + msg)
}

func Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Info().Str("module", getModule()).Msg(InfoSym + " " + msg)
}

func Success(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Info().Str("module", getModule()).Msg(SuccessSym + " " + msg)
}

func Warn(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Warn().Str("module", getModule()).Msg(WarnSym + " " + msg)
}

// fail-level logging with optional error
func Fail(format string, args ...interface{}) {
	logWithError(ErrorSym, log.Error(), format, args...)
}

// fatal-level logging with optional error
func Fatal(format string, args ...interface{}) {
	logWithError(FatalSym, log.Fatal(), format, args...)
}

// panic-level logging with optional error
func Panic(format string, args ...interface{}) {
	logWithError(PanicSym, log.Panic(), format, args...)
}

func logWithError(symbol string, event *zerolog.Event, format string, args ...interface{}) {
	var err error
	var cleanArgs []interface{}

	for _, arg := range args {
		if e, ok := arg.(error); ok && err == nil {
			err = e
		} else {
			cleanArgs = append(cleanArgs, arg)
		}
	}

	msg := fmt.Sprintf(format, cleanArgs...)
	entry := event.Str("module", getModule())
	if err != nil {
		entry = entry.Err(err)
	}
	entry.Msg(symbol + " " + msg)
}

const colorOrange = "\033[38;5;208m" // ANSI extended color for naranja

// Imprime un encabezado con color naranja
func LogTestHeader(t *testing.T) {

	var name string
	pc, _, _, ok := runtime.Caller(1) // 1 = función que llamó a esta
	if !ok {
		name = "unknown"
	}
	fullName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fullName, ".")
	name = parts[len(parts)-1]

	line := strings.Repeat("*", 120)
	header := fmt.Sprintf("\n%s%s\n🟧 %s\n%s%s", colorOrange, line, name, line, colorReset)
	t.Log(header)
}
