package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"kr":  "https://www.google.co.kr/search?q=",
}

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&start=%d&hl=%s&filter=0", googleBase, searchTerm, count, start, languageCode)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		return nil, fmt.Errorf("country (%s) is currently not supported", countryCode)
	}

	return toScrape, nil
}

func getScraperClient(proxyString any) *http.Client {
	client := &http.Client{}
	switch v := proxyString.(type) {
	case string:
		proxyURL, err := url.Parse(v)
		if err != nil {
			return client
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	default:
		return client
	}
	return client
}

func scrapeClientRequest(searchURL string, proxyString any) (*http.Response, error) {
	baseClient := getScraperClient(proxyString)
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", randomUserAgent())

	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("scraper received a non 200 status code: %v", err)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func googleResultParsing(response *http.Response, rank int) ([]*SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	results := []*SearchResult{}
	sel := doc.Find("div.g")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)

		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")

		titleTag := item.Find("h3")
		title := titleTag.Text()

		descTag := item.Find("span")
		desc := descTag.Text()

		link = strings.Trim(link, " ")

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := &SearchResult{
				ResultRank:  rank,
				ResultURL:   link,
				ResultTitle: title,
				ResultDesc:  desc,
			}

			results = append(results, result)
			rank++
		}
	}
	return results, nil
}

func GoogleScrape(searchTerm, countryCode, languageCode string, pages, count, backoff int, proxyString any) ([]*SearchResult, error) {
	results := []*SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}

	for _, page := range googlePages {
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}

		data, err := googleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}

		resultCounter += len(data)
		results = append(results, data...)

		time.Sleep(time.Duration(backoff) * time.Second)
	}

	return results, nil
}

func main() {
	res, err := GoogleScrape("붱철", "kr", "kr", 1, 30, 5, nil)
	if err != nil {
		log.Fatalf("Cannot scrape: %v\n", err)
	}

	fmt.Println("--------------------------------------------------")
	for _, r := range res {
		fmt.Printf("Rank: %d\n", r.ResultRank)
		fmt.Printf("Title: %s\n", r.ResultTitle)
		fmt.Printf("Description: %s\n", r.ResultDesc)
		fmt.Printf("URL: %s\n", r.ResultURL)
		fmt.Println("--------------------------------------------------")
	}
}
