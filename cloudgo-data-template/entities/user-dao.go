package entities

import (
	"github.com/painterdrown/service-computing/cloudgo-data-template/sqlt"
)

type userDao struct {
	sqlt.SQLTemplate
}

var userInsertSQL = "INSERT user SET username = ?, created = ?"
var userFindByIDSQL = "SELECT * FROM user WHERE user.user_id = ?"
var userFindAll = "SELECT * FROM user"

// Save saves a User entity to the database.
func (dao *userDao) Save(u *User) (int, error) {
	return dao.Insert(userInsertSQL, u.Username, u.Created)
}

func (dao *userDao) FindByID(id int) *User {
	u := new(User)
	row := dao.QueryForRow(userFindByIDSQL, id)
	row.Scan(&u.UserID, &u.Username, &u.Balance, &u.Created)
	return u
}

func (dao *userDao) FindAll() ([]User, error) {
	all := make([]User, 0, 0)
	rows, err := dao.QueryForRows(userFindAll)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.UserID, &u.Username, &u.Balance, &u.Created)
		if err != nil {
			return nil, err
		}
		all = append(all, *u)
	}
	return all, err
}
