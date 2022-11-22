package crawler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com.br/maickmachado/crawler-go/models"
	"github.com/PuerkitoBio/goquery"
)

func GoogleTrendsScrape(searchTerm string, proxyString interface{}) ([]models.ResultGoogleSearch, error) {

	results := []models.ResultGoogleSearch{}

	resultCounter := 0
	fmt.Println("passo 1")
	googlePageURL, err := BuildGoogleTrendsURl(searchTerm)
	if err != nil {
		return nil, err
	}
	fmt.Println("passo 2")
	res, err := models.ScrapeClientRequest(googlePageURL, proxyString)
	if err != nil {
		return nil, err
	}

	fmt.Println("passo 3")
	data, err := TrendsResultParsing(res, resultCounter)
	if err != nil {
		return nil, err
	}
	fmt.Println("passo 4")
	resultCounter += len(data)
	for _, result := range data {
		results = append(results, result)
	}
	time.Sleep(time.Duration(10) * time.Second)

	return results, nil
}

//TODO: juntar esse com a do trends
//TODO: retirar o pais e colocar somente brasil
func BuildGoogleTrendsURl(searchTerm string) (string, error) {

	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "%20", -1)

	linkTrends := ("https://trends.google.com.br/trends/explore?geo=BR&q=" + searchTerm)

	return linkTrends, nil
}

func TrendsResultParsing(response *http.Response, rank int) ([]models.ResultGoogleSearch, error) {
	// r := response.Body
	// t, e := goquery.NewDocumentFromReader(r)

	fmt.Println("trendResultsParsing 1")

	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		return nil, err
	}

	fmt.Println("trendResultsParsing 2")

	results := []models.ResultGoogleSearch{}
	sel := doc.Find("div.fe-related-queries.fe-atoms-generic-container")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		// linkTag := item.Find("a")
		// link, _ := linkTag.Attr("href")
		var title string
		item.Find("span").Each(func(index int, item *goquery.Selection) {
			title = item.Text()
			// linkTag := item.Find("a")
			// link, _ := linkTag.Attr("href")
			// fmt.Printf("Post #%d: %s - %s\n", index, title, link)
		})

		result := models.ResultGoogleSearch{
			ResultRank: rank,
			//ResultURL:   link,
			ResultTitle: title,
			//ResultDesc:  desc,
		}
		results = append(results, result)
		rank++

		// descTag := item.Find("span.st")
		// desc := descTag.Text()
		// title := titleTag.Text()
		// link = strings.Trim(link, " ")

		// if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
		// 	result := models.ResultGoogleSearch{
		// 		ResultRank:  rank,
		// 		ResultURL:   link,
		// 		ResultTitle: title,
		// 		ResultDesc:  desc,
		// 	}
		// 	results = append(results, result)
		// 	rank++
		// }
	}
	return results, err

}
