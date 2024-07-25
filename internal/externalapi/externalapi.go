package externalapi

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Operations interface {
	SearchTermWitDatamuse(ctx *appcontext.AppContext, term string) (*DatamuseSearchTermResult, error)
}

type ExternalAPI struct {
	datamuse *resty.Client
}

const (
	datamuseApiEndpoint = "https://api.datamuse.com"
)

func NewExternalAPIClient() *ExternalAPI {
	return &ExternalAPI{
		datamuse: resty.New().
			SetBaseURL(datamuseApiEndpoint).
			SetHeader("Accept", "application/json").
			SetTimeout(30 * time.Second).
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
				return 1, fmt.Errorf("failed to send Datamuse request at %s with status code %d", resp.Request.RawRequest.RequestURI, resp.StatusCode())
			}),
	}
}
