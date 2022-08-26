package weathervn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thanhpp/gopher/pkg/crawlers/weathervn"
)

func TestWeatherCrawlerInfoNow(t *testing.T) {
	w := weathervn.New()
	for i := 0; i < 10; i++ {
		info, err := w.GetHanoiInfo()
		assert.NoError(t, err)
		t.Logf("%+v \n", info)
	}
}
