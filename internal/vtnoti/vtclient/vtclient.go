package vtclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseURL string
	httpC   *http.Client
	bAuth   *basicAuth
}

func NewClient(baseURL, username, password string) *Client {
	return &Client{
		baseURL: baseURL,
		httpC:   http.DefaultClient,
		bAuth:   newBasicAuth(username, password),
	}
}

func (c *Client) GetStates(ctx context.Context, isDone, isFilled bool) ([]StateData, error) {
	url := fmt.Sprintf("%s/state?is_done=%t&fill_state=%t", c.baseURL, isDone, isFilled)

	resp := new(GetStatesResp)

	if err := c.doGet(ctx, url, resp); err != nil {
		return nil, fmt.Errorf("GetStates - %w", err)
	}

	return resp.Data, nil
}

func (c *Client) GetState(ctx context.Context, stateID string) (StateData, error) {
	url := fmt.Sprintf("%s/state/%s", c.baseURL, stateID)

	resp := new(GetStateResp)

	if err := c.doGet(ctx, url, resp); err != nil {
		return StateData{}, fmt.Errorf("GetState - %w", err)
	}

	return resp.Data.StateData, nil
}

func (c *Client) doGet(ctx context.Context, url string, respExpected interface{}) error {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("doGet - new request error: %w", err)
	}
	c.bAuth.addHeader(httpReq)

	resp, err := c.httpC.Do(httpReq)
	if err != nil {
		return fmt.Errorf("doGet error: %w", err)
	}
	defer resp.Body.Close()

	byteRespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("doGet - read response body error: %w", err)
	}

	if err := json.Unmarshal(byteRespBody, respExpected); err != nil {
		return fmt.Errorf("doGet - marshal response error: %w", err)
	}

	return nil
}
