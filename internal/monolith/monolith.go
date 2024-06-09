package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/caching"
	"github.com/namhq1989/vocab-booster-english-hub/internal/config"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monitoring"
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/internal/scraper"
	"github.com/namhq1989/vocab-booster-english-hub/internal/searching"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Searching() *searching.Searching
	Caching() *caching.Caching
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	Monitoring() *monitoring.Monitoring
	Queue() *queue.Queue
	Scraper() *scraper.Scraper
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
