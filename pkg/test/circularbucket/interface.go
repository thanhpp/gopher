package circularbucket

import (
	"errors"
	"time"

	"go.uber.org/zap"
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
	values         []Value
	timeLowerBound time.Time
}

type circularbucketImpl struct {
	buckets        []bucket
	head           int
	dur            time.Duration
	bucketTimesize time.Duration
	logger         *zap.SugaredLogger
}

func New(dur time.Duration, buckets int, logger *zap.SugaredLogger) *circularbucketImpl {
	if logger == nil {
		logger = zap.NewNop().Sugar()
	}

	return &circularbucketImpl{
		buckets:        make([]bucket, buckets),
		head:           0,
		dur:            dur,
		bucketTimesize: dur / time.Duration(buckets),
		logger:         logger,
	}
}

func (c *circularbucketImpl) Insert(timestamp time.Time, data interface{}) {
	// WARN: case (too old timestamp)

	if maxTime := c.buckets[c.head].timeLowerBound.Add(c.dur); maxTime.Before(timestamp) {
		c.makeBucket(timestamp.Add(-c.dur))
	}

	bucketIdx := c.getBucketIdx(timestamp)
	if c.buckets[bucketIdx].timeLowerBound.Before(c.buckets[c.head].timeLowerBound) {
		c.setBucketTimeLowerBound(bucketIdx)
	}
	c.logger.Debugw("circularbucketImpl.Insert", "bucketIdx", bucketIdx,
		"bucketTimeBound", c.buckets[bucketIdx].timeLowerBound.UnixNano(), "timestamp", timestamp.UnixNano(), "data", data)
	// values in each bucket is sorted in asc order
	// old record will shows as a desc order between 2 cache
	for i := 1; i < len(c.buckets[bucketIdx].values); i++ {
		if c.buckets[bucketIdx].values[i].Timestamp.After(c.buckets[bucketIdx].values[i-1].Timestamp) {
			continue
		}

		c.buckets[bucketIdx].values[i].Timestamp = timestamp
		c.buckets[bucketIdx].values[i].Value = data
		return
	}

	// can not replace an exist slot => append a new one
	c.buckets[bucketIdx].values = append(c.buckets[bucketIdx].values, Value{
		Timestamp: timestamp,
		Value:     data,
	})
}

func (c *circularbucketImpl) Get(from, to time.Time) ([]Value, error) {
	if from.Before(c.buckets[c.head].timeLowerBound) {
		return nil, errors.New("too old data")
	}

	if from.After(to) {
		return nil, errors.New("from after to")
	}

	var (
		bucketIdx = c.getBucketIdx(from)
		result    []Value
		skip      = true
	)
	for i := bucketIdx; ; i = (i + 1) % len(c.buckets) {
		if !skip && i == bucketIdx {
			break
		}
		skip = false

		for _, v := range c.buckets[i].values {
			if v.Timestamp.After(to) {
				return result, nil
			}
			if v.Timestamp.After(from) {
				result = append(result, v)
			}
			break // reading invalid cache
		}
	}

	return result, nil
}

func (c *circularbucketImpl) makeBucket(ts time.Time) {
	c.logger.Debugw("circularbucketImpl.invalidBefore - start", "head", c.head, "timestamp", ts.UnixNano())
	defer c.logger.Debugw("circularbucketImpl.invalidBefore - end", "head", c.head, "timestamp", ts.UnixNano())
	prev := c.head
	curr := (c.head + 1) % len(c.buckets)
	skip := true

	for ; ; curr = (curr + 1) % len(c.buckets) {
		if !skip && curr == (c.head+1)%len(c.buckets) {
			break
		}
		skip = false

		if c.buckets[curr].timeLowerBound.After(ts) {
			c.head = prev
			return
		}
		prev = curr
	}

	// no available buckets => reset the cache
	c.head = 0
	c.buckets[c.head].timeLowerBound = ts
}

func (c *circularbucketImpl) getBucketIdx(ts time.Time) int {
	c.logger.Debugw("circularbucketImpl.getBucketIdx", "timestamp", ts.UnixNano(),
		"c.buckets[c.head].timeLowerBound", c.buckets[c.head].timeLowerBound.UnixNano())

	timeDiff := ts.Sub(c.buckets[c.head].timeLowerBound)
	offset := timeDiff / c.bucketTimesize
	idx := (c.head + int(offset)) % len(c.buckets)

	c.logger.Debugw("circularbucketImpl.getBucketIdx", "timeDiff", timeDiff.Milliseconds(),
		"offset", int(offset), "idx", idx)

	return idx
}

func (c *circularbucketImpl) setBucketTimeLowerBound(idx int) {
	var offset int
	if idx >= c.head {
		offset = idx - c.head
	} else {
		offset = idx + (len(c.buckets) - c.head)
	}

	timeLowerBound := c.buckets[c.head].timeLowerBound.Add(time.Duration(offset) * c.bucketTimesize)

	c.logger.Debugw("circularbucketImpl.setBucketTimeLowerBound", "idx", idx, "timeBound", timeLowerBound.UnixNano())
}

func (c *circularbucketImpl) Size() int {
	var (
		skip = true
		size int
	)
	for i := c.head; ; i = (i + 1) % len(c.buckets) {
		if !skip && i == c.head {
			break
		}

		size += len(c.buckets[i].values)
	}

	return size
}
