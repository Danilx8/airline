package usecase

import (
	"app/app/domain"
	"app/app/internal"
	"fmt"
	"time"
)

type AuthUsecase struct {
	userRepository    domain.UserRepository
	contextTimeout    time.Duration
	sessionRepository domain.SessionRepository
}

func NewAuthUsecase(userRepository domain.UserRepository, sessionRepository domain.SessionRepository, timeout time.Duration) AuthUsecase {
	return AuthUsecase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		contextTimeout:    timeout,
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

func (a AuthUsecase) StartLoginSession(user *domain.User) error {
	session := domain.Session{
		UserId:    int(user.ID),
		Date:      time.DateTime,
		LoginTime: time.Now(),
	}
	err := a.sessionRepository.Start(&session)
	if err != nil {
		return fmt.Errorf("couldn't start new session for user with id %d: %w", user.ID, err)
	}
	return nil
}
