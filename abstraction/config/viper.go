package config

import (
	"flag"
	"gobes/abstraction/support/convert"
	"sync"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var (
	once     sync.Once
	instance *ViperConfig
)

type ViperConfig struct {
	vip *viper.Viper
}

// Ensure Viper implements Config
var _ ConfigProvider = (*ViperConfig)(nil)

func NewViperConfig(configFile *ConfigFile) *ViperConfig {
	once.Do(func() {
		if configFile == nil {
			// load default config, running from command
			configFileFromCommand := flag.String("c", "", "configuration file without extension. For config.toml then put \" -c config\"")
			flag.Parse()
			commandToConfigFile := Convert(*configFileFromCommand)
			configFile = &commandToConfigFile
		}

		vip := viper.New()

		vip.SetConfigName(configFile.Name)
		vip.AddConfigPath(configFile.Path)
		err := vip.ReadInConfig()
		if err != nil {
			panic(err)
		}
		instance = &ViperConfig{vip: vip}
	})
	return instance
}

// Env Get config from env.
func (app *ViperConfig) Env(envName string, defaultValue ...any) any {
	value := app.Get(envName, defaultValue...)
	if cast.ToString(value) == "" {
		return convert.Default(defaultValue...)
	}

	return value
}

// Add config to application.
func (app *ViperConfig) Add(name string, configuration any) {
	app.vip.Set(name, configuration)
}

// Get config from application.
func (app *ViperConfig) Get(path string, defaultValue ...any) any {
	if !app.vip.IsSet(path) {
		return convert.Default(defaultValue...)
	}
	return app.vip.Get(path)
}

// GetString get string type config from application.
func (app *ViperConfig) GetString(path string, defaultValue ...string) string {
	if !app.vip.IsSet(path) {
		return convert.Default(defaultValue...)
	}
	return app.vip.GetString(path)
}

// GetInt get int type config from application.
func (app *ViperConfig) GetInt(path string, defaultValue ...int) int {
	if !app.vip.IsSet(path) {
		return convert.Default(defaultValue...)
	}
	return app.vip.GetInt(path)
}

// GetBool get bool type config from application.
func (app *ViperConfig) GetBool(path string, defaultValue ...bool) bool {
	if !app.vip.IsSet(path) {
		return convert.Default(defaultValue...)
	}
	return app.vip.GetBool(path)
}

// GetDuration get time.Duration type config from application
func (app *ViperConfig) GetDuration(path string, defaultValue ...time.Duration) time.Duration {
	if !app.vip.IsSet(path) {
		return convert.Default(defaultValue...)
	}
	return app.vip.GetDuration(path)
}

// GetList get list type config
func (app *ViperConfig) GetList(path string, rawVal interface{}) {
	if !app.vip.IsSet(path) {
		rawVal = nil
	}
	app.vip.UnmarshalKey(path, rawVal)
}
