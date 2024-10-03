package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	DBRootName     string `mapstructure:"MYSQL_ROOT"`
	DBRootPassword string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBTableName    string `mapstructure:"MYSQL_DATABASE"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile("/build/.env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Println("Environment can't be loaded: ", err)
	}

	return &env
}

// TODO: настроить пользователя, чтобы подключаться через него, а не рут
func (e Env) GetCreds() string {
	//creds := fmt.Sprintf(
	//	"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	e.DBRootName, e.DBRootPassword,
	//	e.DBTableName,
	//)
	creds := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		e.DBRootName, e.DBRootPassword,
		e.DBTableName,
	)
	return creds
}
