package discordx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://discord.com/api"
)

type RestClient struct {
	token string
	httpC *http.Client
}

func NewRestClient(token string) *RestClient {
	r := &RestClient{
		token: token,
		httpC: http.DefaultClient,
	}

	return r
}

type RequestCreateMessage struct {
	Content string `json:"content"`
}

func (r *RestClient) CreateContentOnlyMessage(channelID, content string) error {
	url := fmt.Sprintf("%s/channels/%s/messages", baseURL, channelID)
	byteReq, err := json.Marshal(&RequestCreateMessage{content})
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(byteReq))
	if err != nil {
		return err
	}
	r.addAuth(httpReq)
	r.addContentTypeJSON(httpReq)

	httpResp, err := r.httpC.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	byteResp, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] CreateContentOnlyMessage url: %s, code: %d, msg: %s", url, httpResp.StatusCode, byteResp)

	return nil
}

func (r *RestClient) addAuth(httpReq *http.Request) {
	if httpReq == nil {
		return
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bot %s", r.token))
}

func (r *RestClient) addContentTypeJSON(httpReq *http.Request) {
	if httpReq == nil {
		return
	}

	httpReq.Header.Set("Content-Type", "application/json")
}
