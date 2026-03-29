package tests

import (
	"os"
	"testing"

	"newgo/internal/config"
)

func TestMain(m *testing.M) {

	_ = os.Setenv("ENV", "test")

	config.ResetEnvForTests()
	config.Init()
	config.SetModule("test_main")

	code := m.Run()

	ShutdownSuite()

	os.Exit(code)

}
