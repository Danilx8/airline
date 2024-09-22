package repository

import (
	"app/app/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (u *userRepository) Create(c context.Context, user *domain.User) (int64, error) {
	result, err := u.database.Exec("INSERT INTO users (Email, Password, FirstName, LastName, Birthdate, OfficeID, RoleID) VALUES (?, ?, ?)",
		user.Email, user.Password, user.FirstName, user.LastName, user.BirthDate, user.OfficeID, 2)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func (u userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User

	result, err := u.database.Query("SELECT * FROM users WHERE RoleID != 1")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer result.Close()
	for result.Next() {
		var user domain.User
		if err := result.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.OfficeID, &user.BirthDate, &user.Active); err != nil {
			return nil, fmt.Errorf("failed to fetch users: %w", err)
		}
		users = append(users, user)
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	return users, nil
}

func (u userRepository) GetByID(c context.Context, id int64) (domain.User, error) {
	var user domain.User

	row := u.database.QueryRow("SELECT * FROM users WHERE ID = ?", id)
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.OfficeID, &user.BirthDate, &user.Active); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user with id %d not found", id)
		}
		return user, fmt.Errorf("failed to fetch user with id %d: %w", id, err)
	}
	return user, nil
}
