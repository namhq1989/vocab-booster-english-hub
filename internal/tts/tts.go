package tts

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
)

type Operations interface {
	GenerateVocabularyPronunciationSound(ctx *appcontext.AppContext, vocabulary string) (string, error)
}

type TTS struct {
	polly  *polly.Client
	voices []types.Voice
}

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

	t := &TTS{
		polly:  svc,
		voices: resp.Voices,
	}
	t.initDirectories()

	return t
}

func (t TTS) randomVoice() types.Voice {
	rand := manipulation.RandomIntInRange(0, len(t.voices)-1)
	return t.voices[rand]
}

func (t TTS) initDirectories() {
	if err := os.MkdirAll(t.getPronunciationsPath(), 0755); err != nil {
		panic(fmt.Errorf("failed to create pronunciation files directory %s: %s", t.getPronunciationsPath(), err.Error()))
	}
}

func (TTS) getPronunciationsPath() string {
	dir, _ := os.Getwd()
	return path.Join(dir, "files/pronunciations")
}

func (t TTS) generatePronunciationFilePath(fileName string) string {
	return fmt.Sprintf("%s/%s", t.getPronunciationsPath(), fileName)
}
