package service

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/thanhpp/gopher/pkg/crawlers/weathervn"
	"github.com/thanhpp/gopher/pkg/discordx"
)

var (
	triggerHour   = []int{0, 6, 12, 18}
	fetchInterval = time.Minute
)

type HanoiAlert struct {
	weatherC              *weathervn.Crawler
	discordC              *discordx.RestClient
	discordWeatherChannel string

	// should send
	lock     sync.RWMutex
	nextSent time.Time
}

func NewHanoiAlert(weatherC *weathervn.Crawler, discordC *discordx.RestClient, discordChan string) *HanoiAlert {
	return &HanoiAlert{
		weatherC:              weatherC,
		discordC:              discordC,
		discordWeatherChannel: discordChan,
	}
}

func (a *HanoiAlert) Start() {
	var (
		t            = time.NewTicker(fetchInterval)
		lastSentHour int
	)

	defer t.Stop()

	for ; true; <-t.C {
		weatherInfo, err := a.weatherC.GetHanoiInfo()
		if err != nil {
			log.Println("[SKIP ERROR]", err)
			continue
		}

		now := time.Now()

		if h := now.Hour() % 6; h != 0 || (h == 0 && h == lastSentHour) {
			continue
		}

		if err := a.discordC.CreateContentOnlyMessage(
			a.discordWeatherChannel,
			a.formDiscordMessage(weatherInfo),
		); err != nil {
			log.Println("[SKIP ERROR]", err)
			continue
		}

		lastSentHour = now.Hour()
	}
}

func (a *HanoiAlert) formDiscordMessage(info *weathervn.WeatherInfo) string {
	return fmt.Sprintf(`> HANOI WEATHER ALERT
Temp: %s
Humd: %s
Stat: %s
Time: %s`,
		info.Temperature, info.Humidity, info.Status, info.LastUpdated)
}
