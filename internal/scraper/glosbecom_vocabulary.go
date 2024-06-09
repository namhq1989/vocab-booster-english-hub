package scraper

import "github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

const (
	glosbeVocabularyURL        = "https://glosbe.com/en/vi/%s"
	glosbeVocabularyDomainName = "glosbe.com"
)

type GlosbeVocabularyResult struct {
}

func (Scraper) ScrapeGlosbeVocabulary(ctx *appcontext.AppContext, term string) (*GlosbeVocabularyResult, error) {
	ctx.Logger().Info("new Glosbe vocabulary scraping request", appcontext.Fields{"term": term})

	// TODO

	return nil, nil
}
