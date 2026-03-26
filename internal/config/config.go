package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"newgo/internal/logger"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName  string
	Env      string
	Port     int
	LogLevel string
}

var Config AppConfig

var alreadyLoaded bool

func LoadEnvOnce() {
	if alreadyLoaded {
		logger.Trace("LoadEnv skipped (already loaded)")
		return
	}
	alreadyLoaded = true

	loadEnvInternal()
}

func ResetEnvForTests() {
	alreadyLoaded = false
}

func loadEnvInternal() {
	logger.SetModule("config")

	root, err := findProjectRoot()
	if err != nil {
		logger.Fatal("Cannot locate project root: %v", err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	var envFile string
	switch env {
	case "test":
		envFile = filepath.Join(root, "shared", ".env.test")
	case "production":
		envFile = filepath.Join(root, "shared", ".env.prod")
	default:
		envFile = filepath.Join(root, "shared", ".env")
	}

	_ = godotenv.Load(envFile)

	Config.AppName = resolveAppName(root)
	Config.Env = env
	Config.Port = getInt("PORT", 8080)
	Config.LogLevel = getString("LOG_LEVEL", "info")

	logger.Success("Config loaded for [%s] on port %d in [%s] level mode.", Config.Env, Config.Port, Config.LogLevel)
}

func getRequired(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logger.Fatal("Required environment variable %s is missing", key)
	}
	return val
}

func getString(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func getInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		logger.Fatal("Invalid int for %s: %v", key, err)
	}
	return i
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
			break // Ya estás en la raíz del sistema
		}
		current = parent
	}

	return "", os.ErrNotExist
}

func resolveAppName(root string) string {
	if envName := getString("APP_NAME", ""); envName != "" {
		return envName
	}

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
					return filepath.Base(modulePath)
				}
			}
		}
	}

	return filepath.Base(root)
}
