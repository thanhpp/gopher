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
			log.Printf("[DEBUG] state %s. P1 CEX Order: %d", states[i].StateID, len(states[i].P1CEXOrders))
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
				if err := s.notifyDoneState(ctx, &states[i]); err != nil {
					log.Println("[SKIP ERROR]", err)
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

	log.Printf("[DEBUG] cached P1 CEX orders: %d. current P1 CEX orders %d",
		len(cached.P1CEXOrders), len(state.P1CEXOrders))
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

	return nil
}

func (s *StateWatcher) compareAndNotifyState(ctx context.Context, cached, current *vtclient.StateData) error {
	if err := s.compareAndNotifyCEXOrder(ctx, current.StateID, 1, cached.P1CEXOrders, current.P1CEXOrders); err != nil {
		return err
	}

	if err := s.compareAndNotifyCEXOrder(ctx, current.StateID, 2, cached.P2CEXOrders, current.P2CEXOrders); err != nil {
		return err
	}

	if len(current.P2DEXTxs) == 0 {
		return nil
	}
	if len(cached.P2DEXTxs) == 0 {
		return s.notifyDEXTx(ctx, current.StateID, &current.P2DEXTxs[len(current.P2DEXTxs)-1])
	}
	cachedDEXTx := cached.P2DEXTxs[len(cached.P2DEXTxs)-1]
	currentDEXTx := current.P2DEXTxs[len(current.P2DEXTxs)-1]
	if cachedDEXTx.TxHash != currentDEXTx.TxHash {
		if len(current.P2DEXTxs) > 1 {
			if err := s.notifyDEXTx(
				ctx, current.StateID, &current.P2DEXTxs[len(current.P2DEXTxs)-2],
			); err != nil {
				return err
			}
		}
		return s.notifyDEXTx(ctx, current.StateID, &currentDEXTx)
	}

	return nil
}

func (s *StateWatcher) compareAndNotifyCEXOrder(
	ctx context.Context, stateID string, part int, cachedList, currentList []vtclient.CEXOrderData,
) error {
	if len(currentList) == 0 {
		return nil
	}
	if len(cachedList) == 0 { // first CEX order
		return s.notifyCEXOrder(ctx, stateID, part, &currentList[len(currentList)-1])
	}

	cachedP1CEXOrder := cachedList[len(cachedList)-1]
	currentP1CEXOrder := currentList[len(currentList)-1]
	if cachedP1CEXOrder.ID == currentP1CEXOrder.ID {
		if cachedP1CEXOrder.FilledBaseAmount != currentP1CEXOrder.FilledBaseAmount { // latest cex order updated
			return s.notifyCEXOrder(ctx, stateID, part, &currentP1CEXOrder)
		}
		return nil
	}

	if currentP1CEXOrder.FilledBaseAmount != 0 {
		return s.notifyCEXOrder(ctx, stateID, part, &currentP1CEXOrder)
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
Price: %f
Base amount: %f
Filled base amount: %f`,
		stateID, part, order.ID, order.Status, order.Side, order.BaseSymbol, order.QuoteSymbol,
		order.Price, order.BaseAmount, order.FilledBaseAmount)

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

func (s *StateWatcher) notifyDoneState(ctx context.Context, state *vtclient.StateData) error {
	msg := fmt.Sprintf(`> STATE DONE
<@U03LG91301L>
ID: %s
P1 Orders: %d
P2 Orders: %d
P2 Tx: %d`,
		state.StateID, len(state.P1CEXOrders), len(state.P2CEXOrders), len(state.P2DEXTxs))

	if err := s.slackClient.SendWebhookMsg(ctx, msg, s.slackWebhook); err != nil {
		return fmt.Errorf("notify done state: %w", err)
	}

	return nil
}
