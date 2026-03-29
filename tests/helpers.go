package tests

import (
	"os"

	"newgo/internal/config"
)

func InitTestEnv(module string) {

	_ = os.Setenv("ENV", "test")

	config.ResetEnvForTests()
	config.Init()
	config.SetModule(module)
}

func ShutdownSuite() {
	config.Trace("Closing All Databases ...")
}
