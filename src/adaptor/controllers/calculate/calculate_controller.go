package calculate

import (
	"context"
	"fmt"
	"go.elastic.co/apm/v2"
	movecontroller "keyword-generator/src/adaptor/controllers/move_controller"
	"keyword-generator/src/application/shared"
	"keyword-generator/src/application/use_cases/calculate"
	"keyword-generator/src/application/use_cases/move"
	"keyword-generator/src/infrastructure"
)

type Controller struct {
	CalculateUsecase calculate.UseCase
}

func NewController(CalculateUsecase calculate.UseCase) *Controller {
	return &Controller{CalculateUsecase: CalculateUsecase}
}

var moveHandler = movecontroller.NewMoveDataController(move.NewMoveUsecase(infrastructure.AppMongo))

func (controller *Controller) CalculateData(ctx context.Context) shared.StatusMessage {
	span, ctx := apm.StartSpan(ctx, "src/adaptor/controllers/calculate", "CalculateData")
	defer span.End()
	statusMessage := controller.CalculateUsecase.CalculateAll(ctx)
	if statusMessage.GetCode() > 0 {
		return statusMessage
	}
	moveHandler.MoveDataToArchived(ctx)
	fmt.Println("Keyword Generator Finished ==============>")
	return statusMessage
}
