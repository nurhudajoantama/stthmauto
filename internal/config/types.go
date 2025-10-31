package config

import (
	"fmt"
)

type TCPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (t TCPServer) Addr() string {
	return fmt.Sprintf("%s:%s", t.Host, t.Port)
}

type Logging struct {
	Level          string `yaml:"level"`
	Type           string `yaml:"type"`
	LogFileEnabled bool   `yaml:"logFileEnabled"`
	LogFilePath    string `yaml:"logFilePath"`
}

type KVBolt struct {
	Path string `yaml:"path"`
}

type Config struct {
	HTTP TCPServer `yaml:"http"`
	Log  Logging   `yaml:"log"`
	KV   KVBolt    `yaml:"kv"`
}
