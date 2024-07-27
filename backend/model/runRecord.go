package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RunRecordDetail struct {
	Laps                   int     `bson:"laps"`                     // 圈數
	Distance               float64 `bson:"distance"`                 // 單圈距離
	Time                   string  `bson:"time"`                     // 單圈時間
	AvgPace                string  `bson:"avg_pace"`                 // 平均配速
	AvgHR                  int     `bson:"avg_hr"`                   // 平均心率
	AvgRunCadence          int     `bson:"avg_run_cadence"`          // 步頻 (步/分鐘)
	AvgGroundContactTime   int     `bson:"avg_ground_contact_time"`  // 觸地時間(ms)
	AvgStrideLength        float64 `bson:"avg_stride_length"`        // 步幅長度(m)
	AvgVerticalOscillation float64 `bson:"avg_vertical_oscillation"` // 垂直振幅(cm)
	AvgVerticalRatio       float64 `bson:"avg_vertical_ratio"`       // 移動效率(%)
}

type RunRecord struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Date                   time.Time          `bson:"date,omitempty"`           // 名稱
	Name                   string             `bson:"name,omitempty"`           // 時間
	TotalDistance          float64            `bson:"total_distance"`           // 總距離
	TotalTime              string             `bson:"total_time"`               // 總時間
	AvgPace                string             `bson:"avg_pace"`                 // 平均配速
	AvgHR                  int                `bson:"avg_hr"`                   // 平均心率
	AvgRunCadence          int                `bson:"avg_run_cadence"`          // 步頻 (步/分鐘)
	AvgGroundContactTime   int                `bson:"avg_ground_contact_time"`  // 觸地時間(ms)
	AvgStrideLength        float64            `bson:"avg_stride_length"`        // 步幅長度(m)
	AvgVerticalOscillation float64            `bson:"avg_vertical_oscillation"` // 垂直振幅(cm)
	AvgVerticalRatio       float64            `bson:"avg_vertical_ratio"`       // 移動效率(%)
	Temperature            int                `bson:"temperature"`              // 溫度
	TrainingType           string             `bson:"training_type"`            // 訓練類別
	Detail                 []RunRecordDetail  `bson:"detail"`                   // 跑步細節
	Feeling                string             `bson:"feeling"`                  // 訓練體感
}
