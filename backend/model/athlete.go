package model

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