package entities

import (
	"errors"
)

type userDao struct{}

// Save saves a User entity to the database.
func (dao *userDao) Save(u *User) error {
	_, err := engine.Insert(u)
	return err
}

// FindAll finds all the User records in the database.
func (dao *userDao) FindAll() ([]User, error) {
	var all []User
	err := engine.Find(&all)
	return all, err
}

// FindByID finds the User record by UserID.
func (dao *userDao) FindByID(userID int) (*User, error) {
	var u = User{UserID: userID}
	has, err := engine.Get(&u)
	if !has {
		return nil, errors.New("该用户不存在")
	}
	return &u, err
}
