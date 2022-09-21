package circularbucket_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/test/circularbucket"
	"go.uber.org/zap"
)

func TestCircularBucket(t *testing.T) {
	var (
		timeDur = time.Second * 2
		buckets = 2
		// 1s/bucket
	)
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	var (
		dataFeedLength = 2
		ts             = time.Now()
		dataFeed       = make([]circularbucket.Value, dataFeedLength)
	)
	for i := range dataFeed {
		dataFeed[i] = circularbucket.Value{Timestamp: ts, Value: i}
		ts = ts.Add(time.Nanosecond * 300)
	}

	cache := circularbucket.New(timeDur, buckets, logger.Sugar())
	for i := range dataFeed {
		cache.Insert(dataFeed[i].Timestamp, dataFeed[i].Value)
	}
}
