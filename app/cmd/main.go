package main

import (
	"app/app/api/route"
	"app/app/bootstrap"
	gin "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	app, err := bootstrap.App()
	if err != nil {
		log.Fatal("Error while init application: %w", err)
	}

	env := app.Env

	var db *gorm.DB
	db = app.DB

	timeout := time.Duration(env.ContextTimeout) * time.Second

	engine := gin.Default()

	route.Setup(env, timeout, db, engine)

	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler())

	err = engine.Run(env.ServerAddress)
	if err != nil {
		log.Fatal("Error while run server: %w", err)
	}
}
