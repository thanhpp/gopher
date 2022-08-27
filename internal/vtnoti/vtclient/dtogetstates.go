package vtclient

import "time"

const txExecuted = "EXECUTED"

type GetStatesResp struct {
	ResultTemplate
	Data []StateData `json:"data"`
}

type StateData struct {
	StateID               string         `json:"state_id"`
	CEX                   string         `json:"cex"`
	DEX                   string         `json:"dex"`
	DEXChain              string         `json:"dex_chain"`
	DEXWallet             string         `json:"dex_wallet"`
	Token                 string         `json:"token"`
	BaseAmount            float64        `json:"base_amount"`
	Side                  string         `json:"side"`
	P1PriceDiff           float64        `json:"p1_price_diff"`
	P1ProfitableThreshold float64        `json:"p1_profitable_threshold"`
	P1FillableThreshold   float64        `json:"p1_fillable_threshold"`
	P2CancelThreshold     float64        `json:"p2_cancel_threshold"`
	IsDone                bool           `json:"is_done"`
	CreatedTime           time.Time      `json:"created_time"`
	P1CEXOrders           []CEXOrderData `json:"p1_cex_orders,omitempty"`
	P2CEXOrders           []CEXOrderData `json:"p2_cex_orders,omitempty"`
	P2DEXTxs              []DexTxData    `json:"p2_dex_txs,omitempty"`
}

func (d *StateData) CalCEXOrderBaseFilled(part int) float64 {
	var arr []CEXOrderData

	switch part {
	case 1:
		arr = d.P1CEXOrders
	case 2:
		arr = d.P2CEXOrders
	default:
		return 0
	}

	baseFilled := 0.0
	for i := range arr {
		baseFilled += arr[i].FilledBaseAmount
	}

	return baseFilled
}

func (d *StateData) CalCEXOrderAFP(part int) float64 {
	var arr []CEXOrderData

	switch part {
	case 1:
		arr = d.P1CEXOrders
	case 2:
		arr = d.P2CEXOrders
	default:
		return 0
	}

	quoteFilled, baseFilled := 0.0, 0.0

	for i := range arr {
		quoteFilled += arr[i].FilledQuoteAmount
		baseFilled += arr[i].FilledBaseAmount
	}

	if baseFilled == 0 {
		return 0
	}

	return quoteFilled / baseFilled
}

func (d *StateData) CalP2DEXBaseFilled() float64 {
	filled := 0.0

	if d.Side == "buy" {
		for i := range d.P2DEXTxs {
			if d.P2DEXTxs[i].Status == txExecuted {
				filled += d.P2DEXTxs[i].AmountIn
			}
		}

		return filled
	}

	for i := range d.P2DEXTxs {
		if d.P2DEXTxs[i].Status == txExecuted {
			filled += d.P2DEXTxs[i].EstimatedAmountOut
		}
	}

	return filled
}

func (d *StateData) CalP2DEXAFP() float64 {
	tmp := 0.0
	filled := 0.0

	if d.Side == "buy" {
		for i := range d.P2DEXTxs {
			if d.P2DEXTxs[i].Status == "EXECUTED" {
				tmp += d.P2DEXTxs[i].Price * d.P2DEXTxs[i].AmountIn
				filled += d.P2DEXTxs[i].AmountIn
			}
		}
		if filled == 0 {
			return 0
		}
		return tmp / filled
	}

	for i := range d.P2DEXTxs {
		if d.P2DEXTxs[i].Status == "EXECUTED" {
			tmp += d.P2DEXTxs[i].Price * d.P2DEXTxs[i].EstimatedAmountOut
			filled += d.P2DEXTxs[i].EstimatedAmountOut
		}
	}
	if filled == 0 {
		return 0
	}
	return tmp / filled
}

type CEXOrderData struct {
	ID                string  `json:"id"`
	Status            string  `json:"status"`
	BaseSymbol        string  `json:"base_symbol"`
	QuoteSymbol       string  `json:"quote_symbol"`
	Side              string  `json:"side"`
	Price             float64 `json:"price"`
	BaseAmount        float64 `json:"base_amount"`
	FilledBaseAmount  float64 `json:"filled_base_amount"`
	FilledQuoteAmount float64 `json:"filled_quote_amount"`
}

type DexTxData struct {
	StateID            string  `json:"state_id"`
	Nonce              int64   `json:"nonce"`
	TxHash             string  `json:"tx_hash"`
	Status             string  `json:"status"`
	ToWallet           string  `json:"to_wallet"`
	RouterAddr         string  `json:"router_address"`
	TokenIn            string  `json:"token_in"`
	TokenOut           string  `json:"token_out"`
	AmountIn           float64 `json:"amount_in"`
	EstimatedAmountOut float64 `json:"estimated_amount_out"`
	Price              float64 `json:"price"`
	GasPrice           float64 `json:"gas_price"`
	MaxTip             float64 `json:"max_tip"`
}
