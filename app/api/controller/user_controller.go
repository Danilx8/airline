package controller

import (
	"app/app/domain"
	"app/app/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

const (
	ADMIN     = true
	EMPLOYEES = false
)

// GetEmployeeUsers godoc
// @Summary	List of employee users
// @Description get employees
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} domain.ErrorMessage
// @Router /users [get]
func (userController *UserController) GetEmployeeUsers(c *gin.Context) {
	users, err := userController.UserUsecase.FetchAll(c, EMPLOYEES)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "UserController, GetEmployeeUsers", Description: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary	Create of user
// @Tags User
// @Accept json
// @Produce json
// @Param        data    body     domain.User true  "scheme of user"
// @Success 200 {object} domain.User
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /users/create [post]
func (userController *UserController) CreateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Header: "UserController, CreateUser, fail binding model", Description: err.Error()})
		return
	}

	userRes, err := userController.UserUsecase.CreateUser(c, user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "UserController, CreateUser, fail creating user", Description: err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRes)
}

// UpdateUser godoc
// @Summary	Update of user
// @Tags User
// @Accept json
// @Produce json
// @Param        data    body     domain.User true  "scheme of user"
// @Success 200 {object} domain.User
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /users/update [put]
func (userController *UserController) UpdateUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Header: "UserController, UpdateUser, fail binding model", Description: err.Error()})
		return
	}

	err = userController.UserUsecase.UpdateUser(c, user)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "UserController, UpdateUser, fail updating user", Description: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// DeleteUser godoc
// @Summary	Delete of user
// @Tags User
// @Accept json
// @Produce json
// @Param        data    body   domain.UserId true  "scheme of user"
// @Success 200 {object} domain.User
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /users/delete [delete]
func (userController *UserController) DeleteUser(c *gin.Context) {
	var userId domain.UserId
	err := c.BindJSON(&userId)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Header: "UserController, DeleteUser, fail binding id", Description: err.Error()})
		return
	}

	err = userController.UserUsecase.DeleteUser(c, userId.Id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "UserController, DeleteUser, fail deleting user", Description: err.Error()})
		return
	}
	c.JSON(http.StatusOK, userId)
}

// GetUsersSessions godoc
// @Summary	Retrieve user sessions by their id
// @Tags User
// @Accept json
// @Produce json
// @Param        data    body   domain.UserId true  "scheme of user"
// @Success 200 {object} domain.Session
// @Failure 400 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /users/delete [delete]
func (userController *UserController) GetUsersSessions(c *gin.Context) {
	user, exists := c.Get("currentUser")

	if !exists {
		log.Print("Unauthorized user attempted to get user session")
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Header: "UserController, GetUserSessions, fail binding id", Description: "User token wasn't found"})
		return
	}

	session, err := userController.UserUsecase.UserPanel(c, int(user.(domain.User).ID))
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "UserController, GetUserSessions, fail getting user sessions", Description: err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}
