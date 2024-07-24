package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RunRecordDetail struct {
	Laps                   int     `bson:"laps"`
	Distance               float64 `bson:"distance"`
	Time                   string  `bson:"time"`
	AvgPace                string  `bson:"avg_pace"`
	AvgHR                  int     `bson:"avg_hr"`                   // 平均心率
	AvgRunCadence          int     `bson:"avg_run_cadence"`          // 步頻 (步/分鐘)
	AvgGroundContactTime   int     `bson:"avg_ground_contact_time"`  // 觸地時間(ms)
	AvgStrideLength        float64 `bson:"avg_stride_length"`        // 步幅長度(m)
	AvgVerticalOscillation float64 `bson:"avg_vertical_oscillation"` // 垂直振幅(cm)
	AvgVerticalRatio       float64 `bson:"avg_vertical_ratio"`       // 移動效率(%)
}

type RunRecord struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Date                   time.Time          `bson:"date,omitempty"`
	TotalDistance          float64            `bson:"total_distance"`
	TotalTime              string             `bson:"total_time"`
	AvgPace                string             `bson:"avg_pace"`
	AvgHR                  int                `bson:"avg_hr"`                   // 平均心率
	AvgRunCadence          int                `bson:"avg_run_cadence"`          // 步頻 (步/分鐘)
	AvgGroundContactTime   int                `bson:"avg_ground_contact_time"`  // 觸地時間(ms)
	AvgStrideLength        float64            `bson:"avg_stride_length"`        // 步幅長度(m)
	AvgVerticalOscillation float64            `bson:"avg_vertical_oscillation"` // 垂直振幅(cm)
	AvgVerticalRatio       float64            `bson:"avg_vertical_ratio"`       // 移動效率(%)
	Temperature            int                `bson:"temperature"`
	TrainingType           string             `bson:"training_type"`
	Detail                 []RunRecordDetail  `bson:"detail"`
	Feeling                string             `bson:"feeling"`
}
