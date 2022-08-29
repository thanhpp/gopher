package chains

// Chain id for EVM base network
//go:generate enumer -type=Chain -linecomment -json=true -sql -yaml
type Chain int

const (
	Ethereum  Chain = 1     // ethereum
	BSC       Chain = 56    // bsc
	Polygon   Chain = 137   // polygon
	Fantom    Chain = 250   // fantom
	Arbitrum  Chain = 42161 // arbitrum
	Avalanche Chain = 43114 // avalanche
)
