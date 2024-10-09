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

// 取得跑步活動圈數
func GetRunLaps(athleteID uint64, activityID string) (runLaps []model.RunLap) {
	db.Where("athlete_id = ? AND activity_id = ?", athleteID, activityID).Find(&runLaps)

	return
}

// 新增跑步活動
func AddRunActivity(runActivities []data.RunActivities) bool {
	var myRunActivities []model.RunActivity
	for _, runActitvity := range runActivities {
		if runActitvity.SportType != "Run" {
			continue
		}

		runDate, _ := time.Parse(time.RFC3339, runActitvity.Date)

		myRunActivities = append(myRunActivities, model.RunActivity{
			ID:               runActitvity.ActivityID,
			AthleteID:        runActitvity.Athlete.AthleteID,
			Name:             runActitvity.Name,
			Date:             runDate,
			ElapsedTime:      runActitvity.ElapsedTime,
			MovingTime:       runActitvity.MovingTime,
			Distance:         int(math.Round(float64(runActitvity.Distance))),
			AverageSpeed:     runActitvity.AverageSpeed * 3.6,
			MaxSpeed:         runActitvity.MaxSpeed * 3.6,
			AverageCadence:   int(math.Round(float64(runActitvity.AverageCadence * 2))),
			AverageHeartrate: int(math.Round(float64(runActitvity.AverageHeartrate))),
			MaxHeartrate:     int(math.Round(float64(runActitvity.MaxHeartrate))),
			AverageWatts:     int(math.Round(float64(runActitvity.AverageWatts))),
			MaxWatts:         int(math.Round(float64(runActitvity.MaxWatts))),
		})
	}

	if len(myRunActivities) == 0 {
		return true
	}

	return db.Create(&myRunActivities).Error == nil
}

// 新增跑步圈數
func AddRunLaps(runLaps []data.RunLap) bool {
	var myRunLaps []model.RunLap
	for _, runLap := range runLaps {
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
