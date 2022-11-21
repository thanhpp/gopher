// nolint: lll
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	f, _ := os.Create(strconv.FormatInt(time.Now().UnixNano(), 10) + ".csv")
	defer f.Close()

	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()
	for i := 1; i < 1000; i++ {
		tasks := merchbarTask(f, i)
		chromedp.Run(ctx, tasks)
		log.Println("finish", i)
		<-ticker.C
	}
}

func merchbarTask(f *os.File, id int) chromedp.Tasks {
	url := "https://www.merchbar.com/search?idx=Merch_created_desc&cpt=Ubbqvrf%20%26%20Fjrngfuvegf&hMn%5BhierarchicalCategories.lvl0%5D=Sweatshirts&p=" + strconv.Itoa(id)
	log.Println("visiting", url)
	return chromedp.Tasks{
		// go to page
		chromedp.Navigate(
			url,
		),
		// crawl action
		chromedp.ActionFunc(func(ctx context.Context) error {
			// https://viblo.asia/p/su-dung-chromedp-voi-golang-de-crawl-cac-trang-web-co-noi-dung-duoc-tao-boi-javascript-4P856GWLKY3
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				panic(err)
			}
			log.Println("get node OK")

			res, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			if err != nil {
				panic(err)
			}
			log.Println("get res OK")

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
			if err != nil {
				panic(err)
			}
			log.Println("get doc OK")

			doc.Find("#__next > div > div.Layout_contentArea__CEkAb > main > div.Background_background__I3Gag.Background_lightGray__2KYLW > div > div > div.my-2.row > div.SearchInterface_productsContainer__ctL3_.col > section > div:nth-child(4)").
				Children().Each(func(i int, s *goquery.Selection) {
				name := s.Find(".MerchTileV2_titleEllipsis__QwuDG").Text()
				log.Println("get name OK")
				img, _ := s.Find(".MerchTileV2_imageContainer__Ap4zA > img").Attr("src")
				log.Println("get img OK")

				// remove img params
				tmp := strings.Split(img, "?")
				img = tmp[0]

				str := fmt.Sprintf("%d,\"%s\",\"%s\"\n", i, name, img)

				log.Println(str)
				f.WriteString(str)
			})
			return nil
		}),
	}
}
