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

func (s server) CreateCollection(bgCtx context.Context, req *vocabularypb.CreateCollectionRequest) (*vocabularypb.CreateCollectionResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.CreateCollection(ctx, req)
}

func (s server) UpdateCollection(bgCtx context.Context, req *vocabularypb.UpdateCollectionRequest) (*vocabularypb.UpdateCollectionResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.UpdateCollection(ctx, req)
}

func (s server) AddVocabularyToCollection(bgCtx context.Context, req *vocabularypb.AddVocabularyToCollectionRequest) (*vocabularypb.AddVocabularyToCollectionResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.AddVocabularyToCollection(ctx, req)
}

func (s server) RemoveVocabularyFromCollection(bgCtx context.Context, req *vocabularypb.RemoveVocabularyFromCollectionRequest) (*vocabularypb.RemoveVocabularyFromCollectionResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.RemoveVocabularyFromCollection(ctx, req)
}

func (s server) CreateCommunitySentenceDraft(bgCtx context.Context, req *vocabularypb.CreateCommunitySentenceDraftRequest) (*vocabularypb.CreateCommunitySentenceDraftResponse, error) {
	ctx := appcontext.NewGRPC(bgCtx)
	return s.app.CreateCommunitySentenceDraft(ctx, req)
}
