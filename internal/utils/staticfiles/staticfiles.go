package staticfiles

import "fmt"

type Endpoint struct {
	tts string
}

var endpoint = Endpoint{}

func Init(ttsEndpoint string) {
	endpoint.tts = ttsEndpoint
}

func GetVocabularyEndpoint(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint.tts, "audio/vocabulary", fileName)
}

func GetExampleEndpoint(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint.tts, "audio/example", fileName)
}
