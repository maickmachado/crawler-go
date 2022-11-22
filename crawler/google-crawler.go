package crawler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com.br/maickmachado/crawler-go/models"
	"github.com/PuerkitoBio/goquery"
)

func GoogleSearchScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int) ([]models.ResultGoogleSearch, error) {

	results := []models.ResultGoogleSearch{}

	resultCounter := 0

	googlePageURL, err := BuildGoogleSearchURl(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}

	for _, pageURL := range googlePageURL {
		res, err := models.ScrapeClientRequest(pageURL, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := GoogleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}

	return results, nil
}

//TODO: juntar esse com a do trends
//TODO: retirar o pais e colocar somente brasil
func BuildGoogleSearchURl(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := models.GoogleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}

func GoogleResultParsing(response *http.Response, rank int) ([]models.ResultGoogleSearch, error) {
	// r := response.Body
	// t, e := goquery.NewDocumentFromReader(r)
	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		return nil, err
	}

	results := []models.ResultGoogleSearch{}
	sel := doc.Find("div.g")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h3.LC20lb.MBeuO.DKV0Md")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := models.ResultGoogleSearch{
				ResultRank:  rank,
				ResultURL:   link,
				ResultTitle: title,
				ResultDesc:  desc,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, err

}
