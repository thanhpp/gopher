package timeseriescache_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thanhpp/gopher/pkg/test/timeseriescache"
)

func TestTimeSeriesPriceCache(t *testing.T) {
	// insert data
	var (
		expireDur = time.Hour
		dataCount = 10_000
		timestamp = time.Now()
	)
	cache := timeseriescache.New(expireDur)

	// logger, err := zap.NewDevelopment()
	// require.NoError(t, err)
	// cache.SetLogger(logger.Sugar())

	// fill data
	for i := 0; i < dataCount; i++ {
		cache.Put(timestamp, 1)
		timestamp = timestamp.Add(time.Second * time.Duration(rand.Int63n(15)))
	}

	val := cache.Get(timestamp.Add(time.Minute*-50), timestamp)
	t.Log("len(val)", len(val))
	for i := 1; i < len(val); i++ {
		assert.True(t, val[i].Timestamp.Unix() >= val[i-1].Timestamp.Unix())
	}
}

func BenchmarkTimeSeriesPriceCache(b *testing.B) {
	// insert data
	var (
		expireDur = time.Hour
		timestamp = time.Now()
	)

	cache := timeseriescache.New(expireDur)

	// fill data
	for i := 0; i < b.N; i++ {
		cache.Put(timestamp, 1)
		timestamp = timestamp.Add(time.Second * time.Duration(rand.Int63n(180)))
	}
	/*
		[thanhpp@x1carbon] [~/go/src/github.com/thanhpp/gopher/pkg/test/timeseriescache] [main  ]
		[13:34:51] $  go test -bench=TimeSeriesPriceCache -benchtime=100000000x -benchmem
		goos: linux
		goarch: amd64
		pkg: github.com/thanhpp/gopher/pkg/test/timeseriescache
		cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
		BenchmarkTimeSeriesPriceCache-8         100000000              358.3 ns/op           270 B/op          4 allocs/op
		PASS
		ok      github.com/thanhpp/gopher/pkg/test/timeseriescache      35.911s
	*/
}
