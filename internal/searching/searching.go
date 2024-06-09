package searching

import (
	"errors"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

type Searching struct {
	meilisearch *meilisearch.Client
}

func NewSearchingClient(host, apiKey string) *Searching {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   host,
		APIKey: apiKey,
	})

	if !client.IsHealthy() {
		panic(errors.New("meilisearch is not healthy"))
	}

	fmt.Printf("⚡️ [searching]: meilisearch connected \n")

	return &Searching{meilisearch: client}
}
