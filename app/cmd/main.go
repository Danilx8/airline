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

//	@title			Swagger Example API
//	@version		1.0
//	@description	Описание работы сервера
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host		localhost:8080
//	@BasePath	/

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
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
