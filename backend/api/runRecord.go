package api

import (
	"fmt"
	"time"

	"runRecord/database"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var runRecordCache = cache.New(30*time.Minute, 60*time.Minute)

// 取得跑步活動
func GetRunActivities(context *fiber.Ctx) error {
	athleteID, _ := context.ParamsInt("athleteid") // 運動員主鍵
	cacheKey := fmt.Sprintf("runRecord-%d", athleteID)

	if cachedData, found := runRecordCache.Get(cacheKey); found {
		return context.JSON(cachedData)
	}

	myAthlete := database.GetAthleteByAthleteID(uint64(athleteID)) // 依運動員主鍵取得運動員

	_ = FetchStravaActivities(myAthlete)

	runRecord := database.GetRunActivities(myAthlete.ID)
	runRecordCache.Set(cacheKey, runRecord, 30*time.Minute)

	return context.JSON(runRecord)
}

// 取得跑步活動圈數
func GetRunActivityLaps(context *fiber.Ctx) error {
	athleteID, _ := context.ParamsInt("athleteid") // 運動員主鍵
	activityID := context.Params("activityid")     // 跑步活動主鍵

	myAthlete := database.GetAthleteByAthleteID(uint64(athleteID)) // 依運動員主鍵取得運動員

	runLaps := database.GetRunLaps(myAthlete.ID, activityID) // 取得跑步活動圈數

	// 若 沒有跑步圈數 則 爬取後再取得跑步活動圈數
	if len(runLaps) == 0 {
		_ = FetchStravaLaps(myAthlete, activityID)
		runLaps = database.GetRunLaps(myAthlete.ID, activityID) // 取得跑步活動圈數
	}

	return context.JSON(runLaps)
}
