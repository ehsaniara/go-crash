package log

import (
	"github.com/ehsaniara/go-crash/config"
	"go.uber.org/zap"
	"time"
)

var Log *zap.SugaredLogger

func Setup() {
	// Using zap's preset constructors is the simplest way to get a feel for the
	// package, but they don't allow much customization.
	// or NewProduction, or NewDevelopment
	var logger *zap.Logger
	if config.AppConfig.App.LogMode == "Development" {
		logger, _ = zap.NewDevelopment() // or NewProduction, or NewDevelopment
	} else {
		logger, _ = zap.NewProduction() // or NewProduction, or NewDevelopment
	}
	defer logger.Sync()
	//
	//// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	//// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Logger impl test.",
		// Structured context as loosely typed key-value pairs.
		"attempt", 3,
		"backoff", time.Second,
	)
	Log = sugar
}
