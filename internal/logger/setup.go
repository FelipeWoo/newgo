package logger

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	colorReset   = "\033[0m"
	colorGray    = "\033[90m"
	colorBlue    = "\033[34m"
	colorWhite   = "\033[37m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorRed     = "\033[31m"
	colorBoldRed = "\033[1;31m"
)

func SetupLoggerWriters() zerolog.Logger {
	_ = os.MkdirAll("logs", os.ModePerm)
	file, _ := os.OpenFile(filepath.Join("logs", resolveLogFileName()), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	console := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		FormatMessage: func(i interface{}) string {
			msg := fmt.Sprintf("%v", i)
			switch {
			case strings.Contains(msg, "✓"):
				return fmt.Sprintf("%s%s%s", colorGreen, msg, colorReset)
			case strings.Contains(msg, "⚠"):
				return fmt.Sprintf("%s%s%s", colorYellow, msg, colorReset)
			case strings.Contains(msg, "✘"):
				return fmt.Sprintf("%s%s%s", colorRed, msg, colorReset)
			case strings.Contains(msg, "‼‼"):
				return fmt.Sprintf("%s%s%s", colorBoldRed, msg, colorReset)
			case strings.Contains(msg, "‼"):
				return fmt.Sprintf("%s%s%s", colorRed, msg, colorReset)
			case strings.Contains(msg, "→"):
				return fmt.Sprintf("%s%s%s", colorGray, msg, colorReset)
			case strings.Contains(msg, "»"):
				return fmt.Sprintf("%s%s%s", colorBlue, msg, colorReset)
			default:
				return fmt.Sprintf("%s%s%s", colorWhite, msg, colorReset)
			}
		},
	}

	fileWriter := zerolog.ConsoleWriter{
		Out:        file,
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05",
	}

	multi := zerolog.MultiLevelWriter(console, fileWriter)

	// Configura base sin log level aún
	zerolog.TimeFieldFormat = time.RFC3339
	return zerolog.New(multi).With().Timestamp().Logger()
}

func ApplyLogLevelFromEnv(logger zerolog.Logger) zerolog.Logger {
	level, err := zerolog.ParseLevel(strings.ToLower(os.Getenv("LOG_LEVEL")))
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	return logger
}

func resolveLogFileName() string {
	if appName := strings.TrimSpace(os.Getenv("APP_NAME")); appName != "" {
		return appName + ".log"
	}

	root, err := findProjectRoot()
	if err == nil {
		modFile, err := os.Open(filepath.Join(root, "go.mod"))
		if err == nil {
			defer modFile.Close()

			scanner := bufio.NewScanner(modFile)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if strings.HasPrefix(line, "module ") {
					modulePath := strings.TrimSpace(strings.TrimPrefix(line, "module "))
					modulePath = strings.Trim(modulePath, "\"")
					if modulePath != "" {
						return filepath.Base(modulePath) + ".log"
					}
				}
			}
		}
		return filepath.Base(root) + ".log"
	}

	return "app.log"
}

func findProjectRoot() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			return current, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return "", os.ErrNotExist
}

/*
log.Trace().Str("module", "init").Msg("→ loading .env file")
log.Debug().Str("module", "db").Msg("» preparing SQL statement")
log.Info().Str("module", "server").Msg("i server listening on :8080")
log.Warn().Str("module", "auth").Msg("⚠ token missing")
log.Error().Str("module", "api").Err(err).Msg("✘ cannot fetch external data")
log.Fatal().Str("module", "main").Msg("‼ port already in use")
log.Panic().Str("module", "handler").Msg("‼‼ nil pointer dereference")
log.Info().Str("module", "billing").Msg("✓ Payment processed")
*/
