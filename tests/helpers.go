package tests

import (
	"os"

	"ghostnote/internal/config"
	"ghostnote/internal/logger"

	"github.com/rs/zerolog/log"
)

func InitTestEnv(module string) {

	_ = os.Setenv("ENV", "test")

	baseLogger := logger.SetupLoggerWriters()
	log.Logger = baseLogger

	config.ResetEnvForTests()
	config.LoadEnvOnce()

	log.Logger = logger.ApplyLogLevelFromEnv(baseLogger)
	logger.SetModule(module)
}

func ShutdownSuite() {
	logger.Trace("Closing All Databases ...")
	// db.CloseAll()
	// cache.CloseAll()
}
