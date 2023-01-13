package factory

import (
	"github.com/go-faker/faker/v4"
	"go_forum/main/database/entity"
	"log"
	"os"
)

var ThreadID uint = 1

func CreateThreadEntity(user entity.User) entity.Thread {
	//Fill Thread entity with fake data.
	thread := entity.Thread{}
	if err := faker.FakeData(&thread); err != nil {
		log.Println("Thread generation error", err)
		os.Exit(1)
	}

	thread.ID = ThreadID
	thread.UserID = user.ID
	ThreadID++

	return thread
}
