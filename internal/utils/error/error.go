package apperrors

import (
	"context"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	langEn = "en"
	langVi = "vi"
)

func toLang(lang string) string {
	if lang == langVi {
		return langVi
	}

	return langEn
}

var localizers = map[string]*i18n.Localizer{
	langEn: nil,
	langVi: nil,
}

func Init() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	dir, _ := os.Getwd()
	path := filepath.Join(dir, "internal", "utils", "error", "i18n")
	bundle.MustLoadMessageFile(filepath.Join(path, "active.en.toml"))
	bundle.MustLoadMessageFile(filepath.Join(path, "active.vi.toml"))

	localizers[langEn] = i18n.NewLocalizer(bundle, language.English.String(), langEn)
	localizers[langVi] = i18n.NewLocalizer(bundle, language.Vietnamese.String(), langVi)
}

func GetMessage(lang string, err error) (code, msg string) {
	key := err.Error()

	if grpcErr, ok := status.FromError(err); ok {
		key = grpcErr.Message()
	}

	if localizers[lang] == nil {
		return "_", key
	}

	msg, localizeErr := localizers[lang].Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if localizeErr != nil {
		return "-", key
	}

	return key, msg
}

func ToGrpcError(ctx context.Context, err error) error {
	lang := getGrpcLangInContext(ctx)
	code, msg := GetMessage(lang, err)
	return status.Errorf(status.Code(err), "%s | %s", code, msg)
}

func getGrpcLangInContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return langEn
	}

	languages := md["lang"]
	if len(languages) == 0 {
		return langEn
	}
	return toLang(languages[0])
}
