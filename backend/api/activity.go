package api

import (
	"fmt"
	"time"

	"runRecord/database"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

var activityCache = cache.New(30*time.Minute, 60*time.Minute)

// 取得活動
func GetActivities(context *fiber.Ctx) error {
	athleteID, _ := context.ParamsInt("athleteid") // 運動員主鍵
	cacheKey := fmt.Sprintf("activity-%d", athleteID)

	if cachedData, found := activityCache.Get(cacheKey); found {
		return context.JSON(cachedData)
	}

	myAthlete := database.GetAthleteByAthleteID(uint64(athleteID)) // 依運動員主鍵取得運動員

	_ = FetchStravaActivities(myAthlete)

	runActivities := database.GetRunActivities(myAthlete.ID)
	rideActivities := database.GetRideActivities(myAthlete.ID)
	swimActivities := database.GetSwimActivities(myAthlete.ID)

	result := fiber.Map{
		"RunActivities":  runActivities,
		"RideActivities": rideActivities,
		"SwimActivities": swimActivities,
	}
	activityCache.Set(cacheKey, result, 30*time.Minute)

	return context.JSON(result)
}

// 取得活動圈數
func GetActivityLaps(context *fiber.Ctx) error {
	athleteID, _ := context.ParamsInt("athleteid") // 運動員主鍵
	activityID := context.Params("activityid")     // 活動主鍵
	activityType := context.Params("type")         // 活動類型

	myAthlete := database.GetAthleteByAthleteID(uint64(athleteID)) // 依運動員主鍵取得運動員

	if activityType == "Run" {
		runLaps := database.GetRunLaps(myAthlete.ID, activityID) // 取得跑步活動圈數
		// 若 沒有跑步圈數 則 爬取後再取得跑步活動圈數
		if len(runLaps) == 0 {
			_ = FetchStravaLaps(myAthlete, activityID, activityType)
			runLaps = database.GetRunLaps(myAthlete.ID, activityID) // 取得跑步活動圈數
		}
		return context.JSON(runLaps)
	}

	if activityType == "Ride" {
		rideLaps := database.GetRideLaps(myAthlete.ID, activityID) // 取得騎車活動圈數
		// 若 沒有騎車圈數 則 爬取後再取得跑步活動圈數
		if len(rideLaps) == 0 {
			_ = FetchStravaLaps(myAthlete, activityID, activityType)
			rideLaps = database.GetRideLaps(myAthlete.ID, activityID) // 取得騎車活動圈數
		}
		return context.JSON(rideLaps)
	}

	if activityType == "Swim" {
		swimLaps := database.GetSwimLaps(myAthlete.ID, activityID) // 取得游泳活動圈數
		// 若 沒有游泳圈數 則 爬取後再取得跑步活動圈數
		if len(swimLaps) == 0 {
			_ = FetchStravaLaps(myAthlete, activityID, activityType)
			swimLaps = database.GetSwimLaps(myAthlete.ID, activityID) // 取得游泳活動圈數
		}
		return context.JSON(swimLaps)
	}

	return context.SendStatus(fiber.StatusBadRequest)
}
