package internal

import (
	"os/user"
)

func FakeUser(username string) *user.User {
	return &user.User{
		Uid:      "1000",
		Gid:      "1000",
		Username: username,
		Name:     username,
		HomeDir:  "/home/" + username,
	}
}
