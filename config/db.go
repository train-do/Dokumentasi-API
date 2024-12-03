package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/project-sistem-voucher/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", Cfg.Database.Host, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname, Cfg.Database.Port)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	once.Do(func() {
		DB = Db
		fmt.Println("Successfully Connected To Database!")
	})
}

func SyncDB() {
	//nanti disini yta migrate tablenya
	if err := DB.AutoMigrate(&model.Voucher{}, &model.Redeem{}, model.Use{}); err != nil {
		fmt.Printf("AutoMigrate error: %s\n", err)
		panic(err)
	} else {
		fmt.Println("Database migrated successfully!")
	}

}
