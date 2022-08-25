package main

import (
	"context"
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

	stateWatcher := service.NewStateWatcher(
		os.Getenv(alertVirtualTakerWebhook),
		vtclient.NewClient(
			os.Getenv(virtualTakerBaseURL), os.Getenv(virtualTakerUsername), os.Getenv(virtualTakerPassword),
		),
		slackx.NewRestClient(),
	)
	ctx := context.Background()

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
