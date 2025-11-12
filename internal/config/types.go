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

type SQL struct {
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	Port        string `yaml:"port"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
}

func (s SQL) DatabaseUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.User, s.Password, s.Host, s.Port, s.Name)
}

func (s SQL) DataSourceName() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		s.User, s.Password, s.Host, s.Port, s.Name)
}

type MQTT struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func (m MQTT) BrokerURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", m.User, m.Password, m.Host, m.Port)
}

type InternetCheck struct {
	CheckAddress string `yaml:"checkAddress"`
	ModemAddress string `yaml:"modemAddress"`

	Interval string `yaml:"interval"`
}

type Config struct {
	HTTP          TCPServer     `yaml:"http"`
	Log           Logging       `yaml:"log"`
	DB            SQL           `yaml:"db"`
	MQTT          MQTT          `yaml:"mqtt"`
	InternetCheck InternetCheck `yaml:"internetCheck"`
}
