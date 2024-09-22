package domain

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        int64  `bson:"_id"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	BirthDate string `bson:"birthDate"`
	Active    bool   `bson:"active"`
	OfficeID  int64  `bson:"officeId"`
	RoleID    int64  `bson:"roleId"`
}

type UserRepository interface {
	Create(c context.Context, user *User) (int64, error)
	Fetch(c context.Context) ([]User, error)
	GetByID(c context.Context, id int64) (User, error)
}
