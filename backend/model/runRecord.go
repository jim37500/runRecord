package model

import (
	"time"
)

// 運動員
type Athlete struct {
	ID           uint64      `gorm:"primarykey"`
	ClientID     uint        `gorm:"comment:客戶ID"`
	ClientSecret string      `gorm:"comment:客戶密鑰"`
	AccessToken  string      `gorm:"comment:通行權杖"`
	RefreshToken string      `gorm:"comment:換發權杖"`
	RunLap       RunLap      `gorm:"foreignKey:AthleteID"`
	RunActivity  RunActivity `gorm:"foreignKey:AthleteID"`
}

// 跑步活動
type RunActivity struct {
	ID               uint64    `gorm:"primarykey"`
	AthleteID        uint64    `gorm:"comment:運動員主鍵"`
	Name             string    `gorm:"comment:活動名稱"`
	Date             time.Time `gorm:"comment:時間"`
	ElapsedTime      int       `gorm:"comment:持續時間(秒)"`
	MovingTime       int       `gorm:"comment:運動時間(秒)"`
	Distance         int       `gorm:"comment:距離(m)"`
	AverageSpeed     float32   `gorm:"comment:平均速度(km/h)"`
	MaxSpeed         float32   `gorm:"comment:最大速度(km/h)"`
	AverageCadence   int       `gorm:"comment:踏頻(spm)"`
	AverageHeartrate int       `gorm:"comment:平均心率(bpm)"`
	MaxHeartrate     int       `gorm:"comment:最大心率(bpm)"`
	AverageWatts     int       `gorm:"comment:平均功率(bpm)"`
	MaxWatts         int       `gorm:"comment:最大功率(bpm)"`
	RunLap           RunLap    `gorm:"foreignKey:ActivityID"`
}

// 跑步圈數
type RunLap struct {
	ID               uint64  `gorm:"primarykey"`
	ActivityID       uint64  `gorm:"comment:活動主鍵"`
	AthleteID        uint64  `gorm:"comment:運動員主鍵"`
	ElapsedTime      int     `gorm:"comment:持續時間(秒)"`
	MovingTime       int     `gorm:"comment:運動時間(秒)"`
	Distance         int     `gorm:"comment:距離(m)"`
	AverageSpeed     float32 `gorm:"comment:平均速度(km/h)"`
	MaxSpeed         float32 `gorm:"comment:最大速度(km/h)"`
	AverageCadence   int     `gorm:"comment:踏頻(spm)"`
	AverageHeartrate int     `gorm:"comment:平均心率(bpm)"`
	MaxHeartrate     int     `gorm:"comment:最大心率(bpm)"`
	AverageWatts     int     `gorm:"comment:平均功率(bpm)"`
}