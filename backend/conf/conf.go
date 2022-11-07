package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Database struct {
	Address string
}

type Server struct {
	Port string
}

type Config struct {
	Db     Database `toml:"db"`
	Server Server   `toml:"server"`
	Debug  bool
}

var conf Config

func Load() {
	// Read Config File
	data, err := os.ReadFile("conf/conf.toml")
	if err != nil {
		log.Fatal(err)
	}

	// Parse Config File
	if _, err := toml.Decode(string(data), &conf); err != nil {
		log.Fatal(err)
	}
}

func Get() *Config {
	return &conf
}
