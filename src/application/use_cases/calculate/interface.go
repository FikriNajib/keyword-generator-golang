package calculate

import (
	"context"
	"keyword-generator/src/application/repositories"
	"keyword-generator/src/application/shared"
	"keyword-generator/src/infrastructure"
)

type UseCase interface {
	CalculateAll(ctx context.Context) shared.StatusMessage
}

type useCase struct {
	redis      *infrastructure.RedisDatabase
	mongo      *infrastructure.DBMongo
	elastic    *infrastructure.ElasticSearch
}

func NewUseCase(redis *infrastructure.RedisDatabase, mongo *infrastructure.DBMongo, elastic *infrastructure.ElasticSearch) UseCase {
	return &useCase{
		redis:      redis,
		mongo:      mongo,
		elastic:    elastic,
	}
}
