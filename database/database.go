package database

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
		viper.Get("POSTGRES_DB_USER"),
		viper.Get("POSTGRES_DB_USER_PASSWD"),
		viper.Get("POSTGRES_DB_HOST"),
		viper.Get("POSTGRES_DB_NAME"),
	)

	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return DB
}
