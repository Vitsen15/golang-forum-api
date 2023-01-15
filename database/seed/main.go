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
	return []Seed{
		{
			Name: "Threads_With_Users_And_Replies",
			Run: func(db *gorm.DB) error {
				var err error

				//Generate entry user entity, with known password to have ability to authenticate.
				entryUser := factory.CreateSingleUserEntity(entity.User{Email: "email@domain.com"}, "Temp123#")
				//Generate user entities with randomly generated passwords.
				users := factory.CreateMultipleUserEntities(10)
				//Create main thread for entry user.
				thread := factory.CreateThreadEntity(entryUser)
				//Create reply entities.
				replies := factory.CreateReplyEntitiesForEachUser(users, thread)

				//Create entry user in DB.
				if err = CreateUser(db, entryUser); err != nil {
					return err
				}
				//Create users from generated entities above in DB.
				if err = CreateUsers(db, users); err != nil {
					return err
				}

				//Create thread from generated entities above in DB.
				if replyErr := CreateThread(db, thread); replyErr != nil {
					return replyErr
				}

				//Create replies from generated entities above in DB.
				if replyErr := CreateReplies(db, replies); replyErr != nil {
					return replyErr
				}

				return err
			},
		},
	}
}
