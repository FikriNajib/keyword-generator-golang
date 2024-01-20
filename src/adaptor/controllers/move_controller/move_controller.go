package movecontroller

import (
	"context"
	"fmt"
	"go.elastic.co/apm"
	"keyword-generator/src/application/use_cases/move"
	"keyword-generator/src/config"
	"log"
)

type MoveDataController struct {
	uc move.MoveUsecase
}

func NewMoveDataController(uc move.MoveUsecase) *MoveDataController {
	return &MoveDataController{uc: uc}
}

func (mc *MoveDataController) MoveDataAtaPeriodeTime(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "src/adaptor/controllers/move_controller", "MoveDataAtaPeriodeTime")
	defer span.End()
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	statusMessage := mc.uc.MoveMainData(context.TODO(), config.Config.GetString("MONGO_COLLECTION"), config.Config.GetString("MONGO_COLLECTION_PROCESS"))
	m := fmt.Sprintf("MoveMainData Collection from: %s => to : %s", config.Config.GetString("MONGO_COLLECTION"), config.Config.GetString("MONGO_COLLECTION_PROCESS"))
	fmt.Println(m)
	fmt.Println(statusMessage)
	//logger.WithFields(logger.Fields{"component": "move-controller", "action": "move main data usecase"}).
	//	Infof("move data info. %+v", statusMessage)
}

func (mc *MoveDataController) MoveDataToArchived(ctx context.Context) {
	span, ctx := apm.StartSpan(ctx, "src/adaptor/controllers/move_controller", "MoveDataToArchived")
	defer span.End()
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	statusMessage := mc.uc.MoveProcessedData(context.TODO(), config.Config.GetString("MONGO_COLLECTION_PROCESS"), config.Config.GetString("MONGO_COLLECTION_ARCHIVED"))
	m := fmt.Sprintf("MoveProcessedData Collection from: %s => to : %s", config.Config.GetString("MONGO_COLLECTION_PROCESS"), config.Config.GetString("MONGO_COLLECTION_ARCHIVED"))
	fmt.Println(m)
	fmt.Println(statusMessage)
	//logger.WithFields(logger.Fields{"component": "move-controller", "action": "move processed data usecase"}).
	//	Infof("move data info. %+v", statusMessage)
}
