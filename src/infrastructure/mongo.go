package infrastructure

import (
	"context"
	"fmt"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmmongo/v2"
	"keyword-generator/src/config"
	"keyword-generator/src/domain/entities"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBMongo struct {
	Client *mongo.Client
}

var AppMongo *DBMongo

func init() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	var dbMongo DBMongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Config.GetString("MONGO_URL")), options.Client().SetMonitor(apmmongo.CommandMonitor()))
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	dbMongo.Client = client
	AppMongo = &dbMongo
}

func (m *DBMongo) MongoCalculate(ctx context.Context) ([]entities.ListBehaviour, error) {
	var listContentID []entities.ListBehaviour

	coll := m.Client.Database(config.Config.GetString("MONGO_DB")).Collection(config.Config.GetString("MONGO_COLLECTION_PROCESS"))
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"status": "processed"}},
		bson.M{"$unwind": "$action"},
		bson.M{"$project": bson.M{
			"_id":         0,
			"userid":      1,
			"contentid":   "$action.contentid",
			"contenttype": "$action.contenttype",
			"duration":    "$action.duration",
		}},
		bson.M{"$sort": bson.M{"datetime": 1}},
		bson.M{"$group": bson.M{
			"_id": "$userid",
			"action": bson.M{
				"$push": bson.M{
					"contentid":   "$contentid",
					"contenttype": "$contenttype",
					"duration":    "$duration",
				},
			},
		}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return listContentID, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		apm.CaptureError(ctx, err).Send()
		return listContentID, err
	}

	// variable array of struct for store result from mongo
	for _, result := range results {

		var data entities.Data

		bytes, err1 := bson.Marshal(result)
		if err1 != nil {
			apm.CaptureError(ctx, err).Send()
			return listContentID, err1
		}

		err = bson.Unmarshal(bytes, &data)
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			return listContentID, err
		}
		var arrayContent []entities.Content

		// looping to get the difference duration in seconds every content
		for _, item := range data.Action {
			defaultDuration, _ := strconv.Atoi(config.Config.GetString("VIEW_DURATION"))
			defaultdurationFloat := float64(defaultDuration)
			durationContent := item.Duration
			if durationContent > defaultdurationFloat {
				arrayContent = append(arrayContent, entities.Content{
					ContentID:  item.ContentID,
					ContetType: item.ContentType,
				})
			} else {
				continue
			}
		}

		listContentID = append(listContentID, entities.ListBehaviour{
			UserID:  result["_id"],
			Content: arrayContent,
		})

	}
	return listContentID, nil
}

func (m *DBMongo) GetAndCreateNewCollection(ctx context.Context, from, to string) (int, error) {
	span, ctx := apm.StartSpan(ctx, "src/infrastructure/mongo.go", "GetAndCreateNewCollection")
	defer span.End()
	var (
		err              error
		movedNumber      int
		cursor           *mongo.Cursor
		docResult        []interface{}
		insertManyResult *mongo.InsertManyResult
	)

	//set up source and destination collections
	source := m.Client.Database(config.Config.GetString("MONGO_DB")).Collection(from)
	dest := m.Client.Database(config.Config.GetString("MONGO_DB")).Collection(to)

	// use Aggregate to retrieve all document from source collection
	cursor, err = source.Aggregate(ctx, bson.A{bson.M{"$match": bson.M{}}})
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		return movedNumber, err
	}

	for cursor.Next(ctx) {
		var doc bson.M
		if err = cursor.Decode(&doc); err != nil {
			return movedNumber, err
		}
		doc["created_date"] = time.Now()
		docResult = append(docResult, doc)
	}

	collArchived := m.Client.Database(config.Config.GetString("MONGO_DB")).Collection(to)
	duration := config.Config.GetInt("MONGO_TTL")
	ttlIndexModel := mongo.IndexModel{
		Keys:    bson.M{"created_date": 1},
		Options: options.Index().SetExpireAfterSeconds(int32(duration)),
	}
	_, errTTL := collArchived.Indexes().CreateOne(context.TODO(), ttlIndexModel)
	if errTTL != nil {
		fmt.Println("Failed to create TTL index", errTTL)
	}
	batchSize := 2000

	for i := 0; i < len(docResult); i += batchSize {
		end := i + batchSize
		if end > len(docResult) {
			end = len(docResult)
		}
		batch := docResult[i:end]

		insertManyResult, err = dest.InsertMany(context.Background(), batch)
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			fmt.Println(err)
		}

		movedNumber = len(insertManyResult.InsertedIDs)
	}

	return movedNumber, err
}

func (m *DBMongo) DeleteCollectionData(ctx context.Context, collection string) (int, error) {
	span, ctx := apm.StartSpan(ctx, "src/infrastructure/mongo.go", "DeleteCollectionData")
	defer span.End()
	var (
		deletedNumber int
		err           error
		result        *mongo.DeleteResult
	)

	//get collection
	col := m.Client.Database(config.Config.GetString("MONGO_DB")).Collection(collection)

	//delte documents matching filter
	result, err = col.DeleteMany(ctx, bson.M{})

	if err == nil {
		deletedNumber = int(result.DeletedCount)
	}

	return deletedNumber, err
}

func (m *DBMongo) UpdateStatus(ctx context.Context, collection, status string) error {
	span, ctx := apm.StartSpan(ctx, "src/infrastructure/mongo.go", "UpdateStatus")
	defer span.End()
	//get collection
	col := m.Client.Database(config.Config.GetString("MONGO_DB")).Collection(collection)

	filter := bson.D{{}}

	// Define the update to set the status to "done"
	update := bson.D{{"$set", bson.D{{"status", status}}}}

	// Update all documents matching the filter
	_, err := col.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
