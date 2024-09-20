package database

import "runRecord/model"

// 依運動員主鍵取得運動員
func GetAthleteByAthleteID(athleteID uint64) (athlete model.Athlete) {
	db.First(&athlete, athleteID)

	return 
}

// 更新運動員權杖
func UpdateAthleteToken(athleteID uint64, accessToken string, refreshToken string) bool {
	var athlete model.Athlete
	db.First(&athlete, athleteID)

	if athlete.ID == 0 {
		return false
	}

	athlete.AccessToken = accessToken
	athlete.RefreshToken = refreshToken

	return db.Save(&athlete).Error == nil
}