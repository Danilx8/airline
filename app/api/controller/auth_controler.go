package controller

import (
	"app/app/bootstrap"
	"app/app/domain"
	"app/app/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	AuthUsecase usecase.AuthUsecase
	Env         bootstrap.Env
}

func (ac *AuthController) Login(c *gin.Context) {
	var request domain.Login

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"AuthController, Login, fail binding login data": err.Error()})
		return
	}

	user, err := ac.AuthUsecase.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"AuthController, Login, user with given email not found": err.Error()})
		return
	}

	//TODO: начать шифровать пароли в дб (нет в ТЗ)
	//if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
	if user.Password != request.Password {
		c.JSON(http.StatusBadRequest, gin.H{"AuthController, Login": "wrong password"})
		return
	}

	accessToken, err := ac.AuthUsecase.GenerateAccessToken(&user, ac.Env.AccessTokenSecret, ac.Env.AccessTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"AuthController, Login, fail to generate access token": err.Error()})
		return
	}

	refreshToken, err := ac.AuthUsecase.GenerateRefreshToken(&user, ac.Env.RefreshTokenSecret, ac.Env.RefreshTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"AuthController, Login, fail to generate refresh token": err.Error()})
		return
	}

	response := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
