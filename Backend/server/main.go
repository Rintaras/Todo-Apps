package main

import (
	"log"

	"todo-apps/backend/server/Config"
	"todo-apps/backend/server/Models"
	"todo-apps/backend/server/Routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if p := Config.LoadDotEnv(); p != "" {
		log.Printf("db: 読み込んだ .env: %s", p)
	}

	cfg, err := Config.BuildDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("mysql", Config.DbURL(cfg))
	if err != nil {
		log.Fatalf("db: 接続に失敗しました（MySQL が起動しているか、ホスト・ポート・認証が合っているか確認）: %v", err)
	}
	Config.DB = db
	defer func() {
		if err := Config.DB.Close(); err != nil {
			log.Printf("db: Close: %v", err)
		}
	}()

	if err := Config.DB.AutoMigrate(&Models.Todo{}).Error; err != nil {
		log.Fatalf("db: AutoMigrate: %v", err)
	}

	r := Routes.SetupRouter()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
