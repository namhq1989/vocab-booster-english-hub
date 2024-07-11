package config

import "errors"

type (
	Server struct {
		RestPort string
		GRPCPort string

		AppName      string
		Environment  string
		IsEnvRelease bool
		Debug        bool

		// Postgres
		PostgresConn string

		// Redis
		CachingRedisURL string

		// Queue
		QueueRedisURL    string
		QueueUsername    string
		QueuePassword    string
		QueueConcurrency int

		// Sentry
		SentryDSN     string
		SentryMachine string

		// OpenAI
		OpenAIAPIKey string

		// AWS
		AWSAccessKey string
		AWSSecretKey string
		AWSRegion    string

		// NLP
		EndpointNLP string
		EndpointTTS string
	}
)

func Init() Server {
	cfg := Server{
		RestPort: ":3005",
		GRPCPort: ":3006",

		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),
		Debug:       getEnvBool("DEBUG"),

		PostgresConn: getEnvStr("POSTGRES_CONN"),

		CachingRedisURL: getEnvStr("CACHING_REDIS_URL"),

		QueueRedisURL:    getEnvStr("QUEUE_REDIS_URL"),
		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),

		SentryDSN:     getEnvStr("SENTRY_ENGLISH_HUB_DSN"),
		SentryMachine: getEnvStr("SENTRY_MACHINE"),

		OpenAIAPIKey: getEnvStr("OPENAI_API_KEY"),

		AWSAccessKey: getEnvStr("AWS_ACCESS_KEY_ID"),
		AWSSecretKey: getEnvStr("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:    getEnvStr("AWS_REGION"),

		EndpointNLP: getEnvStr("ENDPOINT_NLP"),
		EndpointTTS: getEnvStr("ENDPOINT_TTS"),
	}
	cfg.IsEnvRelease = cfg.Environment == "release"

	// validation
	if cfg.Environment == "" {
		panic(errors.New("missing ENVIRONMENT"))
	}

	if cfg.PostgresConn == "" {
		panic(errors.New("missing POSTGRES_CONN"))
	}

	if cfg.CachingRedisURL == "" {
		panic(errors.New("missing CACHING_REDIS_URL"))
	}

	if cfg.QueueRedisURL == "" {
		panic(errors.New("missing QUEUE_REDIS_URL"))
	}

	if cfg.OpenAIAPIKey == "" {
		panic(errors.New("missing OPENAI_API_KEY"))
	}

	if cfg.AWSAccessKey == "" {
		panic(errors.New("missing AWS_ACCESS_KEY"))
	}
	if cfg.AWSSecretKey == "" {
		panic(errors.New("missing AWS_SECRET_KEY"))
	}
	if cfg.AWSRegion == "" {
		panic(errors.New("missing AWS_REGION"))
	}

	if cfg.EndpointNLP == "" {
		panic(errors.New("missing ENDPOINT_NLP"))
	}
	if cfg.EndpointTTS == "" {
		panic(errors.New("missing ENDPOINT_TTS"))
	}

	return cfg
}
