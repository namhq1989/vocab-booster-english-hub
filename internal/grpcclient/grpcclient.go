package grpcclient

import (
	"github.com/namhq1989/vocab-booster-english-hub/core/appcontext"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-english-hub/internal/genproto/vocabularypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewVocabularyClient(_ *appcontext.AppContext, addr string) (vocabularypb.VocabularyServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return vocabularypb.NewVocabularyServiceClient(conn), nil
}

func NewExerciseClient(_ *appcontext.AppContext, addr string) (exercisepb.ExerciseServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return exercisepb.NewExerciseServiceClient(conn), nil
}
