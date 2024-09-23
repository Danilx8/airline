package repository

import (
	"app/app/domain"
	"fmt"
	"reflect"

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

func (u *userRepository) Create(user *domain.User) (int64, error) {
	result := u.database.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	//TODO: Подумать о смене сигнатуры, так как в юзер присваивается id
	return user.ID, result.Error
}

func (u userRepository) Fetch(users *[]domain.User) error {
	// я видел, что ты возращаешь по бачам, что хорошо, однако у нас будет мало записей,
	// поэтому можно будет просто дергать всю таблицу, но можем обсудить
	result := u.database.Find(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to fetch users: %w", result.Error)
	}

	return nil
}

func (u userRepository) GetByID(id int64) (domain.User, error) {
	user := domain.User{ID: id}

	result := u.database.First(user)
	if result.Error != nil {
		return user, fmt.Errorf("failed to fetch user with id %d: %w", id, result.Error)
	}
	return user, nil
}

func (u userRepository) Update(user *domain.User) error {
	userSearch := &domain.User{Email: user.Email}
	result := u.database.Where("Email = ?", userSearch.Email).First(userSearch)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch user with id %d: %w", user.ID, result.Error)
	}
	userVal := reflect.ValueOf(user).Elem()
	userOldVal := reflect.ValueOf(userSearch).Elem()

	for i := 0; i < userVal.NumField(); i++ {
		value := userVal.Field(i)
		if !value.IsValid() || userVal.Type().Field(i).Name == "ID" {
			continue
		}
		userOldVal.Field(i).Set(value)
	}

	return nil
}
