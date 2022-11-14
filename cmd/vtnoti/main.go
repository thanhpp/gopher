package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thanhpp/gopher/internal/vtnoti/domain/service"
	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
	"github.com/thanhpp/gopher/pkg/bootstrap"
	"github.com/thanhpp/gopher/pkg/slackx"
)

const (
	vtClientEnvFile      = "../../secrets/vtclient.env"
	virtualTakerBaseURL  = "VIRTUAL_TAKER_BASE_URL"
	virtualTakerUsername = "VIRTUAL_TAKER_USERNAME"
	virtualTakerPassword = "VIRTUAL_TAKER_PASSWORD"
)

const (
	slackEnvFile             = "../../secrets/slack.env"
	alertVirtualTakerWebhook = "ALERT_VIRTUAL_TAKER_WEBHOOK"
)

func main() {
	if err := loadEnvFiles(); err != nil {
		log.Fatal("[FATAL] load env files", err)
	}

	slackClient := slackx.NewRestClient()

	stateWatcher := service.NewStateWatcher(
		os.Getenv(alertVirtualTakerWebhook),
		vtclient.NewClient(
			os.Getenv(virtualTakerBaseURL), os.Getenv(virtualTakerUsername), os.Getenv(virtualTakerPassword),
		),
		slackClient,
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := slackClient.SendWebhookMsg(
		ctx, fmt.Sprintf("> Vtnoti started\nVersion: %s\nDesc: %s", version, desc),
		os.Getenv(alertVirtualTakerWebhook)); err != nil {
		log.Fatal("[FATAL] send slack start message error", err) // nolint: gocritic
	}

	go stateWatcher.Start(ctx)

	bootstrap.WaitTerminateSignals()
}

func configFilesArgs() (vtEnv, slackEnv string) {
	args := os.Args
	if len(os.Args) != 3 {
		return vtClientEnvFile, slackEnvFile
	}

	return args[1], args[2]
}

func loadEnvFiles() error {
	vtEnv, slackEnv := configFilesArgs()

	if err := godotenv.Load(vtEnv); err != nil {
		return err
	}
	log.Println("loaded:", vtEnv)

	if err := godotenv.Load(slackEnv); err != nil {
		return err
	}
	log.Println("loaded", slackEnv)

	return nil
}
