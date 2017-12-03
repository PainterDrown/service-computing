package entities

import (
	"github.com/painterdrown/service-computing/cloudgo-data-template/sqlt"
)

//UserAtomicService provides the atomic operations for User.
type UserAtomicService struct{}

//UserService is one entity of UserAtomicService.
var UserService = UserAtomicService{}

// Save saves a User entity to the database.
func (*UserAtomicService) Save(u *User) error {
	dao := userDao{sqlt.NewSQLTemplate(mydb)}
	id, err := dao.Save(u)
	if err != nil {
		return err
	}
	u.UserID = id
	return err
}

// FindAll finds all the User records in the database.
func (*UserAtomicService) FindAll() ([]User, error) {
	dao := userDao{sqlt.NewSQLTemplate(mydb)}
	all, err := dao.FindAll()
	return all, err
}

// FindByID finds the User record by UserID.
func (*UserAtomicService) FindByID(userID int) *User {
	dao := userDao{sqlt.NewSQLTemplate(mydb)}
	u := dao.FindByID(userID)
	return u
}
