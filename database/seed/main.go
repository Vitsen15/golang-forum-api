package seed

import (
	"go_forum/main/entity"
	"go_forum/main/entity/factory"
	"gorm.io/gorm"
)

// Seed is seeding database configuration.
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

// All is running all seeders.
func All() []Seed {
	entryUserEmail := "email@domain.com"
	entryUserPassword := "Temp123#"

	return []Seed{
		{
			Name: "CreateUsers",
			Run: func(db *gorm.DB) error {
				//Generate entry user entity, with known password to have ability to authenticate.
				entryUser := factory.CreateUserEntity(entity.User{Email: entryUserEmail, Hash: entryUserPassword})
				//Generate user entities with randomly generated passwords.
				users := factory.CreateMultipleUserEntities(10)

				var err error

				// Create entry user in DB.
				if err = CreateUser(db, entryUser); err != nil {
					return err
				}

				//Create users from generated entities above in DB.
				if err = CreateUsers(db, users); err != nil {
					return err
				}

				return err
			},
		},
		{
			Name: "CreateThreads",
			Run: func(db *gorm.DB) error {
				var entryUser entity.User
				db.Where("email = ?", entryUserEmail).First(&entryUser)

				// Create main thread for entry user.
				thread := factory.CreateThreadEntity(entryUser)
				// Create thread from generated entity.
				return CreateThread(db, thread)

			},
		},
		{
			Name: "CreateReplies",
			Run: func(db *gorm.DB) error {
				//Retrieve generated users from DB except main.
				var users []entity.User
				db.Where("email != ?", entryUserEmail).Find(&users)
				//Retrieve first generated thread from DB.
				var thread entity.Thread
				db.First(&thread)

				// Create reply entities.
				replies := factory.CreateReplyEntitiesForEachUser(users, thread)
				// Create replies from generated entities above in DB.
				return CreateReplies(db, replies)
			},
		},
	}
}
