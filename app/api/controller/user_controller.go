package controller

import (
	"app/app/domain"
	"app/app/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

const (
	ADMIN     = true
	EMPLOYEES = false
)

// GetEmployeeUsers airline
//
//	@Summary		Get all users in system that are not administrators
//	@Description	get non-administrator users entities
//	@ID				get-non-administrator-users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.User		"ok"
//	@Failure		500	{object}	web.APIError	"error"
//	@Router			/users [get]
func (userController *UserController) GetEmployeeUsers(c *gin.Context) {
	users, err := userController.UserUsecase.FetchAll(c, EMPLOYEES)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (userController *UserController) CreateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userId int64
	userId, err = userController.UserUsecase.CreateUser(c, user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (userController *UserController) UpdateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = userController.UserUsecase.UpdateUser(c, user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (userController *UserController) DeleteUser(c *gin.Context) {
	var id int64
	err := c.BindJSON(&id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = userController.UserUsecase.DeleteUser(c, id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
