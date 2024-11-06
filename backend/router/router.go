package router

import (
	"runRecord/api"

	"github.com/gofiber/fiber/v2"
)

const (
	bodyLimit = 1 * 1024 * 1024 // 上傳大小限制1MB
)

var HttpApplication *fiber.App

// 監聽網頁服務
func Run() {
	HttpApplication = fiber.New(fiber.Config{BodyLimit: bodyLimit, UnescapePath: true, ErrorHandler: notFoundHandler})

	setupRoute() // 設定路由

	_ = HttpApplication.Listen(":61018")
}

// 未找到路徑轉到首頁
func notFoundHandler(context *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// 若 未找到路徑 則 回到首頁
	if code == fiber.StatusNotFound {
		return context.Redirect("/")
	}

	return context.Status(code).SendString(err.Error())
}

// 設定路由
func setupRoute() {
	HttpApplication.Get("/api/activities/:athleteid", api.GetActivities)                          // 取得所有活動紀錄
	HttpApplication.Get("/api/activities/laps/:athleteid/:activityid", api.GetActivityLaps) // 取得活動紀錄圈數
}
