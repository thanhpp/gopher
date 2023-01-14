package gcal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	allDayTimeFormat = "2006-01-02"
)

type Client struct {
	httpC *http.Client
	srv   *calendar.Service
}

func NewClient(ctx context.Context, credentialsFile, tokensFile string) (Client, error) {
	// read credentials
	cred, err := os.ReadFile(credentialsFile)
	if err != nil {
		return Client{}, fmt.Errorf("read credentials file error: %w", err)
	}

	// get gcal config
	scopes := []string{calendar.CalendarEventsScope, calendar.CalendarScope, calendar.CalendarEventsReadonlyScope,
		calendar.CalendarReadonlyScope}
	cfg, err := google.ConfigFromJSON(cred, scopes...)
	if err != nil {
		return Client{}, fmt.Errorf("get config from json error: %w", err)
	}
	tok, err := getAndSetToken(ctx, cfg, tokensFile)
	if err != nil {
		return Client{}, fmt.Errorf("get and set token error: %w", err)
	}

	// create client
	client := cfg.Client(ctx, tok)
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return Client{}, fmt.Errorf("create calendar service error: %w", err)
	}

	return Client{
		httpC: client,
		srv:   srv,
	}, nil
}

// PUBLIC - READ

func (c *Client) ListCalendars() ([]Calendar, error) {
	calList, err := c.srv.CalendarList.List().Do()
	if err != nil {
		return nil, err
	}

	cals := make([]Calendar, len(calList.Items))
	for i := range cals {
		cals[i] = c.transformCalendar(calList.Items[i])
	}

	return cals, nil
}

func (c *Client) GetCurrentMonthEvent(calID string) ([]Event, error) {
	y, m, _ := time.Now().Date()
	tMin := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	tMax := tMin.AddDate(0, 1, -1)
	events, err := c.srv.Events.List(calID).ShowDeleted(false).SingleEvents(true).
		TimeMin(tMin.Format(time.RFC3339)).TimeMax(tMax.Format(time.RFC3339)).MaxResults(2_500).
		OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("get event list error: %w", err)
	}

	iEvents := make([]Event, len(events.Items))
	for i := range iEvents {
		iEv, err := c.transformEvent(events.Items[i], calID)
		if err != nil {
			return nil, fmt.Errorf("transform internal event error: %w", err)
		}
		iEvents[i] = iEv
	}

	return iEvents, nil
}

// PUBLIC - WRITE

func (c *Client) AddEvent(ev *Event, calID string) error {
	_, err := c.srv.Events.Insert(calID, &calendar.Event{
		Summary: ev.Name,
		Start: &calendar.EventDateTime{
			DateTime: ev.Start.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: ev.End.Format(time.RFC3339),
		},
		Description: ev.Desc,
	}).Do()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteEvent(evID, calID string) error {
	return c.srv.Events.Delete(calID, evID).Do()
}

// PRIVATE - TRANSFORM

func (c *Client) transformEvent(event *calendar.Event, calID string) (Event, error) {
	var (
		start, end time.Time
		err        error
	)
	// all day event
	if len(event.Start.Date) != 0 {
		startDate, err := time.Parse(allDayTimeFormat, event.Start.Date)
		if err != nil {
			return Event{}, fmt.Errorf("parse all day start date error: %w", err)
		}
		y, m, d := startDate.Date()
		start = time.Date(y, m, d, 0, 0, 0, 0, time.Local)
		end = time.Date(y, m, d, 23, 59, 59, 0, time.Local)
	} else {
		start, err = time.Parse(time.RFC3339, event.Start.DateTime)
		if err != nil {
			return Event{}, fmt.Errorf("parse event start datetime error: %w", err)
		}
	}

	if len(event.End.Date) != 0 {
		endDate, err := time.Parse(allDayTimeFormat, event.Start.Date)
		if err != nil {
			return Event{}, fmt.Errorf("parse all day start date error: %w", err)
		}
		y, m, d := endDate.Date()
		end = time.Date(y, m, d, 23, 59, 59, 0, time.Local)
	} else if end.IsZero() {
		end, err = time.Parse(time.RFC3339, event.End.DateTime)
		if err != nil {
			return Event{}, fmt.Errorf("parse event end datetime error: %w", err)
		}
	}

	iEv := Event{
		ID:    event.Id,
		CalID: calID,
		Name:  event.Summary,
		Start: start,
		End:   end,
		Desc:  event.Description,
	}

	return iEv, nil
}

func (c *Client) transformCalendar(cal *calendar.CalendarListEntry) Calendar {
	return Calendar{
		ID:   cal.Id,
		Name: cal.Summary,
	}
}

// PRIVATE - SET UP CLIENT
func getAndSetToken(ctx context.Context, config *oauth2.Config, tokensFile string) (*oauth2.Token, error) {
	tok := new(oauth2.Token)

	// get token from file
	data, err := os.ReadFile(tokensFile)
	if err == nil {
		if err1 := json.NewDecoder(bytes.NewBuffer(data)).Decode(tok); err1 != nil {
			return nil, fmt.Errorf("decode token error: %w", err1)
		}
		return tok, nil
	}

	// get token from web
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Printf("Go to the following authURL:\n%s\nthen type the authorization code:", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("scan auth code error: %w", err)
	}
	tok, err = config.Exchange(ctx, authCode)
	if err != nil {
		return nil, fmt.Errorf("exchange auth code error: %w", err)
	}

	// save exchange token to file
	f, err := os.OpenFile(tokensFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600) // nolint
	if err != nil {
		return nil, fmt.Errorf("open token file error: %w", err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(tok); err != nil {
		return nil, fmt.Errorf("encode error: %w", err)
	}

	return tok, nil
}
