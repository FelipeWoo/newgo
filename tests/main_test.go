package tests

import (
	"os"
	"testing"

	"newgo/internal/config"
	// "newgo/internal/db"
	"newgo/internal/logger"

	"github.com/rs/zerolog/log"
)

func TestMain(m *testing.M) {
	// 0. Setear entorno antes que nada
	_ = os.Setenv("ENV", "test")

	// 1. Estética primero
	baseLogger := logger.SetupLoggerWriters()
	log.Logger = baseLogger // Temporal: sin log level aún

	// 2. Cargar entorno (ya tendrá color, símbolo, salida a archivo)
	config.LoadEnvOnce()

	// 3. Aplicar nivel (debug/info) una vez que LOG_LEVEL ya está disponible
	log.Logger = logger.ApplyLogLevelFromEnv(baseLogger)
	logger.SetModule("test_main")

	// // 4. DB
	// conn := db.ConnectNamed("MAIN")
	// if err := db.ApplyMigrations(conn, true); err != nil {
	// 	logger.Fatal("failed to apply migrations: %v", err)
	// }
	// if _, err := conn.Exec(`TRUNCATE TABLE users RESTART IDENTITY CASCADE`); err != nil {
	// 	logger.Fatal("failed to truncate users table: %v", err)
	// }
	// _ = conn.Close()

	// 5. Run tests
	code := m.Run()

	ShutdownSuite()

	os.Exit(code)

}
