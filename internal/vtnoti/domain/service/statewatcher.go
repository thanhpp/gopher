package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/thanhpp/gopher/internal/vtnoti/vtclient"
	"github.com/thanhpp/gopher/pkg/slackx"
)

const (
	stateWatchInterval = time.Second * 5
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

func (s *StateWatcher) Start(ctx context.Context) { // nolint: gocognit
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
		// filter empty state
		var fileredStates []vtclient.StateData
		for i := range states {
			if len(states[i].P1CEXOrders) != 0 {
				fileredStates = append(fileredStates, states[i])
			}
		}
		log.Println("[DEBUG] get states", len(fileredStates))

		for i := range fileredStates {
			log.Printf("[DEBUG] state %s. P1 CEX Order: %d", fileredStates[i].StateID, len(fileredStates[i].P1CEXOrders))
			if err := s.checkAndNotifyState(ctx, &fileredStates[i]); err != nil {
				log.Println("[SKIP ERROR]", err)
				continue
			}
		}

		for k := range s.statesCache {
			isFound := false
			for i := range fileredStates {
				if fileredStates[i].StateID == k {
					isFound = true
					break
				}
			}
			if !isFound {
				state, err := s.vtClient.GetState(ctx, k)
				if err != nil {
					log.Println("[SKIP ERROR]", err)
					continue
				}
				if err := s.notifyDoneState(ctx, &state); err != nil {
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
		order.ActualPrice, order.BaseAmount, order.FilledBaseAmount)

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
		stateID, dexTx.TxHash, dexTx.Status)

	if err := s.slackClient.SendWebhookMsg(ctx, msg, s.slackWebhook); err != nil {
		return fmt.Errorf("notify dex tx error: %w", err)
	}

	return nil
}

func (s *StateWatcher) notifyDoneState(ctx context.Context, state *vtclient.StateData) error {
	msg := fmt.Sprintf(`> STATE DONE
<@U03LG91301L>
ID: %s
Side: %s

P1 Orders: %d
P1 Filled: %f
P1 AFP: %f

P2 Orders: %d
P2 CEX Filled: %f
P2 CEX AFP: %f 

P2 Txs: %d
P2 DEX Filled: %f
P2 DEX AFP: %f

%s`,
		state.StateID, state.Side,
		len(state.P1CEXOrders), state.CalCEXOrderBaseFilled(1), state.CalCEXOrderAFP(1),
		len(state.P2CEXOrders), state.CalCEXOrderBaseFilled(2), state.CalCEXOrderAFP(2),
		len(state.P2DEXTxs), state.CalP2DEXBaseFilled(), state.CalP2DEXAFP(),
		stringtifyMap("Asset change with fee", state.AssetChangeWithFee))

	if err := s.slackClient.SendWebhookMsg(ctx, msg, s.slackWebhook); err != nil {
		return fmt.Errorf("notify done state: %w", err)
	}

	return nil
}

func stringtifyMap[K comparable, V any](name string, in map[K]V) string {
	strB := new(strings.Builder)
	strB.WriteString(name + "\n")
	for k, v := range in {
		strB.WriteString(fmt.Sprintf("  - %v: %v\n", k, v))
	}
	strB.WriteString("\n")

	return strB.String()
}
