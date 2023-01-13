package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go_forum/main/database"
	"go_forum/main/helper/seeder"
	"go_forum/main/migrations"
	"log"
	"net/http"
	"os"
)

func init() {
	initLogger()
	initViper()                                   //Configs engine.
	migrations.AutoMigrate(database.Connection()) //AutoMigrate models.
	seeder.Run()                                  // Run seeds only once.
}

func main() {
	os.Exit(run())
}

func run() int {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API server is working!",
		})
	})

	if err := r.Run(":8080"); err != nil {
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
