package bootstrap

import (
	"newgo/internal/config"
	"newgo/internal/logger"

	"github.com/rs/zerolog/log"
)

func InitAll() {
	// 1. Estética primero
	baseLogger := logger.SetupLoggerWriters()
	log.Logger = baseLogger // Temporal: sin log level aún

	// 2. Cargar entorno (ya tendrá color, símbolo, salida a archivo)
	config.LoadEnvOnce()

	// 3. Aplicar nivel (debug/info) una vez que LOG_LEVEL ya está disponible
	log.Logger = logger.ApplyLogLevelFromEnv(baseLogger)

	logger.Success("✓ All services initialized")
}

func Shutdown() {
	// ...

}
