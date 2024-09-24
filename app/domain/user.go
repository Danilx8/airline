package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	CollectionUser = "users"
)

// TODO: add validation
type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Email     string `gorm:"column:Email"`
	Password  string `gorm:"column:Password"`
	BirthDate string `gorm:"column:BirthDate"`
	Active    bool   `gorm:"column:Active"`
	OfficeID  int64  `gorm:"column:OfficeID"`
	//Office    Office //`gorm:"column:OfficeID;foreignKey:ID"`
	RoleID int64 `gorm:"column:RoleID"`
	//Role      Role   //`gorm:"column:"foreignKey:RoleID;references:ID"`
}

type UserRepository interface {
	Create(user *User) (int64, error)
	Fetch(users *[]User) error
	Update(user *User) error
	Delete(id int64) error
}
