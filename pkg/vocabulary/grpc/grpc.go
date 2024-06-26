package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	vocabularypb.UnimplementedVocabularyServiceServer
}

var _ vocabularypb.VocabularyServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, app application.App) error {
	vocabularypb.RegisterVocabularyServiceServer(registrar, server{app: app})
	return nil
}

func (s server) SearchVocabulary(bgCtx context.Context, req *vocabularypb.SearchVocabularyRequest) (*vocabularypb.SearchVocabularyResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.SearchVocabulary(ctx, req)
}

func (s server) CreateUserVocabularyCollection(bgCtx context.Context, req *vocabularypb.CreateUserVocabularyCollectionRequest) (*vocabularypb.CreateUserVocabularyCollectionResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.CreateUserVocabularyCollection(ctx, req)
}

func (s server) UpdateUserVocabularyCollection(bgCtx context.Context, req *vocabularypb.UpdateUserVocabularyCollectionRequest) (*vocabularypb.UpdateUserVocabularyCollectionResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.UpdateUserVocabularyCollection(ctx, req)
}
