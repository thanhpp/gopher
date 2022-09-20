package timeseriescircularcache

import (
	"errors"
	"time"
)

type TimeSeriesCacheValue interface {
	Timestamp() time.Time
	Data() interface{}
}

type Value struct {
	Ts  time.Time
	Dat interface{}
}

func (v *Value) Timestamp() time.Time {
	return v.Ts
}
func (v *Value) Data() interface{} {
	return v.Dat
}

type TimeSeriesCache interface {
	// Insert invalidate old data of timestamp - oldestTimestamp > dur
	Insert(timestamp time.Time, data interface{})
	// Get returns an error if from < oldestTimestamp
	Get(from, to time.Time) ([]TimeSeriesCacheValue, error)
	RemoveBefore(timestamp time.Time)
}

type cacheImpl struct {
	values       [][]TimeSeriesCacheValue
	startIdx     int64
	startMinTime time.Time
	slotTimeSize time.Duration
	cacheDur     time.Duration
}

func New(duration time.Duration, slotCount int64) *cacheImpl {
	slotTimeSize := duration.Nanoseconds() / slotCount

	return &cacheImpl{
		values:       make([][]TimeSeriesCacheValue, slotCount),
		startIdx:     0,
		startMinTime: time.Now(),
		slotTimeSize: time.Duration(slotTimeSize),
		cacheDur:     duration,
	}
}

func (c *cacheImpl) Insert(timestamp time.Time, data interface{}) {
	// invalid old cache to store new data
	if timestamp.After(c.startMinTime.Add(c.cacheDur)) {
		c.RemoveBefore(timestamp.Add(-c.cacheDur))
	}

	slotIdx := c.getSlotIdx(timestamp)
	for i := len(c.values[slotIdx]) - 1; i >= 1; i-- {
		curr := c.values[slotIdx][i]
		before := c.values[slotIdx][i-1]

		// the timestamp must follow asc order -> can find old cache
		if curr.Timestamp().Before(before.Timestamp()) {
			c.values[slotIdx][i] = &Value{
				Ts:  timestamp,
				Dat: data,
			}
			return
		}
	}

	c.values[slotIdx] = append(c.values[slotIdx], &Value{
		Ts:  timestamp,
		Dat: data,
	})
	// log.Println("[DEBUG] inserted to slot", slotIdx, "value", data, "timestamp", timestamp)
}

func (c *cacheImpl) Get(from, to time.Time) ([]TimeSeriesCacheValue, error) {
	if from.Before(c.startMinTime) {
		return nil, errors.New("too old range")
	}
	// log.Println("[DEBUG] GET", "from", from.UnixNano(), "to", to.UnixNano(), "startMinTime", c.startMinTime.UnixNano())

	slotIdx := c.getSlotIdx(from)
	// log.Println("[DEBUG] GET", "slotIdx", slotIdx)
	skip := true
	result := make([]TimeSeriesCacheValue, 0)

	for i := slotIdx; ; i = (i + 1) % int64(len(c.values)) {
		// log.Println("[DEBUG] GET - loop", "slotIdx", i)
		if !skip && i == slotIdx {
			break
		}
		skip = false

		for _, v := range c.values[i] {
			if v.Timestamp().After(to) {
				// log.Println("[DEBUG] GET - return", "slotIdx", i, "cache timestamp",
				// v.Timestamp().UnixNano(), "to", to.UnixNano())
				return result, nil
			}
			if v.Timestamp().Equal(from) || v.Timestamp().After(from) {
				result = append(result, v)
			}
			// log.Println("[DEBUG] GET", "slotIdx", i, "cache timestamp", v.Timestamp().UnixNano(), "from", from.UnixNano())
		}
	}

	return result, nil
}

func (c *cacheImpl) RemoveBefore(timestamp time.Time) {
	if timestamp.Before(c.startMinTime) {
		return
	}

	skip := true
	idx := c.startIdx
	for i := 0; ; i++ {
		if !skip && idx == c.startIdx {
			break
		}
		skip = false
		idx = (idx + 1) % int64(len(c.values))

		minTime := c.startMinTime.Add((time.Duration(i) * c.slotTimeSize))
		if timestamp.After(minTime) {
			continue
		}

		c.startIdx = idx
		c.startMinTime = c.startMinTime.Add((time.Duration(i) * c.slotTimeSize))
	}
}

func (c *cacheImpl) getSlotIdx(timestamp time.Time) int64 {
	i64Dur := timestamp.UnixNano() - c.startMinTime.UnixNano()

	return (i64Dur / c.slotTimeSize.Nanoseconds()) % int64(len(c.values))
}
