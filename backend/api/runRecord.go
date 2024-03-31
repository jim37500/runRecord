package api

import (
	"runRecord/database"
	"runRecord/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 取得跑步紀錄
func GetRunRecord(context *fiber.Ctx) error {
	return context.JSON(database.GetRunRecord())
}

// 加入跑步紀錄
func AddRunRecord(context *fiber.Ctx) error {
	var myRunRecord model.RunRecord
	_ = context.BodyParser(&myRunRecord)


	if database.AddRunRecord(myRunRecord) {
		return context.SendStatus(fiber.StatusOK)
	}
	
	return context.SendStatus(fiber.StatusInternalServerError)
}

// 更新跑步紀錄
func UpdateRunRecord(context *fiber.Ctx) error {
	var myRunRecord model.RunRecord
	_ = context.BodyParser(&myRunRecord)

	idString := context.Params("id")
	id, _ := primitive.ObjectIDFromHex(idString)


	if database.UpdateRunRecord(id, myRunRecord) {
		return context.SendStatus(fiber.StatusOK)
	}
	
	return context.SendStatus(fiber.StatusInternalServerError)
}

// 刪除跑步紀錄
func DeleteRunRecord(context *fiber.Ctx) error {
	idString := context.Params("id")

	id, _ := primitive.ObjectIDFromHex(idString)

	if database.DeleteRunRecord(id) {
		return context.SendStatus(fiber.StatusOK)
	}
	
	return context.SendStatus(fiber.StatusInternalServerError)
}