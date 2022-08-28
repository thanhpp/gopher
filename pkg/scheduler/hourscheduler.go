package scheduler

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"time"
)

// It will try to trigger once (success) in a defined hour list in a day.
type HourScheduler struct {
	hourList          []int
	nextRangeStartIdx int
	nextRangeEndIdx   int
}

// NewHourScheduler have a bug when len(hours) == 1.
func NewHourScheduler(hours ...int) (*HourScheduler, error) {
	if len(hours) == 0 {
		return nil, errors.New("empty hour list")
	}

	for i := range hours {
		if hours[i] < 0 || hours[i] >= 24 {
			return nil, errors.New("invalid hour, range [0, 23]")
		}
	}

	sort.SliceStable(hours, func(i, j int) bool {
		return hours[i] < hours[j]
	})

	for i := 1; i < len(hours); i++ {
		if hours[i-1] == hours[i] {
			return nil, fmt.Errorf("duplicate hour: %d", hours[i])
		}
	}

	if len(hours) == 1 {
		hours = append(hours, []int{hours[0]}...)
	}

	return &HourScheduler{
		hourList:          hours,
		nextRangeStartIdx: -1,
		nextRangeEndIdx:   -1,
	}, nil
}

func (s *HourScheduler) ShouldTrigger(t time.Time) bool {
	h := t.Hour()

	// find appropriate range
	if s.nextRangeStartIdx == -1 {
		for i := 0; i < len(s.hourList)-1; i++ {
			if h >= s.hourList[i] && h < s.hourList[i+1] {
				s.nextRangeStartIdx = i + 1
				s.nextRangeEndIdx = (s.nextRangeStartIdx + 1) % len(s.hourList)

				return true
			}
		}

		s.nextRangeStartIdx = len(s.hourList) - 1
		s.nextRangeEndIdx = 0

		return true
	}

	if s.hourList[s.nextRangeStartIdx] < s.hourList[s.nextRangeEndIdx] {
		if h >= s.hourList[s.nextRangeStartIdx] && h < s.hourList[s.nextRangeEndIdx] {
			return true
		}

		return false
	}

	if h >= s.hourList[s.nextRangeStartIdx] || h < s.hourList[s.nextRangeEndIdx] {
		return true
	}

	return false
}

func (s *HourScheduler) SetTriggered() {
	s.nextRangeStartIdx = (s.nextRangeStartIdx + 1) % len(s.hourList)
	s.nextRangeEndIdx = (s.nextRangeEndIdx + 1) % len(s.hourList)
}

func (s *HourScheduler) Debug() {
	log.Printf("HourScheduler: %+v", s)
}
