package slackx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RestClient struct {
	httpC *http.Client
}

func NewRestClient() *RestClient {
	return &RestClient{
		httpC: http.DefaultClient,
	}
}

func (r *RestClient) SendWebhookMsg(ctx context.Context, msg, webhookURL string) error {
	req := &SendWebhookMsqReq{
		Text: msg,
	}
	byteReq, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal webhook message error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, webhookURL, bytes.NewReader(byteReq))
	if err != nil {
		return fmt.Errorf("create webhook request error: %w", err)
	}

	resp, err := r.httpC.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send webhook request error: %w", err)
	}
	defer resp.Body.Close()

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read send webhook msg response error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("send webhook msg not status OK, status: %d, resp: %s", resp.StatusCode, byteResp)
	}

	log.Printf("[DEBUG] response %s", byteResp)

	return nil
}
