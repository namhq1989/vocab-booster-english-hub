package tts

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Operations interface {
	GenerateVocabularySound(ctx *appcontext.AppContext, vocabulary string) (string, error)
	GenerateVocabularyExampleSound(ctx *appcontext.AppContext, exampleID, exampleContent string) (string, error)
}

type TTS struct {
	polly  *polly.Client
	voices []types.Voice
}

var pickedVoices = []string{"Joey", "Kendra", "Salli", "Ruth", "Stephen", "Gregory"}

func NewTTSClient(awsAccessKey, awsSecretKey, awsRegion string) *TTS {
	var (
		ctx = context.Background()
	)

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")),
	)
	if err != nil {
		panic(err)
	}
	svc := polly.NewFromConfig(cfg)

	fmt.Printf("⚡️ [tts]: polly connected \n")

	input := &polly.DescribeVoicesInput{LanguageCode: types.LanguageCodeEnUs, Engine: types.EngineNeural}
	resp, err := svc.DescribeVoices(ctx, input)
	if err != nil {
		panic(err)
	}
	if len(resp.Voices) == 0 {
		panic(errors.New("no available voices of polly"))
	}

	voices := make([]types.Voice, 0)

	for _, voice := range resp.Voices {
		index := slices.IndexFunc(pickedVoices, func(v string) bool {
			return v == *voice.Name
		})
		if index != -1 {
			voices = append(voices, voice)
		}
	}

	t := &TTS{
		polly:  svc,
		voices: voices,
	}
	t.initDirectories()

	return t
}

func (t TTS) randomVoice() types.Voice {
	rand := manipulation.RandomIntInRange(0, len(t.voices)-1)
	return t.voices[rand]
}

func (t TTS) initDirectories() {
	if err := os.MkdirAll(t.getVocabularyPath(), 0755); err != nil {
		panic(fmt.Errorf("failed to create vocabulary files directory %s: %s", t.getVocabularyPath(), err.Error()))
	}

	if err := os.MkdirAll(t.getExamplePath(), 0755); err != nil {
		panic(fmt.Errorf("failed to create example files directory %s: %s", t.getExamplePath(), err.Error()))
	}
}

func (TTS) getVocabularyPath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, "files/tts/vocabulary")
}

func (TTS) getExamplePath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, "files/tts/example")
}

func (t TTS) generateVocabularyFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", t.getVocabularyPath(), fileName)
}

func (t TTS) generateExampleFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", t.getExamplePath(), fileName)
}
