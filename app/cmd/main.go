package main

import (
	"fmt"
	"net/http"

	"app/app/bootstrap"
	"app/app/domain"
	"app/app/repository"

	"github.com/gin-gonic/gin"
)

// TODO: rewrite return error from base string to more details
func setupRouter(app bootstrap.Application) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	userRepository := repository.NewUserRepository(app.DB)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "pong"})
	})

	// Get user value
	r.GET("/user/fetch", func(c *gin.Context) {
		var users []domain.User
		if err := userRepository.Fetch(&users); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "Not found!"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	// Create user
	r.POST("/user/create", func(c *gin.Context) {
		var user domain.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Can't parse!"})
			return
		}
		id, err := userRepository.Create(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request!"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
		fmt.Printf("Post %+v\n", user)
	})

	// Update user
	r.PUT("/user/update", func(c *gin.Context) {
		var user domain.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Can't parse!"})
			return
		}

		if err := userRepository.Update(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Error update of row"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Complete update of row"})
	})

	// TODO: Delete of user
	r.DELETE("/user/delete", func(c *gin.Context) {
		var user domain.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Can't parse!"})
			return
		}

		if err := userRepository.Delete(user.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Error delete of row"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Complete delete of row"})
	})

	return r
}

func main() {
	app, err := bootstrap.App()
	if err != nil {
		fmt.Println("Error while init application: %w", err)
	}
	// TODO: вынести роутеры в другую папку
	r := setupRouter(app)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
