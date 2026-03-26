package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type AppConfig struct {
	Name     string
	Env      string
	LogLevel string
	Port     string
	Root     string
}

func Load() (AppConfig, error) {
	_ = loadDotEnv(".env")

	root, err := getRoot()
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Name:     getEnv("APP_NAME", "template_go"),
		Env:      getEnv("APP_ENV", "production"),
		LogLevel: strings.ToUpper(getEnv("LOG_LEVEL", "INFO")),
		Port:     getEnv("APP_PORT", "8000"),
		Root:     root,
	}, nil
}

func loadDotEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "" {
			continue
		}

		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}

	return scanner.Err()
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func getRoot() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if info, err := os.Stat(filepath.Join(current, ".git")); err == nil && info.IsDir() {
			return current, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			return "", errors.New(".git directory not found in any parent directory")
		}

		current = parent
	}
}
