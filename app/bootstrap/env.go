package bootstrap

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
	viper.SetConfigFile(".env")

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
func (e Env) GetCreds() string {
	creds := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		e.DBUserName, e.DBUserPassword,
		e.DBHost, e.DBPort, e.DBTableName,
	)
	return creds
}
