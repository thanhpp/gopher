package vtclient_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
)

const (
	envFile              = "../../../secrets/vtclient.env"
	virtualTakerBaseURL  = "VIRTUAL_TAKER_BASE_URL"
	virtualTakerUsername = "VIRTUAL_TAKER_USERNAME"
	virtualTakerPassword = "VIRTUAL_TAKER_PASSWORD"
)

func init() {
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("[WARN] [vtclient_test] load env file %s error: %v", envFile, err)
	}
}

func TestGetStates(t *testing.T) {
	c := vtclient.NewClient(
		os.Getenv(virtualTakerBaseURL), os.Getenv(virtualTakerUsername), os.Getenv(virtualTakerPassword))

	states, err := c.GetStates(context.Background(), true, true)
	require.NoError(t, err)

	t.Logf("got %d states", len(states))
}
