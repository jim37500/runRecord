package api

import (
	"runRecord/database"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// 取得跑步活動
func GetRunActivities(context *fiber.Ctx) error {
	athleteID, _ := context.ParamsInt("athleteid")                 // 運動員主鍵
	myAthlete := database.GetAthleteByAthleteID(uint64(athleteID)) // 依運動員主鍵取得運動員

	_ = FetchStravaActivities(myAthlete)

	return context.JSON(database.GetRunActivities(myAthlete.ID))
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

// // 取得跑步紀錄
// func GetRunRecord(context *fiber.Ctx) error {
// 	return context.JSON(database.GetRunRecord())
// }

// // 加入跑步紀錄
// func AddRunRecord(context *fiber.Ctx) error {
// 	var myRunRecord model.RunRecord
// 	_ = context.BodyParser(&myRunRecord)

// 	if database.AddRunRecord(myRunRecord) {
// 		return context.SendStatus(fiber.StatusOK)
// 	}

// 	return context.SendStatus(fiber.StatusInternalServerError)
// }

// // 更新跑步紀錄
// func UpdateRunRecord(context *fiber.Ctx) error {
// 	var myRunRecord model.RunRecord
// 	_ = context.BodyParser(&myRunRecord)

// 	idString := context.Params("id")
// 	id, _ := primitive.ObjectIDFromHex(idString)

// 	if database.UpdateRunRecord(id, myRunRecord) {
// 		return context.SendStatus(fiber.StatusOK)
// 	}

// 	return context.SendStatus(fiber.StatusInternalServerError)
// }

// // 刪除跑步紀錄
// func DeleteRunRecord(context *fiber.Ctx) error {
// 	idString := context.Params("id")

// 	id, _ := primitive.ObjectIDFromHex(idString)

// 	if database.DeleteRunRecord(id) {
// 		return context.SendStatus(fiber.StatusOK)
// 	}

// 	return context.SendStatus(fiber.StatusInternalServerError)
// }
