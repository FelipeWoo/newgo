package bootstrap

import (
	// "newgo/internal/cache"
	"newgo/internal/config"
	// "newgo/internal/db"
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

	// // 4. Conectar servicios
	// db.ConnectNamed("MAIN")
	// cache.ConnectNamed("CACHE")

	logger.Success("✓ All services initialized")
}

func Shutdown() {
	// db.CloseAll()
	// cache.CloseAll()
}
