package model

import (
	"gorm.io/gorm"
)

// 自動遷移資料庫
func AutoMigrate(db *gorm.DB) {
	migrateTable(db, &Athlete{})
	migrateTable(db, &Activity{})
	migrateTable(db, &Lap{})

	checkTableData(db)
}

// 轉移資料表結構
func migrateTable(db *gorm.DB, structure interface{}) {
	if !db.Migrator().HasTable(structure) {
		// _ = db.Set("gorm:table_options", " COMMENT='"+tableName+"'").Migrator().CreateTable(structure)
		_ = db.Migrator().CreateTable(structure)
	}
	_ = db.AutoMigrate(structure)
}

// 檢查資料表有無資料
func checkTableData(db *gorm.DB) {
	// 若 沒運動員
	var athlete Athlete
	if db.First(&athlete).RowsAffected == 0 {
		role := []Athlete{{ID: 108845218, ClientID: 134888, ClientSecret: "a613adf2b923051df828c21c1101895288334333", AccessToken: "dacc2edac2a5ced86883014ad2dec4c797cd065b", RefreshToken: "cb36f0f9e76a74b35d4df32a8a36693256cea1e9"}}
		db.Create(role)
	}
}
