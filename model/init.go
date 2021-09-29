package model

import (
"github.com/gin-gonic/gin"
"time"

"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/sqlite"
"log"
)
var DB *gorm.DB

const DatabaseType = "sqlite3"
const DatabasePath = "lottery.db"

func Database() error {

	db, err := gorm.Open(DatabaseType, DatabasePath)
	if err != nil {
		log.Println("open database failed:",err.Error())
		panic("failed to connect database")
		return  err
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//idle
	db.DB().SetMaxIdleConns(20)

	//open
	db.DB().SetMaxOpenConns(100)

	//timeout
	db.DB().SetConnMaxLifetime(time.Second * 30)

	//defer db.Close()

	DB = db

	migration()

	return nil
}