package entities

import (
	"time"
)

// User is the User Schema in the database.
type User struct {
	UserID   int
	Username string
	Balance  float32
	Created  *time.Time
}

// NewUser creates User entity.
func NewUser(username string) *User {
	if len(username) == 0 {
		panic("用户名不能为空")
	}
	u := User{}
	t := time.Now()
	u.Username = username
	u.Created = &t
	return &u
}
