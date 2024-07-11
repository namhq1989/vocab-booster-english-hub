package main

import (
	"crypto/subtle"
	"fmt"

	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/staticfiles"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
	"github.com/namhq1989/vocab-booster-english-hub/core/logger"
	"github.com/namhq1989/vocab-booster-english-hub/internal/ai"
	"github.com/namhq1989/vocab-booster-english-hub/internal/caching"
	"github.com/namhq1989/vocab-booster-english-hub/internal/config"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monitoring"
	"github.com/namhq1989/vocab-booster-english-hub/internal/monolith"
	"github.com/namhq1989/vocab-booster-english-hub/internal/nlp"
	"github.com/namhq1989/vocab-booster-english-hub/internal/queue"
	"github.com/namhq1989/vocab-booster-english-hub/internal/scraper"
	"github.com/namhq1989/vocab-booster-english-hub/internal/tts"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/waiter"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary"
)

func main() {
	var err error

	// config
	cfg := config.Init()

	// logger
	logger.Init(cfg.Environment)

	// app error
	apperrors.Init()

	// static files
	staticfiles.Init(cfg.EndpointTTS)

	// server
	a := app{}
	a.cfg = cfg

	// rest
	a.rest = initRest(cfg)

	// grpc
	a.rpc = initRPC()

	// database
	a.database = database.NewDatabaseClient(cfg.PostgresConn)

	// queue
	a.queue = queue.Init(cfg.QueueRedisURL, cfg.QueueConcurrency)

	// init queue's dashboard
	a.rest.Any(fmt.Sprintf("%s/*", queue.DashboardPath), echo.WrapHandler(queue.EnableDashboard(cfg.QueueRedisURL)), middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if !cfg.IsEnvRelease {
			return true, nil
		}
		return subtle.ConstantTimeCompare([]byte(username), []byte(cfg.QueueUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.QueuePassword)) == 1, nil
	}))

	// caching
	a.caching = caching.NewCachingClient(cfg.CachingRedisURL)

	// monitoring
	a.monitoring = monitoring.Init(a.Rest(), cfg.SentryDSN, cfg.SentryMachine, cfg.Environment)

	// scraper
	a.scraper = scraper.NewScraper()

	// tts
	a.tts = tts.NewTTSClient(cfg.AWSAccessKey, cfg.AWSSecretKey, cfg.AWSRegion)

	// ai
	a.ai = ai.NewAIClient(cfg.OpenAIAPIKey)

	// nlp
	a.nlp = nlp.NewNLPClient(cfg.EndpointNLP)

	// waiter
	a.waiter = waiter.New(waiter.CatchSignals())

	// modules
	a.modules = []monolith.Module{
		&vocabulary.Module{},
		&exercise.Module{},
	}

	// start
	if err = a.startupModules(); err != nil {
		panic(err)
	}

	fmt.Println("--- started vocab-booster-english-hub application")
	defer fmt.Println("--- stopped vocab-booster-english-hub application")

	// wait for other service starts
	a.waiter.Add(
		a.waitForRest,
		a.waitForRPC,
	)
	if err = a.waiter.Wait(); err != nil {
		panic(err)
	}
}
