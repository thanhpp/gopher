package timeseriescache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/test/timeseriescache"
)

func TestCacheV2(t *testing.T) {
	cacheV2 := timeseriescache.CacheV2{}
	cacheV2.Init(time.Second.Milliseconds(), nil)

	var (
		dataFeed []timeseriescache.CacheVal
		test     []timeseriescache.CacheVal
	)
	ts := time.Now()

	for i := 0; i < 1000; i++ {
		dataFeed = append(dataFeed, timeseriescache.CacheVal{ts.UnixMilli(), 1})
		if i > 800 {
			test = append(test, timeseriescache.CacheVal{ts.UnixMilli(), 1})
		}
		ts = ts.Add(time.Second / 500)
	}

	for i := range dataFeed {
		cacheV2.Insert(dataFeed[i].Ts, dataFeed[i].Data)
	}

	t.Log("test cache retrieve", "from", test[0].Ts, "to", test[len(test)-1].Ts)
	cached, err := cacheV2.Retrieve(test[0].Ts, test[len(test)-1].Ts)
	require.NoError(t, err)
	assert.Equal(t, len(test), len(cached))
}
