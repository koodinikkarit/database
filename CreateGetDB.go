package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateGetDB(
	dbUser string,
	dbPass string,
	dbIP string,
	dbPort string,
	dbName string,
) func() *gorm.DB {
	var db *gorm.DB
	return func() *gorm.DB {
		if db == nil {
			Migrate(
				dbUser,
				dbPass,
				dbIP,
				dbPort,
				dbName,
			)

			userStr := dbUser
			if dbPass != "" {
				userStr = userStr + ":" + dbPass
			}

			var err error
			db, err = gorm.Open(
				"mysql",
				userStr+"@tcp("+dbIP+":"+dbPort+")/"+dbName+"?parseTime=True&loc=Local",
			)
			if err != nil {
				log.Fatalf("Creting db connection failed error %v", err)
			}
		}
		return db.Debug()
	}
}
