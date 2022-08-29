package bscscan

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Crawler struct {
	collector *colly.Collector
}

func New() *Crawler {
	return &Crawler{
		collector: colly.NewCollector(),
	}
}

func (c *Crawler) GetTxBNBPrice(txHash string) (float64, error) {
	var (
		url         = fmt.Sprintf("https://bscscan.com/tx/%s", txHash)
		collector   = colly.NewCollector()
		strBNBPrice string
	)

	collector.OnHTML("#ContentPlaceHolder1_spanClosingPrice", func(h *colly.HTMLElement) {
		strBNBPrice = h.Text
	})
	collector.Visit(url)

	if len(strBNBPrice) == 0 {
		return 0, errors.New("bscscan crawler - 0 tx bnb price")
	}

	strBNBPrice = strings.ReplaceAll(strBNBPrice, "$", "")
	strBNBPrice = strings.ReplaceAll(strBNBPrice, " / BNB", "")

	f64BNBPrice, err := strconv.ParseFloat(strBNBPrice, 64)
	if err != nil {
		return 0, fmt.Errorf("bscscan crawler - convert bnb price error: %w, raw: %s", err, strBNBPrice)
	}

	return f64BNBPrice, nil
}
