package database

import (
	"math"
	"time"

	"runRecord/data"
	"runRecord/model"
)

// 取得運動員最近一次的跑步活動
func GetLatestRunActivity(athleteID uint64) (runActivity model.RunActivity) {
	db.Where("athlete_id = ?", athleteID).Order("date DESC").First(&runActivity)

	return
}

// 取得跑步活動
func GetRunActivities(athleteID uint64) (runActivities []model.RunActivity) {
	db.Where("athlete_id = ?", athleteID).Order("date DESC").Find(&runActivities)

	return
}

// 取得騎車活動
func GetRideActivities(athleteID uint64) (rideActivities []model.RideActivity) {
	db.Where("athlete_id = ?", athleteID).Order("date DESC").Find(&rideActivities)

	return
}

// 取得游泳活動
func GetSwimActivities(athleteID uint64) (swimActivities []model.SwimActivity) {
	db.Where("athlete_id = ?", athleteID).Order("date DESC").Find(&swimActivities)

	return
}

// 取得跑步活動圈數
func GetRunLaps(athleteID uint64, activityID string) (runLaps []model.RunLap) {
	db.Where("athlete_id = ? AND activity_id = ?", athleteID, activityID).Find(&runLaps)

	return
}

// 取得騎車活動圈數
func GetRideLaps(athleteID uint64, activityID string) (rideLaps []model.RideLap) {
	db.Where("athlete_id = ? AND activity_id = ?", athleteID, activityID).Find(&rideLaps)

	return
}

// 取得游泳活動圈數
func GetSwimLaps(athleteID uint64, activityID string) (swimLaps []model.SwimLap) {
	db.Where("athlete_id = ? AND activity_id = ?", athleteID, activityID).Find(&swimLaps)

	return
}

// 依運動員與活動主鍵取得圈數
func GetLapsByAthleteAndActivityID[T any](athleteID uint64, activityID string , lap T) []T {
	var laps []T
	db.Debug().Where("athlete_id = ? AND activity_id = ?", athleteID, activityID).Find(&laps)

	return laps
}

// 新增活動
func AddActivity(activities []data.Activities) bool {
	var myRunActivities []model.RunActivity
	var myRideActivities []model.RideActivity
	var mySwimActivities []model.SwimActivity
	result := true

	for _, actitvity := range activities {
		date, _ := time.Parse(time.RFC3339, actitvity.Date)

		if actitvity.SportType == "Run" {
			myRunActivities = append(myRunActivities, model.RunActivity{
				ID:                 actitvity.ActivityID,
				AthleteID:          actitvity.Athlete.AthleteID,
				Name:               actitvity.Name,
				Date:               date,
				ElapsedTime:        actitvity.ElapsedTime,
				MovingTime:         actitvity.MovingTime,
				Distance:           int(math.Round(float64(actitvity.Distance))),
				TotalElevationGain: int(math.Round(float64(actitvity.TotalElevationGain))),
				AverageSpeed:       actitvity.AverageSpeed * 3.6,
				MaxSpeed:           actitvity.MaxSpeed * 3.6,
				AverageCadence:     int(math.Round(float64(actitvity.AverageCadence * 2))),
				AverageHeartrate:   int(math.Round(float64(actitvity.AverageHeartrate))),
				MaxHeartrate:       int(math.Round(float64(actitvity.MaxHeartrate))),
				AverageWatts:       int(math.Round(float64(actitvity.AverageWatts))),
				MaxWatts:           int(math.Round(float64(actitvity.MaxWatts))),
			})
		}

		if actitvity.SportType == "Ride" {
			myRideActivities = append(myRideActivities, model.RideActivity{
				ID:                 actitvity.ActivityID,
				AthleteID:          actitvity.Athlete.AthleteID,
				Name:               actitvity.Name,
				Date:               date,
				ElapsedTime:        actitvity.ElapsedTime,
				MovingTime:         actitvity.MovingTime,
				Distance:           int(math.Round(float64(actitvity.Distance))),
				TotalElevationGain: int(math.Round(float64(actitvity.TotalElevationGain))),
				AverageSpeed:       actitvity.AverageSpeed * 3.6,
				MaxSpeed:           actitvity.MaxSpeed * 3.6,
				AverageHeartrate:   int(math.Round(float64(actitvity.AverageHeartrate))),
				MaxHeartrate:       int(math.Round(float64(actitvity.MaxHeartrate))),
			})
		}

		if actitvity.SportType == "Swim" {
			mySwimActivities = append(mySwimActivities, model.SwimActivity{
				ID:               actitvity.ActivityID,
				AthleteID:        actitvity.Athlete.AthleteID,
				Name:             actitvity.Name,
				Date:             date,
				ElapsedTime:      actitvity.ElapsedTime,
				MovingTime:       actitvity.MovingTime,
				Distance:         int(math.Round(float64(actitvity.Distance))),
				AverageSpeed:     actitvity.AverageSpeed * 3.6,
				MaxSpeed:         actitvity.MaxSpeed * 3.6,
				AverageHeartrate: int(math.Round(float64(actitvity.AverageHeartrate))),
				MaxHeartrate:     int(math.Round(float64(actitvity.MaxHeartrate))),
			})
		}

	}

	if (len(myRunActivities) > 0 && db.Create(&myRunActivities).Error != nil) ||
		(len(myRideActivities) > 0 && db.Create(&myRideActivities).Error != nil) ||
		(len(mySwimActivities) > 0 && db.Create(&mySwimActivities).Error != nil) {
		result = false
	}

	return result
}

// 新增圈數
func AddLaps(laps []data.Lap, activityType string) bool {
	if activityType == "Run" {
		var myRunLaps []model.RunLap
		for _, runLap := range laps {
			myRunLaps = append(myRunLaps, model.RunLap{
				ID:                 runLap.LapID,
				ActivityID:         runLap.Activity.ActivityID,
				AthleteID:          runLap.Athlete.AthleteID,
				ElapsedTime:        runLap.ElapsedTime,
				MovingTime:         runLap.MovingTime,
				Distance:           int(math.Round(float64(runLap.Distance))),
				AverageSpeed:       runLap.AverageSpeed * 3.6,
				MaxSpeed:           runLap.MaxSpeed * 3.6,
				AverageCadence:     int(math.Round(float64(runLap.AverageCadence * 2))),
				AverageHeartrate:   int(math.Round(float64(runLap.AverageHeartrate))),
				MaxHeartrate:       int(math.Round(float64(runLap.MaxHeartrate))),
				AverageWatts:       int(math.Round(float64(runLap.AverageWatts))),
				TotalElevationGain: int(math.Round(float64(runLap.TotalElevationGain * 10))),
			})
		}
		return db.Create(&myRunLaps).Error == nil
	}

	if activityType == "Ride" {
		var myRideLaps []model.RideLap
		for _, runLap := range laps {
			myRideLaps = append(myRideLaps, model.RideLap{
				ID:                 runLap.LapID,
				ActivityID:         runLap.Activity.ActivityID,
				AthleteID:          runLap.Athlete.AthleteID,
				ElapsedTime:        runLap.ElapsedTime,
				MovingTime:         runLap.MovingTime,
				Distance:           int(math.Round(float64(runLap.Distance))),
				AverageSpeed:       runLap.AverageSpeed * 3.6,
				MaxSpeed:           runLap.MaxSpeed * 3.6,
				AverageHeartrate:   int(math.Round(float64(runLap.AverageHeartrate))),
				MaxHeartrate:       int(math.Round(float64(runLap.MaxHeartrate))),
				TotalElevationGain: int(math.Round(float64(runLap.TotalElevationGain * 10))),
			})
		}
		return db.Create(&myRideLaps).Error == nil
	}

	if activityType == "Swim" {
		var mySwimLaps []model.SwimLap
		for _, runLap := range laps {
			mySwimLaps = append(mySwimLaps, model.SwimLap{
				ID:                 runLap.LapID,
				ActivityID:         runLap.Activity.ActivityID,
				AthleteID:          runLap.Athlete.AthleteID,
				ElapsedTime:        runLap.ElapsedTime,
				MovingTime:         runLap.MovingTime,
				Distance:           int(math.Round(float64(runLap.Distance))),
				AverageSpeed:       runLap.AverageSpeed * 3.6,
				MaxSpeed:           runLap.MaxSpeed * 3.6,
				AverageHeartrate:   int(math.Round(float64(runLap.AverageHeartrate))),
				MaxHeartrate:       int(math.Round(float64(runLap.MaxHeartrate))),
			})
		}
		return db.Create(&mySwimLaps).Error == nil
	}

	return false
}
