package crawlers

import "github.com/thanhpp/gopher/pkg/crawlers/weathervn"

func NewWeatherVN() *weathervn.Crawler {
	return weathervn.New()
}
