package discordx_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpp/gopher/pkg/discordx"
)

func init() {
	if err := godotenv.Load("../../secrets/discord.env"); err != nil {
		log.Println("[WARN] load env error", err)
	}
}

func TestCreateMessage(t *testing.T) {
	c := discordx.NewRestClient(os.Getenv("DISCORD_BOT_TOKEN"))

	err := c.CreateContentOnlyMessage(os.Getenv("DISCORD_BOT_TEST_CHANNEL"), "test message")
	assert.NoError(t, err)
}
