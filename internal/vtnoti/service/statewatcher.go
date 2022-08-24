package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
	"github.com/thanhpp/gopher/pkg/slackx"
)

const (
	stateWatchInterval = time.Second
)

type StateWatcher struct {
	slackWebhook string
	vtClient     *vtclient.Client
	slackClient  *slackx.RestClient
	statesCache  map[string]*vtclient.StateData // stateID -> stateData
}

func NewStateWatcher(slackWebhook string, vtClient *vtclient.Client, slackClient *slackx.RestClient) *StateWatcher {
	return &StateWatcher{
		slackWebhook: slackWebhook,
		vtClient:     vtClient,
		slackClient:  slackClient,
		statesCache:  make(map[string]*vtclient.StateData),
	}
}

func (s *StateWatcher) Start(ctx context.Context) {
	t := time.NewTicker(stateWatchInterval)
	defer t.Stop()

	for ; true; <-t.C {
		if err := ctx.Err(); err != nil {
			log.Println("State watcher stopped", err)
			return
		}
		// get undone states
		states, err := s.vtClient.GetStates(ctx, false, true)
		if err != nil {
			log.Println("[SKIP ERROR]", err)
			continue
		}
		log.Println("[DEBUG] get states", len(states))

		for i := range states {
			if err := s.checkAndNotifyState(ctx, &states[i]); err != nil {
				log.Println("[SKIP ERROR]", err)
				continue
			}
		}

		for k := range s.statesCache {
			for i := range states {
				if states[i].StateID == k {
					continue
				}
				delete(s.statesCache, k)
			}
		}
	}
}

func (s *StateWatcher) checkAndNotifyState(ctx context.Context, state *vtclient.StateData) error {
	cached, ok := s.statesCache[state.StateID]
	defer func() { s.statesCache[state.StateID] = state }()
	if !ok {
		log.Println("[DEBUG] new state", state.StateID)
		if err := s.notifyNewState(ctx, state); err != nil {
			return err
		}
		return nil
	}

	if err := s.compareAndNotifyState(ctx, cached, state); err != nil {
		return err
	}

	return nil
}

func (s *StateWatcher) notifyNewState(ctx context.Context, state *vtclient.StateData) error {
	msg := fmt.Sprintf(`> NEW STATE
ID: %s
P1 CEX Orders: %d
P2 CEX Orders: %d
P2 DEX Txs: %d
	`, state.StateID, len(state.P1CEXOrders), len(state.P2CEXOrders), len(state.P2DEXTxs))

	if err := s.slackClient.SendWebhookMsg(ctx, msg, s.slackWebhook); err != nil {
		return fmt.Errorf("notify new state error: %w", err)
	}

	if len(state.P1CEXOrders) == 0 {
		return nil
	}
	if err := s.notifyCEXOrder(ctx, state.StateID, 1, &state.P1CEXOrders[len(state.P1CEXOrders)-1]); err != nil {
		return err
	}

	if len(state.P2CEXOrders) == 0 {
		return nil
	}
	if err := s.notifyCEXOrder(ctx, state.StateID, 2, &state.P2CEXOrders[len(state.P2CEXOrders)-1]); err != nil {
		return err
	}

	if len(state.P2DEXTxs) == 0 {
		return nil
	}
	if err := s.notifyDEXTx(ctx, state.StateID, &state.P2DEXTxs[len(state.P2DEXTxs)-1]); err != nil {
		return err
	}

	return nil
}

func (s *StateWatcher) compareAndNotifyState(ctx context.Context, cached, current *vtclient.StateData) error {
	if len(current.P1CEXOrders) == 0 {
		return nil
	}
	if len(cached.P1CEXOrders) == 0 {
		return s.notifyCEXOrder(ctx, current.StateID, 1, &current.P1CEXOrders[len(current.P1CEXOrders)-1])
	}
	cachedP1CEXOrder := cached.P1CEXOrders[len(cached.P1CEXOrders)-1]
	currentP1CEXOrder := current.P1CEXOrders[len(current.P1CEXOrders)-1]
	if cachedP1CEXOrder.FilledBaseAmount != currentP1CEXOrder.FilledBaseAmount {
		return s.notifyCEXOrder(ctx, current.StateID, 1, &currentP1CEXOrder)
	}

	if len(current.P2CEXOrders) == 0 {
		return nil
	}
	if len(cached.P2CEXOrders) == 0 {
		return s.notifyCEXOrder(ctx, current.StateID, 2, &current.P2CEXOrders[len(current.P2CEXOrders)-1])
	}
	cachedP2CEXOrder := cached.P2CEXOrders[len(cached.P2CEXOrders)-1]
	currentP2CEXOrder := current.P2CEXOrders[len(current.P2CEXOrders)-1]
	if cachedP2CEXOrder.FilledBaseAmount != currentP2CEXOrder.FilledBaseAmount {
		return s.notifyCEXOrder(ctx, current.StateID, 2, &currentP2CEXOrder)
	}

	if len(current.P2DEXTxs) == 0 {
		return nil
	}
	if len(cached.P2DEXTxs) == 0 {
		return s.notifyDEXTx(ctx, current.StateID, &current.P2DEXTxs[len(current.P2DEXTxs)-1])
	}
	cachedDEXTx := cached.P2DEXTxs[len(cached.P2DEXTxs)-1]
	currentDEXTx := current.P2DEXTxs[len(current.P2DEXTxs)-1]
	if cachedDEXTx.Status != currentDEXTx.Status {
		return s.notifyDEXTx(ctx, current.StateID, &currentDEXTx)
	}

	return nil
}

func (s *StateWatcher) notifyCEXOrder(
	ctx context.Context, stateID string, part int, order *vtclient.CEXOrderData,
) error {
	msg := fmt.Sprintf(`> CEX ORDER
State ID: %s
Part: %d
ID: %s
Status: %s
Side: %s
Symbol: %s/%s
Base amount: %f
Filled base amount: %f`,
		stateID, part, order.ID, order.Status, order.Side, order.BaseSymbol, order.QuoteSymbol,
		order.BaseAmount, order.FilledBaseAmount)

	if err := s.slackClient.SendWebhookMsg(ctx, msg, s.slackWebhook); err != nil {
		return fmt.Errorf("notify cex order error: %w", err)
	}

	return nil
}

func (s *StateWatcher) notifyDEXTx(ctx context.Context, stateID string, dexTx *vtclient.DexTxData) error {
	msg := fmt.Sprintf(`> DEX Transaction
State ID: %s
Tx hash: %s
Status: %s`,
		stateID, dexTx.TxHash, dexTx.TxHash)

	if err := s.slackClient.SendWebhookMsg(ctx, msg, s.slackWebhook); err != nil {
		return fmt.Errorf("notify dex tx error: %w", err)
	}

	return nil
}
