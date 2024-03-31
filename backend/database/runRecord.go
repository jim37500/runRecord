package database

import (
	"context"
	"runRecord/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 取得跑步紀錄
func GetRunRecord() (results []model.RunRecord) {
	data := client.Database("runRecord").Collection("run_records")

	cursor, _ := data.Find(context.TODO(), bson.M{})

	for cursor.Next(context.Background()) {
        var record model.RunRecord
        err := cursor.Decode(&record)
        if err != nil {
            return nil
        }
        results = append(results, record)
    }

	return
}

// 加入跑步紀錄
func AddRunRecord(newRecord model.RunRecord) bool{
	collection := client.Database("runRecord").Collection("run_records")
	_, err := collection.InsertOne(context.TODO(), newRecord)

	return err == nil
}

// 更新跑步紀錄
func UpdateRunRecord(id primitive.ObjectID, newRecord model.RunRecord) bool{
	collection := client.Database("runRecord").Collection("run_records")
	result, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": id}, newRecord)

	return err == nil && result.ModifiedCount != 0
}

// 刪除跑步紀錄
func DeleteRunRecord(id primitive.ObjectID) bool{
	collection := client.Database("runRecord").Collection("run_records")
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	return err == nil && result.DeletedCount != 0
}