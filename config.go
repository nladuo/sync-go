package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Username  string
	Password  string
	Host      string
	Port      int
	RemoteDir string
}

func (self Config) Save() {
	rs, err := json.Marshal(self)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
	fmt.Println(string(rs))
	ioutil.WriteFile("sync-go-config.json", rs, 0666)
}

func LoadConfig() Config {

	data, err := ioutil.ReadFile("sync-go-config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
