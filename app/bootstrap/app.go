package bootstrap

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() (Application, error) {
	envFile := NewEnv()
	db, err := gorm.Open(mysql.Open(envFile.GetCreds()), &gorm.Config{})
	if err != nil {
		return Application{}, err
	}
	app := &Application{Env: envFile, DB: db}
	return *app, nil
}
