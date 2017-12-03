package entities

import (
	"time"
)

// User is the User Schema in the database.
type User struct {
	UserID   int        `xorm:"'user_id' INT pk autoincr notnull unique"` // 语义标签
	Username string     `xorm:"'username' VARCHAR(64)"`
	Balance  float32    `xorm:"'balance' FLOAT default 0"`
	Created  *time.Time `xrom:"'created' DATETIME created"`
}

func init() {
	err := engine.Sync2(new(User))
	checkErr(err)
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
