package factory

import (
	"github.com/go-faker/faker/v4"
	"go_forum/main/entity"
	"go_forum/main/helper"
	"go_forum/main/security"
	"log"
	"os"
)

var UserID uint = 1

func CreateMultipleUserEntities(count int) []entity.User {
	var users []entity.User

	for i := 0; i < count; i++ {
		password := helper.RandomString(8)
		user := createUserEntity(entity.User{Hash: generateHash(password)})
		users = append(users, user)
		log.Printf("Generated user entity with email:%s, and password: '%s'\n", user.Email, password)
	}

	return users
}

func CreateSingleUserEntity(template entity.User, password string) entity.User {
	template.Hash = generateHash(password)
	return createUserEntity(template)
}

func createUserEntity(template entity.User) entity.User {
	//Fill User entity with fake data.
	user := entity.User{}
	if err := faker.FakeData(&user); err != nil {
		log.Println("Error user seed", err)
		os.Exit(1)
	}

	user.ID = UserID
	user.Hash = template.Hash
	UserID++

	return user
}

func generateHash(password string) string {
	hash, err := security.HashPassword(password)
	if err != nil {
		log.Println("Couldn't generate password hash for user", err)
		os.Exit(1)
	}

	return hash
}
