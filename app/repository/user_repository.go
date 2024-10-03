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
	//TODO: Подумать о смене сигнатуры, так как в юзер присваивается id по ссылке
	return user.ID, result.Error
}

func (u userRepository) Fetch(users *[]domain.User) error {
	// я видел, что ты возращаешь по бачам, что хорошо, однако у нас будет мало записей,
	// поэтому можно будет просто дергать всю таблицу, но можем обсудить
	result := u.database.Table("users").Joins("LEFT JOIN roles ON users.RoleID = roles.ID").Where("roles.Title <> ?", "Administrator").Find(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to fetched users: %w", result.Error)
	}

	return nil
}

func (u userRepository) FetchAdmins(users *[]domain.User) error {
	result := u.database.Table("users").Joins("LEFT JOIN roles ON users.RoleID = roles.ID").Where("roles.Title = ?", "Administrator").Find(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to fetched users: %w", result.Error)
	}

	return nil
}

func (u userRepository) FetchByEmail(email string, user *domain.User) error {
	result := u.database.Table("users").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to fetched user: %w", result.Error)
	}

	return nil
}

func (u userRepository) Update(user *domain.User) error {
	userOld := &domain.User{}
	result := u.database.Where("ID = ?", user.ID).First(userOld)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch user with id %d: %w", user.ID, result.Error)
	}
	userVal := reflect.ValueOf(user).Elem()
	userOldVal := reflect.ValueOf(userOld).Elem()

	for i := 0; i < userVal.NumField(); i++ {
		value := userVal.Field(i)
		if !value.IsValid() || userVal.Type().Field(i).Name == "ID" {
			continue
		}
		userOldVal.Field(i).Set(value)
	}

	result = u.database.Save(userOld)
	if result.Error != nil {
		return fmt.Errorf("failed to update user with id %d: %w", user.ID, result.Error)
	}

	return nil
}
func (u userRepository) Delete(id int64) error {
	var user domain.User
	result := u.database.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to found user with id %d: %w", id, result.Error)
	}
	result = u.database.Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user with id %d: %w", id, result.Error)
	}

	return nil
}
