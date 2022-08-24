package slackx_test

import (
	"log"
	"os"
	"testing"

	"context"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/slackx"
)

const (
	envFile                  = "../../secrets/slack.env"
	alertVirtualTakerWebhook = "ALERT_VIRTUAL_TAKER_WEBHOOK"
)

func init() {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("[WARN] [slackx_test] load env %s error: %v", envFile, err)
	}
}

func TestSendWebhookMsg(t *testing.T) {
	var (
		ctx     = context.Background()
		c       = slackx.NewRestClient()
		testMsg = "test message"
	)

	err := c.SendWebhookMsg(ctx, testMsg, os.Getenv(alertVirtualTakerWebhook))
	require.NoError(t, err)
}
