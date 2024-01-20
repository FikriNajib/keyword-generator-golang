package move

import (
	"context"
	"keyword-generator/src/application/shared"
	"keyword-generator/src/infrastructure"
)

type MoveUsecase interface {
	MoveMainData(ctx context.Context, from, to string) shared.StatusMessage
	MoveProcessedData(ctx context.Context, from, to string) shared.StatusMessage
}

type moveUsecase struct {
	mongo *infrastructure.DBMongo
}

func NewMoveUsecase(mongo *infrastructure.DBMongo) MoveUsecase {
	return &moveUsecase{
		mongo: mongo,
	}
}
