package config

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func InitializeConfig(path string) (Config, error) {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	// fang.SetEnvPrefix(envPrefix)
	fang.SetConfigType("yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	if err := fang.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return Config{}, err
	}
	// Load configuration
	c := Config{}
	if err = fang.Unmarshal(&c); err != nil {
		return Config{}, err
	}
	return c, nil
}
