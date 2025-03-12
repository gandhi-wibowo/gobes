package main

import (
	"flag"
	"gobes/abstraction/config"
	"gobes/abstraction/logger"
)

func main() {
	// load logger
	log := logger.NewWriter()
	log.In("main").Info("logger was running")

	// load config
	configFile := flag.String("c", "", "configuration file without extension. For config.toml then put \" -c config\"")
	flag.Parse()

	cfg := config.NewViperConfig(config.Convert(*configFile))
	if cfg == nil {
		log.In("main").Error("viper config not loaded")
	}

	serverEnv := cfg.Get("Server.ENV", "undefine")
	log.In("main").Infof("server was running on env %s", serverEnv)
}
