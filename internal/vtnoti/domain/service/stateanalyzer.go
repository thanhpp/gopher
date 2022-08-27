package service

import (
	"context"
	"sort"

	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
	"github.com/thanhpp/gopher/pkg/slackx"
)

type StateAnalyzer struct {
	vtClient    *vtclient.Client
	slackClient *slackx.RestClient
}

func NewStateAnalyzer(vtClient *vtclient.Client, slackClient *slackx.RestClient) *StateAnalyzer {
	return &StateAnalyzer{
		vtClient:    vtClient,
		slackClient: slackClient,
	}
}

func (a *StateAnalyzer) getDoneStatesSortedByTime(ctx context.Context) ([]vtclient.StateData, error) {
	states, err := a.vtClient.GetStates(ctx, true, true)
	if err != nil {
		return nil, err
	}

	sort.Slice(states, func(i, j int) bool {
		return states[i].CreatedTime.Before(states[j].CreatedTime)
	})

	return states, nil
}
