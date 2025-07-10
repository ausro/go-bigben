package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Token     string `json:"token"`
	ChannelId string `json:"channel"`
}

func readConfig() Config {
	dat, err := os.ReadFile("config.json")
	if err != nil {
		log.Printf("Failed to read config: %s. Creating new.\n", err)
		createConfig()
		return readConfig()
	}

	conf := Config{}

	err = json.Unmarshal(dat, &conf)
	if err != nil {
		log.Fatalf("Failed to parse json data: %s", err)
	}

	return conf
}

func createConfig() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Bot Token: ")
	reader.Scan()
	token := reader.Text()

	fmt.Print("Channel Id: ")
	reader.Scan()
	channelId := reader.Text()

	conf := Config{
		Token:     token,
		ChannelId: channelId,
	}

	json, err := json.Marshal(conf)
	if err != nil {
		log.Fatalf("Failed to parse inputs as json: %s", err)
	}

	os.WriteFile("config.json", json, 0644)
}
