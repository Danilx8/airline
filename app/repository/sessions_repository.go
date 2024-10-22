package repository

import (
	"app/app/domain"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

type sessionRepository struct {
	database *gorm.DB
}

func NewSessionRepository(db *gorm.DB) domain.SessionRepository {
	return &sessionRepository{
		database: db,
	}
}

func (s sessionRepository) Start(session *domain.Session) error {
	result := s.database.Create(&session)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (s sessionRepository) Update(session *domain.Session) error {
	oldSession := &domain.Session{}
	result := s.database.Where("userId = ?", session.UserId).Last(&oldSession)

	if result.Error != nil {
		return fmt.Errorf("could not set logout for user with id %d: %s", session.UserId, result.Error)
	}

	sessionVal := reflect.ValueOf(session).Elem()
	oldSessionVal := reflect.ValueOf(oldSession).Elem()
	for i := 0; i < sessionVal.NumField(); i++ {
		value := sessionVal.Field(i)
		if !value.IsValid() || value.Type().Field(i).Name == "id" {
			continue
		}
		oldSessionVal.Field(i).Set(value)
	}

	result = s.database.Save(oldSession)
	if result.Error != nil {
		return fmt.Errorf("could not set logout for user with id %d: %s", session.UserId, result.Error)
	}
	return nil
}
