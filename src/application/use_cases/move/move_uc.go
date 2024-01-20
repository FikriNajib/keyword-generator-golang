package move

import (
	"context"
	"go.elastic.co/apm"
	"keyword-generator/src/application/shared"
	"log"
)

func (m *moveUsecase) MoveMainData(ctx context.Context, from, to string) shared.StatusMessage {
	span, ctx := apm.StartSpan(ctx, "src/application/use_cases/move/move_uc.go", "MoveMainData")
	defer span.End()
	if from == "" || to == "" {
		return shared.INVALID_INPUT
	}

	many, err := m.mongo.GetAndCreateNewCollection(ctx, from, to)
	if err != nil {
		log.Println(err)
		return shared.FAILED_COPY_DATA
	}
	status := "processed"
	err = m.mongo.UpdateStatus(ctx, to, status)
	if err != nil {
		log.Println(err)
		return shared.FAILED_COPY_DATA
	}
	log.Println("get and create new collection success copy data", many)

	many, err = m.mongo.DeleteCollectionData(ctx, from)
	if err != nil {
		log.Println(err)
		return shared.FAILED_DELETE_DATA
	}

	log.Println("get and create new collection success copy data", many)
	return shared.SUCCESS

}

func (m *moveUsecase) MoveProcessedData(ctx context.Context, from, to string) shared.StatusMessage {
	if from == "" || to == "" {
		return shared.INVALID_INPUT
	}

	many, err := m.mongo.GetAndCreateNewCollection(ctx, from, to)
	if err != nil {
		log.Println(err)
		return shared.FAILED_COPY_DATA
	}
	status := "done"
	err = m.mongo.UpdateStatus(ctx, to, status)
	if err != nil {
		log.Println(err)
		return shared.FAILED_COPY_DATA
	}
	log.Println("get and create new collection archived ,success copy data", many)

	many, err = m.mongo.DeleteCollectionData(ctx, from)
	if err != nil {
		log.Println(err)
		return shared.FAILED_DELETE_DATA
	}
	log.Println("delete collection processed")

	return shared.SUCCESS

}
