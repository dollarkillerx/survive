package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var CONFIG *Configuration

type Configuration struct {
	Nodes      []Node `json:"nodes"`
	ListenAddr string `json:"listen_addr"`
}

type Node struct {
	Name    string `json:"name"`
	Token   string `json:"token"`
	Timeout int64  `json:"timeout"` // sec
}

func InitConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	var conf Configuration
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalln(err)
	}

	CONFIG = &conf
}
