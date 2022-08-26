package weathervn

import "github.com/gocolly/colly"

type Crawler struct {
	forecastClt *colly.Collector
}

func New() *Crawler {
	return &Crawler{
		forecastClt: colly.NewCollector(),
	}
}

// const (
// 	todayForeCast = "https://nchmf.gov.vn/Kttvsite/vi-VN/1/thoi-tiet-dat-lien-24h-12h2-15.html"
// )
