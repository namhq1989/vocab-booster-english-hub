package domain

import "github.com/namhq1989/vocab-booster-english-hub/core/appcontext"

type ScraperRepository interface {
	IsTermValid(ctx *appcontext.AppContext, term string) (bool, []string, error)
}
