package trends

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Teste(searchTrend string) {
	fmt.Println("passo 1")

	searchTrend = strings.Trim(searchTrend, " ")
	searchTrend = strings.Replace(searchTrend, " ", "%20", -1)

	fmt.Println(searchTrend)

	linkTrends := ("https://trends.google.com.br/trends/explore?geo=BR&q=" + searchTrend)

	c := colly.NewCollector(
	// colly.Async(true),
	// colly.Debugger(&debug.LogDebugger{}),
	)

	// c.Limit(&colly.LimitRule{
	// 	Parallelism: 1,
	// 	RandomDelay: 5 * time.Second,
	// })

	c.SetRequestTimeout(120 * time.Second)

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Scraping:", r.URL)
	// })

	c.OnRequest(func(r *colly.Request) {

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}

		fmt.Println(r.Method)
	})

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Status:", r.StatusCode)
	// })

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("-----------------------------")

		fmt.Println(r.StatusCode)

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	})

	// c.OnError(func(r *colly.Response, e error) {
	// 	fmt.Println("Got this error:", e)
	// })

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	c.OnHTML("div.widget-template", func(h *colly.HTMLElement) {
		position := h.ChildText("div.label-line-number")
		name := h.ChildText("span.bidiText")
		nameRate := h.ChildText("div.rising-value")
		fmt.Println(position, name, nameRate)
	})

	c.Visit(linkTrends)

	// c.Wait()

}
