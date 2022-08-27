package scheduler_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thanhpp/gopher/pkg/scheduler"
)

func TestHourScheduler(t *testing.T) {
	s, err := scheduler.NewHourScheduler(1, 2, 3)
	require.NoError(t, err)

	ti, err := time.Parse(time.Kitchen, "2:01AM")
	require.NoError(t, err)

	for i := 0; i < 60*48; ti, i = ti.Add(time.Minute), i+1 {
		if s.ShouldTrigger(ti) {
			t.Logf("triggered at %d:%d", ti.Hour(), ti.Minute())
			s.SetTriggered()
			s.Debug()
		}
	}
}
