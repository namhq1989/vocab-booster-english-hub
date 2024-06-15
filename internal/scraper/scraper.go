package scraper

import (
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
)

type Operations interface {
	IsTermValid(ctx *appcontext.AppContext, term string) (bool, []string, error)
}

type Scraper struct {
	cambridgeCollector  *colly.Collector
	dictionaryCollector *colly.Collector
}

func NewScraper() *Scraper {
	return &Scraper{
		cambridgeCollector:  newCollyCollector(cambridgeDictionaryDomainName, 10, 1*time.Second),
		dictionaryCollector: newCollyCollector(dictionaryComDomainName, 10, 1*time.Second),
	}
}

func newCollyCollector(domainName string, parallelism int, delay time.Duration) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domainName),
	)
	c.AllowURLRevisit = true
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  domainName,
		Parallelism: parallelism,
		Delay:       delay,
	})

	return c
}
