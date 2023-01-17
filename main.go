package main

import (
	"github.com/spf13/viper"
	"go_forum/main/database"
	"go_forum/main/database/migrations"
	"go_forum/main/database/seed"
	"go_forum/main/routes"
	"log"
	"os"
)

func init() {
	initLogger()
	initViper()                                   //Configs engine.
	migrations.AutoMigrate(database.Connection()) //AutoMigrate models.
	seed.Run()                                    // Run seeds only once.
}

func main() {
	os.Exit(run())
}

func run() int {
	if err := routes.Start(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

// initViper - sets up viper configs engine.
func initViper() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// initLogger - configure logger
func initLogger() {
	file, err := os.OpenFile("./log/go/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
