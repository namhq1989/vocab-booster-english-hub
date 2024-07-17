package staticfiles

import "fmt"

type Endpoint struct {
	tts    string
	images string
}

var endpoint = Endpoint{}

func Init(ttsEndpoint, imagesEndpoint string) {
	endpoint.tts = ttsEndpoint
	endpoint.images = imagesEndpoint
}

func GetVocabularyEndpoint(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint.tts, "audio/vocabulary", fileName)
}

func GetExampleEndpoint(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint.tts, "audio/example", fileName)
}

func GetExerciseCollectionsEndpoint(fileName string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint.images, "image", fileName)
}
