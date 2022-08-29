package kyberswapbsc

import (
	"errors"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ContractAddress             = "0x617dee16b86534a5d792a4d7a62fb491b544111e"
	KyberswapBSCSwappedLogTopic = "0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8"
)

func GetSwappedLog(kyberSwapBSC *Kyberswapbsc, logs []*types.Log) (*KyberswapbscSwapped, error) {
	if kyberSwapBSC == nil {
		return nil, errors.New("nil kyberswap bsc")
	}

	if logs == nil {
		return nil, errors.New("kyberswapbsc - get swapped log - nil logs")
	}

	for _, l := range logs {
		for _, t := range l.Topics {
			if t.String() == KyberswapBSCSwappedLogTopic {
				return kyberSwapBSC.ParseSwapped(*l)
			}
		}
	}

	return nil, errors.New("kyberswap bsc - swapped log not found")
}
