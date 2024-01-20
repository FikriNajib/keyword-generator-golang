package wathcer

import (
	"context"
	"go.elastic.co/apm"
	movecontroller "keyword-generator/src/adaptor/controllers/move_controller"
	calculate3 "keyword-generator/src/adaptor/handlers/calculate"
	"keyword-generator/src/application/use_cases/move"
	"keyword-generator/src/config"
	"keyword-generator/src/infrastructure"
	"log"
)

func RunMoveWatcher() {
	tx := apm.DefaultTracer.StartTransaction("src/infrastructure/watcher/move_watcher.go", "RunMoveWatcher")

	ctx := apm.ContextWithTransaction(context.Background(), tx)
	span, ctx := apm.StartSpan(ctx, "src/infrastructure/watcher/move_watcher", "RunMoveWatcher")
	defer span.End()

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	moveHandler := movecontroller.NewMoveDataController(move.NewMoveUsecase(infrastructure.AppMongo))
	calculateHandler := calculate3.NewCalculateHandler()

	moveHandler.MoveDataAtaPeriodeTime(ctx)
	calculateHandler.CalculateData(ctx)
	tx.End()

}
