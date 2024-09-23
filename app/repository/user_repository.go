package repository

import (
	"app/app/domain"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (u *userRepository) Create(c context.Context, user *domain.User) (int64, error) {
	result := u.database.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	//TODO: Подумать о смене сигнатуры, так как в юзер присваивается id
	return user.ID, result.Error
}

func (u userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	// я видел, что ты возращаешь по бачам, что хорошо, однако у нас будет мало записей,
	// поэтому можно будет просто дергать всю таблицу, но можем обсудить
	result := u.database.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", result.Error)
	}

	return users, nil
}

func (u userRepository) GetByID(c context.Context, id int64) (domain.User, error) {
	user := domain.User{ID: id}

	result := u.database.First(user)
	if result.Error != nil {
		return user, fmt.Errorf("failed to fetch user with id %d: %w", id, result.Error)
	}
	return user, nil
}
