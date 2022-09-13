package timeseriescache

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	bucketTimeSize = 60 // minute in seconds
)

type CacheElem struct {
	Timestamp time.Time
	Val       uint64
}

type TimeSeriesPriceCache struct {
	cache         map[int64][]CacheElem // bucket key = minute (unix)
	oldestRecord  time.Time
	lastestRecord time.Time
	expireDur     time.Duration

	slicePool sync.Pool
	l         sync.RWMutex
	logger    *zap.SugaredLogger
}

func New(expireDur time.Duration) *TimeSeriesPriceCache {
	return &TimeSeriesPriceCache{
		cache: make(map[int64][]CacheElem),
		slicePool: sync.Pool{
			New: func() any {
				return make([]CacheElem, 0)
			},
		},
		expireDur: expireDur,
		logger:    zap.NewNop().Sugar(),
	}
}

func (t *TimeSeriesPriceCache) genBucketKey(ti time.Time) int64 {
	unixTime := ti.UTC().Unix()
	return unixTime - unixTime%bucketTimeSize
}

func (t *TimeSeriesPriceCache) unsafeInvalidateCache() {
	validTime := t.lastestRecord.Add(-1 * t.expireDur)

	if validTime.Before(t.oldestRecord) { // no invalid cache
		return
	}

	validBucketKey := t.genBucketKey(validTime)
	oldestBucketKey := t.genBucketKey(t.oldestRecord)

	for k := oldestBucketKey; k <= validBucketKey; k += bucketTimeSize {
		if _, ok := t.cache[k]; ok {
			t.logger.Debugw("unsafeInvalidateCache - remove", "key", k)
			t.slicePool.Put(t.cache[k])
			delete(t.cache, k)
			t.oldestRecord = time.Unix(k, 0)
		}
	}
}

func (t *TimeSeriesPriceCache) SetLogger(logger *zap.SugaredLogger) {
	if logger == nil {
		return
	}

	const loggerName = "TimeSeriesPriceCache"
	t.logger = logger.Named(loggerName)
}

func (t *TimeSeriesPriceCache) Put(timestamp time.Time, val uint64) {
	t.l.Lock()
	defer t.l.Unlock()
	if t.oldestRecord.IsZero() {
		t.oldestRecord = timestamp
	}

	bucketKey := t.genBucketKey(timestamp)
	if timestamp.After(t.lastestRecord) {
		t.lastestRecord = timestamp
	}
	if timestamp.Before(t.oldestRecord) {
		t.oldestRecord = timestamp
	}
	t.logger.Debugw("Put", "lastestRecord", t.lastestRecord, "oldestRecord", t.oldestRecord)

	if _, ok := t.cache[bucketKey]; ok {
		t.logger.Debugw("Put - bucket found", "key", bucketKey, "timestamp", timestamp, "val", val)
		t.cache[bucketKey] = append(t.cache[bucketKey], CacheElem{
			Timestamp: timestamp,
			Val:       val,
		})
		return
	}

	t.unsafeInvalidateCache()

	poolElem := t.slicePool.Get()
	elem, ok := poolElem.([]CacheElem)
	if !ok {
		panic("wrong elem pool type")
	}

	elem = append(elem, CacheElem{
		Timestamp: timestamp,
		Val:       val,
	})

	t.logger.Debugw("Put - new bucket", "key", bucketKey)
	t.cache[bucketKey] = elem
}

func (t *TimeSeriesPriceCache) Get(from, to time.Time) []CacheElem {
	t.l.RLock()
	defer t.l.RUnlock()

	fromBucketKey := t.genBucketKey(from)
	toBucketKey := t.genBucketKey(to)

	var result []CacheElem
	for k := fromBucketKey; k <= toBucketKey; k += bucketTimeSize {
		if v, ok := t.cache[k]; ok {
			t.logger.Debugw("Get - found bucket", "key", k, "from", from, "to", to)
			for i := range v {
				if v[i].Timestamp.After(from) && v[i].Timestamp.Before(to) {
					result = append(result, v[i])
				}
			}
		}
	}

	return result
}
