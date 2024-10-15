package model

import (
	"time"
)

// 游泳活動
type SwimActivity struct {
	ID                 uint64    `gorm:"primarykey"`
	AthleteID          uint64    `gorm:"comment:運動員主鍵"`
	Name               string    `gorm:"comment:活動名稱"`
	Date               time.Time `gorm:"comment:時間"`
	ElapsedTime        int       `gorm:"comment:持續時間(秒)"`
	MovingTime         int       `gorm:"comment:運動時間(秒)"`
	Distance           int       `gorm:"comment:距離(m)"`
	AverageSpeed       float32   `gorm:"comment:平均速度(km/h)"`
	MaxSpeed           float32   `gorm:"comment:最大速度(km/h)"`
	AverageHeartrate   int       `gorm:"comment:平均心率(bpm)"`
	MaxHeartrate       int       `gorm:"comment:最大心率(bpm)"`
	SwimLap            SwimLap   `gorm:"foreignKey:ActivityID"`
	// AverageWatts       int       `gorm:"comment:平均功率(bpm)"`
	// MaxWatts           int       `gorm:"comment:最大功率(bpm)"`
}

// 游泳圈數
type SwimLap struct {
	ID                 uint64  `gorm:"primarykey"`
	ActivityID         uint64  `gorm:"comment:活動主鍵"`
	AthleteID          uint64  `gorm:"comment:運動員主鍵"`
	ElapsedTime        int     `gorm:"comment:持續時間(秒)"`
	MovingTime         int     `gorm:"comment:運動時間(秒)"`
	Distance           int     `gorm:"comment:距離(m)"`
	AverageSpeed       float32 `gorm:"comment:平均速度(km/h)"`
	MaxSpeed           float32 `gorm:"comment:最大速度(km/h)"`
	AverageHeartrate   int     `gorm:"comment:平均心率(bpm)"`
	MaxHeartrate       int     `gorm:"comment:最大心率(bpm)"`
	// AverageWatts       int     `gorm:"comment:平均功率(bpm)"`
}
