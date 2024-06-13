package infrastructure

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/scraper"
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
