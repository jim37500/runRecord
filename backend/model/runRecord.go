package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RunRecordDetail struct {
	Laps     int    `bson:"laps"`
	Distance int    `bson:"distance"`
	Time     string `bson:"time"`
	AvgPace  string `bson:"avg_pace"`
}

type RunRecord struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Date          time.Time         `bson:"date,omitempty"`
	TotalDistance int               `bson:"total_distance"`
	TotalTime     string            `bson:"total_time"`
	AvgPace       string            `bson:"avg_pace"`
	Temperature   int               `bson:"temperature"`
	TrainingType  string            `bson:"training_type"`
	Detail        []RunRecordDetail `bson:"detail"`
	Feeling       string            `bson:"feeling"`
}
