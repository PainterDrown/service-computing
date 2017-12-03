package entities

//UserAtomicService provides the atomic operations for User.
type UserAtomicService struct{}

//UserService is one entity of UserAtomicService.
var UserService = UserAtomicService{}

// Save saves a User entity to the database.
func (*UserAtomicService) Save(u *User) {
	dao := userDao{}
	err := dao.Save(u)
	checkErr(err)
}

// FindAll finds all the User records in the database.
func (*UserAtomicService) FindAll() []User {
	dao := userDao{}
	all, err := dao.FindAll()
	checkErr(err)
	return all
}

// FindByID finds the User record by UserID.
func (*UserAtomicService) FindByID(userID int) *User {
	dao := userDao{}
	u, err := dao.FindByID(userID)
	checkErr(err)
	return u
}
