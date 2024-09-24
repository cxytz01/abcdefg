package utils

import (
	"syscall"

	"github.com/jinzhu/configor"
)

// var CNF TomlConfig

type TomlConfig struct {
	Title    string
	Producer Producer `toml:"producer"`
}

type Producer struct {
	Addr     string `toml:"listen"`
	Pg       string `toml:"pg"`
	CSVStore string `toml:"csvstore"`
	Kafka    Kafka  `toml:"kafka"`
}

type Kafka struct {
	BrokerList string `toml:"broker-list"`
	Topic      string `toml:"topic"`
	Ack        int    `toml:"ack"`
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
