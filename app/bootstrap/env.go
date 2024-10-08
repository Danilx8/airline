package bootstrap

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	DBRootName              string `mapstructure:"MYSQL_ROOT"`
	DBRootPassword          string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	DBUserName              string `mapstructure:"MYSQL_USER"`
	DBUserPassword          string `mapstructure:"MYSQL_PASSWORD"`
	DBTableName             string `mapstructure:"MYSQL_DATABASE"`
	DBHost                  string `mapstructure:"DB_HOST"`
	DBPort                  string `mapstructure:"DB_PORT"`
	ContextTimeout          int    `mapstructure:"CONTEXT_TIMEOUT"`
	ServerAddress           string `mapstructure:"SERVER_ADDRESS"`
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHours  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOURS"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN_SECRET"`
	RefreshTokenExpiryHours int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOURS"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Can't find the file .env : ", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Println("Environment can't be loaded: ", err)
		os.Exit(1)
	}

	return &env
}
func (e Env) GetCreds() string {
	//creds := fmt.Sprintf(
	//	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	e.DBUserName, e.DBUserPassword,
	//	e.DBHost, e.DBPort, e.DBTableName,
	//)
	creds := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		e.DBUserName, e.DBUserPassword,
		e.DBTableName,
	)
	return creds
}
