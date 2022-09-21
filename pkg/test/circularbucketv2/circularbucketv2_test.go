package circularbucketv2_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/test/circularbucketv2"
)

func TestCircularBucket(t *testing.T) {
	var (
		// 1 sec/bucket
		cacheDur     = time.Second * 2
		bucketsCount = 2
	)
	cache := circularbucketv2.New(cacheDur, int64(bucketsCount))

	var (
		dataCount = 13
		dataFeed  = make([]circularbucketv2.Value, dataCount)
		ts        = time.Now()
	)

	for i := range dataFeed {
		dataFeed[i].Timestamp = ts
		dataFeed[i].Value = i
		ts = ts.Add(time.Second / 3)
	}

	to := dataFeed[len(dataFeed)-1].Timestamp
	from := to.Add(-cacheDur / 2)
	var test []circularbucketv2.Value
	for i := 0; i < len(dataFeed); i++ {
		if dataFeed[i].Timestamp.After(from) {
			test = append(test, dataFeed[i])
		}
	}

	for i := range dataFeed {
		cache.Insert(dataFeed[i].Timestamp, dataFeed[i].Value)
	}
	cache.DebugPrint()

	t.Log("from", test[0].Timestamp.UnixMilli(), "to", test[len(test)-1].Timestamp.UnixMilli())
	vals, err := cache.Get(test[0].Timestamp, test[len(test)-1].Timestamp)
	require.NoError(t, err)
	assert.Equal(t, len(test), len(vals))

	cache.DebugPrint()
	log.Println("---------")
}

func BenchmarkCircularBucketInsert(b *testing.B) {
	var (
		dataCount = 1_000_000
		dataFeed  = make([]circularbucketv2.Value, dataCount)
		ts        = time.Now()
	)

	for i := range dataFeed {
		dataFeed[i].Timestamp = ts
		dataFeed[i].Value = i
		ts = ts.Add(time.Second)
	}

	for i := 0; i < b.N; i++ {
		var (
			// 1 sec/bucket
			cacheDur     = time.Hour
			bucketsCount = 100
		)
		cache := circularbucketv2.New(cacheDur, int64(bucketsCount))

		for j := range dataFeed {
			cache.Insert(dataFeed[j].Timestamp, dataFeed[j].Value)
		}
	}
}
