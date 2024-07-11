package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/internal/scraper"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ScraperRepository struct {
	scraper scraper.Operations
}

func NewScraperRepository(scraper scraper.Operations) ScraperRepository {
	return ScraperRepository{
		scraper: scraper,
	}
}

func (r ScraperRepository) IsTermValid(ctx *appcontext.AppContext, term string) (bool, []string, error) {
	return r.scraper.IsTermValid(ctx, term)
}
