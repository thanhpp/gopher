package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/thanhpp/gopher/internal/calsync/infra/gcal"
)

func main() {
	ctx := context.Background()
	gcalC, err := gcal.NewClient(ctx, os.Getenv("CREDENTIALS"), os.Getenv("TOKENS_FILE"))
	if err != nil {
		log.Fatal(err)
	}

	cals, err := gcalC.ListCalendars()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", cals)

	events, err := gcalC.GetCurrentMonthEvent(
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := range events {
		events[i].Desc = fmt.Sprintf("ref: %s-%s\n%s", events[i].CalID, events[i].ID, events[i].Desc)
		if err := gcalC.AddEvent(
			&events[i],
			"",
		); err != nil {
			log.Fatal(err)
		}
	}

	// events2, err := gcalC.GetCurrentMonthEvent(
	// 	"",
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for i := range events2 {
	// 	if err := gcalC.DeleteEvent(events2[i].ID,
	// 		""); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
