package domain

import "github.com/namhq1989/vocab-booster-utilities/appcontext"

type ScraperRepository interface {
	IsTermValid(ctx *appcontext.AppContext, term string) (bool, []string, error)
}
