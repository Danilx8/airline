package main

import (
	"app/app/api/route"
	"app/app/bootstrap"
	"fmt"
	gin "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		fmt.Println("Error while init application: %w", err)
	}

	env := app.Env

	var db *gorm.DB
	db = app.DB

	timeout := time.Duration(env.ContextTimeout) * time.Second

	engine := gin.Default()

	route.Setup(env, timeout, db, engine)

	err = engine.Run(env.ServerAddress)
	if err != nil {
		fmt.Println("Error while run server: %w", err)
	}
}
