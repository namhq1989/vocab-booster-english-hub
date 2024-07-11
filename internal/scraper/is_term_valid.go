package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gocolly/colly/v2"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s Scraper) IsTermValid(ctx *appcontext.AppContext, term string) (bool, []string, error) {
	var (
		suggestions     = make([]string, 0)
		vocabularyFound = false
		scrapeError     = false
		scrapeURL       = fmt.Sprintf(dictionaryComBrowseURL, term)
	)

	var retryTimes = 1

	for {
		ctx.Logger().Info("new is term valid scraping request", appcontext.Fields{"term": term, "url": scrapeURL})

		var collector = s.dictionaryCollector.Clone()

		if retryTimes > 3 {
			break
		}

		collector.OnHTML("section[data-type='dictionary-headword-module']", func(e *colly.HTMLElement) {
			vocabularyFound = true
		})

		collector.OnHTML("section[data-type='spell-suggestions-module']", func(e *colly.HTMLElement) {
			e.ForEach("h2", func(_ int, el *colly.HTMLElement) {
				fmt.Println("Found h2 content:", el.Text)
			})
		})

		collector.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", gofakeit.UserAgent())
			ctx.Logger().Info("visiting url ...", appcontext.Fields{"url": r.URL.String()})
		})

		collector.OnScraped(func(r *colly.Response) {
			ctx.Logger().Text("url scraped successfully")
		})

		collector.OnError(func(r *colly.Response, _ error) {
			if r.StatusCode >= 300 {
				vocabularyFound = false
				scrapeError = true
			}

			// parse HTML of response body
			ctx.Logger().Text("page not found, parse the response body for spell suggestions")
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
			if err != nil {
				ctx.Logger().Error("failed to parse response body", err, appcontext.Fields{"body": string(r.Body)})
			} else {
				doc.Find("section[data-type='spell-suggestions-module']").Each(func(i int, s *goquery.Selection) {
					suggestions = append(suggestions, s.Find("h2 a").Text())
				})
			}
		})

		if err := collector.Visit(scrapeURL); err != nil {
			ctx.Logger().Error("failed to visit url", err, appcontext.Fields{"url": scrapeURL})
			retryTimes++
			time.Sleep(1 * time.Second)
		} else if vocabularyFound {
			return vocabularyFound, suggestions, nil
		}

		if scrapeError {
			break
		}
	}

	return vocabularyFound, suggestions, nil
}
