package circularbucketv2

import (
	"errors"
	"log"
	"time"
)

type Value struct {
	Timestamp time.Time
	Value     interface{}
}

type CircularBucketCache interface {
	// Insert invalidate old data of timestamp - oldestTimestamp > dur
	Insert(timestamp time.Time, data interface{})
	// Get returns an error if from < oldestTimestamp
	Get(from, to time.Time) ([]Value, error)
}

type bucket struct {
	values  []Value
	minTime int64
}

type cacheImpl struct {
	buckets        []bucket
	headIdx        int64
	cacheDur       int64
	bucketsCount   int64
	bucketTimeSize int64
}

func New(cacheDuration time.Duration, bucketsCount int64) *cacheImpl {
	i64CacheDur := durToI64(cacheDuration)
	// log.Println("[DEBUG] bucketTimeSize: i64CacheDur / bucketsCount", i64CacheDur/bucketsCount,
	// 	"i64CacheDur", i64CacheDur, "bucketsCount", bucketsCount)
	return &cacheImpl{
		buckets:        make([]bucket, bucketsCount),
		headIdx:        0,
		cacheDur:       i64CacheDur,
		bucketsCount:   bucketsCount,
		bucketTimeSize: i64CacheDur / bucketsCount,
	}
}

func (c *cacheImpl) rrNext(v, total int64) int64 {
	return (v + 1) % total
}

func (c *cacheImpl) oldestTime() int64 {
	return c.buckets[c.headIdx].minTime
}

func (c *cacheImpl) makeBucket(ti int64) {
	var (
		skip = true
	)

	for i := c.headIdx; ; i = c.rrNext(i, c.bucketsCount) {
		if !skip && i == c.headIdx {
			break
		}
		skip = false

		if c.buckets[i].minTime+c.cacheDur > ti {
			c.headIdx = i
			// log.Println("[DEBUG] invalidBefore - found", ti, "head", c.headIdx)
			return
		}
	}

	// no avaible slot -> cache reset
	c.headIdx = 0
	c.buckets[c.headIdx].minTime = ti - (ti % c.bucketTimeSize)
	// log.Println("[DEBUG] invalidBefore - reset", ti, "head", c.headIdx)
	return
}

func (c *cacheImpl) getBucketIdx(ti int64) int64 {
	diff := ti - c.oldestTime()
	offset := diff / c.bucketTimeSize
	idx := (c.headIdx + offset) % c.bucketsCount
	// log.Println("[DEBUG] getBucketIdx", "headIdx", c.headIdx, "c.oldestTime()", c.oldestTime(),
	// 	"ti", ti, "diff", diff, "offset", offset, "idx", idx)

	return idx
}

func (c *cacheImpl) setBucketMinTime(idx int64) {
	var offset int64
	if idx > c.headIdx {
		offset = idx - c.headIdx
	} else {
		offset = c.bucketsCount - c.headIdx + idx
	}

	c.buckets[idx].minTime = c.buckets[c.headIdx].minTime + (offset * c.bucketTimeSize)
}

func (c *cacheImpl) Insert(timestamp time.Time, data interface{}) {
	i64Ts := timeToI64(timestamp)
	if i64Ts > (c.oldestTime() + c.cacheDur) {
		c.makeBucket(i64Ts)
	}

	bucketIdx := c.getBucketIdx(i64Ts)
	// the bucket minTime is outdate -> update it
	if (c.buckets[bucketIdx].minTime + c.bucketTimeSize) < i64Ts {
		c.setBucketMinTime(bucketIdx)
	}

	// first element is outdate
	if len(c.buckets[bucketIdx].values) > 0 {
		if timeToI64(c.buckets[bucketIdx].values[0].Timestamp) < c.buckets[bucketIdx].minTime {
			c.buckets[bucketIdx].values[0].Timestamp = timestamp
			c.buckets[bucketIdx].values[0].Value = data
			return
		}
	}

	for i := len(c.buckets[bucketIdx].values) - 1; i > 0; i-- {
		// values of each bucket must follow asc order, if not -> current index is outdated
		if timeToI64(c.buckets[bucketIdx].values[i].Timestamp) < timeToI64(c.buckets[bucketIdx].values[i-1].Timestamp) {
			c.buckets[bucketIdx].values[i].Timestamp = timestamp
			c.buckets[bucketIdx].values[i].Value = data
			return
		}
	}

	// can not replace -> append
	c.buckets[bucketIdx].values = append(c.buckets[bucketIdx].values, Value{
		Timestamp: timestamp,
		Value:     data,
	})
}

func (c *cacheImpl) Get(from, to time.Time) ([]Value, error) {
	if from.After(to) {
		return nil, errors.New("from after to")
	}

	i64From := timeToI64(from)
	if c.oldestTime() > i64From {
		log.Println("c.oldestTime()", c.oldestTime(), "from", i64From)
		return nil, errors.New("to old data")
	}

	i64To := timeToI64(to)
	if (i64To - i64From) > c.cacheDur {
		return nil, errors.New("out of cache range")
	}

	bucketIdx := c.getBucketIdx(i64From)
	skip := true
	result := make([]Value, 0)

	for i := bucketIdx; ; i = c.rrNext(i, c.bucketsCount) {
		log.Println("[DEBUG] cacheImpl.Get", "bucketIdx", i)
		if !skip && i == bucketIdx {
			break
		}
		skip = false

		for _, v := range c.buckets[i].values {
			if v.Timestamp.After(to) {
				log.Println("[DEBUG] cacheImpl.Get - return", "v.Timestamp", v.Timestamp, "to", to)
				return result, nil
			}
			if v.Timestamp.After(from) || v.Timestamp.Equal(from) {
				log.Println("[DEBUG] cacheImpl.Get - append", "timestamp", v.Timestamp.UnixMilli())
				result = append(result, v)
				continue
			}
		}
	}

	return result, nil
}

func (c *cacheImpl) DebugPrint() {
	skip := true

	for i := c.headIdx; ; i = c.rrNext(i, c.bucketsCount) {
		if !skip && i == c.headIdx {
			break
		}
		skip = false

		log.Printf("head %d, idx %d, minTime %d", c.headIdx, i, c.buckets[i].minTime)
		for _, v := range c.buckets[i].values {
			log.Printf("\t ts %d, val %+v", timeToI64(v.Timestamp), v.Value)
		}
	}
}

func durToI64(dur time.Duration) int64 {
	return dur.Milliseconds()
}

func timeToI64(ti time.Time) int64 {
	return ti.UnixMilli()
}
