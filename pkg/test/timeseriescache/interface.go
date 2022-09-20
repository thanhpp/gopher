package timeseriescache

import (
	"errors"
	"log"
	"time"
)

type CacheData struct {
	minTime int64
	vals    []CacheVal
}

type CacheVal struct {
	Ts   int64
	Data interface{}
}

type Interface interface {
	Init(duration int64, data []CacheData) error
	// Insert new data into cache, this method should be call with non-decreasing timestamp
	Insert(timestamp int64, data interface{}) error
	// Retrieve data in range of cache, most cases 'to' ~ now()
	Retrieve(from int64, to int64) ([]CacheData, error)
	Expire(before int64) error // Remove old data to save memory space, this function can be called by a ticker
}

type CacheV2 struct {
	arr      [200]CacheData
	dur      int64
	startIdx int
}

func (c *CacheV2) Init(duration int64, data []CacheVal) error {
	c.dur = duration
	c.startIdx = 0

	if len(data) == 0 {
		c.arr[c.startIdx].minTime = time.Now().UnixMilli()
	}

	for i := range data {
		c.Insert(data[i].Ts, data[i].Data)
	}

	return nil
}

func (c *CacheV2) Insert(timestamp int64, data interface{}) {
	startMinTime := c.arr[c.startIdx].minTime
	if timestamp-startMinTime > c.dur {
		log.Println("[INSERT] exec expire", "before", timestamp-c.dur)
		c.Expire(timestamp - c.dur)
	}

	slotIdx := c.getSlotIdx(timestamp)
	c.setSlotMinTime(slotIdx)
	slotMinTime := c.arr[slotIdx].minTime

	for _, v := range c.arr[slotIdx].vals {
		// outdated
		if v.Ts < slotMinTime {
			v.Ts = timestamp
			v.Data = data
			return
		}
	}

	c.arr[slotIdx].vals = append(c.arr[slotIdx].vals, CacheVal{timestamp, data})

	return
}

func (c *CacheV2) Retrieve(from int64, to int64) ([]CacheVal, error) {
	startMinTime := c.arr[c.startIdx].minTime
	if from < startMinTime {
		return nil, errors.New("cache not found, too old data")
	}

	slotIdx := c.getSlotIdx(from)
	var (
		result []CacheVal
		skip   = true
	)

	log.Println("[DEBUG] Retrieve", "slot idx", slotIdx)

	for i := slotIdx; ; i = (i + 1) % 200 {
		if !skip && i == slotIdx {
			log.Println("[DEBUG] Retrieve", "start idx", slotIdx, "skip", skip)
			break
		}
		skip = false

		log.Println("[DEBUG] Retrieve", "slot", slotIdx, "vals", len(c.arr[i].vals))
		for _, v := range c.arr[i].vals {
			if v.Ts > to {
				return result, nil
			}
			log.Println("[DEBUG] Retrieve", "slot", slotIdx, "val ts", v.Ts)
			if v.Ts > from {
				result = append(result, v)
			}
		}
	}

	return result, nil
}

func (c *CacheV2) Expire(before int64) {
	if before < c.arr[c.startIdx].minTime {
		return
	}

	var (
		newIdx int
		skip   = true
	)
	for i := c.startIdx; ; i = (i + 1) % 200 {
		if !skip && i == c.startIdx {
			break
		}
		skip = false

		if c.arr[i].minTime > before {
			break
		}
		newIdx = i
		continue
	}

	c.startIdx = newIdx

	return
}

func (c *CacheV2) getSlotIdx(ts int64) int64 {
	startMinTime := c.arr[c.startIdx].minTime
	slotOffset := (ts - startMinTime) / (c.dur / 200)
	slotIdx := (slotOffset + 200) % 200

	return slotIdx
}

func (c *CacheV2) setSlotMinTime(slotIdx int64) {
	c.arr[slotIdx].minTime = c.arr[c.startIdx].minTime + (c.dur/200)*slotIdx
}

func (c *CacheV2) Debug(withVals bool) {
	log.Println("debug cacheV2", "dur:", c.dur, "startIdx:", c.startIdx)

	skip := true
	for i := c.startIdx; ; i = (i + 1) % 200 {
		log.Println("[DEBUG] idx", i)
		if !skip && i == c.startIdx {
			break
		}
		skip = false

		log.Println("debug cacheV2 - slot", "idx", i, "minTime", c.arr[i].minTime)
		if withVals {
			for _, v := range c.arr[i].vals {
				log.Println("debug cacheV2 - val", "ts", v.Ts, "data", v.Data)
			}
		}
	}
}
