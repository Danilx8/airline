package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = make(map[string]string)

// TODO: перенести всю логику работы с env в другой файл
func initEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("Not found env file")
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}

func getCreds() string {
	initEnv()
	userName := getEnv("MYSQL_USER")
	userPassword := getEnv("MYSQL_PASSWORD")
	dbHost := getEnv("DB_HOST")
	dbPort := getEnv("DB_PORT")
	tableName := getEnv("MYSQL_DATABASE")
	creds := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, userPassword,
		dbHost, dbPort, tableName,
	)
	//fmt.Println(creds)
	return creds
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	creds := getCreds()
	_, err := gorm.Open(mysql.Open(creds), &gorm.Config{})
	if err != nil {
		fmt.Errorf("Error while init connection with db: %w", err)
	}
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "pong"})
	})

	// Get user value
	//r.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := db[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)

	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}

	//	if c.Bind(&json) == nil {
	//		db[user] = json.Value
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
