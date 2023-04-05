package main

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestBSCClient(t *testing.T) {
	ethClient, err := ethclient.Dial("http://127.0.0.1:12342")
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		id, err := ethClient.BlockNumber(context.Background())
		require.NoError(t, err)

		t.Logf("%v", id)

		time.Sleep(time.Second)
	}
}

func TestIPInfo(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:12342")
	require.NoError(t, err)

	t.Logf("%+v", resp.Status)

	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	require.NoError(t, err)

	t.Logf("data: %s", string(data))

	req, err := http.NewRequest(http.MethodGet, "https://ipinfo.io", nil)
	require.NoError(t, err)
	t.Logf("test req: %+v", req)
}
