package factory

import (
	"github.com/go-faker/faker/v4"
	"go_forum/main/database/entity"
	"log"
	"os"
)

var ReplyID uint = 1

func CreateReplyEntity(user entity.User, thread entity.Thread) entity.Reply {
	//Fill Thread entity with fake data.
	reply := entity.Reply{}
	if err := faker.FakeData(&reply); err != nil {
		log.Println("Reply generation error", err)
		os.Exit(1)
	}

	reply.ID = ReplyID
	reply.ThreadID = thread.ID
	reply.UserID = user.ID
	ReplyID++

	return reply
}

func CreateReplyEntitiesForEachUser(users []entity.User, thread entity.Thread) []entity.Reply {
	var replies []entity.Reply

	for _, user := range users {
		reply := CreateReplyEntity(user, thread)
		replies = append(replies, reply)
	}

	return replies
}
