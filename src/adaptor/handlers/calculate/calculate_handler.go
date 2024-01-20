package calculate

import (
	"keyword-generator/src/adaptor/controllers/calculate"
	"keyword-generator/src/adaptor/repositories/story"
	"keyword-generator/src/adaptor/repositories/videos"
	collector_uc "keyword-generator/src/application/use_cases/calculate"
	"keyword-generator/src/infrastructure"
)

func NewCalculateHandler() *calculate.Controller {
	redisConf := infrastructure.AppRedis
	mongoConf := infrastructure.AppMongo
	elasticConf := infrastructure.AppElasticSearch

	collectorUC := collector_uc.NewUseCase(redisConf, mongoConf, elasticConf)

	return calculate.NewController(collectorUC)
}
