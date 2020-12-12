package Tool

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

)

type Config struct {
	EngineHost  string `json:"engine_host"`
	EnginePort  string  `json:"engine_port"`
	DatabaseConfig *DatabaseConfig `json:"database_config"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Dbname   string `json:"db_name"`
}

var _cfg *Config = nil

func ParseConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)

	err = decoder.Decode(&_cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return _cfg

}
