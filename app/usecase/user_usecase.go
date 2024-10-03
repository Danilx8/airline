package usecase

import (
	"app/app/domain"
	"github.com/gin-gonic/gin"
	"time"
)

type UserUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

// TODO: Realize timeout logic in each function
func (userUsecase *UserUsecase) FetchAll(c *gin.Context, isAdmin bool) ([]domain.User, error) {
	var users []domain.User
	var err error
	if !isAdmin {
		err = userUsecase.userRepository.Fetch(&users)
	} else {
		err = userUsecase.userRepository.FetchAdmins(&users)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userUsecase *UserUsecase) CreateUser(c *gin.Context, user domain.User) (int64, error) {
	userId, err := userUsecase.userRepository.Create(&user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (userUsecase *UserUsecase) UpdateUser(c *gin.Context, user domain.User) error {
	err := userUsecase.userRepository.Update(&user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase *UserUsecase) DeleteUser(c *gin.Context, userID int64) error {
	err := userUsecase.userRepository.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}
