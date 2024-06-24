package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/ai"
	"github.com/namhq1989/vocab-booster-english-hub/internal/caching"
	"github.com/namhq1989/vocab-booster-english-hub/internal/config"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monitoring"
	"github.com/namhq1989/vocab-booster-english-hub/internal/nlp"
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/internal/scraper"
	"github.com/namhq1989/vocab-booster-english-hub/internal/tts"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Caching() *caching.Caching
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	Monitoring() *monitoring.Monitoring
	Queue() *queue.Queue
	Scraper() *scraper.Scraper
	TTS() *tts.TTS
	AI() *ai.AI
	NLP() *nlp.NLP
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
