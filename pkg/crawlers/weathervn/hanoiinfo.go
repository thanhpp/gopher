package weathervn

import (
	"errors"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	hanoiWeather = "https://nchmf.gov.vn/Kttvsite/vi-VN/1/ha-noi-w28.html"
)

// #wrapper > div > div.uk-container > section > div > div > article > div >
// div > div.content-news.fix-content-news > div > div > div:nth-child(2) >
// div > ul > li:nth-child(2) > div > div.uk-width-3-4
type WeatherInfo struct {
	LastUpdated string `selector:"div.time-update"`
	Temperature string `selector:"ul > li:nth-child(1) > div > div.uk-width-3-4"`
	Status      string `selector:"ul > li:nth-child(2) > div > div.uk-width-3-4"`
	Humidity    string `selector:"ul > li:nth-child(3) > div > div.uk-width-3-4"`
}

func (wi *WeatherInfo) beautifyString(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, ":", ""))
}

func (wi *WeatherInfo) Beautify() {
	wi.LastUpdated = strings.TrimSpace(
		strings.ReplaceAll(
			strings.Join(strings.Fields(wi.LastUpdated), " "), "Cập nhật: ", "",
		))
	wi.Temperature = wi.beautifyString(wi.Temperature)
	wi.Status = wi.beautifyString(wi.Status)
	wi.Humidity = wi.beautifyString(wi.Humidity)
}

func (wi *WeatherInfo) IsRaining() bool {
	return strings.Contains(strings.ToLower(wi.Status), "có mưa")
}

func (wi *WeatherInfo) IsEmpty() bool {
	return len(wi.Status) == 0
}

func (w *Crawler) GetHanoiInfo() (*WeatherInfo, error) {
	var (
		collector = colly.NewCollector()
		infoNow   = new(WeatherInfo)
	)

	// #wrapper > div > div.uk-container > section > div > div > article > div > div >
	// div.content-news.fix-content-news > div > div > div:nth-child(2) > div
	collector.OnHTML(
		"div.content-news.fix-content-news > div > div > div:nth-child(2) > div",
		func(h *colly.HTMLElement) {
			if err := h.Unmarshal(infoNow); err != nil {
				return
			}
		},
	)
	if err := collector.Visit(hanoiWeather); err != nil {
		return nil, err
	}

	infoNow.Beautify()

	if infoNow.IsEmpty() {
		return nil, errors.New("empty hanoi weather info")
	}

	return infoNow, nil
}
