package utils

import (
	"syscall"

	"github.com/jinzhu/configor"
)

var CNF TomlConfig

type TomlConfig struct {
	Title    string
	Consumer Consumer `toml:"consumer"`
}

type Consumer struct {
	Pg    string `toml:"pg"`
	Kafka Kafka  `toml:"kafka"`
}

type Kafka struct {
	BootstrapServer string `toml:"bootstrap-server"`
	ConsumerGroup   string `toml:"consumer-group"`
	Topic           string `toml:"topic"`
	Ack             int    `toml:"ack"`
}

func InitConfFile(file string, cf *TomlConfig) error {
	err := syscall.Access(file, syscall.O_RDONLY)
	if err != nil {
		return err
	}
	err = configor.Load(cf, file)
	if err != nil {
		return err
	}

	return nil
}
