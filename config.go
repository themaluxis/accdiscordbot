package accdiscordbot

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	General struct {
		AccServerPath string `envconfig:"ACCSERVER_PATH"`
	}
	Discord struct {
		BotToken  string `envconfig:"DISCORD_TOKEN"`
		ChannelID string `envconfig:"DISCORD_CHANNEL"`
	}
}

func LoadConfig(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Println("Error while grabbing env variables.")
	}
}
