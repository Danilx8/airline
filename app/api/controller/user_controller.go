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

func (userController *UserController) GetEmployeeUsers(c *gin.Context) {
	users, err := userController.UserUsecase.FetchAll(c, EMPLOYEES)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"UserController, GetEmployeeUsers": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (userController *UserController) CreateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"UserController, CreateUser, fail binding model": err.Error()})
		return
	}

	var userId int64
	userId, err = userController.UserUsecase.CreateUser(c, user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"UserController, CreateUser, fail creating user": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (userController *UserController) UpdateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"UserController, UpdateUser, fail binding model": err.Error()})
		return
	}

	err = userController.UserUsecase.UpdateUser(c, user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"UserController, UpdateUser, fail updating user": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (userController *UserController) DeleteUser(c *gin.Context) {
	var id int64
	err := c.BindJSON(&id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"UserController, DeleteUser, fail binding id": err.Error()})
	}

	err = userController.UserUsecase.DeleteUser(c, id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"UserController, DeleteUser, fail deleting user": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
