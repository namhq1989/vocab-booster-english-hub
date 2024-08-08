package grpc

import (
	"context"

	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	apperrors "github.com/namhq1989/vocab-booster-english-hub/internal/utils/error"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/vocabulary/application"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
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

func (s server) BookmarkVocabulary(bgCtx context.Context, req *vocabularypb.BookmarkVocabularyRequest) (*vocabularypb.BookmarkVocabularyResponse, error) {
	resp, err := s.app.BookmarkVocabulary(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserBookmarkedVocabularies(bgCtx context.Context, req *vocabularypb.GetUserBookmarkedVocabulariesRequest) (*vocabularypb.GetUserBookmarkedVocabulariesResponse, error) {
	resp, err := s.app.GetUserBookmarkedVocabularies(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetWordOfTheDay(bgCtx context.Context, req *vocabularypb.GetWordOfTheDayRequest) (*vocabularypb.GetWordOfTheDayResponse, error) {
	resp, err := s.app.GetWordOfTheDay(appcontext.NewGRPC(bgCtx), req)
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

func (s server) GetCommunitySentences(bgCtx context.Context, req *vocabularypb.GetCommunitySentencesRequest) (*vocabularypb.GetCommunitySentencesResponse, error) {
	resp, err := s.app.GetCommunitySentences(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetCommunitySentence(bgCtx context.Context, req *vocabularypb.GetCommunitySentenceRequest) (*vocabularypb.GetCommunitySentenceResponse, error) {
	resp, err := s.app.GetCommunitySentence(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}

func (s server) GetUserCommunitySentenceDrafts(bgCtx context.Context, req *vocabularypb.GetUserCommunitySentenceDraftsRequest) (*vocabularypb.GetUserCommunitySentenceDraftsResponse, error) {
	resp, err := s.app.GetUserCommunitySentenceDrafts(appcontext.NewGRPC(bgCtx), req)
	if err != nil {
		return nil, apperrors.ToGrpcError(bgCtx, err)
	}
	return resp, nil
}
