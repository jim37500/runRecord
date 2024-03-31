package main

import (
	"log"
	"runRecord/database"
	"runRecord/router"
)

func main() {
	log.Println("Opening Project DB...")

	// 連線資料庫
	database.Open()

	log.Println("Starting Project...")

	// 監聽網頁服務
	router.Run()

	log.Println("Stoping Project...")
}