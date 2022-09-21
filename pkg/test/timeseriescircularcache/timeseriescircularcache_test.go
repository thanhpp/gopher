package timeseriescircularcache_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/test/timeseriescircularcache"
)

func TestCache(t *testing.T) {
	c := timeseriescircularcache.New(time.Second*10, 20) // 20 slots -> 0.5s/slot

	dataFeed := make([]timeseriescircularcache.Value, 110)
	ts := time.Now()
	for i := range dataFeed {
		dataFeed[i] = timeseriescircularcache.Value{Ts: ts, Dat: 1}
		ts = ts.Add(time.Second / 10)
	}

	for i := range dataFeed {
		c.Insert(dataFeed[i].Timestamp(), dataFeed[i].Data())
	}
	t.Log("cache size", c.Size())

	var (
		start = 101
		end   = len(dataFeed) - 1
	)
	cached, err := c.Get(dataFeed[start].Timestamp(), dataFeed[end].Ts)
	require.NoError(t, err)
	t.Log("cached length", len(cached))

	for i := range cached {
		assert.Equal(t, dataFeed[start+i].Ts, cached[i].Timestamp())
	}
}

func BenchmarkInsert(b *testing.B) {
	c := timeseriescircularcache.New(time.Minute, 100) // 100 slots -> 0.1s/slot

	dataFeed := make([]timeseriescircularcache.Value, 300)
	ts := time.Now()
	for i := range dataFeed {
		dataFeed[i] = timeseriescircularcache.Value{Ts: ts, Dat: 1}
		ts = ts.Add(time.Second / 10)
	}

	for i := 0; i < b.N; i++ {
		for i := range dataFeed {
			c.Insert(dataFeed[i].Timestamp(), dataFeed[i].Data())
		}
	}
}
