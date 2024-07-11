package tts

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (t TTS) GenerateVocabularyExampleSound(ctx *appcontext.AppContext, exampleID, exampleContent string) (string, error) {
	var (
		fileName = fmt.Sprintf("%s.ogg", exampleID)
		voice    = t.randomVoice()
	)

	output, err := t.polly.SynthesizeSpeech(ctx.Context(), &polly.SynthesizeSpeechInput{
		OutputFormat: types.OutputFormatOggVorbis,
		Text:         aws.String(exampleContent),
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

	file, err := os.Create(t.generateExampleFilePath(fileName))
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
