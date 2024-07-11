package tts

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (t TTS) GenerateVocabularySound(ctx *appcontext.AppContext, vocabulary string) (string, error) {
	var (
		slug     = manipulation.Slugify(vocabulary)
		fileName = fmt.Sprintf("%s.ogg", slug)
		voice    = t.randomVoice()
	)

	output, err := t.polly.SynthesizeSpeech(ctx.Context(), &polly.SynthesizeSpeechInput{
		OutputFormat: types.OutputFormatOggVorbis,
		Text:         aws.String(vocabulary),
		VoiceId:      voice.Id,
		Engine:       types.EngineNeural,
		LanguageCode: types.LanguageCodeEnUs,
		TextType:     types.TextTypeText,
	})
	if err != nil {
		ctx.Logger().Error("failed to synthesize speech from Polly", err, appcontext.Fields{})
		return "", err
	}
	defer func() { _ = output.AudioStream.Close() }()

	file, err := os.Create(t.generateVocabularyFilePath(fileName))
	if err != nil {
		ctx.Logger().Error("failed to create file from Polly response", err, appcontext.Fields{})
		return "", err
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, output.AudioStream)
	if err != nil {
		ctx.Logger().Error("failed to write file from Polly response", err, appcontext.Fields{})
		return "", err
	}

	return fileName, nil
}
