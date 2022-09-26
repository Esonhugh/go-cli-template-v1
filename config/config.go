package config

import (
	"fast_cli_template/utils/Error"
	"github.com/spf13/viper"
)

// Config struct is a wrapper of viper
type Config struct {
	*viper.Viper
}

// GlobalConfig default Global Variable for Config
var GlobalConfig *Config

// DefaultInit func is default Init Config file without config file information
func DefaultInit() {
	GlobalConfig = &Config{
		viper.New(),
	}

	// Config filename (no suffix)
	GlobalConfig.SetConfigName("config")

	// Config Type
	GlobalConfig.SetConfigType("yaml")

	// Searching Path
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("../") // For Debug
	GlobalConfig.AddConfigPath("./config")

	// Reading Config
	err := GlobalConfig.ReadInConfig()
	Error.HandleFatal(err)
}

// SpecificInit func is Init with specific Config file
func SpecificInit(file string) {
	GlobalConfig = &Config{
		viper.New(),
	}
	GlobalConfig.SetConfigFile(file)
	err := GlobalConfig.ReadInConfig()
	Error.HandleFatal(err)
}
