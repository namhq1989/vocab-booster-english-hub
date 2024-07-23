package nlp

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"github.com/namhq1989/vocab-booster-utilities/language"
)

type Operations interface {
	AnalyzeSentence(_ *appcontext.AppContext, sentence string) (*SentenceAnalysisResult, error)
	Translate(_ *appcontext.AppContext, term string) (result *language.Multilingual, err error)
	EvaluateSentence(_ *appcontext.AppContext, sentence, tense string, vocabulary []string) (result *EvaluateSentenceResult, err error)
	GrammarCheck(_ *appcontext.AppContext, sentence string) (result *GrammarCheckResult, err error)
}

type NLP struct {
	httpClient *resty.Client
}

func NewNLPClient(endpoint string) *NLP {
	return &NLP{
		httpClient: resty.New().
			SetBaseURL(endpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send NLP request at %s with status code %d", endpoint, resp.StatusCode())
			}),
	}
}
