package service

import (
	"fmt"
	"log"
	"time"

	"github.com/thanhpp/gopher/pkg/crawlers/weathervn"
	"github.com/thanhpp/gopher/pkg/discordx"
	"github.com/thanhpp/gopher/pkg/scheduler"
)

var (
	fetchInterval = time.Minute
)

type HanoiAlert struct {
	weatherC              *weathervn.Crawler
	discordC              *discordx.RestClient
	discordWeatherChannel string

	// should send
	sched *scheduler.HourScheduler
}

func NewHanoiAlert(weatherC *weathervn.Crawler, discordC *discordx.RestClient, discordChan string) *HanoiAlert {
	sched, err := scheduler.NewHourScheduler(6, 8, 12, 18)
	if err != nil {
		panic(err)
	}

	return &HanoiAlert{
		weatherC:              weatherC,
		discordC:              discordC,
		discordWeatherChannel: discordChan,
		sched:                 sched,
	}
}

func (a *HanoiAlert) Start() {
	var (
		t = time.NewTicker(fetchInterval)
	)

	defer t.Stop()

	for ; true; <-t.C {
		weatherInfo, err := a.weatherC.GetHanoiInfo()
		if err != nil {
			log.Println("[SKIP ERROR]", err)
			continue
		}

		now := time.Now()

		a.sched.Debug()
		if !a.sched.ShouldTrigger(now) {
			continue
		}

		if err := a.discordC.CreateContentOnlyMessage(
			a.discordWeatherChannel,
			a.formDiscordMessage(weatherInfo),
		); err != nil {
			log.Println("[SKIP ERROR]", err)
			continue
		}

		a.sched.SetTriggered()
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
