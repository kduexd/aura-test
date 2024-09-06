package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
	log "go.uber.org/zap"
)

// EnvForge the instance of env.yaml via viper
func EnvForge() *viper.Viper {
	envInst := viper.New()
	envInst.SetConfigName("env")
	envInst.SetConfigType("yml")
	envInst.AddConfigPath(".")
	envInst.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	envInst.AutomaticEnv()

	if err := envInst.ReadInConfig(); err != nil {
		log.Error(errors.New("Could not find env configuration file, error: " + err.Error()))
	}

	return envInst
}

// ConForge the instance of config.json via viper
func ConForge() *viper.Viper {
	cfgInst := viper.New()
	cfgInst.SetConfigName("config")
	cfgInst.SetConfigType("json")
	cfgInst.AddConfigPath(".")
	cfgInst.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfgInst.AutomaticEnv()

	if err := cfgInst.ReadInConfig(); err != nil {
		log.Error(errors.New("Could not find env configuration file, error: " + err.Error()))
	}

	return cfgInst
}
