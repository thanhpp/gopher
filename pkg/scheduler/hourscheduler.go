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
	hourList      []int
	rangeStartIdx int
	rangeEndIdx   int
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
		hourList:      hours,
		rangeStartIdx: -1,
		rangeEndIdx:   -1,
	}, nil
}

func (s *HourScheduler) ShouldTrigger(t time.Time) bool {
	h := t.Hour()

	// find appropriate range
	if s.rangeStartIdx == -1 {
		for i := 0; i < len(s.hourList)-1; i++ {
			if h >= s.hourList[i] && h < s.hourList[i+1] {
				s.rangeStartIdx = i
				s.rangeEndIdx = (i + 1) % len(s.hourList)

				return true
			}
		}

		s.rangeStartIdx = len(s.hourList) - 1
		s.rangeEndIdx = 0

		return true
	}

	if s.hourList[s.rangeStartIdx] < s.hourList[s.rangeEndIdx] {
		if h >= s.hourList[s.rangeStartIdx] && h < s.hourList[s.rangeEndIdx] {
			return true
		}

		return false
	}

	if h >= s.hourList[s.rangeStartIdx] || h < s.hourList[s.rangeEndIdx] {
		return true
	}

	return false
}

func (s *HourScheduler) SetTriggered() {
	s.rangeStartIdx = (s.rangeStartIdx + 1) % len(s.hourList)
	s.rangeEndIdx = (s.rangeEndIdx + 1) % len(s.hourList)
}

func (s *HourScheduler) Debug() {
	log.Printf("HourScheduler: %+v", s)
}
