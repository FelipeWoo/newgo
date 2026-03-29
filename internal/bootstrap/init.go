package bootstrap

import (
	"newgo/internal/config"
)

func InitAll() {
	config.Init()
}

func Shutdown() {
	config.Shutdown()
}
