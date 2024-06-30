package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/core/error"
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
	resp, err := s.app.SearchVocabulary(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) CreateCollection(bgCtx context.Context, req *vocabularypb.CreateCollectionRequest) (*vocabularypb.CreateCollectionResponse, error) {
	resp, err := s.app.CreateCollection(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) UpdateCollection(bgCtx context.Context, req *vocabularypb.UpdateCollectionRequest) (*vocabularypb.UpdateCollectionResponse, error) {
	resp, err := s.app.UpdateCollection(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) AddVocabularyToCollection(bgCtx context.Context, req *vocabularypb.AddVocabularyToCollectionRequest) (*vocabularypb.AddVocabularyToCollectionResponse, error) {
	resp, err := s.app.AddVocabularyToCollection(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) RemoveVocabularyFromCollection(bgCtx context.Context, req *vocabularypb.RemoveVocabularyFromCollectionRequest) (*vocabularypb.RemoveVocabularyFromCollectionResponse, error) {
	resp, err := s.app.RemoveVocabularyFromCollection(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) CreateCommunitySentenceDraft(bgCtx context.Context, req *vocabularypb.CreateCommunitySentenceDraftRequest) (*vocabularypb.CreateCommunitySentenceDraftResponse, error) {
	resp, err := s.app.CreateCommunitySentenceDraft(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) UpdateCommunitySentenceDraft(bgCtx context.Context, req *vocabularypb.UpdateCommunitySentenceDraftRequest) (*vocabularypb.UpdateCommunitySentenceDraftResponse, error) {
	resp, err := s.app.UpdateCommunitySentenceDraft(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) PromoteCommunitySentenceDraft(bgCtx context.Context, req *vocabularypb.PromoteCommunitySentenceDraftRequest) (*vocabularypb.PromoteCommunitySentenceDraftResponse, error) {
	resp, err := s.app.PromoteCommunitySentenceDraft(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) LikeCommunitySentence(bgCtx context.Context, req *vocabularypb.LikeCommunitySentenceRequest) (*vocabularypb.LikeCommunitySentenceResponse, error) {
	resp, err := s.app.LikeCommunitySentence(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}
