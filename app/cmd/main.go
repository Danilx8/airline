package main

import (
	"app/app/api/route"
	"app/app/bootstrap"
	"fmt"
	"log"
	"os"
	"time"

	gin "github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	log.SetOutput(file)

	app, err := bootstrap.App()
	if err != nil {
		log.Fatalf("Error while init application: %s\n", err.Error())
	}

	env := app.Env

	db := app.DB

	timeout := time.Duration(env.ContextTimeout) * time.Second

	engine := gin.Default()

	route.Setup(env, timeout, db, engine)

	err = engine.Run(env.ServerAddress)
	if err != nil {
		log.Fatalf("Error while run server: %s\n", err)
	}
}
