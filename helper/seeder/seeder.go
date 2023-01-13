package seeder

import (
	"go_forum/main/database"
	"go_forum/main/database/seed"
	"log"
	"os"
)

func Run() {
	isDatabaseSeededFileName := "./isSeeded"

	if _, err := os.Stat(isDatabaseSeededFileName); err != nil {
		log.Println("File doesn't exist")
		log.Println("Seeding")
		for _, seeder := range seed.All() {
			if err := seeder.Run(database.Connection()); err != nil {
				log.Printf("Running seeder '%s', failed with error: %s\n", seeder.Name, err)
				os.Exit(1)
			}
		}

		if _, fileErr := os.Create(isDatabaseSeededFileName); fileErr != nil {
			log.Println("Error during creation of seeds flag file.", fileErr)
			os.Exit(1)
		}

		log.Println("Created seeds flag file.")
	}
}
