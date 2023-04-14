package benchpool

import "strings"

type Pool struct {
	Address      string       `json:"-"`
	ReserveUsd   float64      `json:"reserveUsd,omitempty"`
	AmplifiedTvl float64      `json:"amplifiedTvl,omitempty"`
	SwapFee      float64      `json:"swapFee,omitempty"`
	Exchange     string       `json:"exchange,omitempty"`
	Type         string       `json:"type,omitempty"`
	Timestamp    int64        `json:"timestamp,omitempty"`
	Reserves     PoolReserves `json:"reserves,omitempty"`
	Tokens       []*PoolToken `json:"tokens,omitempty"`
	Extra        string       `json:"extra,omitempty"`
	StaticExtra  string       `json:"staticExtra,omitempty"`
	TotalSupply  string       `json:"totalSupply,omitempty"`
}

type PoolReserves []string

func (r PoolReserves) Encode() string {
	return strings.Join(r, ":")
}

type PoolToken struct {
	Address   string `json:"address,omitempty"`
	Name      string `json:"name,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	Decimals  uint8  `json:"decimals,omitempty"`
	Weight    uint   `json:"weight,omitempty"`
	Swappable bool   `json:"swappable,omitempty"`
}

type PoolTokens []*PoolToken
