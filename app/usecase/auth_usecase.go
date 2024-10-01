package usecase

import (
	"app/app/domain"
	"app/app/internal"
	"time"
)

type AuthUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(userRepository domain.UserRepository, timeout time.Duration) AuthUsecase {
	return AuthUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (a AuthUsecase) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := a.userRepository.FetchByEmail(email, &user); err != nil {
		return user, err
	}
	return user, nil
}

func (a AuthUsecase) GenerateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return internal.GenerateAccessToken(user, secret, expiry)
}

func (a AuthUsecase) GenerateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return internal.GenerateRefreshToken(user, secret, expiry)
}
