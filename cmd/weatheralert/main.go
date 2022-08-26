package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thanhpp/gopher/internal/weatheralert/domain/service"
	"github.com/thanhpp/gopher/pkg/crawlers"
	"github.com/thanhpp/gopher/pkg/discordx"
)

func main() {
	discordC, discordMeta, err := setupDiscordClient()
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	a := service.NewHanoiAlert(crawlers.NewWeatherVN(), discordC, discordMeta.weatherChannel)

	a.Start()
}

const (
	discordEnvFile   = "../../secrets/discord.env"
	discordToken     = "DISCORD_BOT_TOKEN"
	discordWeatherID = "DISCORD_BOT_WEATHER_CHANNEL_ID"
)

type discordMetaData struct {
	weatherChannel string
}

func setupDiscordClient() (*discordx.RestClient, *discordMetaData, error) {
	if err := godotenv.Load(discordEnvFile); err != nil {
		return nil, nil, fmt.Errorf("setup discord client - load env %s error: %w", discordEnvFile, err)
	}

	r := discordx.NewRestClient(os.Getenv(discordToken))

	return r, &discordMetaData{
		weatherChannel: os.Getenv(discordWeatherID),
	}, nil
}
