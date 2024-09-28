package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"app/cron/scripts/utils"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Update struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	TypeUpdate string    `gorm:"column:TypeUpdate"`
	DateUpdate time.Time `gorm:"DateUpdate"`
}

var errorLog *log.Logger = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var infoLog *log.Logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

func checkLastUpdate(db *gorm.DB) error {
	var lastUpdate Update
	result := db.Last(&lastUpdate)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			infoLog.Println("Not found row")
			return gorm.ErrRecordNotFound
		}
		return fmt.Errorf("While get last row from update_table get error: %v", result.Error)
	}
	infoLog.Println("Found last update row")
	currentNumDay := time.Now().Add(-time.Hour * 3).Day()
	oldNumDay := lastUpdate.DateUpdate.Day()
	infoLog.Printf("Old day - %d, Current day - %d\n", oldNumDay, currentNumDay)
	if currentNumDay != oldNumDay {
		return fmt.Errorf("Need to update")
	}
	return nil
}

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	infoLog.Println("Start update admin_panel of Age")

	env := utils.NewEnv()
	gormConfig := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(env.GetCreds()), gormConfig)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Update{})
	updateFlag := checkLastUpdate(db)
	if updateFlag != nil {
		errorLog.Printf("Get error from update flag: %s", updateFlag)
		infoLog.Println("Start SQL query")
		db.Exec("UPDATE admin_panel INNER JOIN users ON admin_panel.id = users.id SET AGE = TIMESTAMPDIFF(YEAR, users.Birthdate, CURDATE())")
		result := db.Create(&Update{
			TypeUpdate: "update_age",
			DateUpdate: time.Now(),
		})
		if result.Error != nil {
			errorLog.Println("After update get error while create new row from update_table: %w", result.Error)
		} else {
			infoLog.Println("Create new row in updaters and update admin panel of Age")
		}
	} else {
		infoLog.Println("Nothings need to update")
	}
}
