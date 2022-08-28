package scheduler_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/scheduler"
)

func TestHourScheduler(t *testing.T) {
	s, err := scheduler.NewHourScheduler(1, 5, 13)
	require.NoError(t, err)

	ti, err := time.Parse(time.Kitchen, "0:00AM")
	require.NoError(t, err)

	for i := 0; i < 60*48; ti, i = ti.Add(time.Hour), i+1 {
		log.Println("current hour", ti.Hour())
		s.Debug()
		if s.ShouldTrigger(ti) {
			t.Logf("triggered at %d:%d", ti.Hour(), ti.Minute())
			s.SetTriggered()
		}
	}
}
