package benchconcurrentmap

import "sync"

type Pool struct {
	Address      string       `json:"-"`
	ReserveUsd   float64      `json:"reserveUsd,omitempty"`
	AmplifiedTvl float64      `json:"amplifiedTvl,omitempty"`
	SwapFee      float64      `json:"swapFee,omitempty"`
	Exchange     string       `json:"exchange,omitempty"`
	Type         string       `json:"type,omitempty"`
	Timestamp    int64        `json:"timestamp,omitempty"`
	Reserves     PoolReserves `json:"reserves,omitempty"`
	Tokens       []PoolToken  `json:"tokens,omitempty"`
	Extra        string       `json:"extra,omitempty"`
	StaticExtra  string       `json:"staticExtra,omitempty"`
	TotalSupply  string       `json:"totalSupply,omitempty"`
}

func (p Pool) Clone() Pool {
	tokens := make([]PoolToken, len(p.Tokens))
	copy(tokens, p.Tokens)

	return Pool{
		Address:      p.Address,
		ReserveUsd:   p.ReserveUsd,
		AmplifiedTvl: p.AmplifiedTvl,
		SwapFee:      p.SwapFee,
		Exchange:     p.Exchange,
		Type:         p.Type,
		Timestamp:    p.Timestamp,
		Reserves:     p.Reserves.Clone(),
		Tokens:       tokens,
		Extra:        p.Extra,
		StaticExtra:  p.StaticExtra,
		TotalSupply:  p.TotalSupply,
	}
}

type PoolReserves []string

func (pr PoolReserves) Clone() PoolReserves {
	newPR := make([]string, len(pr))
	copy(newPR, pr)

	return pr
}

type PoolToken struct {
	Address   string `json:"address,omitempty"`
	Name      string `json:"name,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	Decimals  uint8  `json:"decimals,omitempty"`
	Weight    uint   `json:"weight,omitempty"`
	Swappable bool   `json:"swappable,omitempty"`
}

type SimpleCache struct {
	rw sync.RWMutex
	m  map[string]Pool
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		m: make(map[string]Pool),
	}
}

func (c *SimpleCache) Get(keys ...string) ([]Pool, []string) {
	var (
		result = make([]Pool, 0, len(keys))
		missed []string
	)
	c.rw.RLock()
	defer c.rw.RUnlock()

	for i := range keys {
		p, ok := c.m[keys[i]]
		if !ok {
			missed = append(missed, keys[i])
			continue
		}
		result = append(result, p)
	}

	return result, missed
}

func (c *SimpleCache) Set(pools ...Pool) {
	c.rw.Lock()
	defer c.rw.Unlock()

	for i := range pools {
		c.m[pools[i].Address] = pools[i]
	}
}
