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

func loadEnvFiles() error {
	if err := godotenv.Load(vtClientEnvFile); err != nil {
		return err
	}
	if err := godotenv.Load(slackEnvFile); err != nil {
		return err
	}

	return nil
}
