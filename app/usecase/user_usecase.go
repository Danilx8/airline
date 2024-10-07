package usecase

import (
	"app/app/domain"
	"time"

	"github.com/gin-gonic/gin"
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

func (userUsecase *UserUsecase) CreateUser(c *gin.Context, user domain.User) (*domain.User, error) {
	userRes, err := userUsecase.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}
	return userRes, nil
}

func (userUsecase *UserUsecase) UpdateUser(c *gin.Context, user domain.User) error {
	err := userUsecase.userRepository.Update(&user)
	if err != nil {
		return err
	}
	return nil
}

func (userUsecase *UserUsecase) DeleteUser(c *gin.Context, id int) error {
	err := userUsecase.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
