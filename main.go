package main

import (
	"gobes/abstraction/config"
	"gobes/abstraction/logger"
)

func main() {
	// load logger
	log := logger.NewWriter()
	log.In("main").Info("logger was running")

	cfg := config.NewViperConfig(nil)
	if cfg == nil {
		log.In("main").Error("viper config not loaded")
	}

	serverEnv := cfg.Get("Server.ENV", "undefine")
	log.In("main").Infof("server was running on env %s", serverEnv)
}
