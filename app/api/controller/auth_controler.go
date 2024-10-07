package controller

import (
	"app/app/bootstrap"
	"app/app/domain"
	"app/app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUsecase usecase.AuthUsecase
	Env         bootstrap.Env
}

// Login godoc
// @Summary	Login of user
// @Tags Login
// @Accept json
// @Produce json
// @Param        data    body   domain.Login true  "scheme of login"
// @Success 200 {object} domain.LoginResponse
// @Failure 400 {object} domain.ErrorMessage
// @Failure 404 {object} domain.ErrorMessage
// @Failure 500 {object} domain.ErrorMessage
// @Router /login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var request domain.Login

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Header: "AuthController, Login, fail binding login data", Description: err.Error()})
		return
	}

	user, err := ac.AuthUsecase.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorMessage{Header: "AuthController, Login, user with given email not found", Description: err.Error()})
		return
	}

	//TODO: начать шифровать пароли в дб (нет в ТЗ)
	//if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
	if user.Password != request.Password {
		c.JSON(http.StatusBadRequest, domain.ErrorMessage{Header: "AuthController, Login", Description: "wrong password"})
		return
	}

	accessToken, err := ac.AuthUsecase.GenerateAccessToken(&user, ac.Env.AccessTokenSecret, ac.Env.AccessTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "AuthController, Login, fail to generate access token", Description: err.Error()})
		return
	}

	refreshToken, err := ac.AuthUsecase.GenerateRefreshToken(&user, ac.Env.RefreshTokenSecret, ac.Env.RefreshTokenExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorMessage{Header: "AuthController, Login, fail to generate refresh token", Description: err.Error()})
		return
	}

	response := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
