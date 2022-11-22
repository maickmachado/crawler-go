package main

import (
	"fmt"

	"github.com.br/maickmachado/crawler-go/crawler"
)

func main() {
	//TODO: colocar declaração explicita
	// res, err := crawler.GoogleSearchScrape("submarino", "com", "br", nil, 1, 30, 10)
	// if err == nil {
	// 	for _, res := range res {
	// 		fmt.Println(res)
	// 	}
	// }

	resTeste, err := crawler.GoogleTrendsScrape("bolo de chocolate", nil)
	if err == nil {
		for _, res := range resTeste {
			fmt.Println(res)
		}
	}

	// searchTrend := "lancha de passeio"

	// trends.Teste(searchTrend)

	// reader := bufio.NewReader(os.Stdin)
	// searchTrend, _ := reader.ReadString('\n')

}
